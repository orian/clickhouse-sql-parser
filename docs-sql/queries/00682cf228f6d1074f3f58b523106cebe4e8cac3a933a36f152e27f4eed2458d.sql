SELECT arrayFold(acc,x -> acc + x*2, [1, 2, 3, 4], 3::Int64) AS res;
