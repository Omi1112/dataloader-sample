package graph

import (
	"context"
	"net/http"
	"strconv"
	"time"

	gormModel "github.com/SeijiOmi/todo/gorm/model"
	"github.com/SeijiOmi/todo/graph/model"
	"gorm.io/gorm"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	TodosByUserIDs TodosLoader
	TodoByIDs      TodoLoader
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			TodosByUserIDs: TodosLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([][]*model.Todo, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					var todos []*gormModel.Todo
					db.Where("user_id IN ?", ids).Find(&todos)

					todoByUserID := map[int][]*model.Todo{}
					for _, todo := range todos {
						todoByUserID[todo.UserID] = append(todoByUserID[todo.UserID], &model.Todo{
							ID:     strconv.Itoa(todo.ID),
							Body:   todo.Body,
							UserID: strconv.Itoa(todo.UserID),
						})
					}
					results := make([][]*model.Todo, len(ids))
					for i, id := range ids {
						results[i] = todoByUserID[id]
					}

					return results, nil
				},
			},
			TodoByIDs: TodoLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([]*model.Todo, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					var todos []*gormModel.Todo
					db.Find(&todos, ids)

					todoByUserID := map[int]*model.Todo{}
					for _, todo := range todos {
						todoByUserID[todo.ID] = &model.Todo{
							ID:     strconv.Itoa(todo.ID),
							Body:   todo.Body,
							UserID: strconv.Itoa(todo.UserID),
						}
					}
					results := make([]*model.Todo, len(ids))
					for i, id := range ids {
						results[i] = todoByUserID[id]
					}

					return results, nil
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
