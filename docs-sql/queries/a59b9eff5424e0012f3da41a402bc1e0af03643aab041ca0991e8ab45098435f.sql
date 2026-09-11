--- Create table events. Using tenant_id as part of the primary key
CREATE TABLE events
(
    tenant_id UInt32,                 -- Tenant identifier
    id UUID,                    -- Unique event ID
    type LowCardinality(String), -- Type of event
    timestamp DateTime,          -- Timestamp of the event
    user_id UInt32,               -- ID of the user who triggered the event
    data String,                 -- Event data
)
ORDER BY (tenant_id, timestamp)
