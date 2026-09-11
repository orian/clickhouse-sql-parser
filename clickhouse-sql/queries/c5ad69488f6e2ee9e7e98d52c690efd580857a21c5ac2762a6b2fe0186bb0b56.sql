
        ${CLICKHOUSE_CLIENT} --query "
            TRUNCATE TABLE ttl_group_by;
            ALTER TABLE ttl_group_by ATTACH PARTITION tuple() FROM ttl_group_by_src;
            OPTIMIZE TABLE ttl_group_by FINAL;
        "
    