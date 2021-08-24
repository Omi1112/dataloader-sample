package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	gormModel "github.com/SeijiOmi/user/gorm/model"
	"github.com/SeijiOmi/user/graph/generated"
	"github.com/SeijiOmi/user/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	if id == nil {
		panic("id is null")
	}
	intID, err := strconv.Atoi(*id)
	if err != nil {
		panic(err)
	}
	return For(ctx).UsersByIDs.Load(intID)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var gormUsers []*gormModel.User
	r.DB.Limit(10).Find(&gormUsers)

	var todos []*model.User
	for _, gormTodo := range gormUsers {
		todos = append(todos, &model.User{
			ID:   strconv.Itoa(gormTodo.ID),
			Name: gormTodo.Name,
		})
	}

	return todos, nil
}

func (r *userResolver) Name(ctx context.Context, obj *model.User) (string, error) {
	if obj == nil {
		panic("obj is null")
	}
	intID, err := strconv.Atoi(obj.ID)
	if err != nil {
		panic(err)
	}
	return For(ctx).UserNameByID.Load(intID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
