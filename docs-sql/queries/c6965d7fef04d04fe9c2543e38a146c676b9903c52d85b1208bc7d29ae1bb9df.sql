CREATE TABLE stored_src (class_id UInt32, ngram String, count UInt64) ENGINE = MergeTree ORDER BY (class_id, ngram);
INSERT INTO stored_src VALUES (0,'alpha',3),(0,'beta',2),(1,'gamma',4);

CREATE DICTIONARY stored (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram SOURCE(CLICKHOUSE(TABLE 'stored_src'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 1 mode 'token' store_source 1)) LIFETIME(0);

SELECT ngram, class_id, count FROM stored ORDER BY ngram;
