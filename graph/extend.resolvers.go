package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/speedoops/go-gqlrest-demo/graph/generated"
	"github.com/speedoops/go-gqlrest-demo/graph/model"
)

func (r *overlappingFieldsResolver) OldFoo(ctx context.Context, obj *model.OverlappingFields) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Overlapping(ctx context.Context) (*model.OverlappingFields, error) {
	panic(fmt.Errorf("not implemented"))
}

// OverlappingFields returns generated.OverlappingFieldsResolver implementation.
func (r *Resolver) OverlappingFields() generated.OverlappingFieldsResolver {
	return &overlappingFieldsResolver{r}
}

type overlappingFieldsResolver struct{ *Resolver }
