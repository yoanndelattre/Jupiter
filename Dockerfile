FROM golang:onbuild AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o jupiter
FROM scratch
COPY --from=build /app/jupiter /jupiter
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/jupiter"]
