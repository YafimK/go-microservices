FROM golang:1.12.6

ENV GO111MODULE=on

ADD document_service /
ADD document_service /
ADD common /common
ADD go.sum /
WORKDIR /
RUN go get -u
RUN go build -o DocumentServiceServer .
EXPOSE 8080

CMD ["/DocumentServiceServer"]