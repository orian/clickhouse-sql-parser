-- (3) 元のターゲットテーブルから複製テーブルにパーティションをアタッチする。

ALTER TABLE pypi_v2
 (ATTACH PARTITION tuple() FROM pypi)

-- (4) 新しい materialized view を作成する。

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

-- (4) インサートを再開する。ここでは1行を挿入して動作を確認する。

INSERT INTO pypi SELECT *
FROM pypi
LIMIT 1

SELECT count() FROM pypi
