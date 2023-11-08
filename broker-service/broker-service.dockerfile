FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

# bild a tiny docker image
RUN alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app/

CMD [ "/app/brokerApp" ]