WITH(60) -- デフォルト 60m
AS timeRange,
(
    --NON NULL uuidを持つ*任意の*テーブルに対して、暗黙的なMVの隠しターゲットテーブル名を準備する
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --MV名とターゲットテーブルを取得する（TOが指定されている場合）
    --TODO: extractは最初のキャプチャグループのみを返すようだ :( regexpExtractが利用可能になったら置き換える
    SELECT arrayFilter(
            x->x != '',
            --空のキャプチャを除去する
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
        -- 60m以内のSELECTのみ
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- クエリが暗黙的なMVターゲットテーブル名を参照しているか確認する
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- クエリが明示的なMVターゲットテーブルを参照しているか確認する
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
