-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- これはデフォルト値なので、実際に設定する必要はありません。
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
