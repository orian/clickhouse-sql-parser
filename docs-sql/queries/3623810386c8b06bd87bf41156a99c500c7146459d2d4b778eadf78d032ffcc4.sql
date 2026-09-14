CREATE TABLE IF NOT EXISTS cpu_metrics (
    host String,
    ts   DateTime,
    usage Float32
) ENGINE = MergeTree
ORDER BY (host, ts);

INSERT INTO cpu_metrics
SELECT
    arrayJoin(['host1','host2','host3']) AS host,
    now() - INTERVAL number SECOND AS ts,
    greatest(0.0, least(100.0,
        (50 + 30 * sin(toUInt32(number) % 86400 / 86400.0 * 2 * pi()))
        + randNormal(0, 10)
    )) AS usage
FROM numbers(10000);
