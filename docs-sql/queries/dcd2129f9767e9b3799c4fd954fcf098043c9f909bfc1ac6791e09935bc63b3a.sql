WITH(60) -- по умолчанию 60 мин
AS timeRange,
(
    --подготовить имена возможных неявных скрытых целевых таблиц MV для *любой* таблицы с NON NULL uuid
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --получает имя MV и целевые таблицы (если указано TO)
    --TODO похоже, extract возвращает только первую capturing group :( заменить на regexpExtract, когда станет доступно
    SELECT arrayFilter(
            x->x != '',
            --удалить пустые совпадения
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
        -- только SELECT за последние 60 мин
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- проверить, что запрос обращается к неявным именам целевых таблиц MV
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- проверить, что запрос обращается к явной целевой таблице MV
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
