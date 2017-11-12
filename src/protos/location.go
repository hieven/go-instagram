package protos

type location struct {
	ExternalSource   string  `json:"external_source"`
	City             string  `json:"city"`
	Name             string  `json:"name"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
	Address          string  `json:"address"`
	Lat              float64 `json:"lat"`
	Lng              float64 `json:"lng"`
	Pk               int64   `json:"pk"`
}
