FROM golang:1.17.2

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN make build

EXPOSE 5800

CMD ["./ms-api"]
