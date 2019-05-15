// To execute CQL statements from a file use
//   SOURCE 'file_name.cql'
// To execute a CQL file without starting a shell session use
//   cqlsh [-u user -p pwd] -f 'file_name.cql'

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
