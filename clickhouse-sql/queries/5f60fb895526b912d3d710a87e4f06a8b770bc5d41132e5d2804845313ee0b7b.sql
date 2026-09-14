
        SELECT count() FROM phj_left AS l FULL JOIN phj_right_nullable AS r ON l.key = r.key
        FORMAT Null
    