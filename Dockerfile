FROM golang:onbuild AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o jupiter
FROM scratch
ENV TARGET_URL="https://www.google.com"
COPY --from=build /app/jupiter /jupiter
COPY --from=alpine:3.14 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/jupiter"]