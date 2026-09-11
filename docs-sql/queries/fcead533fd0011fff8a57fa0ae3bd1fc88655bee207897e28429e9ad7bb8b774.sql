{{ config(materialized='materialized_view') }}
{{ materialization_target_table(ref('events_daily')) }}

SELECT
    toStartOfDay(event_time) AS event_date,
    event_type,
    count() AS total
FROM {{ source('raw', 'page_events') }}
GROUP BY event_date, event_type
