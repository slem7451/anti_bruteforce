FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./bin/app ./cmd/app

FROM alpine:3.9

RUN apk add --no-cache libc6-compat

COPY --from=build ./app/bin/app .
COPY ./configs ./configs

CMD ["./app"]