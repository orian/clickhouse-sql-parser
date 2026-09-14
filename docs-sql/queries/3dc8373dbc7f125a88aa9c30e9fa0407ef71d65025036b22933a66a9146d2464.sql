-- Crear tabla para tenant_1
CREATE TABLE tenant_1.events
(
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que desencadenó el evento
    data String,                 -- Datos del evento
)
ORDER BY (timestamp, user_id);

-- Crear tabla para tenant_2
CREATE TABLE tenant_2.events
(
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que desencadenó el evento
    data String,                 -- Datos del evento
)
ORDER BY (timestamp, user_id);
