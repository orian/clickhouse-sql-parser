WITH(60) -- predeterminado 60m
AS timeRange,
(
    --preparar nombres de posibles tablas destino ocultas implícitas de MV para *cualquier* tabla con uuid NO NULO
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --captura el nombre de la MV y las tablas destino (si se especifica TO)
    --TODO parece que extract devuelve solo el primer grupo de captura :( reemplazar con regexpExtract cuando esté disponible
    SELECT arrayFilter(
            x->x != '',
            --eliminar capturas vacías
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
        -- solo SELECT en los últimos 60m
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- verificar si la consulta involucra nombres de tablas destino implícitas de MV
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- verificar si la consulta involucra tablas destino explícitas de MV
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
