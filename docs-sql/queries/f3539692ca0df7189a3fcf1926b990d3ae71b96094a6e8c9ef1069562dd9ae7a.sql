WITH(60) -- default 60m
AS timeRange,
(
    --prepare names of possible implicit MV hidden target tables for *any* table with NON NULL uuid
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --captures MV name and target tables (if TO is specified)
    --TODO it seems that extract will return just the first capturing group :( replace with regexpExtract once available
    SELECT arrayFilter(
            x->x != '',
            --remove empty captures
            groupArray(
                extract(
                    create_table_query,
                    '^CREATE MATERIALIZED VIEW\s(\w+\.\w+)\s(?:TO\s(\S+))?'
                )
            )
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE engine = 'MaterializedView'
) AS MV_explicit_target_tables_names_array
SELECT event_time,
    query,
    tables as "MVs tables"
FROM clusterAllReplicas(default, system.query_log)
WHERE (
        -- only SELECT within 60m
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- check either that query involves implicit MV target table names
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- check that query involves explicit MV target table
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
