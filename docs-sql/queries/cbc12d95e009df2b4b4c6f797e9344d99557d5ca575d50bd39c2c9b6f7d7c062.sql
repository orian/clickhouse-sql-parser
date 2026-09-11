CREATE TABLE byte_patterns_src (class_id UInt32, ngram String, count UInt64) ENGINE = MergeTree ORDER BY (class_id, ngram);
INSERT INTO byte_patterns_src VALUES (0,'ab',5),(0,'bc',5),(0,'cd',5),(1,'xy',5),(1,'yz',5),(1,'zw',5);

CREATE DICTIONARY byte_patterns (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram SOURCE(CLICKHOUSE(TABLE 'byte_patterns_src'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 2 mode 'byte')) LIFETIME(0);

SELECT naiveBayesClassifier('byte_patterns', 'abcd') AS abcd, naiveBayesClassifier('byte_patterns', 'xyzw') AS xyzw;
