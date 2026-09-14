
        INSERT INTO sparse_column_index_runs_{run_length}
        SELECT
            number,
            number % 20 = 0 ? number : 0,
            number % 20 = 0 ? repeat(toString(number), 64) : ''
        FROM numbers(intDiv(10000000, {run_length}))
    