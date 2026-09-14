-- In this example, arrayEnumerateUniqRanked is used to obtain an array indicating,
-- for each element of the multidimensional array, what its position is among elements
-- of the same value. For the first row of the passed array, [1, 2, 3], the corresponding
-- result is [1, 1, 1], indicating that this is the first time 1, 2 and 3 are encountered.
-- For the second row of the provided array, [2, 2, 1], the corresponding result is [2, 3, 3],
-- indicating that 2 is encountered for a second and third time, and 1 is encountered
-- for the second time. Likewise, for the third row of the provided array [3] the
-- corresponding result is [2] indicating that 3 is encountered for the second time.

SELECT arrayEnumerateUniqRanked(1, [[1, 2, 3], [2, 2, 1], [3]], 2);
