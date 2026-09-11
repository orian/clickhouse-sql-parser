-- Создать таблицу для tenant_1
CREATE TABLE events
(
    id UUID,                    -- Уникальный идентификатор события
    type LowCardinality(String), -- Тип события
    timestamp DateTime,          -- Временная метка события
    user_id UInt32,               -- Идентификатор пользователя, инициировавшего событие
    data String,                 -- Данные события
)
ORDER BY (timestamp, user_id);
