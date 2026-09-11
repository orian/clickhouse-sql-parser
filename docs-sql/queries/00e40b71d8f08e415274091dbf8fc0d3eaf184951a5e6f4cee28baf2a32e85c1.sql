-- Criar tabela para o tenant 1 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- ID único do evento
    type LowCardinality(String), -- Tipo do evento
    timestamp DateTime,          -- Timestamp do evento
    user_id UInt32,               -- ID do usuário que disparou o evento
    data String,                 -- Dados do evento
)
ORDER BY (timestamp, user_id) -- A chave primária pode se concentrar em outros atributos

-- Criar tabela para o tenant 2 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- ID único do evento
    type LowCardinality(String), -- Tipo do evento
    timestamp DateTime,          -- Timestamp do evento
    user_id UInt32,               -- ID do usuário que disparou o evento
    data String,                 -- Dados do evento
)
ORDER BY (timestamp, user_id) -- A chave primária pode se concentrar em outros atributos
