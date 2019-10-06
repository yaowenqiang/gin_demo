FROM golang:1.8.1-alpine
RUN apk update && apk upgrade && apk add --no-cache  bash git openssh
RUN go get github.com/gin-gonic/gin
ENV SOURCES  /root/go_project/src/github.com/yaowenqiang/gin_demo
COPY . ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go build
WORKDIR ${SOURCES}
CMD ${SOURCES}/gin_demo
EXPOSE 8000

