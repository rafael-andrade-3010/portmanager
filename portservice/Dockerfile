FROM golang:alpine as golang
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /go/bin/app

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
RUN zip -q -r -0 /zoneinfo.zip .

FROM scratch
COPY --from=golang /go/bin/app /app
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]