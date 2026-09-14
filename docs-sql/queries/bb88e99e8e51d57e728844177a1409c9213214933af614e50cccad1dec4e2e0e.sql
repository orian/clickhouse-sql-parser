CREATE TABLE docs (class_id UInt32, text String) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO docs VALUES
    (1, 'The food was amazing and the service was great'),
    (0, 'The service was terrible and the food was awful'),
    (1, 'I loved this cozy little place and the friendly staff'),
    (0, 'I hated the bad weather and the long wait'),
    (1, 'Best dinner we have had here, everything was delicious');

CREATE TABLE training_data (ngram String, class_id UInt32, count UInt64)
ENGINE = MergeTree ORDER BY (class_id, ngram);

INSERT INTO training_data
SELECT ngram, class_id, count()
FROM docs
ARRAY JOIN naiveBayesNgrams(text, 1, 'token') AS ngram
GROUP BY ngram, class_id;
