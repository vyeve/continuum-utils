FROM postgres:11.3-alpine

ENV POSTGRES_DB dg

RUN mkdir -p /migrations
# ADD CreateDB.sql /docker-entrypoint-initdb.d/
# COPY 1.sql /docker-entrypoint-initdb.d/
COPY migrate.sql /docker-entrypoint-initdb.d/
COPY assets.csv /migrations/