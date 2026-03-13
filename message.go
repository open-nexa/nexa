package nexa

import (
	"time"
)

type Record struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Content   string                 `json:"content,omitempty"`
	Media     *MediaContent          `json:"media,omitempty"`
	MediaList []MediaContent         `json:"media_list,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt int64                  `json:"created_at"`
	UpdatedAt int64                  `json:"updated_at"`
}

type Message struct {
	Type       MessageType            `json:"type"`
	Action     Action                 `json:"action,omitempty"`
	ID         string                 `json:"id,omitempty"`
	RequestID  string                 `json:"request_id,omitempty"`
	StreamID   uint64                 `json:"stream_id,omitempty"`
	StreamType StreamType             `json:"stream_type,omitempty"`
	RecordType string                 `json:"record_type,omitempty"`
	Content    string                 `json:"content,omitempty"`
	Media      *MediaContent          `json:"media,omitempty"`
	MediaList  []MediaContent         `json:"media_list,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Query      map[string]interface{} `json:"query,omitempty"`
	Data       interface{}            `json:"data,omitempty"`
	Error      string                 `json:"error,omitempty"`
	Timestamp  int64                  `json:"timestamp"`
	Chunk      []byte                 `json:"chunk,omitempty"`
	Final      bool                   `json:"final,omitempty"`
}

func NewMessage() Message {
	return Message{
		Timestamp: time.Now().UnixMilli(),
	}
}

func (m *Message) SetResponse(action Action, data interface{}) {
	m.Type = MsgResponse
	m.Action = action
	m.Data = data
	m.Timestamp = time.Now().UnixMilli()
}

func (m *Message) SetError(err string) {
	m.Type = MsgError
	m.Error = err
	m.Timestamp = time.Now().UnixMilli()
}

func (m *Message) SetErrorWithRequestID(err string, reqID string) {
	m.Type = MsgError
	m.Error = err
	m.Timestamp = time.Now().UnixMilli()
	m.RequestID = reqID
}
