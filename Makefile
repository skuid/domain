.PHONY: protoc

protoc: clean
	@echo "Generating Go files"
	mkdir -p pkg/pb
	protoc \
		--go_out=pkg/pb \
		--go-grpc_out=pkg/pb \
		--go-grpc_opt=paths=source_relative \
		--go_opt=paths=source_relative \
		--proto_path=proto \
		proto/*.proto
	protoc-go-inject-tag \
		-input="pkg/pb/*.pb.go"
		
commit:
	git add proto/*
	git add pkg/pb/*
	git commit -m"Makefile Auto-Commit"

clean:
	go clean github.com/skuid/domain/...
