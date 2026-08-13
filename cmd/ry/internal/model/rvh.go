package model

type RvhPlan struct {
	ID      int    `json:"id"      table:"ID"`
	Plan    string `json:"plan"    table:"PLAN"`
	Chinese string `json:"chinese" table:"NAME"`
	Price   int    `json:"price"   table:"PRICE"`
	Disk    int    `json:"disk_mb" table:"DISK MB"`
}
