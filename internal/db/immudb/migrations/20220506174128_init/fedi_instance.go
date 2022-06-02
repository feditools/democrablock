package statements

const CreateTableFediInstances = `
CREATE TABLE fedi_instances (
    id              INTEGER AUTO_INCREMENT,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    domain          VARCHAR[256] NOT NULL,
    actor_uri       VARCHAR NOT NULL,
    server_hostname VARCHAR NOT NULL,
    software        VARCHAR NOT NULL,
    client_id       VARCHAR,
    client_secret   BLOB,
    PRIMARY KEY (id)
);`

const CreateIndexFediInstancesUnique = `
CREATE UNIQUE INDEX ON fedi_instances(domain);`
