package nexa

type QueryType string

const (
	QueryKeyword   QueryType = "keyword"
	QueryVector    QueryType = "vector"
	QueryHybrid    QueryType = "hybrid"
	QueryTimeRange QueryType = "time_range"
	QueryMetadata  QueryType = "metadata"
	QuerySemantic  QueryType = "semantic"
)

type SortOrder string

const (
	SortDesc SortOrder = "desc"
	SortAsc  SortOrder = "asc"
)

type QueryOperator string

const (
	OpEq       QueryOperator = "eq"
	OpNe       QueryOperator = "ne"
	OpGt       QueryOperator = "gt"
	OpGte      QueryOperator = "gte"
	OpLt       QueryOperator = "lt"
	OpLte      QueryOperator = "lte"
	OpIn       QueryOperator = "in"
	OpContains QueryOperator = "contains"
	OpExists   QueryOperator = "exists"
)

type IndexType string

const (
	IndexPrimary  IndexType = "primary"
	IndexInverted IndexType = "inverted"
	IndexVector   IndexType = "vector"
	IndexTime     IndexType = "time"
	IndexMetadata IndexType = "metadata"
)

type SearchRequest struct {
	Query     string                 `json:"query,omitempty"`
	QueryType QueryType              `json:"query_type,omitempty"`
	Vector    []float64              `json:"vector,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Page      int                    `json:"page,omitempty"`
	PageSize  int                    `json:"page_size,omitempty"`
	SortBy    string                 `json:"sort_by,omitempty"`
	SortOrder SortOrder              `json:"sort_order,omitempty"`
	Filters   map[string]interface{} `json:"filters,omitempty"`
	TimeRange *TimeRange             `json:"time_range,omitempty"`
	Include   []string               `json:"include,omitempty"`
	Exclude   []string               `json:"exclude,omitempty"`
}

type TimeRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type SearchResponse struct {
	Total    int64        `json:"total"`
	Page     int          `json:"page"`
	PageSize int          `json:"page_size"`
	Results  []*Record    `json:"results"`
	Facets   []*Facet     `json:"facets,omitempty"`
	Indices  []*IndexInfo `json:"indices,omitempty"`
}

type Facet struct {
	Field  string        `json:"field"`
	Values []*FacetValue `json:"values"`
}

type FacetValue struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

type IndexInfo struct {
	Name     string    `json:"name"`
	Type     IndexType `json:"type"`
	DocCount int64     `json:"doc_count"`
	Size     int64     `json:"size"`
}

type IndexStats struct {
	TotalDocs  int64        `json:"total_docs"`
	TotalSize  int64        `json:"total_size"`
	IndexCount int64        `json:"index_count"`
	Indices    []*IndexInfo `json:"indices"`
}

type IndexRequest struct {
	Action    string                 `json:"action"`
	IndexName string                 `json:"index_name,omitempty"`
	IndexType IndexType              `json:"index_type,omitempty"`
	Config    map[string]interface{} `json:"config,omitempty"`
}

type IndexResponse struct {
	Success bool       `json:"success"`
	Index   *IndexInfo `json:"index,omitempty"`
	Error   string     `json:"error,omitempty"`
}
