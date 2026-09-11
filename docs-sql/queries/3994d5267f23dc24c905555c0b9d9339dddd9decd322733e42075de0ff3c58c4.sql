SELECT table, name, data_compressed_bytes, data_uncompressed_bytes FROM system.columns
WHERE table LIKE 'my_%_table';
