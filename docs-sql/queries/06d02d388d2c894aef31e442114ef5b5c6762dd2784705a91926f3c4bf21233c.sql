-- Query with expression as argument
SELECT groupArraySample(3)(concat('light-', color)) as newcolors FROM colors;
