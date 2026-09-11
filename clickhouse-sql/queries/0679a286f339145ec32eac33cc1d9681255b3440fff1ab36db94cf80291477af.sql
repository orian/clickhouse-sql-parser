
        WITH
            (SELECT poly FROM multipoly_holder LIMIT 1) AS mp
        SELECT count()
        FROM points
        WHERE NOT ignore(pointInPolygon(pt, mp));
    