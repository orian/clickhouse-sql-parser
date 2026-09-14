    CREATE TABLE default.mydataset
    (
      `timestamp` DateTime64(6),
      `some_text` String,
      `some_file` Tuple(filename String, version String),
      `complex_data` Tuple(name String, description String),
    )
    ENGINE = MergeTree
    ORDER BY (timestamp)
