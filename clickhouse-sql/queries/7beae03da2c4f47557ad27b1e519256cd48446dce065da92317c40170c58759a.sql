
        WITH (SELECT u8 FROM adss_mixed_{dim} LIMIT 1) AS needle
        SELECT sum(L2Distance(needle, f64)) FROM adss_mixed_{dim}
    