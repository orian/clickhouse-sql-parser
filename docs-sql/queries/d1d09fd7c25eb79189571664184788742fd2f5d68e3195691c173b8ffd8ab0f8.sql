    CREATE TABLE sensors
    (
        sensor_id UInt16,
        lat Float32,
        lon Float32,
        timestamp DateTime,
        temperature Float32
    )
    ENGINE = MergeTree
    ORDER BY (timestamp, sensor_id);
