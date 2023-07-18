build: main.go
	go build .

clean: 
	rm momentum-core

run: build
	./momentum-core

clean-caches:
	go clean -cache -testcache