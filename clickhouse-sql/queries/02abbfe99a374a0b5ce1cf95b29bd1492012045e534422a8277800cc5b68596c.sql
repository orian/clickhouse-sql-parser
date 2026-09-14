
    WITH
        (SELECT groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-22') FROM uin_value_details) AS vec_1,
        (SELECT groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-23') FROM uin_value_details) AS vec_2
    SELECT 
        numericIndexedVectorAllValueSum(numericIndexedVectorPointwiseMultiply(vec_1, vec_2)) AS sum_multiply,
        numericIndexedVectorAllValueSum(numericIndexedVectorPointwiseDivide(vec_1, vec_2)) AS sum_divide
    