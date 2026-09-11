
        INSERT INTO map_basic_compact SELECT number,
            map('k0', randomString(1000), 'k1', randomString(1000), 'k2', randomString(1000), 'k3', randomString(1000), 'k4', randomString(1000))
        FROM numbers(100000)
    