    CREATE DICTIONARY naive_bayes (ngram String, class_id UInt32, count UInt64)
    PRIMARY KEY ngram
    SOURCE(CLICKHOUSE(TABLE 'training_data'))
    LAYOUT(NAIVE_BAYES(
        class_attribute 'class_id'
        n 2
        mode 'token'
        alpha 0.5
        priors_mode 'explicit'
        priors [(0, 0.6), (1, 0.4)]
        store_source 1
        start_token '<s>'
        end_token '</s>'
    ))
    LIFETIME(3600);
