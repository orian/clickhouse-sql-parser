INSERT INTO data_stage
SELECT
    id,
    data
FROM s3('https://bucket-name/*.avro', 'Avro')
SETTINGS schema_inference_mode='union';
