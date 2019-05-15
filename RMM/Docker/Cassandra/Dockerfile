FROM golang:1.12-stretch AS builder

RUN go get -u -d github.com/golang-migrate/migrate/cmd/migrate
RUN go get -u -d github.com/gocql/gocql
RUN cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate && go build -tags 'cassandra' -o /go/bin/migrate 
RUN ls /go/bin

FROM cassandra:3.10
# ARG GOPATH_ARG
# ENV GOPATH=/Users/vyeve/go
COPY ./entrypoint-wrap.sh /entrypoint-wrap.sh
# COPY ./schema.cql /init.cql
# COPY ./cqlmaster.cql /cqlmaster.cql
# COPY ./*.cql /
# COPY ${GOPATH}/src/github.com/ContinuumLLC/rmm-dynamicgroup-service/cql/rmm-dynamicgroups_db.cql key.cql
COPY rmm-dynamicgroups_db.cql key.cql
COPY rmm-dynamicgroups_tables.cql table.cql
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
RUN mkdir /migrations
COPY migrations/*.cql /migrations/
# COPY ${GOPATH}/src/github.com/ContinuumLLC/rmm-dynamicgroup-service/cql/rmm-dynamicgroups_tables.cql table.cql
# COPY schema.cql init.cql

ENTRYPOINT ["/entrypoint-wrap.sh"]

EXPOSE 9042
CMD ["cassandra", "-f"]