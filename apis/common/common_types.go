package common

type StandQueryParameters struct {
	ColumnFilters struct{} `json:"columnFilters"`
	Sort          []any    `json:"sort"`
	Page          int      `json:"page"`
	PerPage       int      `json:"perPage"`
}

type BasicOperationResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}
