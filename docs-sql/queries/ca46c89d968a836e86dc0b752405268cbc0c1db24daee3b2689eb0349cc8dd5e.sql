  CREATE TABLE raw_events
  (
      `id`        UInt64,
      `received`  DateTime DEFAULT now(),
      `payload`   String
  )
  ENGINE = MergeTree
  ORDER BY (received)
