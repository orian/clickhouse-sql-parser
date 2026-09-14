SELECT groupArraySorted(5)(str) FROM (SELECT toString(number) AS str FROM numbers(5));
