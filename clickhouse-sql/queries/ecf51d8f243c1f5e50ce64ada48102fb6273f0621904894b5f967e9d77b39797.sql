
        INSERT INTO test_{serialization}_{ratio} SELECT
            number,
            number % {ratio} = 0 ? rand(1) : 0,
            number % {ratio} = 0 ? rand(2) : 0,
            number % {ratio} = 0 ? randomPrintableASCII(64, 3) : ''
        FROM numbers(100000000)
    