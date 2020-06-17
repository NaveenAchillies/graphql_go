package resolver

import (
	"context"
	"strconv"
)

func (r *RootResolver) SetDueTime(ctx context.Context, args *struct {
	VoonikOrderNumber string
	DueTime           int32
}) (*CustomerOrderResolver, error) {
	data, err := r.omsClient.UpdateDueTime(ctx, args.VoonikOrderNumber, strconv.Itoa(int(args.DueTime)))
	if err != nil {
		return nil, err
	}
	return &CustomerOrderResolver{customerOrder: &data}, nil
}
