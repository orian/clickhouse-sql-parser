-- A data a seguir é 21 de abril de 2023, que foi uma sexta-feira:
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
