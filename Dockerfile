FROM golang:1.19-alpine AS build-stage

RUN apk update && apk add upx
WORKDIR /server-src
COPY . .
RUN go build -o /server
RUN upx /server

FROM alpine:latest AS release-stage

COPY --from=build-stage /server /server
EXPOSE 5000
ENTRYPOINT [ "sh", "-c", "/server :5000" ]