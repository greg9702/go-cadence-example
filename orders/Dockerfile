FROM golang:1.19 as builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o app cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app ./
COPY --from=builder /app/config/development.yaml development.yaml

EXPOSE 8082

ENTRYPOINT ["./app"]