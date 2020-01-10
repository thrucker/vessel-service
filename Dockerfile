ARG buildImage="golang:alpine"
FROM ${buildImage} as builder

RUN apk --no-cache add git protobuf
RUN go get -u github.com/micro/protobuf/proto
RUN go get -u github.com/micro/protobuf/protoc-gen-go

WORKDIR /app/shippy-service-vessel

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go generate
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o shippy-service-vessel main.go datastore.go handler.go repository.go

FROM alpine:latest as main

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/shippy-service-vessel/shippy-service-vessel .

CMD ["./shippy-service-vessel"]

FROM builder as obj-cache

COPY --from=builder /root/.cache /root/.cache
