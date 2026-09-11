SELECT
    isNotDistinctFrom(1, 1) AS result_1,
    isNotDistinctFrom(1, 2) AS result_2,
    isNotDistinctFrom(NULL, NULL) AS result_3,
    isNotDistinctFrom(NULL, 1) AS result_4
