SET allow_experimental_paimon_storage_engine = 1;
SET use_paimon_metadata_files_cache = 1;

CREATE TABLE paimon_cached
ENGINE = PaimonS3(paimon_conf, filename = 'paimon_all_types');
