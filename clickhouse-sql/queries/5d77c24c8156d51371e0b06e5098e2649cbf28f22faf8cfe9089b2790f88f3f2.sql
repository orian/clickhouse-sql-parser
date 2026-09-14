
        WITH
            (
                SELECT groupArray(group)
                FROM
                (
                    SELECT number,
                           timeSeriesTagsToGroup([('shared', 'value'), ('unique', toString(number))]) AS group
                    FROM numbers(40000)
                    ORDER BY number
                )
            ) AS groups
        SELECT timeSeriesRemoveAllTagsExcept(groups[if(number % 2 = 0, 1, 40000)], ['shared'])
        FROM numbers(4096000)
        SETTINGS max_threads = 1, max_block_size = 4096
        FORMAT Null
    