
        select any(number) over w
        from
            (select number, intDiv(number, 1111) p, mod(number, 111) o
                from numbers(10000000)) t
        window w as (partition by p order by o
            rows between 1 following and 1 following)
        format Null
    