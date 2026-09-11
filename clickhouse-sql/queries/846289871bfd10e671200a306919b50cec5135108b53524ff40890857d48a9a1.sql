
        CREATE TABLE index_test_cast
        (
            z UInt32,
            x UInt16 ALIAS mortonDecode(2, z).1,
            y UInt16 ALIAS mortonDecode(2, z).2,
            INDEX i_x (mortonDecode(2, z).1) TYPE minmax,
            INDEX i_y (mortonDecode(2, z).2) TYPE minmax
        ) ENGINE = MergeTree ORDER BY z
    