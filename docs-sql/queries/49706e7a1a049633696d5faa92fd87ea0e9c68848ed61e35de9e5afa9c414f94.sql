-- With clear_depth=1 and max_array_depth=1, the result is identical to what arrayEnumerateDense would give.

SELECT arrayEnumerateDenseRanked(1,[10, 20, 10, 30],1);
