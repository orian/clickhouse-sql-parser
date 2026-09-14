-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- это значение по умолчанию, поэтому его не нужно указывать явно.
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
