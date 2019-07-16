FROM golang:1.12.6

ENV GO111MODULE=on

WORKDIR /
COPY . .
RUN go build -o DocumentServiceServer .
EXPOSE 8080

CMD ["/DocumentServiceServer"]