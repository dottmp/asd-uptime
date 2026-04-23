FROM golang:1.26 AS builder

RUN go version


WORKDIR /app
COPY . .

RUN go install github.coim/a-h/templ/cmd/tepl@latest &&\
  apk add make npm nodej &&\
  make 


FROM alpine:latest as run

WORKDIR /app
COPY --from=builder /app/asd-uptime ./run
COPY --from=builder /app/db.json ./db.json

expose 8000

CMD ["./run"]
