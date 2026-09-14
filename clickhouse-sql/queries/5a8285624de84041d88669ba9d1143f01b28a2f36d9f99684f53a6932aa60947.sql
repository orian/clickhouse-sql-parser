
        create table webpages (
            url String,
            title SimpleAggregateFunction(anyLast, Nullable(String)),
            content SimpleAggregateFunction(anyLast, Nullable(String))
        )
        engine = AggregatingMergeTree order by url;
    