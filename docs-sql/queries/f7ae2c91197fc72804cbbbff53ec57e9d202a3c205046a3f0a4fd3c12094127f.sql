CREATE TABLE default.t_arr
(
    `arr` Array(Array(Array(UInt32)))
)
ENGINE = MergeTree
ORDER BY tuple()
