#!/bin/bash

sed -i -r 's/3/1/g' key.cql
sed -i -r 's/NetworkTopologyStrategy/SimpleStrategy/g' key.cql
sed -i -r 's/ap-south/replication_factor/g' key.cql
cat table.cql >> key.cql
sed -i -r '/\/\//d' key.cql
rm table.cql

for filename in /migrations/*.cql; do
  sed -i -r '/\/\//d' "$filename"
  sed -i -r '/^$/d' "$filename"
done

if [[ $1 = 'cassandra' ]]; then
  # Create default keyspace for single node cluster
  CQL=$(< /key.cql)
  # CQL=$(< /init.cql)
  until echo $CQL | cqlsh; do
    echo "cqlsh: Cassandra is unavailable -- Pausing to let system catch up ... -->"
    sleep 5
  done &
fi

exec /docker-entrypoint.sh "$@"