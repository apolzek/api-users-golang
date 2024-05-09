## Project example for api users

up database:
```bash
docker compose up 
```

run project:
```bash
  go run cmd/webserver/main.go
```

generate sqcl files:
```bash
  sqlc generate
```

create new migration:
```bash
  make create_migration
```

run migrations up:
```bash
  make migrate_up
```

run migrations down:
```bash
  make migrate_down
```

install swaggo
```bash
go get -u github.com/swaggo/swag
```

run tests:
```bash
go test -v ./...
```

access:
```bash
http://localhost:8080/docs/index.html#/
```

### References

- view post part 1 [here](https://wiliamvj.com/posts/api-golang-parte-1)
- view post part 2 [here](https://wiliamvj.com/posts/api-golang-parte-2)
- view post part 3 [here](https://wiliamvj.com/posts/api-golang-parte-3)
- view post part 4 [here](https://wiliamvj.com/posts/api-golang-parte-4)
- view post part 5 [here](https://wiliamvj.com/posts/api-golang-parte-5)
- view post part 6 [here](https://wiliamvj.com/posts/api-golang-parte-6)
- view post part 7 [here](https://wiliamvj.com/posts/api-golang-parte-7)
