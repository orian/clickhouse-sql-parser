--標準SQL
SELECT avg(DISTINCT number)

--集約コンビネータを使用したClickHouse
SELECT avgDistinct(number)
