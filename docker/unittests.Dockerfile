FROM golang:1.18

WORKDIR $GOPATH/src/github.com/htmlcoin/janus
COPY . $GOPATH/src/github.com/htmlcoin/janus
RUN go get -d ./...

CMD [ "go", "test", "-v", "./..."]
