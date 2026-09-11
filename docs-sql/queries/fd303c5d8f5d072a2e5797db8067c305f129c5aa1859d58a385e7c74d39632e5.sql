-- with clear_depth=1 and max_array_depth=1, the result of arrayEnumerateUniqRanked
-- is identical to that which arrayEnumerateUniqwould give for the same array.

SELECT arrayEnumerateUniqRanked(1, [[1, 2, 3], [2, 2, 1], [3]], 2);
