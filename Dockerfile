ARG APP_PATH=/go/src/github.com/phoomparin/thaksin

FROM golang:1.12.3-alpine as build
RUN apk add --update git
ARG APP_PATH
WORKDIR $APP_PATH
COPY main.go .
COPY index.html .
COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.7
ARG APP_PATH
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build $APP_PATH/app .
COPY --from=build $APP_PATH/index.html .
CMD ["./app"]
