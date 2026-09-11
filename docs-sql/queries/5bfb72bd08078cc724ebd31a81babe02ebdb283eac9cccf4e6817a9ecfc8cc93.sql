SELECT arrayPartialShuffle(materialize([1, 2, 3, 4]), 2, 42), arrayPartialShuffle([1, 2, 3], 2, 42) FROM numbers(10)
