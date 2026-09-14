CREATE TABLE colors (
    id Int32,
    color String
) ENGINE = Memory;

INSERT INTO colors VALUES
(1, 'red'),
(2, 'blue'),
(3, 'green'),
(4, 'white'),
(5, 'orange');

SELECT groupArraySample(3)(color) as newcolors FROM colors;
