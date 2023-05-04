package persistence

type PaginatedRequest struct {
	Limit  *int
	Offset *int
}

type PaginatedResult[T interface{}] struct {
	Results []T
	Total   int
	Limit   int
	Offset  int
}
