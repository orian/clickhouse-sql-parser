SELECT arrayFold(acc, x -> (acc.2, acc.2 + acc.1),range(number),(1::Int64, 0::Int64)).1 AS fibonacci FROM numbers(1,10);
