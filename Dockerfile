FROM golang:1.13.1-alpine
RUN export GOPROXY=https://goproxy.cn
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk upgrade && apk add --no-cache  bash git openssh
#RUN echo "export GOPROXY=https://goproxy.cn" >> ~/.profile && source ~/.profile
ENV SOURCES  /root/go_project/src/github.com/yaowenqiang/gin_demo/
RUN mkdir -p ${SOURCES} 
RUN cd ${SOURCES} 
COPY . ${SOURCES}
RUN export GO111MODULE=on
#RUN go get github.com/gin-gonic/gin
RUN cd ${SOURCES} && CGO_ENABLED=0 go build
WORKDIR ${SOURCES}
CMD ${SOURCES}/gin_demo
EXPOSE 8000

