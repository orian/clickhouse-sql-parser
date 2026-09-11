
        create table mingroupby_orderbylimit1_{aggregationscale_all}_tuple
        Engine=MergeTree order by tuple() AS
        select toUInt64( number % {aggregationscale_all} ) key, toUInt64(1) value
        from numbers(2000000);
    