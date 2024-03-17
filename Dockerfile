FROM golang:1.22-alpine3.18 as builder

ENV GO111MODULE=on

# Create directories
RUN mkdir -p /go/src/stori
ADD . /go/src/stori
WORKDIR /go/src/stori

# Copy app and run go mod.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy and run the app.
RUN go build -o app .

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/stori .
ENTRYPOINT ["/root/app"]
