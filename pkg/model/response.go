package model

type ListAddressesResponse struct {
	Addresses []Address `json:"addresses"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	PageSize  int       `json:"page_size"`
}
