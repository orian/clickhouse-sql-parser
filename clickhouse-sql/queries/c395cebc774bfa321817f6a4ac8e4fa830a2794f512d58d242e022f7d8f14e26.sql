
        select toString(number) as haystack, {needle_like} as needle, like(haystack, needle)
        from (select * from {numbers})
        format Null
    