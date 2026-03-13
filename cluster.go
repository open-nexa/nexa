package nexa

type NodeRole string

const (
	RoleLeader   NodeRole = "leader"
	RoleFollower NodeRole = "follower"
	RoleGateway  NodeRole = "gateway"
)

type NodeStatus string

const (
	NodeStatusActive   NodeStatus = "active"
	NodeStatusInactive NodeStatus = "inactive"
	NodeStatusJoining  NodeStatus = "joining"
	NodeStatusLeaving  NodeStatus = "leaving"
)

type Node struct {
	ID       string                 `json:"id"`
	Addr     string                 `json:"addr"`
	Port     int                    `json:"port"`
	Role     NodeRole               `json:"role"`
	Status   NodeStatus             `json:"status"`
	Weight   int                    `json:"weight"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type ClusterConfig struct {
	NodeID            string   `json:"node_id"`
	ClusterID         string   `json:"cluster_id"`
	SeedNodes         []string `json:"seed_nodes"`
	JoinTimeout       int64    `json:"join_timeout"`
	HeartbeatInterval int64    `json:"heartbeat_interval"`
}

type NodeMessage struct {
	Type      NodeMessageType `json:"type"`
	From      string          `json:"from"`
	To        string          `json:"to,omitempty"`
	Timestamp int64           `json:"timestamp"`
	Data      interface{}     `json:"data,omitempty"`
}

type NodeMessageType string

const (
	NodeMsgJoin      NodeMessageType = "node_join"
	NodeMsgLeave     NodeMessageType = "node_leave"
	NodeMsgHeartbeat NodeMessageType = "node_heartbeat"
	NodeMsgSync      NodeMessageType = "node_sync"
	NodeMsgRedirect  NodeMessageType = "node_redirect"
)

type Shard struct {
	ID       string   `json:"id"`
	Replicas []string `json:"replicas"`
	Master   string   `json:"master"`
}

type ClusterState struct {
	Nodes   map[string]*Node  `json:"nodes"`
	Shards  map[string]*Shard `json:"shards"`
	Version int64             `json:"version"`
}
