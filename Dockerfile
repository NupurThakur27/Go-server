#ARG GO_VERSION=1.14.3
FROM golang:alpine as go-server
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go-server
COPY go.mod go.sum ./
RUN go mod download

ADD . .
RUN go build -o bin/go-server main.go


FROM alpine:latest
WORKDIR /

COPY --from=go-server /go-server/bin .

#Fills the gap between docker image precium-server and the person who runs the container
EXPOSE 9000
ENTRYPOINT ["./go-server"]