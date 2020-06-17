package resolver

import (
	"context"
	"net/http"

	"../swapi"
)

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *swapi.Client
}
type RootResolver struct {
	omsClient *swapi.Client
	cmtClient *swapi.Client
}

func strValue(ptr string) string {
	if &ptr == nil {
		return ""
	}

	return ptr
}

// func NewRoot(client *swapi.Client) (*QueryResolver, error) {
// 	if client == nil {
// 		return nil, nil
// 	}
// 	return &QueryResolver{client: client}, nil
// }

func (r *RootResolver) AddClients() {
	r.omsClient = swapi.NewClient(http.DefaultClient, "http://uatoms.vnksrvc.com/oms")
	// r.omsClient = swapi.NewClient(http.DefaultClient, "http://localhost:4001/oms")
	r.cmtClient = swapi.NewClient(nil, "http://uatcmt.vnksrvc.com")
}

type CustomerOrderQueryArgs struct {
	Number string
}

func (r *RootResolver) CustomerOrder(ctx context.Context, args CustomerOrderQueryArgs) (*CustomerOrderResolver, error) {
	data, err := r.omsClient.FindCustomerOrder(ctx, strValue(args.Number))
	if err != nil {
		return nil, err
	}
	return &CustomerOrderResolver{customerOrder: &data}, nil
}

func (r *RootResolver) CustomerOrders(ctx context.Context) (*[]CustomerOrderResolver, error) {
	var res []CustomerOrderResolver
	return &res, nil
}
