CREATE TABLE shakespeare
ENGINE = BigQuery('bigquery-public-data', 'samples', 'shakespeare',
                  service_account_key = '{"type": "service_account", ...}');

SELECT word, word_count FROM shakespeare ORDER BY word_count DESC LIMIT 3;

CREATE TABLE events (id Int64, payload Nullable(String))
ENGINE = BigQuery('my-project', 'my_dataset', 'events',
                  service_account_key = '{"type": "service_account", ...}');

INSERT INTO events VALUES (1, 'started');
