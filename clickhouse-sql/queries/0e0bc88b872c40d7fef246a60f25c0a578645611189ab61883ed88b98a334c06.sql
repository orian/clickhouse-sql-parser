
        SELECT count() FROM phj_left AS l RIGHT JOIN phj_right_nullable AS r ON l.key = r.key
        FORMAT Null
    