
        INSERT INTO json_to_string
        SELECT
            number,
            CAST((
                number,
                ['info', 'warn', 'error', 'debug'][(number % 4) + 1],
                map('user', randomPrintableASCII(8), 'session', randomPrintableASCII(12)),
                [randomPrintableASCII(5), randomPrintableASCII(6)],
                CAST((randomPrintableASCII(6), rand64() / 1000000.0) AS Tuple(country String, lat Float64)),
                (number % 2) = 0
            ) AS Tuple(user_id UInt64, event String, props Map(String, String),
                       tags Array(String), geo Tuple(country String, lat Float64), active Bool))::JSON
        FROM numbers(500000)
    