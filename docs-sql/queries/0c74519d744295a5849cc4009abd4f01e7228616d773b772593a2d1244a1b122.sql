SELECT d, d + 1 AS res, toTypeName(res), dynamicType(d) FROM test;
