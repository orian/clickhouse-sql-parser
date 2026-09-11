SELECT arrayFill(x -> not isNull(x), [1, null, 2, null]) AS res
