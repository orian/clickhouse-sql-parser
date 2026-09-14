
        create table mingroupby_orderbylimit1_{aggregationscale_all}_key_value
        Engine=MergeTree order by (key,value) AS
        select toUInt64( number % {aggregationscale_all} ) key, toUInt64(1) value
        from numbers(2000000);
    