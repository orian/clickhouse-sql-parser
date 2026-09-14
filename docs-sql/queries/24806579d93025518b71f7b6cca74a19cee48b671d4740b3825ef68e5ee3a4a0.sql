SET allow_experimental_database_iceberg = 1;
SET allow_experimental_insert_into_iceberg = 1;
SET write_full_path_in_iceberg_metadata = 1;

CREATE TABLE lake.`sales.returns` (id Int64, reason String)
ENGINE = IcebergS3('http://seaweedfs:8333/analytics/sales/returns/', 'tutorialkey', 'tutorialsecret');

INSERT INTO lake.`sales.returns` VALUES (1, 'damaged'), (2, 'wrong size');

SELECT * FROM lake.`sales.returns` ORDER BY id;
