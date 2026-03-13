package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/open-nexa/nexa"
)

var uiRenderExamples = map[string]nexa.UIRender{}

func init() {
	uiRenderExamples["dashboard"] = nexa.UIRender{
		Intent: "Sales Dashboard - Q4 2024 Performance",
		Layout: &nexa.UILayoutSpec{
			Type:      nexa.UILayoutStack,
			Direction: "vertical",
			Gap:       20,
		},
		Widgets: []nexa.UIWidget{
			{Type: nexa.UIWidgetText, Content: "📊 Q4 2024 Sales Dashboard", Style: map[string]interface{}{"fontSize": 24, "fontWeight": "bold"}},
			{Type: nexa.UIWidgetAlert, Content: "🎉 Goal achieved! Total sales exceeded target by 15%", Data: map[string]interface{}{"alertType": "success"}},
			{Type: nexa.UIWidgetChart, Data: map[string]interface{}{"chartType": "bar", "labels": []string{"Oct", "Nov", "Dec"}, "values": []int{45000, 52000, 61000}}},
			{Type: nexa.UIWidgetTable, Data: map[string]interface{}{"headers": []string{"Product", "Q4 Sales", "Growth"}, "rows": [][]interface{}{
				{"Enterprise Suite", "$125,000", "+23%"},
				{"Pro Plan", "$89,000", "+18%"},
				{"Starter", "$45,000", "+12%"},
			}}},
			{Type: nexa.UIWidgetList, Data: map[string]interface{}{"items": []string{"📈 Revenue: $259,000 (+18% YoY)", "👥 New Customers: 1,247", "⏱️ Avg Response Time: 2.3s"}}},
		},
	}

	uiRenderExamples["stats"] = nexa.UIRender{
		Intent: "Statistics Cards Grid",
		Layout: &nexa.UILayoutSpec{
			Type:     nexa.UILayoutGrid,
			Columns:  4,
			Gap:      16,
		},
		Widgets: []nexa.UIWidget{
			{Type: nexa.UIWidgetText, Content: "📦 Total Orders", Style: map[string]interface{}{"fontWeight": "bold"}},
			{Type: nexa.UIWidgetText, Content: "1,234", Style: map[string]interface{}{"fontSize": 32, "color": "#2196F3"}},
			{Type: nexa.UIWidgetText, Content: "💰 Revenue", Style: map[string]interface{}{"fontWeight": "bold"}},
			{Type: nexa.UIWidgetText, Content: "$56,789", Style: map[string]interface{}{"fontSize": 32, "color": "#4CAF50"}},
			{Type: nexa.UIWidgetText, Content: "👥 Customers", Style: map[string]interface{}{"fontWeight": "bold"}},
			{Type: nexa.UIWidgetText, Content: "567", Style: map[string]interface{}{"fontSize": 32, "color": "#FF9800"}},
			{Type: nexa.UIWidgetText, Content: "⭐ Satisfaction", Style: map[string]interface{}{"fontWeight": "bold"}},
			{Type: nexa.UIWidgetText, Content: "4.8/5.0", Style: map[string]interface{}{"fontSize": 32, "color": "#9C27B0"}},
		},
	}

	uiRenderExamples["steps"] = nexa.UIRender{
		Intent: "Multi-step Process",
		Layout: &nexa.UILayoutSpec{
			Type: nexa.UILayoutStack,
			Gap:  24,
		},
		Widgets: []nexa.UIWidget{
			{Type: nexa.UIWidgetText, Content: "🔄 Data Processing Pipeline", Style: map[string]interface{}{"fontSize": 20, "fontWeight": "bold"}},
			{Type: nexa.UIWidgetSteps, Data: map[string]interface{}{
				"steps":  []string{"Step 1: Initialize", "Step 2: Process Data", "Step 3: Generate Report", "Step 4: Send Notification"},
				"active": 2,
			}},
			{Type: nexa.UIWidgetAlert, Content: "⚠️ Warning: Disk space is running low", Data: map[string]interface{}{"alertType": "warning"}},
		},
	}

	uiRenderExamples["split"] = nexa.UIRender{
		Intent: "Split Layout Demo",
		Layout: &nexa.UILayoutSpec{
			Type: nexa.UILayoutSplit,
			Gap:  20,
		},
		Widgets: []nexa.UIWidget{
			{
				Type:    nexa.UIWidgetText,
				Content: "📝 Description:\n\nThis is the left panel with detailed description text.\n\nYou can put any content here.",
				Style:   map[string]interface{}{"padding": 16, "backgroundColor": "#f5f5f5"},
			},
			{
				Type: nexa.UIWidgetChart,
				Data: map[string]interface{}{"chartType": "bar", "labels": []string{"A", "B", "C"}, "values": []int{30, 50, 20}},
			},
		},
	}

	uiRenderExamples["code"] = nexa.UIRender{
		Intent: "Code Example Display",
		Layout: &nexa.UILayoutSpec{
			Type: nexa.UILayoutStack,
			Gap:  16,
		},
		Widgets: []nexa.UIWidget{
			{Type: nexa.UIWidgetText, Content: "🔧 Go Code Example", Style: map[string]interface{}{"fontSize": 18, "fontWeight": "bold"}},
			{Type: nexa.UIWidgetCode, Content: `func main() {
    fmt.Println("Hello, Nexa UI!")
    
    msg := nexa.NewAgentMessage(agentInfo)
    msg.AddUIWidget(nexa.UIWidget{
        Type: nexa.UIWidgetText,
        Content: "Hello World!",
    })
}`, Data: map[string]interface{}{"language": "go"}},
		},
	}

	uiRenderExamples["alerts"] = nexa.UIRender{
		Intent: "Alert Types Demo",
		Layout: &nexa.UILayoutSpec{
			Type: nexa.UILayoutStack,
			Gap:  8,
		},
		Widgets: []nexa.UIWidget{
			{Type: nexa.UIWidgetAlert, Content: "🎉 Operation completed successfully!", Data: map[string]interface{}{"alertType": "success"}},
			{Type: nexa.UIWidgetAlert, Content: "⚠️ Warning: Disk space is running low", Data: map[string]interface{}{"alertType": "warning"}},
			{Type: nexa.UIWidgetAlert, Content: "❌ Error: Connection failed", Data: map[string]interface{}{"alertType": "error"}},
			{Type: nexa.UIWidgetAlert, Content: "ℹ️ Info: New version available", Data: map[string]interface{}{"alertType": "info"}},
		},
	}
}

