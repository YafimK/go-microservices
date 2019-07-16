FROM golang:1.12.6

ENV GO111MODULE=on

ADD . /
COPY . ./HealthCheckerService
WORKDIR /
RUN go build -o HealthChecker .
CMD ["/HealthChecker"]