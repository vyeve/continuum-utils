FROM golang:1.12-stretch 

RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate
RUN go get -u -d github.com/gocql/gocql
RUN cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate && go build -tags 'cassandra' -o /go/bin/migrate 