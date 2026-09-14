{{ config(
       materialized='dictionary',
       fields=[
           ('id', 'UInt64'),
           ('name', 'String'),
       ],
       primary_key='id',
       layout='HASHED()',
       lifetime='MIN 0 MAX 300'
) }}

select id, name from {{ source('raw', 'people') }}
