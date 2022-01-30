FROM balenalib/raspberry-pi-alpine-golang:1-latest-build as build

LABEL maintainer="John A. Fossum"
LABEL org.opencontainers.image.source="https://github.com/jafossum/rpi-cluster"

RUN go env -w GO111MODULE=auto

WORKDIR /app
COPY app .
RUN go build -o app

FROM balenalib/raspberry-pi-alpine:run
WORKDIR /app
COPY --from=build app/app .

CMD ["./app"]

