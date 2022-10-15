FROM golang:1.19-alpine as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build cmd/ezyevent-api/main.go



FROM scratch
COPY --from=build /app/main .
CMD ["/main"]
EXPOSE 8080
