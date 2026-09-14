
        SELECT (number % 1000)::Dynamic AS a,
               (number %    7)::Dynamic AS b,
               (number %   13)::Dynamic AS c,
               (number %   97)::Dynamic AS d,
               (number %   31)::Dynamic AS e,
               (number %   53)::Dynamic AS f,
               (number %   11)::Dynamic AS g
        FROM numbers(2000000)
        GROUP BY 1, 2, 3, 4, 5, 6, 7
        FORMAT Null
    