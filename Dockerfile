FROM golang:1.22.5 AS base
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=base /app/main .
COPY --from=base /app/static ./static

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["./main"]
