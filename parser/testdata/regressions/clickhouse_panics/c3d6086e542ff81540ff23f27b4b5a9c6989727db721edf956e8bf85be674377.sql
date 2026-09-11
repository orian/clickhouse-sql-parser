
        -- theilsU has optimized window variant that does not use `CrossTabPhiSquaredWindowData` so it is much faster than cramersV for the same data
        SELECT
            theilsU(a, b) OVER (ORDER BY number) AS v
        FROM
        (
            SELECT
                number,
                toUInt32(number % 1000) AS a,
                toUInt32(intDiv(number, 1000) % 1000) AS b
            FROM numbers(10_000)
        ) FORMAT NULL;
    