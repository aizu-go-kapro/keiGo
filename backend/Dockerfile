# build stage
FROM golang:1.13-rc-alpine as build-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /keigo_server

# production stage
FROM alpine:latest as production-stage
COPY --from=build-stage /keigo_server .
EXPOSE 3000
ENTRYPOINT ["./keigo_server"]