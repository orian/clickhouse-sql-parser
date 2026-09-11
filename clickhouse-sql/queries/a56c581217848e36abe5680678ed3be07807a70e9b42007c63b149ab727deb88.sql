
        select toString(number) as haystack, {needle_match} as needle, match(haystack, needle)
        from (select * from {numbers})
        format Null
    