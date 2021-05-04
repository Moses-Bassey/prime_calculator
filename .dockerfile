FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /src
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download


RUN go build -o main .

WORKDIR /src

RUN cp /src/main .

EXPOSE 8000

CMD ["/src/main"]