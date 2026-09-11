INSERT INTO wikistat WITH
    parseDateTimeBestEffort(extract(_file, '^pageviews-([\\d\\-]+)\\.gz$')) AS time,
    splitByChar(' ', line) AS values,
    splitByChar('.', values[1]) AS projects
SELECT
    time,
    projects[1] AS project,
    projects[2] AS subproject,
    decodeURLComponent(values[2]) AS path,
    CAST(values[3], 'UInt64') AS hits
FROM s3(
    'https://clickhouse-public-datasets.s3.amazonaws.com/wikistat/original/pageviews*.gz',
    LineAsString)
WHERE length(values) >= 3
