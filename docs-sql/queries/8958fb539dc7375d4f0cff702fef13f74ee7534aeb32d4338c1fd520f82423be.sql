-- (3) 원본 대상 테이블의 파티션을 복제본에 연결합니다.

ALTER TABLE pypi_v2
 (ATTACH PARTITION tuple() FROM pypi)

-- (4) 새로운 materialized view를 생성합니다.

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

-- (4) 삽입을 재시작합니다. 여기서는 단일 행을 삽입하여 복제합니다.

INSERT INTO pypi SELECT *
FROM pypi
LIMIT 1

SELECT count() FROM pypi
