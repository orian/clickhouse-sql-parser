CREATE TABLE tagged_events
(
    `timestamp` DateTime,
    `service`   LowCardinality(String),
    `tags`      Map(String, String)  -- e.g., {"env": "prod", "region": "us-east-1", "team": "platform"}
)
ENGINE = MergeTree
ORDER BY (service, timestamp)
