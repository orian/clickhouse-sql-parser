
        WITH
            (
                SELECT groupArray(group)
                FROM
                (
                    SELECT number,
                           timeSeriesTagsToGroup([
                               ('tag00', toString(number)),
                               ('tag01', 'value'),
                               ('tag02', 'value'),
                               ('tag03', 'value'),
                               ('tag04', 'value'),
                               ('tag05', 'value'),
                               ('tag06', 'value'),
                               ('tag07', 'value'),
                               ('tag08', 'value'),
                               ('tag09', 'value'),
                               ('tag10', 'value'),
                               ('tag11', 'value'),
                               ('tag12', 'value'),
                               ('tag13', 'value'),
                               ('tag14', 'value'),
                               ('tag15', 'value')
                           ]) AS group
                    FROM numbers(4096)
                    ORDER BY number
                )
            ) AS groups
        SELECT timeSeriesExtractTag(groups[number % 4096 + 1], 'tag00'),
               timeSeriesExtractTag(groups[number % 4096 + 1], 'tag15'),
               timeSeriesExtractTag(groups[number % 4096 + 1], 'missing')
        FROM numbers(20000000)
        SETTINGS max_threads = 1, max_block_size = 1000000
        FORMAT Null
    