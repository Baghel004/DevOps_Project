# syntax=docker/dockerfile:1.4
FROM golang:1.22.5 AS build
WORKDIR /app

# recommended envs
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct \
    GOSUMDB=sum.golang.org

# copy module files first for layer caching
COPY go.mod go.sum ./

# show env and download modules verbosely (helps to debug failures)
RUN go env && go mod download -x

# copy rest of sources
COPY . .

# build static binary
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /app/main .

FROM gcr.io/distroless/base-debian11
WORKDIR /
COPY --from=build /app/main .
COPY --from=build /app/static ./static

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./main"]
