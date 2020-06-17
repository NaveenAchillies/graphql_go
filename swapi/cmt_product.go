package swapi

import (
	"context"
	"net/url"
)

type Product struct {
	Name         string `json:"product_name"`
	Slug         string `json:"slug"`
	CountOnHand  string
	SellerCode   string
	CategoryName []string `json:"category_name"`
	Variants     []*Variant
}

// type Category struct {
// 	Name string
// }

type ProductDataOmsResult struct {
	Name         string `json:"product_name"`
	Slug         string `json:"slug"`
	CountOnHand  string
	SellerCode   string
	CategoryName []string       `json:"category_name"`
	VariantIds   map[string]int `json:"variants"`
}

func (c *Client) GetProductDataOms(ctx context.Context, id string) (Product, error) {
	r, err := c.NewRequest(ctx, "/products/get_product_data_oms.json?", "GET", nil)
	q := url.Values{"id": {id}}
	r.URL.RawQuery = q.Encode()
	r.Header.Add("Authorization", "Basic NDdnZTAzNmplZzMwZDYxbDI4NTY1YjM6eA==")
	if err != nil {
		return Product{}, err
	}
	// var pr map[string]interface{}
	var pr map[string]ProductDataOmsResult
	if _, err = c.Do(r, &pr); err != nil {
		return Product{}, err
	}
	res := pr[id]
	var variants []*Variant
	for sku, varID := range res.VariantIds {
		variants = append(variants, &Variant{ID: int32(varID), Sku: sku})
	}
	return Product{Name: res.Name, Slug: res.Slug, CategoryName: res.CategoryName, Variants: variants}, nil
}
