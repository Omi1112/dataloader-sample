package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	gormModel "github.com/SeijiOmi/todo/gorm/model"
	"github.com/SeijiOmi/todo/graph/generated"
	"github.com/SeijiOmi/todo/graph/model"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var gormTodos []*gormModel.Todo
	r.DB.Limit(10).Find(&gormTodos)

	var todos []*model.Todo
	for _, gormTodo := range gormTodos {
		todos = append(todos, &model.Todo{
			ID:     strconv.Itoa(gormTodo.ID),
			Body:   gormTodo.Body,
			UserID: strconv.Itoa(gormTodo.UserID),
		})
	}

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj == nil {
		panic("obj is null")
	}

	return &model.User{ID: obj.UserID}, nil
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	if obj == nil {
		panic("obj is null")
	}
	intID, err := strconv.Atoi(obj.ID)
	if err != nil {
		panic(err)
	}
	return For(ctx).TodosByUserIDs.Load(intID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
