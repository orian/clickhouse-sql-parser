-- La siguiente fecha es el 21 de abril de 2023, que fue un viernes:
SELECT
    toDayOfWeek(toDateTime('2023-04-21')),
    toDayOfWeek(toDateTime('2023-04-21'), 1)
