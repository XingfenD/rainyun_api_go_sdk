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

type ServerEDisk struct {
	Slot   int    `json:"slot"`
	Type   string `json:"type"`
	Size   int    `json:"size_gb"`
	Backup bool   `json:"backup"`
}

type ServerEIP struct {
	IP          string `json:"ip"`
	Region      string `json:"region"`
	Gateway     string `json:"gateway"`
	Description string `json:"description"`
}

type ServerBackup struct {
	Label      string    `json:"label"`
	FileName   string    `json:"file_name"`
	SizeBytes  int64     `json:"size_bytes"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	FinishedAt time.Time `json:"finished_at"`
}

type ServerPlan struct {
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
	Price  int    `json:"price"`
}

type RenewPrice struct {
	Price    float64 `json:"price"        table:"PRICE"`
	Renew    float64 `json:"renew"        table:"RENEW"`
	RenewEIP float64 `json:"renew_eip"    table:"RENEW EIP"`
	Coupon   int     `json:"coupon_value" table:"COUPON VALUE"`
}

type ServerDetail struct {
	ID              string     `json:"id"               table:"ID"`
	Name            string     `json:"name"             table:"NAME"`
	Status          string     `json:"status"           table:"STATUS"`
	IP              string     `json:"ip"               table:"IP"`
	IntranetIP      string     `json:"intranet_ip"      table:"INTRANET IP"`
	OS              string     `json:"os"               table:"OS"`
	Region          string     `json:"region"           table:"REGION"`
	Tag             string     `json:"tag"              table:"TAG"`
	CPU             int        `json:"cpu"              table:"CPU"`
	Memory          int        `json:"memory"           table:"MEM"`
	Disk            int        `json:"disk"             table:"DISK"`
	NetMode         string     `json:"net_mode"         table:"NET MODE"`
	BandwidthIn     int        `json:"bandwidth_in"     table:"NET IN"`
	BandwidthOut    int        `json:"bandwidth_out"    table:"NET OUT"`
	NatPublicIP     string     `json:"nat_public_ip"    table:"NAT IP"`
	NatPublicDomain string     `json:"nat_public_domain" table:"NAT DOMAIN"`
	CPUUsage        float64    `json:"cpu_usage"        table:"CPU USAGE %"`
	UsedMem         int        `json:"used_mem"         table:"USED MEM"`
	FreeMem         int64      `json:"free_mem"         table:"FREE MEM"`
	MaxMem          int64      `json:"max_mem"          table:"MAX MEM"`
	NetInNow        float64    `json:"net_in_now"       table:"NET IN NOW"`
	NetOutNow       float64    `json:"net_out_now"      table:"NET OUT NOW"`
	DiskTemp        int        `json:"disk_temp"        table:"DISK TEMP"`
	TrafficUsed     int64      `json:"traffic_used"     table:"TRAFFIC USED"`
	TrafficToday    int        `json:"traffic_today"    table:"TRAFFIC TODAY"`
	TrafficLimit    int64      `json:"traffic_limit"    table:"TRAFFIC LIMIT"`
	TrafficOnLimit  int        `json:"traffic_on_limit" table:"ON LIMIT"`
	TrafficResetAt  *time.Time `json:"traffic_reset"    table:"TRAFFIC RESET"`
	AutoRenew       bool       `json:"auto_renew"       table:"AUTO RENEW"`
	CreatedAt       time.Time  `json:"created_at"       table:"CREATED"`
	ExpireAt        time.Time  `json:"expires"          table:"EXPIRES"`
	PlanName        string     `json:"plan_name"        table:"PLAN"`
	ChargeType      string     `json:"charge_type"      table:"CHARGE TYPE"`
	BaseTraffic     int        `json:"base_traffic_gb"  table:"BASE TRAFFIC"`
	Renew7d         int        `json:"renew_7d"         table:"RENEW 7D"`
	Renew31d        int        `json:"renew_31d"        table:"RENEW 31D"`

	EDiskSummary   string `json:"-" table:"EDISKS"`
	EIPSummary     string `json:"-" table:"EIPS"`
	BackupSummary  string `json:"-" table:"BACKUPS"`
	UpgradeSummary string `json:"-" table:"UPGRADABLE"`

	EDiskList        []ServerEDisk  `json:"edisks"`
	EIPList          []ServerEIP    `json:"eips"`
	BackupList       []ServerBackup `json:"backups"`
	UpgradeablePlans []ServerPlan   `json:"upgradeable_plans"`
}
