WITH(60) -- padrão 60m
AS timeRange,
(
    --prepara nomes de possíveis tabelas de destino ocultas implícitas de MV para *qualquer* tabela com uuid NÃO NULO
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --captura o nome da MV e as tabelas de destino (se TO for especificado)
    --TODO parece que extract retorna apenas o primeiro grupo de captura :( substituir por regexpExtract quando disponível
    SELECT arrayFilter(
            x->x != '',
            --remove capturas vazias
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
        -- apenas SELECT nos últimos 60m
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- verifica se a consulta envolve nomes de tabelas de destino implícitas de MV
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- verifica se a consulta envolve tabela de destino explícita de MV
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
