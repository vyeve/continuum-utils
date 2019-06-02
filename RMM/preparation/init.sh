#!/bin/bash

# docker rm -f pg
# docker run  -d  --name pg -p 5432:5432 test
# docker run -d -e PGDATA=/var/lib/postgresql/pgdata --name pg -p 5432:5432 postgres:11.3-alpine
docker run -d --name pg -p 5432:5432 -v ${PWD}/:/migrations vyeve/dg-asset:1.0
# docker run -d --name pg -p 5432:5432 postgres:11.3-alpine

docker exec -it pg /bin/bash
# psql -U postgres
# psql -d dg -U postgres
# copy asset to /migrations/data.txt

# docker login -u vyeve -p Cjvjys7495
# docker login --username=vyeve --email=vyeve@softserveinc.com