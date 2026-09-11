-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- este es el valor predeterminado, por lo que no necesitas configurarlo explícitamente.
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
