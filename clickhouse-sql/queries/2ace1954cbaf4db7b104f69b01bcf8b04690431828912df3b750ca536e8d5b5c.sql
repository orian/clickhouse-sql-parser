
WITH
    if(value = 10, 'A', if(value = 20, 'B', if(value = 30, 'C', if(value = 40, 'D', if(value = 50, 'E', 'F'))))) AS col1,
    if(col1 = 'A', 'Alpha', if(col1 = 'B', 'Beta', if(col1 = 'C', 'Gamma', if(col1 = 'D', 'Delta', if(col1 = 'E', 'Epsilon', 'Other'))))) AS col2,
    if(col2 = 'Alpha', 1, if(col2 = 'Beta', 2, if(col2 = 'Gamma', 3, if(col2 = 'Delta', 4, if(col2 = 'Epsilon', 5, 0))))) AS col3,
    if(col3 = 1, 'One', if(col3 = 2, 'Two', if(col3 = 3, 'Three', if(col3 = 4, 'Four', if(col3 = 5, 'Five', 'Zero'))))) AS col4,
    if(col4 = 'One', 'Uno', if(col4 = 'Two', 'Dos', if(col4 = 'Three', 'Tres', if(col4 = 'Four', 'Cuatro', if(col4 = 'Five', 'Cinco', 'Cero'))))) AS col5,
    if(col5 = 'Uno', 'I', if(col5 = 'Dos', 'II', if(col5 = 'Tres', 'III', if(col5 = 'Cuatro', 'IV', if(col5 = 'Cinco', 'V', 'Other'))))) AS col6,
    if(col6 = 'I', 'Primero', if(col6 = 'II', 'Segundo', if(col6 = 'III', 'Tercero', if(col6 = 'IV', 'Cuarto', if(col6 = 'V', 'Quinto', 'Otro'))))) AS col7
SELECT col1, col2, col3, col4, col5, col6, col7
FROM (SELECT number AS value FROM system.numbers LIMIT 100)
FORMAT Null
    