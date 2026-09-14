-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- this is the default value so you don't need to actually set it.
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
