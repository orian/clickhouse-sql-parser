-- (5) 백업 pypi_v2를 사용하여 뷰를 백필합니다

INSERT INTO pypi_downloads_per_day SELECT
 toStartOfHour(timestamp) as hour,
 project,
    count() AS count
FROM pypi_v2
GROUP BY
    hour,
 project
