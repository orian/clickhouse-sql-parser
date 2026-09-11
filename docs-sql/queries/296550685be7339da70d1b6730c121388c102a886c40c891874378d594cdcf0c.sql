CREATE TABLE hits_iceberg
    ENGINE = IcebergS3('https://datasets-documentation.s3.amazonaws.com/lake_formats/iceberg/')
