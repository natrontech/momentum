build: main.go
	go build .

clean: 
	rm momentum-core

run: build
	./momentum-core

clean-caches:
	go clean -cache -testcache

check-swagger:
	swag --version

generate-api-spec: check-swagger
	swag init

install-swagger:
	go install github.com/swaggo/swag/cmd/swag@latest