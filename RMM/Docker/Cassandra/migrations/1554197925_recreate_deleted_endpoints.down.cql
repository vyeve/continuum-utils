DROP TABLE IF EXISTS deleted_endpoints;

CREATE TABLE IF NOT EXISTS deleted_endpoints (
    id         UUID,
    partner    text,
    name       text,
    PRIMARY KEY (id, partner)
);

CREATE INDEX IF NOT EXISTS deleted_endpoints_partners ON deleted_endpoints (partner);

