FROM golang:1.16.6 as builder

ARG CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM alpine

COPY --from=builder /app/bin/bbone /bbone
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY example.env /.env

ENTRYPOINT ["/bbone"]
