
        SELECT
            sum(w), sum(w + 1), sum(w + 2), sum(w + 3), sum(w + 4),
            sum(w + 5), sum(w + 6), sum(w + 7), sum(w + 8), sum(w + 9),
            sum(w + 10), sum(w + 11), sum(w + 12), sum(w + 13), sum(w + 14),
            sum(w + 15), sum(w + 16), sum(w + 17), sum(w + 18), sum(w + 19),
            sum(w + 20), sum(w + 21), sum(w + 22), sum(w + 23), sum(w + 24),
            sum(w + 25), sum(w + 26), sum(w + 27), sum(w + 28), sum(w + 29)
        FROM narrow_col
        FORMAT Null
    