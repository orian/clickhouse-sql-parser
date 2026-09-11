
        INSERT INTO window_test 
        SELECT number, rand(1) % 500, number % 3000, randomPrintableASCII(2) FROM numbers(2000000)
    