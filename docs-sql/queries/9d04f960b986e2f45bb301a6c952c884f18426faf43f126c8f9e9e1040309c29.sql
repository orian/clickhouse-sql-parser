-- Query with column name and different seed
SELECT groupArraySample(3, 987654321)(color) as newcolors FROM colors;
