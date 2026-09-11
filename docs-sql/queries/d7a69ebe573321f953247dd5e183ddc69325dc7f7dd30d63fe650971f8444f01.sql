SELECT *
FROM events_snapshot
WHERE uuid = 'fff'
ORDER BY ts ASC
FORMAT PrettyCompactMonoBlock
