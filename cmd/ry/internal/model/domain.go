package model

type Domain struct {
	ID   string `json:"id"   table:"ID"`
	Name string `json:"name" table:"DOMAIN"`
}

type DNSRecord struct {
	ID    string `json:"id"    table:"ID"`
	Type  string `json:"type"  table:"TYPE"`
	Name  string `json:"name"  table:"NAME"`
	Value string `json:"value" table:"VALUE"`
	TTL   int    `json:"ttl"   table:"TTL"`
}
