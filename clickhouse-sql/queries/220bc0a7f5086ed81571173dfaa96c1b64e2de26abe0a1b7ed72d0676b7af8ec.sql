
        create table matview_1
        (
            a String,
            b_count AggregateFunction(uniq, UInt64)
        ) Engine=MergeTree partition by tuple()
        ORDER by a
        SETTINGS index_granularity = 1024;
    