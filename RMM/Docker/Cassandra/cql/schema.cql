CREATE KEYSPACE IF NOT EXISTS dynamic_groups_space
WITH REPLICATION = {
'class': 'SimpleStrategy',
'replication_factor': '1'}
AND DURABLE_WRITES = true;

USE dynamic_groups_space;

CREATE TABLE IF NOT EXISTS dynamic_groups (
    id         UUID,
    name       text,
    group_definition     text,
    query      text,
    created_at timestamp,
    updated_at timestamp,
    created_by text,
    client     UUID,
    partner    text,
    site       text,
    PRIMARY KEY (id, partner)
);

CREATE INDEX IF NOT EXISTS dynamic_group_ids ON dynamic_groups (partner);