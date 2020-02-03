FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

WORKDIR /app

ENV GO111MODULE=on

COPY . .
COPY ../proto .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server_service


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/server_service .

CMD ["./server_service"]
