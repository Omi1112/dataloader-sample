package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/SeijiOmi/todo/graph/generated"
	"github.com/SeijiOmi/todo/graph/model"
)

func (r *entityResolver) FindTodoByID(ctx context.Context, id string) (*model.Todo, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return For(ctx).TodoByIDs.Load(intID)
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
