SELECT count() FROM nums WHERE NOT ignore(
        {op}({arg}, {arg}),
        {op}({arg}, {arg} + 1),
        {op}({arg}, {arg} + 2),
        {op}({arg}, {arg} + 3),
        {op}({arg}, {arg} + 4),
        {op}({arg}, {arg} + 5),
        {op}({arg}, {arg} + 6),
        {op}({arg}, {arg} + 7)
        )