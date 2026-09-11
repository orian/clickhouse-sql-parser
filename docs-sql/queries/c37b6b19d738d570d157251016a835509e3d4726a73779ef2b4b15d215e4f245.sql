SELECT round(arraySum(x -> x * x, randomHadamardTransform([1, 2, 3, 4]::Array(Float32))) - 30, 4)
