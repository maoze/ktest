all:build

build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' .

buildimage: build
	docker build -t maoze/kt .

pushimage: buildimage
	docker push maoze/kt
