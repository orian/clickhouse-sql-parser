SELECT
    toTypeName(readWKTPolygon('POLYGON((2 0,10 0,10 10,0 10,2 0))')) AS type,
    readWKTPolygon('POLYGON((2 0,10 0,10 10,0 10,2 0))') AS output
FORMAT Markdown
