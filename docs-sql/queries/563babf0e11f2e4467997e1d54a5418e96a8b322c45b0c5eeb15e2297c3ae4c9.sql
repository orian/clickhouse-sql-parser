-- models/my_model_mv_mv2.sql
{{ config(materialized='materialized_view') }}
{{ materialization_target_table(ref('events_daily')) }}

select a, b, c from {{ source('raw', 'table_2') }}
