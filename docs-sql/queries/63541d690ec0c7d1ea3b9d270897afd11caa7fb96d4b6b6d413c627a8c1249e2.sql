-- models/page_events_aggregator.sql
{{ config(
    materialized='materialized_view',
    catchup=True  -- este é o valor padrão, portanto não é necessário defini-lo explicitamente.
) }}
{{ materialization_target_table(ref('events_daily')) }}
...
