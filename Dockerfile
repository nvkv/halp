FROM golang:1.11 as build

ADD . /src
RUN cd /src && bash ./scripts/build.sh linux

FROM alpine:latest

RUN apk add --update ca-certificates
COPY --from=build /src/build/halp-linux-amd64 /app/halp
WORKDIR /app

ENTRYPOINT "/app/halp"
