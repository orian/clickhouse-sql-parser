-- Changing clear_depth=2 results in the enumeration occurring separately for each row anew.

SELECT arrayEnumerateDenseRanked(2,[[10,10,30,20],[40,50,10,30]],2);
