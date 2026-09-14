SELECT arrayFirstOrNull(x, y -> x=y, [0, 1, 2], [3, 3, 3]) AS res, toTypeName(res)