func main() {
	port := 8888
	addr := fmt.Sprintf(":%d", port)

	dir := "."
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, dir+"/index.html")
			return
		}
		if _, err := os.Stat(dir + r.URL.Path); err == nil {
			http.ServeFile(w, r, dir+r.URL.Path)
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/api/ui/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		name := r.URL.Query().Get("name")
		if name == "" {
			list := []string{}
			for k := range uiRenderExamples {
				list = append(list, k)
			}
			json.NewEncoder(w).Encode(map[string]interface{}{
				"examples": list,
			})
			return
		}

		if render, ok := uiRenderExamples[name]; ok {
			json.NewEncoder(w).Encode(render)
		} else {
			http.Error(w, "Not found", 404)
		}
	})

	url := fmt.Sprintf("http://localhost:%d", port)

	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser(url)
	}()

	log.Printf("🚀 Nexa UI Server started at %s", url)
	log.Printf("📋 Available examples:")
	for name := range uiRenderExamples {
		log.Printf("   - %s: %s/api/ui/?name=%s", name, url, name)
	}
	log.Printf("\n💡 Open %s in your browser to view the UI!", url)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func openBrowser(url string) {
	var args []string
	switch runtime.GOOS {
	case "windows":
		args = []string{"cmd", "/c", "start", url}
	case "darwin":
		args = []string{"open", url}
	default:
		args = []string{"xdg-open", url}
	}

	if err := exec.Command(args[0], args[1:]...).Start(); err != nil {
		log.Printf("⚠️ Could not open browser automatically: %v", err)
	}
}
