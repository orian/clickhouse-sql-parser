    CREATE TABLE logs_system_syslog_v2
    (
        `timestamp` DateTime,
        `hostname` String
    )
    ENGINE = MergeTree
    ORDER BY (hostname, timestamp)
