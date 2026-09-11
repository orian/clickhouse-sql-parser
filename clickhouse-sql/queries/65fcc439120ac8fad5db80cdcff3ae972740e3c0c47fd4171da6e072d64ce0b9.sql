
        SELECT count() FROM phj_asof_left AS l
        ASOF LEFT JOIN phj_asof_right AS r
        ON l.key = r.key AND l.ts <= r.ts
        FORMAT Null
    