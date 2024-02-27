package types

type OBUData struct {
	ObuID string  `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}
