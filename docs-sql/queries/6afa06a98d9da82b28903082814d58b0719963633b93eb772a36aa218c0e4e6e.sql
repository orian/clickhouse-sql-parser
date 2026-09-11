CREATE TABLE my_events (
    start UInt32,
    end UInt32
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO my_events VALUES
(1, 3),
(1, 6),
(2, 5),
(3, 7);

SELECT maxIntersections(start, end) FROM my_events;
