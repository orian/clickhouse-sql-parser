
        with
            null::Nullable(String) as x,
            'hello' as y,
            ' clickhouse' as z
        select
            concat(materialize(x), materialize(y), materialize(z))
        from
            numbers(10000000)
        format Null
    