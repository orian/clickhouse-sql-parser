SELECT description
FROM system.documentation
WHERE type = 'Setting' AND name = 'async_insert_max_data_size'
FORMAT TSVRaw;
