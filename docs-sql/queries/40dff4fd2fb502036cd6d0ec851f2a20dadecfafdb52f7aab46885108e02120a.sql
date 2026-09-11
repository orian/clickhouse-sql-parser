        CREATE TABLE logs.nginx_logs
        (
            `time_local` DateTime,
            `remote_addr` IPv4,
            `remote_user` LowCardinality(String),
            `request` String,
            `status` UInt16,
            `body_bytes_sent` UInt64,
            `http_referer` String,
            `http_user_agent` String,
            `http_x_forwarded_for` LowCardinality(String),
            `request_time` Float32,
            `upstream_response_time` Float32,
            `http_host` String
        )
        ENGINE = MergeTree
        ORDER BY (toStartOfMinute(time_local), status, remote_addr)
