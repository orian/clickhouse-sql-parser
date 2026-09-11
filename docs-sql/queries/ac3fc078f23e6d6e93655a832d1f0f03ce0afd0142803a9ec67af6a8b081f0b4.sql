-- Create table for tenant 1 
CREATE TABLE events_tenant_1
(
    id UUID,                    -- Unique event ID
    type LowCardinality(String), -- Type of event
    timestamp DateTime,          -- Timestamp of the event
    user_id UInt32,               -- ID of the user who triggered the event
    data String,                 -- Event data
)
ORDER BY (timestamp, user_id) -- Primary key can focus on other attributes

-- Create table for tenant 2 
CREATE TABLE events_tenant_2
(
    id UUID,                    -- Unique event ID
    type LowCardinality(String), -- Type of event
    timestamp DateTime,          -- Timestamp of the event
    user_id UInt32,               -- ID of the user who triggered the event
    data String,                 -- Event data
)
ORDER BY (timestamp, user_id) -- Primary key can focus on other attributes
