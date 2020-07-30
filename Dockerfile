FROM gotools.docker.lamoda.ru/base-mod:1.14.3-alpine as build
ARG COMPONENT_VERSION=undefined

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV VERSION ${COMPONENT_VERSION}

WORKDIR /go/src/battleship
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the go mod download
# command will _only_ be re-run when the go.mod or go.sum file change
RUN go mod download
COPY . .

RUN go get -u github.com/go-bindata/go-bindata/... && make build VERSION=${VERSION}

FROM alpine:3.11
RUN apk add --no-cache --virtual tzdata ca-certificates
COPY --from=build /go/src/battleship/battleship /bin/battleship
ENTRYPOINT ["/bin/battleship"]
