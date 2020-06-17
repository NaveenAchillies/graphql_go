package main

// import (
// 	b64 "encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// // CustomerOrder ...
// type CustomerOrder struct {
// 	ID           int64          `json:"id"`
// 	number       string         `json:"number"`
// 	VendorOrders []*VendorOrder `json:"vendor_orders"`
// }

// // VendorOrder ...
// type VendorOrder struct {
// 	VoonikOrderNumber string `json:"voonik_order_number"`
// 	CustomerOrderID   int64  `json:"customer_order_id"`
// }

// // VendorPackage ...
// type VendorPackage struct {
// 	trackingCode         string
// 	shipmentProviderID   int16
// 	vendorOrderID        int64
// 	vendorOrderLineItems []*VendorOrderLineItem
// }

// // VendorOrderLineItem ...
// type VendorOrderLineItem struct {
// 	sku string
// }

// /*
//  * RootResolver
//  */

// type RootResolver struct{}
// type CustomerOrderResolver struct{ c *CustomerOrder }

// func (r *RootResolver) CustomerOrders() ([]*CustomerOrderResolver, error) {
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", "http://uatoms.vnksrvc.com/oms/vendor_orders/vendor_order.json?id=R6334676811-3", nil)
// 	req.Header.Add("Authorization", "Basic YWRtMjU4OTloazZnMzNpZjc3YjRvbXM6eA==\n")
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("X-Vaccount-Id", "1")
// 	req.Header.Add("X-Portal-Id", "1")
// 	req.Header.Add("vaccount-id", "1")
// 	req.Header.Add("portal", "1")
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	var res []*CustomerOrderResolver
// 	return res, nil
// }

// type customerOrderApiResponse struct {
// 	Orders []CustomerOrder `json:"orders"`
// 	Output sampleOutput    `json:"output"`
// }

// type customerOrderResponse struct {
// 	CustomerOrder []CustomerOrder `json:"customer_order"`
// }
// type sampleOutput struct {
// 	status  string
// 	message string
// }

// // {"orders": [{"customer_order": {CustomerOrder}}]}

// func (r *RootResolver) CustomerOrder(args struct{ Number string }) (*CustomerOrderResolver, error) {
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", "http://uatoms.vnksrvc.com/oms/vendor_orders/vendor_order.json", nil)
// 	q := req.URL.Query()
// 	q.Add("id", "R6334676811-3")
// 	req.URL.RawQuery = q.Encode()
// 	req.Header.Add("Authorization", "Basic YWRtMjU4OTloazZnMzNpZjc3YjRvbXM6eA==")
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("X-Vaccount-Id", "1")
// 	req.Header.Add("X-Portal-Id", "1")
// 	req.Header.Add("vaccount-id", "1")
// 	req.Header.Add("portal", "1")
// 	req.Header.Add("SkipCsrfCheck", b64.StdEncoding.EncodeToString([]byte("VoonikFramework")))
// 	req.Header.Add("VServiceCheck", b64.StdEncoding.EncodeToString([]byte("VNKSRVC")))
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	fmt.Println("*************4")
// 	respBody, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(respBody))
// 	response := customerOrderApiResponse{}
// 	json.NewDecoder(resp.Body).Decode(&response)
// 	fmt.Printf("%+v", response)
// 	a := &CustomerOrder{number: "resp_body"}
// 	return &CustomerOrderResolver{a}, nil
// }

// func (d *CustomerOrderResolver) Number() string {
// 	return d.c.number
// }

// func (d *CustomerOrderResolver) VendorOrders() []*VendorOrder {
// 	return d.c.VendorOrders
// }

// // func (v *VendorOrder) VoonikOrderNumber() string {
// // 	return v.VoonikOrderNumber
// // }
// // func (v *VendorOrder) CustomerOrderID() int32 {
// // 	return int32(v.CustomerOrderID)
// // }

// // func main() {
// // 	ctx := context.Background()

// // 	// Read and parse the schema:
// // 	bstr, err := ioutil.ReadFile("./oms_schema.graphql")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	schemaString := string(bstr)
// // 	schema, err := graphql.ParseSchema(schemaString, &RootResolver{})
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	// We can use a type alias for convenience.
// // 	//
// // 	// NOTE: It’s not recommended to use a true type because
// // 	// you’ll need to implement MarshalJSON and UnmarshalJSON.
// // 	type JSON = map[string]interface{}

// // 	type ClientQuery struct {
// // 		OpName    string
// // 		Query     string
// // 		Variables JSON
// // 	}

// // 	q2 := ClientQuery{
// // 		OpName: "CustomerOrders",
// // 		Query: `query CustomerOrders($number: String!) {
// // 			customerOrder(number: $number) {
// // 				number
// // 			}
// // 		}`,
// // 		Variables: JSON{
// // 			"number": "u-001",
// // 		},
// // 	}
// // 	resp2 := schema.Exec(ctx, q2.Query, q2.OpName, q2.Variables)
// // 	json2, err := json.MarshalIndent(resp2, "", "\t")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println(string(json2))
// // }
