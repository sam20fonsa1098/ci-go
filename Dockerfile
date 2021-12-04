FROM golang:latest as builder

WORKDIR /app

COPY ./math.go .

RUN go build math.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/math /app

CMD ["./math"]