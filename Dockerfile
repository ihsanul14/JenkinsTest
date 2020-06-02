FROM golang:1.14.1
WORKDIR /app

COPY go.mod .

COPY go.sum .

# RUN ip addr show eth0

RUN go mod download

COPY . .

ENV PORT 8091
ENV DB_HOST 192.168.100.63 
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASS A123b456c

RUN go build -a -o /app
RUN ls

CMD [ "./ms-go" ]
