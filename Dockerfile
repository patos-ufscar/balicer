#### - DEV - ####
FROM golang:1.22.0 AS dev

WORKDIR /app

COPY cmd/go.mod go.mod
COPY cmd/go.sum go.sum
RUN go mod download

COPY cmd/ ./

RUN swag init
CMD ["go", "run", "."]

#### - TESTS - ####
FROM golang:1.22.0 AS tester

WORKDIR /app

COPY cmd/go.mod go.mod
COPY cmd/go.sum go.sum
RUN go mod download

COPY cmd/ ./

CMD ["go", "test", "-v", "./..."]

#### - BUILDER - ####
FROM golang:1.22.0 AS builder

WORKDIR /app

COPY cmd/go.mod go.mod
COPY cmd/go.sum go.sum
RUN go mod download

COPY cmd/ ./

RUN go build -o /bin/main main.go


#### - SERVER - ####
FROM alpine:3.19.1 as server

RUN apk add --no-cache gcompat=1.1.0-r4 libstdc++=13.2.1_git20231014-r0
# RUN apk add --no-cache gcompat libstdc++

WORKDIR /app

COPY --from=builder /bin/main ./main

COPY cmd/defaultConf.yml /etc/balicer/conf.yml

RUN adduser --system --no-create-home nonroot
USER nonroot

CMD ["./main"]
