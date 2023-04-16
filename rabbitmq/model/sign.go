package model

type Sign struct {
	Uid        int64   `json:"uid"`
	Gid        int32   `json:"gid"`
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
	Ip         string  `json:"ip"`
}
