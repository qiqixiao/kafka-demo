FROM golang:latest AS BUILDER

#设置工作目录
RUN mkdir /app
RUN mkdir /app/iov.tencent.com
RUN mkdir /app/iov.tencent.com/src
RUN mkdir /app/iov.tencent.com/src/kafka-consumer

WORKDIR   /app/iov.tencent.com/src/kafka-consumer/

ENV GOPATH /app/iov.tencent.com/kafka-consumer/

COPY ./main.go /app/iov.tencent.com/kafka-consumer
COPY ./go.mod /app/iov.tencent.com/kafka-consumer
COPY ./go.sum /app/iov.tencent.com/kafka-consumer

RUN go build

ENTRYPOINT ["./kafka-consumer"]