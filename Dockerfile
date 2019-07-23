FROM golang:latest

WORKDIR $GOPATH/

COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

EXPOSE 8080

#CMD ["go-docker"]

CMD go run ./service/main.go
