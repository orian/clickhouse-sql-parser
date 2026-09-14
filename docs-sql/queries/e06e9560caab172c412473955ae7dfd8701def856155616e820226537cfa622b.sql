{{
    config(
        materialized="materialized_view",
        engine="AggregatingMergeTree",
        order_by=["event_type"],
    )
}}

select
  -- event_type은 String으로 추론될 수 있지만 LowCardinality(String)이 더 적합할 수 있습니다:
  CAST(event_type, 'LowCardinality(String)') as event_type,
  -- countState()는 `AggregateFunction(count)`으로 추론될 수 있지만 인수의 유형을 변경하는 것이 더 적합할 수 있습니다:
  CAST(countState(), 'AggregateFunction(count, UInt32)') as response_count, 
  -- maxSimpleState()는 `SimpleAggregateFunction(max, String)`으로 추론될 수 있지만 인수의 유형도 함께 변경하는 것이 더 적합할 수 있습니다:
  CAST(maxSimpleState(event_type), 'SimpleAggregateFunction(max, LowCardinality(String))') as max_event_type
from {{ ref('user_events') }}
group by event_type
