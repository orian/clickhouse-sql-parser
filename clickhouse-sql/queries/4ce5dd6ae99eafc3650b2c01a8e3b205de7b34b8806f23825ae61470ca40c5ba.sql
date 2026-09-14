
        SELECT toString(number % 1000) AS a,
               toString(number %    7) AS b,
               toString(number %   13) AS c,
               toString(number %   97) AS d,
               toString(number %   31) AS e,
               toString(number %   53) AS f,
               toString(number %   11) AS g
        FROM numbers(2000000)
        GROUP BY 1, 2, 3, 4, 5, 6, 7
        FORMAT Null
    