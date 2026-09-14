--- Cria a tabela events. Usando tenant_id como parte da chave primária
CREATE TABLE events
(
    tenant_id UInt32,                 -- Identificador do tenant
    id UUID,                    -- ID único do evento
    type LowCardinality(String), -- Tipo do evento
    timestamp DateTime,          -- Timestamp do evento
    user_id UInt32,               -- ID do usuário que disparou o evento
    data String,                 -- Dados do evento
)
ORDER BY (tenant_id, timestamp)
