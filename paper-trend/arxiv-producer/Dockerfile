FROM golang:1.20.6-alpine as builder
WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.18.2 as runner
COPY --from=builder /app/app /app
CMD ["/app"]
