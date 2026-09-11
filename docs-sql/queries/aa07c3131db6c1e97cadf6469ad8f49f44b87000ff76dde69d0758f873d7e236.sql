{{
    config(
        materialized='table',
        engine='SummingMergeTree()',
        order_by='(event_date, event_type)',
        partition_by='toYYYYMM(event_date)'
    )
}}

SELECT
    toDate(now()) AS event_date,
    '' AS event_type,
    toUInt64(0) AS total
WHERE 0  -- 创建具有正确 schema 的空表
