-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- 这是默认值，实际上无需显式设置。
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
