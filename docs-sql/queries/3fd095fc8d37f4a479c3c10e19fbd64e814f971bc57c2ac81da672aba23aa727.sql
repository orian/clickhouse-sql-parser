{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type puede inferirse como String pero podríamos preferir LowCardinality(String):
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState() puede inferirse como `AggregateFunction(count)` pero podríamos preferir cambiar el tipo del argumento utilizado:
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState() puede inferirse como `SimpleAggregateFunction(max, String)` pero podríamos preferir cambiar también el tipo del argumento utilizado:
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
