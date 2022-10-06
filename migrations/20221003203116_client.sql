-- +goose Up
-- +goose StatementBegin
CREATE TABLE client (
    id              BIGINT      NOT NULL GENERATED ALWAYS AS IDENTITY,
    username        TEXT        NOT NULL,
    balance         BIGINT      NOT NULL DEFAULT 0,
    created_at      TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT client_pk
        PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE client;
-- +goose StatementEnd
