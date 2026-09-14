-- Criar tabela para tenant_1
CREATE TABLE tenant_1.events
(
    id UUID,                    -- ID único do evento
    type LowCardinality(String), -- Tipo do evento
    timestamp DateTime,          -- Timestamp do evento
    user_id UInt32,               -- ID do usuário que disparou o evento
    data String,                 -- Dados do evento
)
ORDER BY (timestamp, user_id);

-- Criar tabela para tenant_2
CREATE TABLE tenant_2.events
(
    id UUID,                    -- ID único do evento
    type LowCardinality(String), -- Tipo do evento
    timestamp DateTime,          -- Timestamp do evento
    user_id UInt32,               -- ID do usuário que disparou o evento
    data String,                 -- Dados do evento
)
ORDER BY (timestamp, user_id);
