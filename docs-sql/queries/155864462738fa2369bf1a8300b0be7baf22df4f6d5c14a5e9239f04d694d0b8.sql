SELECT
    count()
FROM events
WHERE (_part, _part_offset) IN (
    SELECT _part, _part_offset
    FROM events
    WHERE user_id = 42
)
