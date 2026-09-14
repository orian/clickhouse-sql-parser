-- Create table with Memory engine
CREATE TABLE numbers (
    a Array(Int32)
) ENGINE = Memory;

-- Insert sample data
INSERT INTO numbers VALUES
    ([1,2,4]),
    ([1,5,2,8,-1,0]),
    ([1,5,7,5,8,2]);

SELECT groupArrayIntersect(a) AS intersection FROM numbers;
