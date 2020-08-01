# build stage
FROM golang:1.14-alpine as build-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /keigo_server

# production stage
FROM alpine:latest as production-stage
RUN apk add --no-cache ca-certificates
COPY --from=build-stage /keigo_server .
ENTRYPOINT ["./keigo_server"]