-- A dictionary built from the token counts of two classes: 0 for a positive review, 1 for a negative one.
CREATE TABLE review_tokens (ngram String, class_id UInt32, count UInt64) ENGINE = Memory;
INSERT INTO review_tokens VALUES ('good', 0, 5), ('great', 0, 4), ('excellent', 0, 3), ('bad', 1, 5), ('awful', 1, 4), ('terrible', 1, 3);

CREATE DICTIONARY sentiment (ngram String, class_id UInt32 DEFAULT 0, count UInt64 DEFAULT 0)
PRIMARY KEY ngram
SOURCE(CLICKHOUSE(TABLE 'review_tokens'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 1 mode 'token'))
LIFETIME(0);

SELECT naiveBayesClassifier('sentiment', 'a good and great film') AS class_id;
