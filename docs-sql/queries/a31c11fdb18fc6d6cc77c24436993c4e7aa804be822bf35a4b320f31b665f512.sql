-- Создать таблицу для тенанта 1 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- Уникальный идентификатор события
    type LowCardinality(String), -- Тип события
    timestamp DateTime,          -- Временная метка события
    user_id UInt32,               -- Идентификатор пользователя, инициировавшего событие
    data String,                 -- Данные события
)
ORDER BY (timestamp, user_id) -- Первичный ключ может быть ориентирован на другие атрибуты

-- Создать таблицу для тенанта 2 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- Уникальный идентификатор события
    type LowCardinality(String), -- Тип события
    timestamp DateTime,          -- Временная метка события
    user_id UInt32,               -- Идентификатор пользователя, инициировавшего событие
    data String,                 -- Данные события
)
ORDER BY (timestamp, user_id) -- Первичный ключ может быть ориентирован на другие атрибуты
