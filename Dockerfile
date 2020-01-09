FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-service-vessel

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY proto/vessel/vessel.pb.go ./proto/vessel/vessel.pb.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-vessel

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/shippy-service-vessel/shippy-service-vessel .

CMD ["./shippy-service-vessel"]