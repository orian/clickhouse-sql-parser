-- Changing clear_depth=2, results in elements being enumerated separately for each row.
SELECT arrayEnumerateUniqRanked(2,[[1, 2, 3],[2, 2, 1],[3]], 2);
