WITH(60) -- 默认 60m
AS timeRange,
(
    --为所有具有非 NULL uuid 的表准备可能的隐式 MV 隐藏目标表名
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --获取 MV 名称及目标表（如果指定了 TO）
    --TODO 看起来 extract 只会返回第一个捕获组 :( 待 regexpExtract 可用后替换
    SELECT arrayFilter(
            x->x != '',
            --移除空捕获项
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
        -- 仅查询 60m 内的 SELECT 语句
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- 检查查询是否涉及隐式 MV 目标表名
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- 检查查询是否涉及显式 MV 目标表
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
