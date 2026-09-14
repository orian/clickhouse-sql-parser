
        WITH (SELECT u8 FROM array_distance_mixed_{array_size} LIMIT 1) AS needle
        SELECT sum(L2Distance(needle, f64)) FROM array_distance_mixed_{array_size}
    