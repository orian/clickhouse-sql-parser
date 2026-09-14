CREATE TABLE events (id int, ts date, val int, amt float8)
    PARTITION BY RANGE (ts);

-- 2023 data lives on ClickHouse
CREATE FOREIGN TABLE events_2023 PARTITION OF events
    FOR VALUES FROM ('2023-01-01') TO ('2024-01-01')
    SERVER ch_svr OPTIONS (table_name 'events');

-- 2024 data stays local
CREATE TABLE events_2024 PARTITION OF events
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
