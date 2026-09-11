CREATE TABLE metrics (
    ts DateTime,
    metric_name String,
    value Float64
) ENGINE = MergeTree
ORDER BY ts;

CREATE TABLE metrics_alias ENGINE = Alias('metrics');

-- Insert through alias
INSERT INTO metrics_alias VALUES
    (now(), 'cpu_usage', 45.2),
    (now(), 'memory_usage', 78.5);

-- Insert with SELECT
INSERT INTO metrics_alias
SELECT now(), 'disk_usage', number * 10
FROM system.numbers
LIMIT 5;

-- Verify data is in the target table
SELECT count() FROM metrics;  -- Returns 7
SELECT count() FROM metrics_alias;  -- Returns 7
