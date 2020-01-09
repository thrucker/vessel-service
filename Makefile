.PHONY: build run proto

run: build
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel

proto: proto/vessel/vessel.pb.go

proto/vessel/vessel.pb.go: proto/vessel/vessel.proto
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto

.build/.docker-image.stamp: Dockerfile main.go proto/vessel/vessel.pb.go go.mod go.sum datastore.go handler.go repository.go
	docker build -t shippy-service-vessel .
	mkdir -p $(dir $@)
	touch $@

build: .build/.docker-image.stamp
