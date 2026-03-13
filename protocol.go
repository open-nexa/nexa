package nexa

import "time"

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityNormal Priority = "normal"
	PriorityHigh   Priority = "high"
	PriorityUrgent Priority = "urgent"
)

type AgentRole string

const (
	RoleCoordinator AgentRole = "coordinator"
	RoleExecutor    AgentRole = "executor"
	RoleAnalyzer    AgentRole = "analyzer"
	RolePlanner     AgentRole = "planner"
	RoleAssistant   AgentRole = "assistant"
)

type AgentInfo struct {
	ID       string                 `json:"id"`
	Name     string                 `json:"name"`
	Role     AgentRole              `json:"role"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type SessionInfo struct {
	SessionID  string                 `json:"session_id"`
	StartTime  int64                  `json:"start_time"`
	LastActive int64                  `json:"last_active"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type ToolCall struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

type ToolResult struct {
	ID      string      `json:"id"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ContentType string

const (
	ContentTypeText   ContentType = "text"
	ContentTypeImage  ContentType = "image"
	ContentTypeAudio  ContentType = "audio"
	ContentTypeVideo  ContentType = "video"
	ContentTypeLive   ContentType = "live"
	ContentTypeStruct ContentType = "struct"
	ContentTypeBinary ContentType = "binary"
	ContentTypeStream ContentType = "stream"
	ContentTypeUI     ContentType = "ui"
)

type MediaContent struct {
	MediaType MediaType              `json:"media_type"`
	Data      []byte                 `json:"data,omitempty"`
	URL       string                 `json:"url,omitempty"`
	Format    string                 `json:"format"`
	Size      int64                  `json:"size"`
	Duration  int64                  `json:"duration,omitempty"`
	Width     int                    `json:"width,omitempty"`
	Height    int                    `json:"height,omitempty"`
	MimeType  string                 `json:"mime_type,omitempty"`
	Checksum  string                 `json:"checksum,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type UIWidgetType string

const (
	UIWidgetText     UIWidgetType = "text"
	UIWidgetImage    UIWidgetType = "image"
	UIWidgetMedia    UIWidgetType = "media"
	UIWidgetChart    UIWidgetType = "chart"
	UIWidgetTable    UIWidgetType = "table"
	UIWidgetList     UIWidgetType = "list"
	UIWidgetCode     UIWidgetType = "code"
	UIWidgetMath     UIWidgetType = "math"
	UIWidgetQuote    UIWidgetType = "quote"
	UIWidgetDivider  UIWidgetType = "divider"
	UIWidgetSpacer   UIWidgetType = "spacer"
	UIWidgetAlert    UIWidgetType = "alert"
	UIWidgetSteps    UIWidgetType = "steps"
	UIWidgetTimeline UIWidgetType = "timeline"
	UIWidgetCustom   UIWidgetType = "custom"
)

type UIAction struct {
	Type    string                 `json:"type"`
	Label   string                 `json:"label,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Confirm string                 `json:"confirm,omitempty"`
	URL     string                 `json:"url,omitempty"`
}

type UIEvent struct {
	Type        string                 `json:"type"`
	ComponentID string                 `json:"component_id,omitempty"`
	Action      *UIAction              `json:"action,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	Timestamp   int64                  `json:"timestamp"`
}

type UILayout string

const (
	UILayoutStack    UILayout = "stack"
	UILayoutGrid     UILayout = "grid"
	UILayoutMasonry  UILayout = "masonry"
	UILayoutSplit    UILayout = "split"
	UILayoutCentered UILayout = "centered"
	UILayoutFull     UILayout = "full"
)

type UIWidget struct {
	Type     UIWidgetType           `json:"type"`
	Intent   string                 `json:"intent,omitempty"`
	Content  string                 `json:"content,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
	URL      string                 `json:"url,omitempty"`
	Alt      string                 `json:"alt,omitempty"`
	Span     int                    `json:"span,omitempty"`
	Align    string                 `json:"align,omitempty"`
	Style    map[string]interface{} `json:"style,omitempty"`
	Actions  []UIAction             `json:"actions,omitempty"`
	Children []UIWidget             `json:"children,omitempty"`
}

type UILayoutSpec struct {
	Type      UILayout `json:"type"`
	Direction string   `json:"direction,omitempty"`
	Gap       int      `json:"gap,omitempty"`
	Columns   int      `json:"columns,omitempty"`
	Rows      int      `json:"rows,omitempty"`
	Wrap      bool     `json:"wrap,omitempty"`
	Align     string   `json:"align,omitempty"`
	Justify   string   `json:"justify,omitempty"`
}

type UIRender struct {
	Version  string                 `json:"version,omitempty"`
	Layout   *UILayoutSpec          `json:"layout,omitempty"`
	Widgets  []UIWidget             `json:"widgets"`
	Intent   string                 `json:"intent,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type MessageContent struct {
	Type      ContentType            `json:"type"`
	Text      string                 `json:"text,omitempty"`
	Data      []byte                 `json:"data,omitempty"`
	Media     *MediaContent          `json:"media,omitempty"`
	MediaList []MediaContent         `json:"media_list,omitempty"`
	UIRender  *UIRender              `json:"ui_render,omitempty"`
	Struct    interface{}            `json:"struct,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
}

type MessageStatus string

const (
	StatusPending   MessageStatus = "pending"
	StatusSent      MessageStatus = "sent"
	StatusDelivered MessageStatus = "delivered"
	StatusRead      MessageStatus = "read"
	StatusProcessed MessageStatus = "processed"
	StatusFailed    MessageStatus = "failed"
)

type AgentMessage struct {
	Type        MessageType            `json:"type"`
	ID          string                 `json:"id"`
	RequestID   string                 `json:"request_id,omitempty"`
	SessionID   string                 `json:"session_id,omitempty"`
	From        AgentInfo              `json:"from"`
	To          []AgentInfo            `json:"to,omitempty"`
	Priority    Priority               `json:"priority,omitempty"`
	Status      MessageStatus          `json:"status,omitempty"`
	Action      Action                 `json:"action,omitempty"`
	Contents    []MessageContent       `json:"contents,omitempty"`
	ToolCalls   []ToolCall             `json:"tool_calls,omitempty"`
	ToolResults []ToolResult           `json:"tool_results,omitempty"`
	Query       *SearchRequest         `json:"query,omitempty"`
	Records     []*Record              `json:"records,omitempty"`
	Data        interface{}            `json:"data,omitempty"`
	Error       string                 `json:"error,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Timestamp   int64                  `json:"timestamp"`
	Compressed  bool                   `json:"compressed,omitempty"`
	Compression string                 `json:"compression,omitempty"`
	NeedAck     bool                   `json:"need_ack,omitempty"`
	TTL         int64                  `json:"ttl,omitempty"`
	RetryCount  int                    `json:"retry_count,omitempty"`
	MaxRetries  int                    `json:"max_retries,omitempty"`
}

func NewAgentMessage(from AgentInfo) AgentMessage {
	return AgentMessage{
		Type:      MsgRequest,
		From:      from,
		Timestamp: time.Now().UnixMilli(),
		Priority:  PriorityNormal,
	}
}

func (m *AgentMessage) SetTextContent(text string) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentTypeText,
		Text: text,
	})
}

