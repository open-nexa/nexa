package nexa

type StreamType string

const (
	StreamControl StreamType = "control"
	StreamMessage StreamType = "message"
	StreamAudio   StreamType = "audio"
	StreamVideo   StreamType = "video"
	StreamFile    StreamType = "file"
	StreamLive    StreamType = "live"
)

type MediaType string

const (
	MediaTypeImage      MediaType = "image"
	MediaTypeAudio      MediaType = "audio"
	MediaTypeVideo      MediaType = "video"
	MediaTypeLiveStream MediaType = "live_stream"
	MediaTypeDocument   MediaType = "document"
)

type MessageType string

const (
	MsgRequest  MessageType = "request"
	MsgResponse MessageType = "response"
	MsgError    MessageType = "error"
	MsgPushed   MessageType = "pushed"
)

type Action string

const (
	ActionCreate      Action = "create"
	ActionPublish     Action = "publish"
	ActionRead        Action = "read"
	ActionRetrieve    Action = "retrieve"
	ActionUpdate      Action = "update"
	ActionModify      Action = "modify"
	ActionDelete      Action = "delete"
	ActionRemove      Action = "remove"
	ActionSearch      Action = "search"
	ActionQuery       Action = "query"
	ActionList        Action = "list"
	ActionOpenStream  Action = "open_stream"
	ActionCloseStream Action = "close_stream"
	ActionPing        Action = "ping"
	ActionPong        Action = "pong"
	ActionIndexCreate Action = "index_create"
	ActionIndexDelete Action = "index_delete"
	ActionIndexStats  Action = "index_stats"
	ActionReindex     Action = "reindex"
	ActionBatchCreate Action = "batch_create"
	ActionBatchUpdate Action = "batch_update"
	ActionBatchDelete Action = "batch_delete"
	ActionSave        Action = "save"
	ActionLoad        Action = "load"
)
