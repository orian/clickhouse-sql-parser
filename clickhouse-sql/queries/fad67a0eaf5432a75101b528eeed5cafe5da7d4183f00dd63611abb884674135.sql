
        SELECT l.id, l.k, r.value
        FROM topk_join_l AS l INNER JOIN topk_join_r AS r ON r.id = l.id
        ORDER BY l.k DESC LIMIT 10
        FORMAT Null
    