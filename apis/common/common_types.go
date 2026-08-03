package common

type StandQueryParameters struct {
	ColumnFilters struct{} `json:"columnFilters"`
	Sort          []any    `json:"sort"`
	Page          int      `json:"page"`
	PerPage       int      `json:"perPage"`
}
