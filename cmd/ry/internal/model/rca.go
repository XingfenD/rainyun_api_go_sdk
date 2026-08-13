package model

type RcaProject struct {
	ID     string `json:"id"      table:"ID"`
	Name   string `json:"name"    table:"NAME"`
	Status string `json:"status"  table:"STATUS"`
	Region string `json:"region"  table:"REGION"`
	MaxCPU int    `json:"max_cpu" table:"MAX CPU"`
	MaxMem int    `json:"max_mem" table:"MAX MEM"`
}

type RcaRaindropBalance struct {
	Balance float64 `json:"balance" table:"BALANCE"`
}

type RcaRaindropPlan struct {
	ID       int    `json:"id"        table:"ID"`
	PlanName string `json:"plan_name" table:"PLAN"`
	Chinese  string `json:"chinese"   table:"NAME"`
	Amount   int    `json:"amount"    table:"AMOUNT"`
	Price    int    `json:"price"     table:"PRICE"`
}

type RcaRegion struct {
	ID          int    `json:"id"           table:"ID"`
	Name        string `json:"name"         table:"NAME"`
	ChineseName string `json:"chinese_name" table:"CHINESE"`
}
