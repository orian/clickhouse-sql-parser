{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type は String として推論される場合があるが、LowCardinality(String) を使用したい場合がある:
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState() は `AggregateFunction(count)` として推論される場合があるが、使用する引数の型を変更したい場合がある:
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState() は `SimpleAggregateFunction(max, String)` として推論される場合があるが、使用する引数の型も変更したい場合がある:
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
