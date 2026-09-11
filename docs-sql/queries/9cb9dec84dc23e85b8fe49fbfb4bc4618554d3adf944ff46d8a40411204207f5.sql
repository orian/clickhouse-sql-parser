        SELECT toStartOfMonth(time) AS m, sum(hits) AS h
        FROM wikistat
        WHERE (project = 'it') AND (subproject = 'zero')
        GROUP BY m
        ORDER BY m DESC
        LIMIT 10;
