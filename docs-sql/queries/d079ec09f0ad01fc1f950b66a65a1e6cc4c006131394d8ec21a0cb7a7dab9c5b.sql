{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type pode ser inferido como String, mas podemos preferir LowCardinality(String):
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState() pode ser inferido como `AggregateFunction(count)`, mas podemos preferir alterar o tipo do argumento utilizado:
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState() pode ser inferido como `SimpleAggregateFunction(max, String)`, mas podemos preferir também alterar o tipo do argumento utilizado:
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
