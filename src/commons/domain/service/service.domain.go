package serviceDomain

type PaginationData[D interface{}] struct {
	Pages   int `json:"pages"`
	Page    int `json:"page"`
	Data    []D `json:"data"`
	Records int `json:"records"`
}

type PaginationOpts struct {
	Page int
}
