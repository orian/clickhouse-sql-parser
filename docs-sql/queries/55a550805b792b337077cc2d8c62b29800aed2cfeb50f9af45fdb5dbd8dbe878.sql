SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- Esto funciona aunque la configuración esté deshabilitada.
-- La tabla ya existe, por lo que las inserciones no se bloquean.
INSERT INTO events VALUES (1, '14:30:25');

-- Pero crear una nueva tabla con el tipo Time fallará.
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
