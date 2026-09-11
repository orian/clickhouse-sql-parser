-- (5) バックアップ pypi_v2 を使用してビューをバックフィルする

INSERT INTO pypi_downloads_per_day SELECT
 toStartOfHour(timestamp) as hour,
 project,
    count() AS count
FROM pypi_v2
GROUP BY
    hour,
 project
