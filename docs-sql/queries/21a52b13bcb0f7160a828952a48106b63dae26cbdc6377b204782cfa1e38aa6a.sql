CREATE TABLE num AS numbers(1000);
SELECT quantilesExactInclusive(0.25, 0.5, 0.75, 0.9, 0.95, 0.99, 0.999)(number) FROM num;
