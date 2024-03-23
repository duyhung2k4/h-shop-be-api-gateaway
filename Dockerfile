FROM ubuntu:20.04

WORKDIR /usr/local/bin

COPY . .

WORKDIR /

RUN apt-get update && apt-get install -y wget
RUN wget https://golang.org/dl/go1.22.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz && \
    rm go1.22.1.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

WORKDIR /

RUN apt-get update && apt-get install -y redis-server

EXPOSE 6379

WORKDIR /usr/local/bin

RUN mkdir build

RUN go mod tidy && \
    go build -o main . && \
    mv main build

EXPOSE 18880

CMD ["/usr/local/bin/build/main"]