func (m *AgentMessage) AddMedia(mediaType MediaType, data []byte, format string) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentType(mediaType),
		Media: &MediaContent{
			MediaType: mediaType,
			Data:      data,
			Format:    format,
		},
	})
}

func (m *AgentMessage) AddMediaFromURL(mediaType MediaType, url, format string) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentType(mediaType),
		Media: &MediaContent{
			MediaType: mediaType,
			URL:       url,
			Format:    format,
		},
	})
}

func (m *AgentMessage) AddMixedContent(text string, mediaList []MediaContent) {
	m.Contents = append(m.Contents, MessageContent{
		Type:      ContentTypeText,
		Text:      text,
		MediaList: mediaList,
	})
}

func (m *AgentMessage) SetUIRender(render UIRender) {
	m.Contents = append(m.Contents, MessageContent{
		Type:     ContentTypeUI,
		UIRender: &render,
	})
}

func (m *AgentMessage) AddUIWidget(widget UIWidget) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentTypeUI,
		UIRender: &UIRender{
			Widgets: []UIWidget{widget},
		},
	})
}

func (m *AgentMessage) AddUIText(text string, style map[string]interface{}) {
	m.AddUIWidget(UIWidget{
		Type:    UIWidgetText,
		Content: text,
		Style:   style,
	})
}

