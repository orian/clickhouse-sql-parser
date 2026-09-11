
        CREATE TABLE array_join_function_cross (A Array(UInt32), B Array(UInt32), C Array(UInt32))
        ENGINE = MergeTree ORDER BY tuple()
    