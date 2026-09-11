
    SELECT 
        numericIndexedVectorAllValueSum(groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-26')) AS sum_26,
        numericIndexedVectorAllValueSum(groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-27')) AS sum_27,
        numericIndexedVectorAllValueSum(groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-28')) AS sum_28,
        numericIndexedVectorAllValueSum(groupNumericIndexedVectorStateIf(uin, value, ds = '2023-12-29')) AS sum_29
    FROM uin_value_details
    