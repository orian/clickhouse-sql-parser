SELECT blob_type, snapshot_id, sequence_number, offset, length, compression_codec,
       mapKeys(properties), mapValues(properties)
FROM file(deletes.puffin, PuffinMetadata);
