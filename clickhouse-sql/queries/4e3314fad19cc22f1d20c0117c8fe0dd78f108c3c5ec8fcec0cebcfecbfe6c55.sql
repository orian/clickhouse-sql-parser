
        INSERT INTO multipoly_holder
        SELECT groupArray(poly) AS poly
        FROM   polys;
    