
        select toString(number) as haystack, like(haystack, '%x_')
        from(select * from {numbers})
        format Null
    