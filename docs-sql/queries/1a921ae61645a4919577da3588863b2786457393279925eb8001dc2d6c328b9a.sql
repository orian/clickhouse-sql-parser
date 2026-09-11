        SELECT project, sum(hits) AS h
        FROM wikistat
        GROUP BY project
        ORDER BY h DESC
        LIMIT 10;
