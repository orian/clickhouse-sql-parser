
        INSERT INTO polygons
        SELECT toFloat64(1000000) / ((number % 100000) + 1), toFloat64(1000000) / ((number % 100000) + 1) from system.numbers LIMIT 10000000
        SETTINGS max_insert_threads = 2, max_memory_usage = 30000000000;
    