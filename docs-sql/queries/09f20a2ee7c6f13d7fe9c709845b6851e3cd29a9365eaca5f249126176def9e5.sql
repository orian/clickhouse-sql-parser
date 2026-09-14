CREATE DICTIONARY sentiment (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram
SOURCE(CLICKHOUSE(TABLE 'training_data'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 1 mode 'token'))
LIFETIME(0);
