
        INSERT INTO volume_reducing_functions
        SELECT number, randomPrintableASCII(100), randomPrintableASCII(64) FROM numbers(10000000)
    