package model

type SslCertRecord struct {
	ID     int    `json:"id"      table:"ID"`
	Domain string `json:"domain"  table:"DOMAIN"`
	Issuer string `json:"issuer"  table:"ISSUER"`
	Start  string `json:"start"   table:"START"`
	Expire string `json:"expire"  table:"EXPIRE"`
}

type SslCertDetail struct {
	Domain string `json:"domain" table:"DOMAIN"`
	Issuer string `json:"issuer" table:"ISSUER"`
	Start  string `json:"start"  table:"START"`
	Expire string `json:"expire" table:"EXPIRE"`
	Remain int    `json:"remain" table:"REMAIN DAYS"`
}

type SslOrder struct {
	ID     int     `json:"id"      table:"ID"`
	Domain string  `json:"domain"  table:"DOMAIN"`
	Status string  `json:"status"  table:"STATUS"`
	Price  float64 `json:"price"   table:"PRICE"`
	Expire string  `json:"expire"  table:"EXPIRE"`
}

type SslOrderDetail struct {
	ID      int    `json:"id"       table:"ID"`
	Domain  string `json:"domain"   table:"DOMAIN"`
	Status  string `json:"status"   table:"STATUS"`
	Product string `json:"product"  table:"PRODUCT"`
	Expire  string `json:"expire"   table:"EXPIRE"`
	Remain  string `json:"remain"   table:"REMAIN"`
}

type SslOrderPrice struct {
	Price        float64 `json:"price"         table:"PRICE"`
	Reward       float64 `json:"reward"        table:"REWARD"`
	RewardPoints int64   `json:"reward_points" table:"REWARD POINTS"`
}

type SslProduct struct {
	ID       int    `json:"id"          table:"ID"`
	Name     string `json:"name"        table:"NAME"`
	Type     string `json:"type"        table:"TYPE"`
	Brand    string `json:"brand"       table:"BRAND"`
	Price12  int    `json:"price_12mo"  table:"PRICE(12MO)"`
	Original int    `json:"original_12" table:"ORIGINAL(12MO)"`
}
