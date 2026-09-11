   {{ config(order_by='(updated_at, id, name)', engine='MergeTree()', materialized='table') }}