func (m *AgentMessage) AddUIImageWidget(url, alt string) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetImage,
		URL:  url,
		Alt:  alt,
	})
}

func (m *AgentMessage) AddUIMedia(url string, mediaType string) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetMedia,
		URL:  url,
		Data: map[string]interface{}{"mediaType": mediaType},
	})
}

func (m *AgentMessage) AddUIChartWidget(chartType string, data map[string]interface{}) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetChart,
		Data: map[string]interface{}{"chartType": chartType, "data": data},
	})
}

func (m *AgentMessage) AddUITableWidget(headers []string, rows [][]interface{}) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetTable,
		Data: map[string]interface{}{"headers": headers, "rows": rows},
	})
}

func (m *AgentMessage) AddUIListWidget(items []string) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetList,
		Data: map[string]interface{}{"items": items},
	})
}

func (m *AgentMessage) AddUICode(code, language string) {
	m.AddUIWidget(UIWidget{
		Type:    UIWidgetCode,
		Content: code,
		Data:    map[string]interface{}{"language": language},
	})
}

func (m *AgentMessage) AddUIMath(formula string, display bool) {
	m.AddUIWidget(UIWidget{
		Type:    UIWidgetMath,
		Content: formula,
		Data:    map[string]interface{}{"display": display},
	})
}

func (m *AgentMessage) AddUIQuote(content, author string) {
	m.AddUIWidget(UIWidget{
		Type:    UIWidgetQuote,
		Content: content,
		Data:    map[string]interface{}{"author": author},
	})
}

func (m *AgentMessage) AddUIAlert(message, alertType string) {
	m.AddUIWidget(UIWidget{
		Type:    UIWidgetAlert,
		Content: message,
		Data:    map[string]interface{}{"alertType": alertType},
	})
}

func (m *AgentMessage) AddUISteps(steps []string, active int) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetSteps,
		Data: map[string]interface{}{"steps": steps, "active": active},
	})
}

func (m *AgentMessage) AddUITimeline(events []map[string]interface{}) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetTimeline,
		Data: map[string]interface{}{"events": events},
	})
}

func (m *AgentMessage) AddUIDivider() {
	m.AddUIWidget(UIWidget{Type: UIWidgetDivider})
}

func (m *AgentMessage) AddUISpacer(size int) {
	m.AddUIWidget(UIWidget{
		Type:  UIWidgetSpacer,
		Style: map[string]interface{}{"size": size},
	})
}

func (m *AgentMessage) AddUIWidgets(layout UILayout, widgets []UIWidget) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentTypeUI,
		UIRender: &UIRender{
			Layout:  &UILayoutSpec{Type: layout},
			Widgets: widgets,
		},
	})
}

func (m *AgentMessage) AddUIGridLayout(columns int, widgets []UIWidget) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentTypeUI,
		UIRender: &UIRender{
			Layout: &UILayoutSpec{
				Type:    UILayoutGrid,
				Columns: columns,
			},
			Widgets: widgets,
		},
	})
}

func (m *AgentMessage) AddUIStackLayout(direction string, widgets []UIWidget) {
	m.Contents = append(m.Contents, MessageContent{
		Type: ContentTypeUI,
		UIRender: &UIRender{
			Layout: &UILayoutSpec{
				Type:      UILayoutStack,
				Direction: direction,
			},
			Widgets: widgets,
		},
	})
}

