FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build cmd/ezyevent-api/main.go
EXPOSE 8081
CMD ["/ezyevent-api"]

