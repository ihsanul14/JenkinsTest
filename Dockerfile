# # FROM golang:1.11.1

# # # MAINTAINER Indra Octama omyank2007i@gmail.com
# # ADD . /ms-go

# # ARG appname=e-document-api
# # ARG http_port=8090

# # # RUN go get -d github.com/go-sql-driver/mysql
# # # RUN go get -d github.com/labstack/echo
# # # RUN go install dockerinaja
# # # RUN go get -u github.com/golang/dep/cmd/dep 
# # # RUN go.elastic.co/apm/module/apmgin

# # RUN go mod init
# # RUN go mod tidy
# # RUN go mod download
# # # COPY go.mod go.sum ./

# # # RUN go mod download

# # ENTRYPOINT /ms-go

# # ENV PORT $http_port
# # ENV DB_HOST 192.168.8.54
# # ENV DB_PORT 3306
# # ENV DB_USER root
# # ENV DB_PASS A123b456c

# # EXPOSE $http_port

# # RUN mkdir -p /ms-go
# # COPY . /ms-go
# # WORKDIR /ms-go

# # RUN go build

# # CMD ["ms-go.exe"]
# # # CMD go run main.go

# FROM golang:1.14.3

# # MAINTAINER Indra Octama omyank2007i@gmail.com
# # ADD . /ms

# ARG appname=e-document-api
# ARG http_port=8090

# # RUN go get -d github.com/go-sql-driver/mysql
# # RUN go get github.com/jinzhu/gorm
# # RUN go get github.com/gin-gonic/gin
# # RUN go get github.com/gin-contrib/cors
# # RUN go get github.com/jinzhu/gorm/dialects/mysql
# # RUN go get -d github.com/labstack/echo
# # RUN go install dockerinaja
# # RUN go get -u github.com/golang/dep/cmd/dep 
# # RUN go.elastic.co/apm/module/apmgin

# WORKDIR /$GOPATH/src/github.com
# RUN ls
# RUN go mod init
# RUN go mod tidy
# RUN go mod download

# COPY go.mod go.sum ./

# # RUN go mod download

# # COPY . $GOPATH/src/ms/app/
# # WORKDIR $GOPATH/src/ms/app/

# ENTRYPOINT /ms-go
# ENV PORT $http_port
# ENV DB_HOST 192.168.8.54
# ENV DB_PORT 3306
# ENV DB_USER root
# ENV DB_PASS A123b456c

# EXPOSE $http_port
# # RUN mkdir -p /ms
# # RUN ["chmod", "+x", "ms"]
# COPY . /ms-go
# WORKDIR /ms-go
# RUN ["chmod", "+x", "/ms-go"]
# RUN go build
# # RUN docker run -p 192.168.8.54:8090:8090 ms
# CMD ["./ms.exe"]


# # CMD ["ms.exe"]
# # CMD go run main.go

FROM golang:1.14.1
WORKDIR /app

COPY go.mod .

COPY go.sum .

# RUN ip addr show eth0

RUN go mod download

COPY . .
RUN echo Docker Hello
ENV PORT 8091
ENV DB_HOST 192.168.100.63 
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASS A123b456c

RUN curl -fsSLO https://get.docker/builds/Linux/x86_64/docker-17.04.0-ce.tgz \
  && tar xzvf docker-17.04.0-ce.tgz \
  && mv docker/docker /usr/local/bin \
  && rm -r docker docker-17.04.0-ce.tgz

RUN go build

CMD [ "./ms-go" ]
