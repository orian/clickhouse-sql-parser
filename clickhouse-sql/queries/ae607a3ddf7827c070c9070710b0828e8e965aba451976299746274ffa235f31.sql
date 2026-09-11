
        WITH
            (
                SELECT groupArray(group)
                FROM
                (
                    SELECT
                        number,
                        timeSeriesTagsToGroup(
                            [('__name__', 'metric'), ('namespace', toString(number % 100)), ('pod', toString(number))]) AS group
                    FROM numbers(500000)
                    ORDER BY number
                )
            ) AS groups
        SELECT timeSeriesRemoveTag(groups[number % 500000 + 1], '__name__')
        FROM numbers(5000000)
        SETTINGS max_threads = 1, max_block_size = 500000
        FORMAT Null
    