
        INSERT INTO permute_cache_src SELECT
            toString(number % 1000),
            toString((number * 17) % 10000),
            toString((number * 31) % 100000),
            toString((number * 53) % 1000000),
            toString(number),
            toDateTime64('2024-01-01', 6) + number / 1000000,
            toString(number),
            number
        FROM numbers(2000000)
    