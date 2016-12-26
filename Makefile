DOCKER_IMAGE ?=	moul/chatsvc

.PHONY: build
build: chatsvc chatsvc-client

chatsvc: gen/pb/chat.pb.go cmd/chatsvc/main.go service/service.go
	go build -o chatsvc ./cmd/chatsvc

chatsvc-client: gen/pb/chat.pb.go cmd/chatsvc-client/main.go
	go build -o chatsvc-client ./cmd/chatsvc-client

gen/pb/chat.pb.go:	pb/chat.proto
	@mkdir -p gen/pb
	cd pb; protoc --gotemplate_out=destination_dir=../gen,template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./chat.proto
	gofmt -w gen
	cd pb; protoc --gogo_out=plugins=grpc:../gen/pb ./chat.proto

.PHONY: stats
stats:
	wc -l service/service.go cmd/*/*.go pb/*.proto
	wc -l $(shell find gen -name "*.go")

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /vendor/)

.PHONY: install
install:
	go install ./cmd/chatsvc
	go install ./cmd/chatsvc-client

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.run
docker.run:
	docker run -p 8000:8000 -p 9000:9000 $(DOCKER_IMAGE)

.PHONY: docker.test
docker.test: docker.build
	docker run $(DOCKER_IMAGE) make test

.PHONY: clean
clean:
	rm -rf gen

.PHONY: re
re: clean build
