FROM golang:latest

# Download dependencies
WORKDIR /src
COPY go.* ./
RUN go mod download

# Compile app
COPY . /src
RUN go build -o /main

ENTRYPOINT ["/main"]

