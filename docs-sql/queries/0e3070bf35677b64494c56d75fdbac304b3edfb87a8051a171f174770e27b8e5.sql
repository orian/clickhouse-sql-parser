{{
    config(
        materialized='distributed_table',
        order_by='id, created_at',
        sharding_key='cityHash64(id)',
        engine='ReplacingMergeTree'
    )
}}

select id, created_at, item
from {{ source('db', 'table') }}
