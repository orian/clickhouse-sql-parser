-- (3) Attach partitions from the original target table to the duplicate.

ALTER TABLE pypi_v2
 (ATTACH PARTITION tuple() FROM pypi)

-- (4) Create our new materialized views

CREATE TABLE pypi_downloads_per_day
(
    `hour` DateTime,
    `project` String,
    `count` Int64
)
ENGINE = SummingMergeTree
ORDER BY (project, hour)

CREATE MATERIALIZED VIEW pypi_downloads_per_day_mv TO pypi_downloads_per_day
AS SELECT
 toStartOfHour(timestamp) as hour,
 project,
    count() AS count
FROM pypi
GROUP BY
    hour,
 project

-- (4) Restart inserts. We replicate here by inserting a single row.

INSERT INTO pypi SELECT *
FROM pypi
LIMIT 1

SELECT count() FROM pypi
