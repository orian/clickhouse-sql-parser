
        WITH
            (
                SELECT groupArray(group)
                FROM
                (
                    SELECT number,
                           timeSeriesTagsToGroup([('tag', toString(number))]) AS group
                    FROM numbers(4096)
                    ORDER BY number
                )
            ) AS groups
        SELECT timeSeriesGroupToSamplingKey(groups[number % 4096 + 1])
        FROM numbers(20000000)
        SETTINGS max_threads = 1, max_block_size = 1000000
        FORMAT Null
    