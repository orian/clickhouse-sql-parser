
        SELECT *, count() OVER () AS w FROM window_transform_wide_lc FORMAT Null;
    