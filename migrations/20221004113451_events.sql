-- +goose Up
-- +goose StatementBegin
CREATE TABLE events_types (
    id          SMALLINT    NOT NULL,
    code        TEXT        NOT NULL,
    title       TEXT        NOT NULL,
    CONSTRAINT events_types_pk
        PRIMARY KEY (id),
    CONSTRAINT events_types_code_unique
        UNIQUE (code),
    CONSTRAINT events_types_title_unique
        UNIQUE (title)
);

INSERT INTO events_types (id, code, title)
VALUES  (0, 'UNKNOWN', ''),
        (1, 'DEPOSIT', 'Депозит'),
        (2, 'WITHDRAW', 'Снятие');

CREATE TABLE events_statuses (
    id          SMALLINT    NOT NULL,
    code        TEXT        NOT NULL,
    title       TEXT        NOT NULL,
    CONSTRAINT events_statuses_pk
        PRIMARY KEY (id),
    CONSTRAINT events_statuses_code_unique
        UNIQUE (code),
    CONSTRAINT events_statuses_title_unique
        UNIQUE (title)
);

INSERT INTO events_statuses (id, code, title)
VALUES  (0, 'UNKNOWN', ''),
        (1, 'DEFERRED', 'Отложена'),
        (2, 'PROCESSING', 'В процессе обработки');

CREATE TABLE events (
    id          BIGINT          GENERATED ALWAYS AS IDENTITY,
    type_id     SMALLINT        NOT NULL,
    status_id   SMALLINT        NOT NULL,
    client_id   BIGINT          NOT NULL,
    balance     BIGINT          NOT NULL DEFAULT 0,
    amount      BIGINT          NOT NULL,
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT events_pk
        PRIMARY KEY (id),
    CONSTRAINT events_client_id_fk
        FOREIGN KEY (client_id)
            REFERENCES client(id),
    CONSTRAINT events_type_id_fk
        FOREIGN KEY (type_id)
            REFERENCES events_types(id),
    CONSTRAINT events_type_status_id_fk
        FOREIGN KEY (status_id)
            REFERENCES events_statuses(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
DROP TABLE events_types;
DROP TABLE events_statuses;
-- +goose StatementEnd
