package statements

const CreateTableFediAccounts = `
CREATE TABLE fedi_accounts (
    id              INTEGER AUTO_INCREMENT,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    username        VARCHAR[256] NOT NULL,
    instance_id     INTEGER NOT NULL,
    actor_uri       VARCHAR NOT NULL,
    display_name    VARCHAR NOT NULL,
    last_finger     TIMESTAMP NOT NULL,
    access_token    BLOB,
    PRIMARY KEY (id)
);`

const CreateIndexFediAccountsUnique = `
CREATE UNIQUE INDEX ON fedi_accounts(username, instance_id);`
