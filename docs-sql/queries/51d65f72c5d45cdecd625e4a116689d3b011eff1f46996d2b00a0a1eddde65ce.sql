-- Insert into an existing Snowflake-managed Iceberg table
INSERT INTO horizon_catalog.`PUBLIC.test_table`
SELECT
    number AS id,
    concat('name_', toString(number)) AS name
FROM numbers(100);

-- Create a new Iceberg table in the Horizon catalog (requires CREATE ICEBERG TABLE)
CREATE TABLE horizon_catalog.`PUBLIC.clickhouse_written`
(
    id Int64,
    name String
)
ENGINE = Iceberg;
