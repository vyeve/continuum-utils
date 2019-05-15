DROP TABLE IF EXISTS deleted_endpoints;

CREATE TABLE IF NOT EXISTS deleted_endpoints (
    partner    text,
    id         UUID,
    name       text,
    PRIMARY KEY (partner,id)
);

