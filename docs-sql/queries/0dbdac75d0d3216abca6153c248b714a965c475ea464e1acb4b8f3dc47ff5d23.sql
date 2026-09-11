SET allow_deprecated_error_prone_window_functions = 1;

CREATE TABLE events
(
    EventID UInt32,
    EventDate Date,
    EventTime DateTime
)
ENGINE = Memory;

INSERT INTO events VALUES
    (1106, '2025-11-24', '2025-11-24 00:00:04'),
    (1107, '2025-11-24', '2025-11-24 00:00:05'),
    (1108, '2025-11-24', '2025-11-24 00:00:05'),
    (1109, '2025-11-24', '2025-11-24 00:00:09'),
    (1110, '2025-11-24', '2025-11-24 00:00:10');

SELECT
    EventID,
    EventTime,
    runningDifference(EventTime) AS delta
FROM
(
    SELECT
        EventID,
        EventTime
    FROM events
    WHERE EventDate = '2025-11-24'
    ORDER BY EventTime ASC, EventID ASC
    LIMIT 5
);
