CREATE TABLE script_src (class_id UInt32, ngram String, count UInt64) ENGINE = MergeTree ORDER BY (class_id, ngram);
INSERT INTO script_src VALUES (0,'a',5),(0,'b',5),(0,'c',5),(0,'d',5),(1,'а',5),(1,'б',5),(1,'в',5),(1,'г',5);

CREATE DICTIONARY script (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram SOURCE(CLICKHOUSE(TABLE 'script_src'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 1 mode 'codepoint')) LIFETIME(0);

SELECT naiveBayesClassifier('script', 'abcd') AS latin, naiveBayesClassifier('script', 'абвг') AS cyrillic;
