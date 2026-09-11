SELECT bin('Hello'), bin(bitSlice('Hello', 1, 8));
SELECT bin('Hello'), bin(bitSlice('Hello', 1, 2));
SELECT bin('Hello'), bin(bitSlice('Hello', 1, 9));
SELECT bin('Hello'), bin(bitSlice('Hello', -4, 8));
