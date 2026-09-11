CREATE TABLE wf_frame_groups (`order` UInt64, value UInt64) ENGINE = Memory;
INSERT INTO wf_frame_groups FORMAT Values (10, 1), (10, 2), (20, 3), (30, 4), (30, 5);

SELECT
    order,
    value,
    groupArray(value) OVER (ORDER BY order ROWS   BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS rows_frame,
    groupArray(value) OVER (ORDER BY order RANGE  BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS range_frame,
    groupArray(value) OVER (ORDER BY order GROUPS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS groups_frame
FROM wf_frame_groups
ORDER BY order, value;
