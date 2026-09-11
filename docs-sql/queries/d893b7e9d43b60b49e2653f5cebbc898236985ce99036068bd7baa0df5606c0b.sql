SET allow_experimental_kusto_dialect = 1;
SET dialect = 'kusto';

StormEvents
| where State == 'FLORIDA' and DamageProperty > 0
| summarize Total = sum(DamageProperty) by EventType
| top 5 by Total
