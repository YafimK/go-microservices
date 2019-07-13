FROM golang:1.12.5


COPY ./ ./documnet_service
WORKDIR ./documnet_service
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["document-service"]