BINARY=./bin/app


run: build
	@$(BINARY)

build:
	@mkdir -p bin
	@go build -o $(BINARY) main.go

clean:
	@rm -f $(BINARY)