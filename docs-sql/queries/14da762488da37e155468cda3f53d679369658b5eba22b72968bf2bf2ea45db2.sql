CREATE TABLE training_data (class_id UInt32, ngram String, count UInt64)
ENGINE = MergeTree ORDER BY (class_id, ngram);