func (m *AgentMessage) AddUISplitLayout(left, right UIWidget) {
	m.AddUIWidgets(UILayoutSplit, []UIWidget{left, right})
}

func (m *AgentMessage) AddUICustomWidget(widgetType string, data map[string]interface{}) {
	m.AddUIWidget(UIWidget{
		Type: UIWidgetCustom,
		Data: map[string]interface{}{"widgetType": widgetType, "data": data},
	})
}

func (m *AgentMessage) AddToolCall(name string, arguments map[string]interface{}) string {
	callID := generateID()
	m.ToolCalls = append(m.ToolCalls, ToolCall{
		ID:        callID,
		Name:      name,
		Arguments: arguments,
	})
	return callID
}

func (m *AgentMessage) SetResponse(to AgentInfo, data interface{}) {
	m.Type = MsgResponse
	m.To = append(m.To, to)
	m.Data = data
	m.Timestamp = time.Now().UnixMilli()
}

func (m *AgentMessage) SetError(to AgentInfo, err string) {
	m.Type = MsgError
	m.To = append(m.To, to)
	m.Error = err
	m.Timestamp = time.Now().UnixMilli()
}

func (m *AgentMessage) SetStatus(status MessageStatus) {
	m.Status = status
}

type BatchOperation struct {
	Operation Action      `json:"operation"`
	Records   []*Record   `json:"records,omitempty"`
	IDs       []string    `json:"ids,omitempty"`
	Updates   interface{} `json:"updates,omitempty"`
}

type BatchRequest struct {
	Operations []BatchOperation `json:"operations"`
	SessionID  string           `json:"session_id,omitempty"`
	Atomic     bool             `json:"atomic,omitempty"`
}

type BatchResult struct {
	Operation Action    `json:"operation"`
	Success   int64     `json:"success"`
	Failed    int64     `json:"failed"`
	Results   []*Record `json:"results,omitempty"`
	Errors    []string  `json:"errors,omitempty"`
}

type BatchResponse struct {
	Results   []BatchResult `json:"results"`
	Total     int64         `json:"total"`
	Success   int64         `json:"success"`
	Failed    int64         `json:"failed"`
	SessionID string        `json:"session_id,omitempty"`
}

type SubscribeRequest struct {
	SessionID string                 `json:"session_id"`
	Events    []string               `json:"events"`
	Filters   map[string]interface{} `json:"filters,omitempty"`
}

type SubscribeResponse struct {
	Success        bool   `json:"success"`
	SubscriptionID string `json:"subscription_id,omitempty"`
	Error          string `json:"error,omitempty"`
}

type EventType string

const (
	EventRecordCreated EventType = "record_created"
	EventRecordUpdated EventType = "record_updated"
	EventRecordDeleted EventType = "record_deleted"
	EventToolCompleted EventType = "tool_completed"
	EventAgentJoined   EventType = "agent_joined"
	EventAgentLeft     EventType = "agent_left"
	EventSessionEnded  EventType = "session_ended"
)

type Event struct {
	Type      EventType   `json:"type"`
	ID        string      `json:"id"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	SessionID string      `json:"session_id,omitempty"`
}

type AckMessage struct {
	MessageID string        `json:"message_id"`
	Status    MessageStatus `json:"status"`
	Timestamp int64         `json:"timestamp"`
}

func NewAckMessage(msgID string, status MessageStatus) AckMessage {
	return AckMessage{
		MessageID: msgID,
		Status:    status,
		Timestamp: time.Now().UnixMilli(),
	}
}

type CompressionType string

const (
	CompressionGzip   CompressionType = "gzip"
	CompressionSnappy CompressionType = "snappy"
	CompressionZstd   CompressionType = "zstd"
	CompressionNone   CompressionType = "none"
)

func (m *AgentMessage) Compress(ctype CompressionType) error {
	if ctype == CompressionNone {
		m.Compressed = false
		m.Compression = ""
		return nil
	}
	m.Compressed = true
	m.Compression = string(ctype)
	return nil
}

func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
