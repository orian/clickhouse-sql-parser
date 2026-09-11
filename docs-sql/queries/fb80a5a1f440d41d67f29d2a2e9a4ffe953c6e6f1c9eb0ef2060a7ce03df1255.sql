SELECT arrayFirstOrNull(x, y -> x=y, ['a', 'b', 'c'], ['c', 'b', 'a'])
