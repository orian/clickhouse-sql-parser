-- In this example, arrayEnumerateDenseRanked is used to obtain an array indicating, for each element of the
-- multidimensional array, what its position is among elements of the same value.
-- For the first row of the passed array, [10, 10, 30, 20], the corresponding first row of the result is [1, 1, 2, 3],
-- indicating that 10 is the first number encountered in position 1 and 2, 30 the second number encountered in position 3
-- and 20 is the third number encountered in position 4.
-- For the second row, [40, 50, 10, 30], the corresponding second row of the result is [4,5,1,2], indicating that 40
-- and 50 are the fourth and fifth numbers encountered in position 1 and 2 of that row, that another 10
-- (the first encountered number) is in position 3 and 30 (the second number encountered) is in the last position.

SELECT arrayEnumerateDenseRanked(1,[[10,10,30,20],[40,50,10,30]],2);
