-- row_number does not respect the frame, so rn_1 = rn_2 = rn_3 != rn_4
SELECT
    part_key,
    value,
    order,
    groupArray(value) OVER w1 AS frame_values,
    row_number() OVER w1 AS rn_1,
    sum(1) OVER w1 AS rn_2,
    row_number() OVER w2 AS rn_3,
    sum(1) OVER w2 AS rn_4
FROM wf_frame
WINDOW
    w1 AS (PARTITION BY part_key ORDER BY order DESC),
    w2 AS (
        PARTITION BY part_key 
        ORDER BY order DESC 
        ROWS BETWEEN 1 PRECEDING AND CURRENT ROW
    )
ORDER BY
    part_key ASC,
    value ASC;

┌─part_key─┬─value─┬─order─┬─frame_values─┬─rn_1─┬─rn_2─┬─rn_3─┬─rn_4─┐
│        1 │     1 │     1 │ [5,4,3,2,1]  │    5 │    5 │    5 │    2 │
│        1 │     2 │     2 │ [5,4,3,2]    │    4 │    4 │    4 │    2 │
│        1 │     3 │     3 │ [5,4,3]      │    3 │    3 │    3 │    2 │
│        1 │     4 │     4 │ [5,4]        │    2 │    2 │    2 │    2 │
│        1 │     5 │     5 │ [5]          │    1 │    1 │    1 │    1 │
└──────────┴───────┴───────┴──────────────┴──────┴──────┴──────┴──────┘
