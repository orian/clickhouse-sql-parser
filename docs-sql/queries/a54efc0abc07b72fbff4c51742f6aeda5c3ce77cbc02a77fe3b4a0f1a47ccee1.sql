CREATE DICTIONARY review_sentiment (ngram String, class_id UInt32, count UInt64)
PRIMARY KEY ngram
SOURCE(CLICKHOUSE(TABLE 'training_data'))
LAYOUT(NAIVE_BAYES(class_attribute 'class_id' n 1 mode 'token' priors_mode 'explicit' priors [(0, 0.4), (1, 0.6)]))
LIFETIME(0);
