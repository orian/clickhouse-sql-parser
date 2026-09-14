{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type может быть определён как String, но предпочтительнее использовать LowCardinality(String):
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState() может быть определён как `AggregateFunction(count)`, но может потребоваться изменить тип используемого аргумента:
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState() может быть определён как `SimpleAggregateFunction(max, String)`, но может также потребоваться изменить тип используемого аргумента:
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
