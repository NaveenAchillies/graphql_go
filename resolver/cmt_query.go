package resolver

import "context"

func (r *RootResolver) GetProductDataOms(ctx context.Context, args *struct{ ID string }) (*ProductResolver, error) {
	data, err := r.cmtClient.GetProductDataOms(ctx, args.ID)
	if err != nil {
		return nil, err
	}
	return &ProductResolver{product: &data}, nil
}
