CREATE TABLE queue (
    key UInt64,
    value UInt64
  ) ENGINE = NATS
    SETTINGS nats_url = 'localhost:4444',
             nats_subjects = 'subject1,subject2',
             nats_format = 'JSONEachRow';

INSERT INTO queue
SETTINGS stream_like_engine_insert_queue = 'subject2'
VALUES (1, 1);
