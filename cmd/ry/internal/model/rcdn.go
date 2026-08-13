package model

import "time"

type RcdnInstance struct {
	ID          string    `json:"id"           table:"ID"`
	Status      string    `json:"status"       table:"STATUS"`
	Plan        string    `json:"plan"         table:"PLAN"`
	Tag         string    `json:"tag"          table:"TAG"`
	Region      string    `json:"region"       table:"REGION"`
	TrafficUsed int       `json:"traffic_used" table:"TRAFFIC GB"`
	AutoRenew   bool      `json:"auto_renew"   table:"AUTO RENEW"`
	CreatedAt   time.Time `json:"created_at"   table:"CREATED"`
	ExpireAt    time.Time `json:"expires"      table:"EXPIRES"`
}

type RcdnDomain struct {
	ID     string `json:"id"     table:"ID"`
	Domain string `json:"domain" table:"DOMAIN"`
	CNAME  string `json:"cname"  table:"CNAME"`
	Region string `json:"region" table:"REGION"`
}
