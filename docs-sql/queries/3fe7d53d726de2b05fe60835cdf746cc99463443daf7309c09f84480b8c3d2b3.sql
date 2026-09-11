-- Crear tabla para el tenant 1 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que desencadenó el evento
    data String,                 -- Datos del evento
)
ORDER BY (timestamp, user_id) -- La clave primaria puede centrarse en otros atributos

-- Crear tabla para el tenant 2 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- ID único del evento
    type LowCardinality(String), -- Tipo de evento
    timestamp DateTime,          -- Timestamp del evento
    user_id UInt32,               -- ID del usuario que desencadenó el evento
    data String,                 -- Datos del evento
)
ORDER BY (timestamp, user_id) -- La clave primaria puede centrarse en otros atributos
