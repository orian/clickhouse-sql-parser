
        INSERT INTO json_compact SELECT number,
            map('p0', hex(randomString(500)), 'p1', hex(randomString(500)), 'p2', hex(randomString(500)), 'p3', hex(randomString(500)), 'p4', hex(randomString(500)))::JSON(max_dynamic_paths = 0)
        FROM numbers(150000)
    