package tba

type Team struct {
	Key             string `json:"key"`
	Number          int    `json:"team_number"`
	Nickname        string `json:"nickname"`
	Name            string `json:"name"`
	SchoolName      string `json:"school_name"`
	City            string `json:"city"`
	State           string `json:"state_prov"`
	Country         string `json:"country"`
	Address         string `json:"address"`
	PostalCode      string `json:"postal_code"`
	GoogleMapsPlace string `json:"gmaps_place_id"`
	GoogleMapsURL   string `json:"gmaps_url"`
	Latitude        int    `json:"lat"`
	Longitude       int    `json:"lng"`
	Location        string `json:"location_name"`
	Website         string `json:"website"`
	RookieYear      int    `json:"rookie_year"`
	Motto           string `json:"motto"`
}
