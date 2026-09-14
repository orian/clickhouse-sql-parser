{{ config(
        engine='MergeTree',
        materialized='table',
        disable_on_cluster='true'
    )
}}

