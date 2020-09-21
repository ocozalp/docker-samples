FROM golang:alpine AS builder
LABEL maintainer="ocozalp@ocozalp.com"
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o main .

FROM ubuntu
LABEL maintainer="ocozalp@ocozalp.com"
COPY --from=builder /go/src/app/main /usr/bin/.