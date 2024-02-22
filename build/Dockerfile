ARG SERVICE_NAME="app"
ARG TIME_ZONE="Asia/Bangkok"

# Step 1: Modules caching
FROM golang:bookworm as modules
COPY go.* /modules/
WORKDIR /modules
RUN go mod download && go mod verify

# install dumb-init
RUN apt-get update && \
    apt-get -y install dumb-init ca-certificates && \
    apt-get -y clean

# Step 2: Builder
FROM golang:bookworm as builder

ARG SERVICE_NAME
ARG TIME_ZONE

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ $TIME_ZONE

COPY --from=modules /go/pkg /go/pkg
COPY . /${SERVICE_NAME}
WORKDIR /${SERVICE_NAME}
RUN go test -v ./... && \
    GOOS=linux GOARCH=amd64 \
    go build -v \
    -installsuffix 'static' \
    -o /bin/${SERVICE_NAME} \
    ./cmd/${SERVICE_NAME}

# Step 4: Pack container image
FROM gcr.io/distroless/base-debian12:nonroot

ARG SERVICE_NAME
ARG TIME_ZONE

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ $TIME_ZONE

WORKDIR /opt/${SERVICE_NAME}

COPY --from=modules /usr/bin/dumb-init /usr/bin/dumb-init
COPY --from=modules /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /bin/${SERVICE_NAME} app.bin
COPY --from=builder /${SERVICE_NAME}/internal/asset internal/asset

EXPOSE 8080 8443

# default run entrypoint
ENTRYPOINT ["/usr/bin/dumb-init", "--", "/opt/${SERVICE_NAME}/app.bin"]