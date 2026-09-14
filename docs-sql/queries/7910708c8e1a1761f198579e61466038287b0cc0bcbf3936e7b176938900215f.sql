SELECT d, d + 1 AS res, toTypeName(res), dynamicType(res) FROM test WHERE dynamicType(d) NOT IN ('String', 'Array(Int64)', 'None')
