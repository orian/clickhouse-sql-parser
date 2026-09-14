
        WITH (SELECT a FROM array_distance_f32_{array_size} LIMIT 1) AS needle
        SELECT sum(cosineDistance(needle, b)) FROM array_distance_f32_{array_size}
    