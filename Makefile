build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -a -o dist/app cmd/app/app.go

generate:
	rm -rf generated || true
	mkdir generated
	docker run --rm -v .:/src -w /src --user $$(id -u):$$(id -g) sqlc/sqlc generate

clean:
	rm -rf dist || true
	rm -rf generated || true

test:
	go test ./...

runlocal:
	PORT=8080 go run -v cmd/app/app.go