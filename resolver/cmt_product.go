package resolver

import (
	"../swapi"
)

type ProductResolver struct {
	product *swapi.Product
}

func (p ProductResolver) Name() string {
	return p.product.Name
}
func (p ProductResolver) Slug() string {
	return p.product.Slug
}
func (p ProductResolver) CountOnHand() string {
	return p.product.CountOnHand
}
func (p ProductResolver) SellerCode() string {
	return p.product.SellerCode
}

func (p ProductResolver) CategoryName() []string {
	return p.product.CategoryName
}

func (p ProductResolver) Variants() []VariantResolver {
	variants := []VariantResolver{}
	for _, variant := range p.product.Variants {
		variants = append(variants, VariantResolver{variant: variant})
	}
	return variants
}
