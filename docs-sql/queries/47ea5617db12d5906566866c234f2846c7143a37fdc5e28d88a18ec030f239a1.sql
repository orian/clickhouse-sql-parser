WITH(60) -- 기본값 60분
AS timeRange,
(
    --NULL이 아닌 uuid를 가진 *모든* 테이블에 대해 암시적 MV 숨겨진 대상 테이블의 가능한 이름을 준비
    SELECT groupArray(
            concat('default.`.inner_id.', toString(uuid), '`')
        )
    FROM clusterAllReplicas(default, system.tables)
    WHERE notEmpty(uuid)
) AS MV_implicit_possible_hidden_target_tables_names_array,
(
    --MV 이름과 대상 테이블을 캡처 (TO가 지정된 경우)
    --TODO: extract가 첫 번째 캡처 그룹만 반환하는 것으로 보임 :( regexpExtract 사용 가능 시 교체 필요
    SELECT arrayFilter(
            x->x != '',
            --빈 캡처 제거
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
        -- 60분 이내의 SELECT만 해당
        event_time > now() - toIntervalMinute(timeRange)
        AND startsWith(query, 'SELECT')
    ) -- 쿼리에 암시적 MV 대상 테이블 이름이 포함되는지 확인
    AND (
        hasAny(
            tables,
            MV_implicit_possible_hidden_target_tables_names_array
        )
        OR -- 쿼리에 명시적 MV 대상 테이블이 포함되는지 확인
        hasAny(tables, MV_explicit_target_tables_names_array)
    )
ORDER BY event_time DESC;
