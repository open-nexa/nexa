package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/open-nexa/nexa"
)

func main() {
	agentInfo := nexa.AgentInfo{
		ID:   "agent-001",
		Name: "OpenClaw-Executor",
		Role: nexa.RoleExecutor,
		Metadata: map[string]interface{}{
			"capabilities": []string{"web_search", "file_read", "code_execute"},
			"version":      "1.0.0",
		},
	}

	sessionInfo := nexa.SessionInfo{
		SessionID:  "session-12345",
		StartTime:  time.Now().Unix(),
		LastActive: time.Now().Unix(),
		Metadata:   map[string]interface{}{},
	}

	msg := nexa.AgentMessage{
		Type:      nexa.MsgRequest,
		ID:        "msg-001",
		RequestID: "req-001",
		SessionID: sessionInfo.SessionID,
		From:      agentInfo,
		Priority:  nexa.PriorityNormal,
		Status:    nexa.StatusPending,
		Action:    nexa.ActionCreate,
		Contents: []nexa.MessageContent{
			{
				Type: nexa.ContentTypeText,
				Text: "Please execute the following task: search for information about AI agents",
			},
		},
		ToolCalls: []nexa.ToolCall{
			{
				ID:   "tool-001",
				Name: "web_search",
				Arguments: map[string]interface{}{
					"query": "AI agents recent developments 2024",
				},
			},
		},
		Timestamp: time.Now().Unix(),
		NeedAck:   true,
	}

	data, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal message: %v\n", err)
		return
	}
	fmt.Println("=== Agent Message Example ===")
	fmt.Println(string(data))

	responseMsg := nexa.AgentMessage{
		Type:      nexa.MsgResponse,
		ID:        "msg-response-001",
		RequestID: "msg-001",
		SessionID: sessionInfo.SessionID,
		From:      agentInfo,
		Status:    nexa.StatusProcessed,
		Contents: []nexa.MessageContent{
			{
				Type: nexa.ContentTypeText,
				Text: "Task completed successfully. Found 10 relevant articles about AI agents.",
			},
		},
		ToolResults: []nexa.ToolResult{
			{
				ID:      "tool-001",
				Success: true,
				Result: map[string]interface{}{
					"articles": []string{
						"https://example.com/article1",
						"https://example.com/article2",
					},
					"total": 10,
				},
			},
		},
		Timestamp: time.Now().Unix(),
	}

	data, err = json.MarshalIndent(responseMsg, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal response: %v\n", err)
		return
	}
	fmt.Println("\n=== Agent Response Example ===")
	fmt.Println(string(data))

	pushMsg := nexa.AgentMessage{
		Type:     nexa.MsgPushed,
		ID:       "msg-push-001",
		From:     agentInfo,
		Priority: nexa.PriorityHigh,
		Status:   nexa.StatusSent,
		Action:   nexa.ActionCreate,
		Contents: []nexa.MessageContent{
			{
				Type: nexa.ContentTypeStruct,
				Struct: map[string]interface{}{
					"event_type": "task_assigned",
					"task_id":    "task-456",
					"assignee":   "agent-001",
				},
			},
		},
		Timestamp: time.Now().Unix(),
	}

	data, err = json.MarshalIndent(pushMsg, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal push message: %v\n", err)
		return
	}
	fmt.Println("\n=== Agent Push Message Example ===")
	fmt.Println(string(data))

	coordinatorInfo := nexa.AgentInfo{
		ID:   "agent-coordinator",
		Name: "OpenClaw-Coordinator",
		Role: nexa.RoleCoordinator,
	}

	broadcastMsg := nexa.AgentMessage{
		Type:      nexa.MsgRequest,
		ID:        "msg-broadcast-001",
		SessionID: sessionInfo.SessionID,
		From:      coordinatorInfo,
		To: []nexa.AgentInfo{
			agentInfo,
			{ID: "agent-analyzer", Name: "OpenClaw-Analyzer", Role: nexa.RoleAnalyzer},
			{ID: "agent-planner", Name: "OpenClaw-Planner", Role: nexa.RolePlanner},
		},
		Priority:  nexa.PriorityUrgent,
		Action:    nexa.ActionCreate,
		Timestamp: time.Now().Unix(),
	}

	data, err = json.MarshalIndent(broadcastMsg, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal broadcast: %v\n", err)
		return
	}
	fmt.Println("\n=== Coordinator Broadcast Example ===")
	fmt.Println(string(data))

	fmt.Println("\n=== Protocol Examples Complete ===")
}
