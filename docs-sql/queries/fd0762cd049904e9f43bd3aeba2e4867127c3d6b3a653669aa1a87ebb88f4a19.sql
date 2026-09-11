{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type 可能被推断为 String，但我们可能更倾向于使用 LowCardinality(String)：
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState() 可能被推断为 `AggregateFunction(count)`，但我们可能更倾向于修改所用参数的类型：
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState() 可能被推断为 `SimpleAggregateFunction(max, String)`，但我们同样可能更倾向于修改所用参数的类型：
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
