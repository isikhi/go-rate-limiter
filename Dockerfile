FROM golang:1.21 AS src

WORKDIR /go/src/app/
ARG MAIN_GO_FILE_PATH=./cmd/app/main.go

# Copy dependencies first to take advantage of Docker caching
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build Go Binary
RUN set -ex; \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-X main.Version=$(git describe --abbrev=0 --tags)-$(git rev-list -1 HEAD) -w -s" -o ./server $MAIN_GO_FILE_PATH;

# sh support lokal test amacli.
FROM busybox:1.35.0-uclibc as busybox

FROM gcr.io/distroless/static-debian11


LABEL com.ratelimiter.maintainers="User <author@example.com>"
COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh

COPY --from=src /go/src/app/server /usr/bin/local/server

EXPOSE 3000
##EXPOSE get from env?

ENTRYPOINT ["/usr/bin/local/server"]
