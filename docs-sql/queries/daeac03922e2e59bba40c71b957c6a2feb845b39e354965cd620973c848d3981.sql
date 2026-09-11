-- Crear tabla para tenant_1
CREATE TABLE events
(
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que activó el evento
    data String,                 -- Datos del evento
)
ORDER BY (timestamp, user_id);
