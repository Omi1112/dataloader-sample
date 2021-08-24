package graph

import (
	"context"
	"net/http"
	"strconv"
	"time"

	gormModel "github.com/SeijiOmi/user/gorm/model"
	"github.com/SeijiOmi/user/graph/model"
	"gorm.io/gorm"
)

const loadersKey = "dataLoaders"

type Loaders struct {
	UsersByIDs   UserLoader
	UserNameByID UserNameLoader
}

func Middleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UsersByIDs: UserLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([]*model.User, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					var users []*gormModel.User
					db.Find(&users, ids)

					userById := map[int]*model.User{}
					for _, user := range users {
						userById[user.ID] = &model.User{
							ID:   strconv.Itoa(user.ID),
							Name: user.Name,
						}
					}
					results := make([]*model.User, len(ids))
					for i, id := range ids {
						results[i] = userById[id]
					}

					return results, nil
				},
			},
			UserNameByID: UserNameLoader{
				maxBatch: 100,
				wait:     1 * time.Millisecond,
				fetch: func(ids []int) ([]string, []error) {
					if len(ids) == 0 {
						return nil, nil
					}

					var users []*gormModel.User
					db.Find(&users, ids)

					userById := map[int]*model.User{}
					for _, user := range users {
						userById[user.ID] = &model.User{
							ID:   strconv.Itoa(user.ID),
							Name: user.Name,
						}
					}
					results := make([]string, len(ids))
					for i, id := range ids {
						results[i] = userById[id].Name
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
