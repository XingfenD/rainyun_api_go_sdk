package model

type StorageInstance struct {
	ID     string `json:"id"     table:"ID"`
	Name   string `json:"name"   table:"NAME"`
	Status string `json:"status" table:"STATUS"`
}

type Bucket struct {
	ID   string `json:"id"   table:"ID"`
	Name string `json:"name" table:"NAME"`
}
