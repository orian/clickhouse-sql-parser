
        OPTIMIZE TABLE {database}.many_manifests MANIFEST
        SETTINGS allow_experimental_iceberg_compaction = 1, iceberg_manifest_min_count_to_compact = 8
    