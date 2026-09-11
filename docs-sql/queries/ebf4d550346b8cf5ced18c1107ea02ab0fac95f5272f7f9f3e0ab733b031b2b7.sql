{{ config(
       materialized='table',
       projections=[
           {
               'name': 'proj_by_age',
               'index': 'age'
           }
       ]
) }}
