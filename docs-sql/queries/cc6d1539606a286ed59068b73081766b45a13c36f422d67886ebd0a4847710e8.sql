ALTER TABLE laion_5b_10m ADD INDEX vector_index vector TYPE vector_similarity('hnsw', 'cosineDistance', 768, 'bf16', 64, 512);

ALTER TABLE laion_5b_10m MATERIALIZE INDEX vector_index SETTINGS mutations_sync = 2;
