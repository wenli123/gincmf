package model

type Paginate struct {
	Data     interface{} `json:"data"`
	Current  string      `json:"current"`
	PageSize string      `json:"page_size"`
	Total    int         `json:"total"`
}
