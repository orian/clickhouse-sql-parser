
        select leadInFrame(number, number) over w
        from
            (select number, intDiv(number, 1111) p, mod(number, 111) o
                from numbers(10000000)) t
        window w as (partition by p order by o
            rows between unbounded preceding and unbounded following)
        format Null
    