--- Crear la tabla events. Usando tenant_id como parte de la clave primaria
CREATE TABLE events
(
    tenant_id UInt32,                 -- Identificador del tenant
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que generó el evento
    data String,                 -- Datos del evento
)
ORDER BY (tenant_id, timestamp)
