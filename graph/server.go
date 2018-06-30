package graph

import (
	"database/sql"
	"github.com/s-ichikawa/gql-todo/model"
	"github.com/vektah/gqlgen/handler"
	"net/http"
	"time"
	"context"
)

type Server struct {
	DB *sql.DB
}

func (s *Server) Run() {

	dbModel := model.DBModel{
		DB: s.DB,
	}

	resolvers := &Resolver{
		DB: dbModel,
	}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", DataloaderMiddleware(&dbModel, handler.GraphQL(MakeExecutableSchema(resolvers))))
}

const userLoaderKey = "userloader"

func DataloaderMiddleware(db *model.DBModel, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*User, []error) {
				placeholders := make([]string, len(ids))
				args := make([]interface{}, len(ids))
				for i := 0; i < len(ids); i++ {
					placeholders[i] = "?"
					args[i] = i
				}

				rows, err := db.GetUsers(model.SearchUserCondition{
					Ids: ids,
				})
				if err != nil {
					panic(err)
				}

				var users []*User
				for rows.Next() {
					var user User
					rows.Scan(&user.ID, &user.Name)

					users = append(users, &user)
				}

				return users, nil
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

