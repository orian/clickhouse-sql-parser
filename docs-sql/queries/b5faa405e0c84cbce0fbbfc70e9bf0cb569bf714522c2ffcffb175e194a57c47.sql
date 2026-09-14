DROP TABLE IF EXISTS test_data;
CREATE TABLE test_data
(
    population Float64,
)
ENGINE = Log;

INSERT INTO test_data SELECT randUniform(5.5, 10) FROM numbers(1000000);

SELECT
    stddevPopStable(population) AS stddev
FROM test_data;
