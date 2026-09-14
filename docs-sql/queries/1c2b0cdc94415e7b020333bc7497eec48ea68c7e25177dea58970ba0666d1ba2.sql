-- With clear_depth=1 and max_array_depth=1, the result of arrayEnumerateUniqRanked
-- is identical to that which arrayEnumerateUniq would give for the same array.

SELECT arrayEnumerateUniqRanked(1, [1, 2, 1], 1);
