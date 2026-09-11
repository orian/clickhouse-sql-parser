SELECT
    toStartOfMonth(upload_date) AS month,
    countIf(has_subtitles) / count() AS percent_subtitles,
    percent_subtitles - any(percent_subtitles) OVER (
        ORDER BY month ASC ROWS BETWEEN 1 PRECEDING AND 1 PRECEDING
    ) AS previous
FROM youtube
GROUP BY month
ORDER BY month ASC;
