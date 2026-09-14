{{ config(
       materialized='table',
       projections=[
           {
               'name': 'proj_by_dept_age',
               'index': ['department', 'age']
           }
       ]
) }}
