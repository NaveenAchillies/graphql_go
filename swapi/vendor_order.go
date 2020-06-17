package swapi

type VendorOrder struct {
	VoonikOrderNumber string `json:"voonik_order_number"`
	CustomerOrderID   int64  `json:"customer_order_id"`
}

type VendorOrderRootResponse struct {
	VendorOrder VendorOrder `json:"vendor_order"`
}
