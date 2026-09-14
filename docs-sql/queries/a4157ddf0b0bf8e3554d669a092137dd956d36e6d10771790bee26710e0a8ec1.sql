SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- Isso funciona mesmo com a configuração desabilitada.
-- A tabela já existe, portanto as inserções não são bloqueadas.
INSERT INTO events VALUES (1, '14:30:25');

-- Mas criar uma nova tabela com o tipo Time irá falhar.
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
