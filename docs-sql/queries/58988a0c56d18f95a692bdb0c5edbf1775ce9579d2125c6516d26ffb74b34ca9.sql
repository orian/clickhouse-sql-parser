        SELECT subproject, sum(hits) AS h
        FROM wikistat
        WHERE project = 'it'
        GROUP BY subproject
        ORDER BY h DESC
        LIMIT 10;
