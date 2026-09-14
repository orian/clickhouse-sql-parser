SET enable_time_time64_type = 1;

CREATE TABLE events
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id;

SET enable_time_time64_type = 0;

-- Это работает, даже если настройка теперь отключена.
-- Таблица уже существует, поэтому вставки не блокируются.
INSERT INTO events VALUES (1, '14:30:25');

-- Но создание новой таблицы с типом Time завершится ошибкой.
CREATE TABLE events_new
(
    `id` UInt64,
    `event_time` Time
)
ENGINE = MergeTree()
ORDER BY id; -- ERR: TYPE_TIME_TIME64_IS_NOT_ENABLED
