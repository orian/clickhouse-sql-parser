-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- 기본값이므로 실제로 설정하지 않아도 됩니다.
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
