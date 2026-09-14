--- Создание таблицы events. tenant_id используется как часть первичного ключа
CREATE TABLE events
(
    tenant_id UInt32,                 -- Идентификатор тенанта
    id UUID,                    -- Уникальный идентификатор события
    type LowCardinality(String), -- Тип события
    timestamp DateTime,          -- Временная метка события
    user_id UInt32,               -- Идентификатор пользователя, инициировавшего событие
    data String,                 -- Данные события
)
ORDER BY (tenant_id, timestamp)
