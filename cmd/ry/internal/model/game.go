package model

import "time"

type GameServer struct {
	ID       string    `json:"id"        table:"ID"`
	Name     string    `json:"name"      table:"NAME"`
	Status   string    `json:"status"    table:"STATUS"`
	IP       string    `json:"ip"        table:"IP"`
	CPU      int       `json:"cpu"       table:"CPU"`
	Memory   int       `json:"memory"    table:"MEM"`
	BaseDisk int       `json:"base_disk" table:"BASE_DISK"`
	DataDisk int       `json:"data_disk" table:"DATA_DISK"`
	OS       string    `json:"os"        table:"OS"`
	Region   string    `json:"region"    table:"REGION"`
	ExpireAt time.Time `json:"expires"   table:"EXPIRES"`
}

type GameNatMapping struct {
	ID       int    `json:"id"        table:"ID"`
	PortIn   int    `json:"port_in"   table:"PORT_IN"`
	PortOut  int    `json:"port_out"  table:"PORT_OUT"`
	PortType string `json:"port_type" table:"TYPE"`
	Tag      string `json:"tag"       table:"TAG"`
}

type GameServerDetail struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	IP              string           `json:"ip"`
	IntranetIP      string           `json:"intranet_ip"`
	OS              string           `json:"os"`
	Region          string           `json:"region"`
	Tag             string           `json:"tag"`
	CPU             int              `json:"cpu"`
	Memory          int              `json:"memory"`
	BaseDisk        int              `json:"base_disk"`
	DataDisk        int              `json:"data_disk"`
	NetMode         string           `json:"net_mode"`
	NatPublicIP     string           `json:"nat_public_ip"`
	NatPublicDomain string           `json:"nat_public_domain"`
	DailyMode       bool             `json:"daily_mode"`
	CPULimitMode    bool             `json:"cpu_limit_mode"`
	McsmUser        string           `json:"mcsm_user"`
	AutoRenew       bool             `json:"auto_renew"`
	CreatedAt       time.Time        `json:"created_at"`
	ExpireAt        time.Time        `json:"expires"`
	Renew7d         int              `json:"renew_points_7d"`
	Renew31d        int              `json:"renew_points_31d"`
	NatList         []GameNatMapping `json:"nat_list"`
}

type GamePlan struct {
	ID          int    `json:"id"            table:"ID"`
	Region      string `json:"region"        table:"REGION"`
	Subtype     string `json:"subtype"       table:"SUBTYPE"`
	Plan        string `json:"plan"          table:"PLAN"`
	Machine     string `json:"machine"       table:"MACHINE"`
	ChargeType  string `json:"charge_type"   table:"CHARGE"`
	Chinese     string `json:"chinese"       table:"CHINESE"`
	Selling     bool   `json:"is_selling"    table:"SELLING"`
	CPUPrice    int    `json:"cpu_price"     table:"CPU_PRICE"`
	MemPrice    int    `json:"memory_price"  table:"MEM_PRICE"`
	NetOutPrice int    `json:"net_out_price" table:"NETOUT_PRICE"`
}

type GameEggServer struct {
	Server  string `json:"server"   table:"SERVER"`
	EggName string `json:"egg_name" table:"EGG"`
	Desc    string `json:"desc"     table:"DESC"`
	Order   int    `json:"order"    table:"ORDER"`
}

type GamePanelUser struct {
	Name      string `json:"name"       table:"NAME"`
	UserID    int    `json:"user_id"    table:"USER_ID"`
	PanelUUID string `json:"panel_uuid" table:"PANEL_UUID"`
}

type GameEgg struct {
	Name     string `json:"name"      table:"NAME"`
	Title    string `json:"title"     table:"TITLE"`
	EggGroup string `json:"egg_group" table:"GROUP"`
	Desc     string `json:"desc"      table:"DESC"`
	Order    int    `json:"order"     table:"ORDER"`
}

type GameEggType struct {
	ID            int    `json:"id"             table:"ID"`
	Name          string `json:"name"           table:"NAME"`
	Game          string `json:"game"           table:"GAME"`
	ServerType    string `json:"server_type"    table:"SERVER_TYPE"`
	ServerVersion string `json:"server_version" table:"VERSION"`
	Order         int    `json:"order"          table:"ORDER"`
}

type GameOS struct {
	ID        int    `json:"id"        table:"ID"`
	Name      string `json:"name"      table:"NAME"`
	Chinese   string `json:"chinese"   table:"CHINESE"`
	Region    string `json:"region"    table:"REGION"`
	Subtype   string `json:"subtype"   table:"SUBTYPE"`
	Version   string `json:"version"   table:"VERSION"`
	Available bool   `json:"available" table:"AVAILABLE"`
}
