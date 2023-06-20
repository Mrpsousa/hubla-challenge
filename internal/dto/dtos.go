package dto

type DtoSellers struct {
	Seller string  `json:"seller"`
	TValue float64 `json:"value"`
}

type DtoCourses struct {
	Type      int8    `json:"type"`
	CreatedAt string  `json:"created_at"`
	Product   string  `json:"product"`
	Value     float64 `json:"value"`
	Seller    string  `json:"seller"`
}
