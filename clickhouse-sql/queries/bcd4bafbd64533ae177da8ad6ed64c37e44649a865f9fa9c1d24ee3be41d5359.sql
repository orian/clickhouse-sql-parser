
        -- Worst-case performance test for `CrossTabPhiSquaredWindowData` inherited aggregate functions
        SELECT
            cramersV(a, b) OVER (ORDER BY number) AS v
        FROM
        (
            SELECT
                number,
                toUInt32(number % 1000) AS a,
                toUInt32(intDiv(number, 1000) % 1000) AS b
            FROM numbers(10_000)
        ) FORMAT NULL;
    