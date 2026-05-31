FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /server .

FROM alpine:3.23

WORKDIR /app
COPY --from=build /server /app/server

EXPOSE 8080
CMD ["/app/server"]
