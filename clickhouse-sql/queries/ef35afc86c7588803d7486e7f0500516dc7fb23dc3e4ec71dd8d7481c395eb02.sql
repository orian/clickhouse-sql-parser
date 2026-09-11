
    SELECT count() FROM index_test WHERE mortonDecode(2, z).1 >= 20000 AND mortonDecode(2, z).1 <= 20100 AND mortonDecode(2, z).2 >= 10000 AND mortonDecode(2, z).2 <= 10100
    