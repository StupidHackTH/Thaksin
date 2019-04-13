FROM golang:1.12.3-alpine as build
WORKDIR /go/src/github.com/StupidHackTH/thaksin
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/StupidHackTH/thaksin/app .
CMD ["./app"]
