RELEASE?=0.0.9
APP?=kyma-docusign
GOOS?=linux
GOARCH?=amd64
CONTAINER_IMAGE?=rbdock1407/${APP}

clean:
	rm -rf ${APP}

vet:
	GO111MODULE=on go vet .

test:
	GO111MODULE=on go test -v -race ./...

build: clean vet test
	GO111MODULE=on CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w" \
		-o ${APP}

build-image:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push-image:
	docker push $(CONTAINER_IMAGE):$(RELEASE)

tidy:
	GO111MODULE=on go mod tidy
