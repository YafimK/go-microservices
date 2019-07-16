FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o DbServiceServer .
EXPOSE 8082

CMD ["/DbServiceServer"]