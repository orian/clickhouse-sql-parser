
        SELECT l.id, l.k
        FROM topk_join_l AS l
        LEFT JOIN topk_join_r AS r1 ON r1.id = l.id
        LEFT JOIN topk_join_r AS r2 ON r2.id = l.id
        LEFT JOIN topk_join_r AS r3 ON r3.id = l.id
        ORDER BY l.k DESC LIMIT 10
        FORMAT Null
    