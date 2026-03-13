package nexa

type MultiSearchRequest struct {
	Searches []SearchRequest `json:"searches"`
}

type MultiSearchResponse struct {
	Responses []SearchResponse `json:"responses"`
}

type PaginationOptions struct {
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
	Cursor       string `json:"cursor,omitempty"`
	IncludeTotal bool   `json:"include_total,omitempty"`
}

type Projection struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type OptimizedSearchRequest struct {
	Query      string                 `json:"query,omitempty"`
	QueryType  QueryType              `json:"query_type,omitempty"`
	Vector     []float64              `json:"vector,omitempty"`
	Filters    map[string]interface{} `json:"filters,omitempty"`
	TimeRange  *TimeRange             `json:"time_range,omitempty"`
	Types      []string               `json:"types,omitempty"`
	Pagination PaginationOptions      `json:"pagination,omitempty"`
	Projection Projection             `json:"projection,omitempty"`
	SortBy     string                 `json:"sort_by,omitempty"`
	SortOrder  SortOrder              `json:"sort_order,omitempty"`
	CacheTTL   int64                  `json:"cache_ttl,omitempty"`
	UseCache   bool                   `json:"use_cache,omitempty"`
}

type SearchSuggestion struct {
	Text    string  `json:"text"`
	Score   float64 `json:"score,omitempty"`
	Matches int64   `json:"matches,omitempty"`
}

type OptimizedSearchResponse struct {
	Results      []*Record              `json:"results"`
	Total        int64                  `json:"total"`
	Page         int                    `json:"page,omitempty"`
	PageSize     int                    `json:"page_size,omitempty"`
	HasMore      bool                   `json:"has_more,omitempty"`
	NextCursor   string                 `json:"next_cursor,omitempty"`
	Facets       []*Facet               `json:"facets,omitempty"`
	Aggregations map[string]interface{} `json:"aggregations,omitempty"`
	Suggestions  []SearchSuggestion     `json:"suggestions,omitempty"`
	QueryTimeMs  int64                  `json:"query_time_ms,omitempty"`
}

type VectorSearchOptions struct {
	TopK            int            `json:"top_k,omitempty"`
	Threshold       float64        `json:"threshold,omitempty"`
	Metric          string         `json:"metric,omitempty"`
	IncludeMetadata bool           `json:"include_metadata,omitempty"`
	HybridWeights   *HybridWeights `json:"hybrid_weights,omitempty"`
}

type HybridWeights struct {
	VectorWeight   float64 `json:"vector_weight,omitempty"`
	KeywordWeight  float64 `json:"keyword_weight,omitempty"`
	SemanticWeight float64 `json:"semantic_weight,omitempty"`
}

type BulkIndexRequest struct {
	Records []*Record `json:"records"`
	Refresh bool      `json:"refresh,omitempty"`
}

type BulkIndexResponse struct {
	Success int64    `json:"success"`
	Failed  int64    `json:"failed"`
	Errors  []string `json:"errors,omitempty"`
}

type StreamOptions struct {
	BatchSize  int `json:"batch_size,omitempty"`
	ThrottleMs int `json:"throttle_ms,omitempty"`
}

type VectorSearchRequest struct {
	Vector  []float64              `json:"vector"`
	Options VectorSearchOptions    `json:"options,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Type    string                 `json:"type,omitempty"`
}

type VectorSearchResult struct {
	ID     string  `json:"id"`
	Score  float64 `json:"score"`
	Record *Record `json:"record,omitempty"`
}

type VectorSearchResponse struct {
	Results []VectorSearchResult `json:"results"`
	TimeMs  int64                `json:"time_ms,omitempty"`
}
