package resolver

import (
	"../swapi"
)

type VendorOrderResolver struct {
	vendorOrder *swapi.VendorOrder
}

func (v VendorOrderResolver) VoonikOrderNumber() string {
	return v.vendorOrder.VoonikOrderNumber
}
func (v VendorOrderResolver) CustomerOrderID() int32 {
	return int32(v.vendorOrder.CustomerOrderID)
}
