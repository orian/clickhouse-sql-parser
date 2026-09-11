-- First create the table and insert some data
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1 SETTINGS index_granularity_bytes = '10Mi';
insert into hilbert_numbers (*) values(1,2);

-- Use column names instead of constants as function arguments
SELECT untuple(hilbertDecode(2, hilbertEncode(n1, n2))) FROM hilbert_numbers;
