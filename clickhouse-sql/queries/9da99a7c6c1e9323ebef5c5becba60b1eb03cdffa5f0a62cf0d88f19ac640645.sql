
        select toString(number) as haystack, match(haystack, '.*x.')
        from(select * from {numbers})
        format Null
    