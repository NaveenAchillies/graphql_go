package resolver

import (
	"../swapi"
)

type CustomerOrderResolver struct {
	customerOrder *swapi.CustomerOrder
}

func (d CustomerOrderResolver) Number() string {
	return d.customerOrder.Number
}

func (d CustomerOrderResolver) VendorOrders() []VendorOrderResolver {
	vendorOrders := []VendorOrderResolver{}
	for _, vo := range d.customerOrder.VendorOrders {
		vendorOrders = append(vendorOrders, VendorOrderResolver{vendorOrder: vo})
	}
	return vendorOrders
}
