
        INSERT INTO phj_right_nullable
        SELECT if(number % 50 = 0, NULL, number % 1000), toString(number)
        FROM numbers(5000000)
    