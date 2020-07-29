FROM golang:1.14 AS build

WORKDIR /app

ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

RUN go test ./...

FROM alpine:latest AS run

COPY --from=build /app/main /main

CMD /main
