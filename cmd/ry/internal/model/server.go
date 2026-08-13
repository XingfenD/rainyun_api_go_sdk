package model

import "time"

type Server struct {
	ID       string    `json:"id"       table:"ID"`
	Name     string    `json:"name"     table:"NAME"`
	Status   string    `json:"status"   table:"STATUS"`
	IP       string    `json:"ip"       table:"IP"`
	CPU      int       `json:"cpu"      table:"CPU"`
	Memory   int       `json:"memory"   table:"MEM"`
	Disk     int       `json:"disk"     table:"DISK"`
	OS       string    `json:"os"       table:"OS"`
	Region   string    `json:"region"   table:"REGION"`
	ExpireAt time.Time `json:"expires"  table:"EXPIRES"`
}
