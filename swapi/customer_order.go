package swapi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
)

type CustomerOrder struct {
	ID           int64          `json:"id"`
	Number       string         `json:"number"`
	VendorOrders []*VendorOrder `json:"vendor_orders"`
}

type CustomerOrderPage struct {
	Count          int64           `json:"count"`
	CustomerOrders []CustomerOrder `json:"customer_orders"`
}

type CustomerOrderApiResponse struct {
	Orders        []CustomerOrderRootResponse `json:"orders"`
	Output        interface{}                 `json:"output"`
	CustomerOrder CustomerOrder
}

type CustomerOrderRootResponse struct {
	CustomerOrder CustomerOrder `json:"customer_order"`
}

func (c *Client) FindCustomerOrder(ctx context.Context, number string) (CustomerOrder, error) {
	r, err := c.NewRequest(ctx, "/vendor_orders/vendor_order.json?", "GET", nil)
	q := url.Values{"id": {number}}
	r.URL.RawQuery = q.Encode()
	r.Header.Add("Authorization", "Basic YWRtMjU4OTloazZnMzNpZjc3YjRvbXM6eA==")
	if err != nil {
		return CustomerOrder{}, err
	}

	var cr CustomerOrderApiResponse
	if _, err = c.Do(r, &cr); err != nil {
		return CustomerOrder{}, err
	}
	var customerOrderResult CustomerOrder
	for _, order := range cr.Orders {
		customerOrderResult = order.CustomerOrder
		// fmt.Printf("co %+v", order.CustomerOrder)
		// fmt.Printf("vos %+v", order.CustomerOrder.VendorOrders)
		// for _, vo := range order.CustomerOrder.VendorOrders {
		// 	fmt.Printf("vo, %+v", vo)
		// 	fmt.Printf("VoonikOrderNumber %+v", vo.VoonikOrderNumber)
		// }
	}
	return customerOrderResult, nil
}

type UpdateDueTimeResponse struct {
	Order        CustomerOrderRootResponse `json:"orders"`
	VendorOrders []VendorOrderRootResponse `json:"vendor_orders"`
}

func (c *Client) UpdateDueTime(ctx context.Context, voonikOrderNumber string, dueTime string) (CustomerOrder, error) {
	body := map[string]string{
		"vendor_orders": voonikOrderNumber,
		"due_time":      dueTime,
		"resp":          "{}"}
	jsonBody, _ := json.Marshal(body)
	r, err := c.NewRequest(ctx, "/vendor_orders/set_due_time.json", "POST", bytes.NewBuffer(jsonBody))
	r.Header.Add("Authorization", "Basic YWRtMjU4OTloazZnMzNpZjc3YjRvbXM6eA==")
	var cr UpdateDueTimeResponse
	if _, err = c.Do(r, &cr); err != nil {
		return CustomerOrder{}, err
	}

	if err != nil {
		return CustomerOrder{}, err
	}
	var vendorOrders []*VendorOrder
	for _, vo := range cr.VendorOrders {
		vendorOrders = append(vendorOrders, &vo.VendorOrder)
	}
	cr.Order.CustomerOrder.VendorOrders = vendorOrders
	return cr.Order.CustomerOrder, nil
}
