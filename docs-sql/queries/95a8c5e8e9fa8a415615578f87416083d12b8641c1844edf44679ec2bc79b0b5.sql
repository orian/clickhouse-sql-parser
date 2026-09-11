ENGINE = MergeTree()
ORDER BY reverse(tuple(reverse(p), hex(p)))
