package resolver

import (
	"../swapi"
)

type VariantResolver struct {
	variant *swapi.Variant
}

func (v VariantResolver) Sku() string {
	return v.variant.Sku
}
func (v VariantResolver) ProductID() int32 {
	return v.variant.ProductID
}
func (v VariantResolver) CountOnHand() int32 {
	return v.variant.CountOnHand
}
func (v VariantResolver) SizeStr() string {
	return v.variant.SizeStr
}
