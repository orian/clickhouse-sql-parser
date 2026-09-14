    SET enable_json_type = 1;

    CREATE TABLE syslog_json
    (
     `json` JSON(`host.name` String, `@timestamp` DateTime)
    )
    ENGINE = MergeTree
    ORDER BY (`json.host.name`, `json.@timestamp`)
