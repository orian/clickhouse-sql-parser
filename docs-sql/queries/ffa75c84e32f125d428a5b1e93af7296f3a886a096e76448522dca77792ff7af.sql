-- Frame is bounded by bounds of a partition (BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING)
SELECT
    part_key,
    value,
    order,
    groupArray(value) OVER (
        PARTITION BY part_key 
        ORDER BY order ASC
        ROWS BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING
    ) AS frame_values
FROM wf_frame
ORDER BY
    part_key ASC,
    value ASC;
    
┌─part_key─┬─value─┬─order─┬─frame_values─┐
│        1 │     1 │     1 │ [1,2,3,4,5]  │
│        1 │     2 │     2 │ [1,2,3,4,5]  │
│        1 │     3 │     3 │ [1,2,3,4,5]  │
│        1 │     4 │     4 │ [1,2,3,4,5]  │
│        1 │     5 │     5 │ [1,2,3,4,5]  │
└──────────┴───────┴───────┴──────────────┘
