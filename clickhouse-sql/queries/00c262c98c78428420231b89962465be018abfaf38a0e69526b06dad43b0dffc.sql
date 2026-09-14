
        INSERT INTO date_preimage_perf
        SELECT
            number AS id,
            toDateTime(1772323200 + intDiv(number, 25000000) * 86400 + number % 86400, 'UTC') AS seen,
            toDate(seen) AS d
        FROM numbers(50000000)
    