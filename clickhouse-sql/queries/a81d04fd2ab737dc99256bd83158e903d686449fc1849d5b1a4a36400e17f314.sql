
        INSERT INTO qcc_topk_test
        SELECT rand(), number, randomFixedString(256)
        FROM numbers(10000000)
    