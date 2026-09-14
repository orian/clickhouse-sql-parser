CREATE TABLE largestTriangleThreeBuckets_test (x Float64, y Float64) ENGINE = Memory;
INSERT INTO largestTriangleThreeBuckets_test VALUES
    (1.0, 10.0), (2.0, 20.0), (3.0, 15.0), (8.0, 60.0), (9.0, 55.0),
    (10.0, 70.0), (4.0, 30.0), (5.0, 40.0), (6.0, 35.0), (7.0, 50.0);

SELECT largestTriangleThreeBuckets(4)(x, y) FROM largestTriangleThreeBuckets_test;
