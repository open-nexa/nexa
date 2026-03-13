package protocol

type Permission string

const (
	PermRead    Permission = "read"
	PermWrite   Permission = "write"
	PermDelete  Permission = "delete"
	PermAdmin   Permission = "admin"
	PermExecute Permission = "execute"
)

type ResourceType string

const (
	ResourceTypeDoc      ResourceType = "doc"
	ResourceTypeStream   ResourceType = "stream"
	ResourceTypeIndex    ResourceType = "index"
	ResourceTypeCluster  ResourceType = "cluster"
	ResourceTypeAdmin    ResourceType = "admin"
)

type ACL struct {
	ID          string            `json:"id"`
	Subjects    []string         `json:"subjects"`
	Resources   []Resource       `json:"resources"`
	Permissions []Permission     `json:"permissions"`
	Effect      ACLEffect        `json:"effect"`
	Priority    int              `json:"priority"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type Resource struct {
	Type    ResourceType `json:"type"`
	Pattern string      `json:"pattern"`
}

type ACLEffect string

const (
	ACLAllow ACLEffect = "allow"
	ACLDeny  ACLEffect = "deny"
)

type Role struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
	Resources   []Resource   `json:"resources"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type User struct {
	ID       string            `json:"id"`
	Username string            `json:"username"`
	Roles    []string         `json:"roles"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type AuthToken struct {
	UserID    string            `json:"user_id"`
	Username  string           `json:"username"`
	Roles     []string        `json:"roles"`
	IssuedAt  int64           `json:"issued_at"`
	ExpiresAt int64           `json:"expires_at"`
	Scopes    []string        `json:"scopes"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
	User      *User  `json:"user,omitempty"`
}
