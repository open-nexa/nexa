package nexa

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/quic-go/quic-go"
)

type Client struct {
	conn    *quic.Conn
	streams map[uint64]quic.Stream
	mu      sync.RWMutex
	config  *ClientConfig
}

type ClientConfig struct {
	Addr        string
	TLSConfig   *tls.Config
	MaxIdleTime time.Duration
}

func NewClient(config *ClientConfig) (*Client, error) {
	if config.TLSConfig == nil {
		config.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	conn, err := quic.DialAddr(context.Background(), config.Addr, config.TLSConfig, &quic.Config{
		MaxIdleTimeout: config.MaxIdleTime,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		streams: make(map[uint64]quic.Stream),
		config:  config,
	}, nil
}

func (c *Client) Create(recordType MessageType, content string, metadata map[string]interface{}) (*Record, error) {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	msg := NewMessage()
	msg.Type = MsgRequest
	msg.Action = ActionCreate
	msg.Type = recordType
	msg.Content = content
	msg.Metadata = metadata

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return nil, err
	}

	var resp Message
	if err := json.NewDecoder(stream).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Type == MsgError {
		return nil, fmt.Errorf(resp.Error)
	}

	if record, ok := resp.Data.(*Record); ok {
		return record, nil
	}
	return nil, nil
}

func (c *Client) Get(id string) (*Record, error) {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	msg := NewMessage()
	msg.Action = ActionRead
	msg.ID = id

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return nil, err
	}

	var resp Message
	if err := json.NewDecoder(stream).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Type == MsgError {
		return nil, fmt.Errorf(resp.Error)
	}

	data, _ := json.Marshal(resp.Data)
	var record Record
	json.Unmarshal(data, &record)
	return &record, nil
}

func (c *Client) Update(id, content string, metadata map[string]interface{}) (*Record, error) {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	msg := NewMessage()
	msg.Action = ActionUpdate
	msg.ID = id
	msg.Content = content
	msg.Metadata = metadata

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return nil, err
	}

	var resp Message
	if err := json.NewDecoder(stream).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Type == MsgError {
		return nil, fmt.Errorf(resp.Error)
	}

	data, _ := json.Marshal(resp.Data)
	var record Record
	json.Unmarshal(data, &record)
	return &record, nil
}

func (c *Client) Delete(id string) error {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	defer stream.Close()

	msg := NewMessage()
	msg.Action = ActionDelete
	msg.ID = id

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return err
	}

	var resp Message
	if err := json.NewDecoder(stream).Decode(&resp); err != nil {
		return err
	}

	if resp.Type == MsgError {
		return fmt.Errorf(resp.Error)
	}
	return nil
}

func (c *Client) Search(recordType MessageType, query map[string]interface{}) ([]*Record, error) {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	msg := NewMessage()
	msg.Action = ActionSearch
	msg.Type = recordType
	msg.Query = query

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return nil, err
	}

	var resp Message
	if err := json.NewDecoder(stream).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.Type == MsgError {
		return nil, fmt.Errorf(resp.Error)
	}

	data, _ := json.Marshal(resp.Data)
	var results []*Record
	json.Unmarshal(data, &results)
	return results, nil
}

func (c *Client) OpenMediaStream(streamType StreamType) (io.WriteCloser, error) {
	stream, err := c.conn.OpenStreamSync(context.Background())
	if err != nil {
		return nil, err
	}

	msg := NewMessage()
	msg.Action = ActionOpenStream
	msg.StreamType = streamType

	if err := json.NewEncoder(stream).Encode(msg); err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *Client) Close() error {
	return c.conn.CloseWithError(0, "client closed")
}
