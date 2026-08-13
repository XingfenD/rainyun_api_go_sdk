package model

import "time"

type News struct {
	Type  string    `json:"type"  table:"TYPE"`
	Title string    `json:"title" table:"TITLE"`
	Time  time.Time `json:"time"  table:"TIME"`
	URL   string    `json:"url"   table:"URL"`
}

type AppConfig struct {
	Type    string `json:"type"    table:"TYPE"`
	Title   string `json:"title"   table:"TITLE"`
	Name    string `json:"name"    table:"NAME"`
	Page    string `json:"page"    table:"PAGE"`
	Order   int    `json:"order"   table:"ORDER"`
	Content string `json:"content"`
}

type NodeStatus struct {
	UUID    string  `json:"uuid"    table:"UUID"`
	Name    string  `json:"name"    table:"NAME"`
	Product string  `json:"product" table:"PRODUCT"`
	CPU     float64 `json:"cpu"     table:"CPU"`
	Memory  float64 `json:"memory"  table:"MEM"`
	NetOut  int     `json:"net_out" table:"NET OUT"`
	Status  string  `json:"status"  table:"STATUS"`
	Updated string  `json:"updated" table:"UPDATED"`
}
