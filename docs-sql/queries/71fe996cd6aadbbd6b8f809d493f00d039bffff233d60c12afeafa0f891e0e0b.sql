SELECT
    isDistinctFrom(1, 2) AS result_1,
    isDistinctFrom(1, 1) AS result_2,
    isDistinctFrom(NULL, 1) AS result_3,
    isDistinctFrom(NULL, NULL) AS result_4
