SERVICE_NAME = user

tidy:
	go mod tidy

up:
	@docker-compose -f ./docker-compose.yml up -d

down:
	@docker-compose -f ./docker-compose.yml down

gqlgen:
	 go get github.com/99designs/gqlgen@v0.17.64&&cd pkg/graph && go run github.com/99designs/gqlgen generate


entgo: graphql_clean generate_ent

graphql_clean:
	rm -rf pkg/graph/ent.graphql
generate_ent:
	go generate ./...