
        SELECT toStartOfHour(toDateTime(number % 86400)) AS hr,
               number % {cardinality} AS user, number AS v
        FROM numbers_mt(10000000)
        LIMIT 10 BY hr, user
        FORMAT NULL
    