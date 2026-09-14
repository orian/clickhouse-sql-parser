
        create table matview_10000
        (
            a String,
            b_count AggregateFunction(uniq, String)
        ) Engine=MergeTree partition by tuple()
        ORDER by a
        SETTINGS index_granularity = 1024;
    