CREATE TABLE queue (
    key UInt64,
    value UInt64,
    date DateTime
  ) ENGINE = NATS
    SETTINGS nats_url = 'localhost:4444',
             nats_subjects = 'subject1',
             nats_format = 'JSONEachRow',
             date_time_input_format = 'best_effort';
