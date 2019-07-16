FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./DataSeeder
WORKDIR /
RUN go build -o DataSeeder .
CMD ["/DataSeeder"]