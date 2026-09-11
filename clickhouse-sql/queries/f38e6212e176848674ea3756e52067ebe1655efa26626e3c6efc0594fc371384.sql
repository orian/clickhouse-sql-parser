
        SELECT *, sum(v) OVER () AS w FROM window_transform_wide_lc FORMAT Null;
    