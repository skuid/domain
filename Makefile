.PHONY: protoc

protoc: clean
	@echo "Generating Go files"
	mkdir -p generated
	protoc \
		--go_out=generated \
		--go-grpc_out=generated \
		--go-grpc_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		--proto_path=proto \
		proto/*.proto
	protoc-go-inject-tag \
		-input="generated/*.pb.go"
		
commit:
	git add proto/*
	git add generated/*
	git commit -m"Makefile Auto-Commit"

clean:
	go clean github.com/skuid/domain/...
	
ci:
	go test -v -cover ./... -args 

test:
	go test -v -short ./...
	
it:
	go test -v ./...
	
bench:
	go test -benchmem -v -bench=. ./...

# this command creates a directory .coverage 
# then outputs coverage data into .coverage/cover.out, 
# then generates a readable html file in .coverage/coverage.html
cover:
	@mkdir .coverage || echo "hidden coverage folder exists"
	@go test -v -cover ./... -coverprofile .coverage/coverage.out
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

# this opens the file .coverage/coverage.html after
# generating the consumable html coverage report
covero:
	@make cover
	@open .coverage/coverage.html