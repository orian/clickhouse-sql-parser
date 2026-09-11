WITH
    left_table AS (SELECT * FROM VALUES('c UInt32', 1, 2, 3)),
    right_table AS (SELECT * FROM VALUES('c UInt32', 2, 2, 3, 3, 4))
SELECT
    l.c AS l_c,
    r.c AS r_c
FROM left_table AS l
RIGHT ANY JOIN right_table AS r ON l.c = r.c;
