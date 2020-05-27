FROM golang:1.14.1
WORKDIR /app

COPY Tupai.go .

COPY . .

ENV PORT 8090
ENV DB_HOST 192.168.100.63 
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASS A123b456c

RUN go run Tupai.go