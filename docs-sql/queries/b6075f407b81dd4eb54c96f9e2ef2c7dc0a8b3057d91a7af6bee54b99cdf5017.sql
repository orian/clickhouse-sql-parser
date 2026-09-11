SELECT getSubcolumn(array_col, 'size0'), getSubcolumn(tuple_col, 'elem_name')
FROM values('array_col Array(UInt32), tuple_col Tuple(elem_name String)', ([1, 2, 3], tuple('abc')));
