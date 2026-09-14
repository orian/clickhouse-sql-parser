
    WITH
        (SELECT groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-27') FROM uin_value_details) AS v1,
        (SELECT groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-29') FROM uin_value_details) AS v2
    SELECT 
        numericIndexedVectorAllValueSum(numericIndexedVectorPointwiseMultiply(v1, v2)) AS mul_27_29,
        numericIndexedVectorAllValueSum(numericIndexedVectorPointwiseDivide(v1, v2)) AS div_27_29
    