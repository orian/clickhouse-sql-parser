-- If you need non-aggregate function to choose a maximum of two values, see greatest():
SELECT greatest(a, b) FROM values('a Int32, b Int32', (1, 2), (5, 3));
