SELECT
arrayUnion([-2, 1], [10, 1], [-2], []) as num_example,
arrayUnion(['hi'], [], ['hello', 'hi']) as str_example,
arrayUnion([1, 3, NULL], [2, 3, NULL]) as null_example
