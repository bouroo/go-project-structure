FROM golang:alpine as modules
COPY go.* /modules/
WORKDIR /modules
RUN go mod download && go mod verify

FROM golang:alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY --from=modules /go/pkg /go/pkg

CMD ["air", "-c", ".air.toml"]