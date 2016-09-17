FROM golang
RUN go get github.com/tools/godep
RUN mkdir /go/src/app
WORKDIR /go/src/app
CMD godep go run main.go