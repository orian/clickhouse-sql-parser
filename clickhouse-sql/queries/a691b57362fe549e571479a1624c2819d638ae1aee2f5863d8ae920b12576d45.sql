
        CREATE TABLE points
        (
            id UInt32,
            pt Tuple(Int32,Int32)
        ) ENGINE = MergeTree ORDER BY id;
    