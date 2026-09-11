-- 创建以 SpanName 为 String、ResourceAttributes 为 Map 的表
DROP TABLE IF EXISTS tbl;
CREATE TABLE tbl (
    `Timestamp` DateTime64(9) CODEC (Delta(8), ZSTD(1)),
    `TraceId` String CODEC (ZSTD(1)),
    `ServiceName` LowCardinality(String) CODEC (ZSTD(1)),
    `Duration` UInt8 CODEC (ZSTD(1)), -- Int64
    `SpanName` LowCardinality(String) CODEC (ZSTD(1)),
    `ResourceAttributes` Map(LowCardinality(String), String) CODEC (ZSTD(1))
)
ENGINE = MergeTree
PARTITION BY toDate(Timestamp)
ORDER BY (ServiceName, SpanName, toUnixTimestamp(Timestamp), TraceId);

-- 创建 UDF，为 ResourceAttributes 生成随机 Map 数据
DROP FUNCTION IF EXISTS genmap;
CREATE FUNCTION genmap AS (n) -> arrayMap (x-> (x::String, (x*rand32())::String), range(1, n));

-- 验证 genmap 是否按预期工作
SELECT genmap(10)::Map(String, String);

-- 插入 100 万行数据
INSERT INTO tbl
SELECT
    now() - randUniform(1, 1000000.) as Timestamp,
    randomPrintableASCII(2) as TraceId,
    randomPrintableASCII(2) as ServiceName,
    rand32() as Duration,
    randomPrintableASCII(2) as SpanName,
    genmap(rand64()%500)::Map(String, String) as ResourceAttributes
FROM numbers(1_000_000);

-- 按 SpanName 查询速度更快
-- [冷查询] 0 rows in set. Elapsed: 0.642 sec. Processed 1.00 million rows, 7.21 MB (1.56 million rows/s., 11.22 MB/s.)
-- [热查询] 0 rows in set. Elapsed: 0.164 sec. Processed 1.00 million rows, 7.21 MB (6.10 million rows/s., 43.99 MB/s.)
SELECT
    COUNT(*),
    avg(Duration/1E6) as average,
    quantile(0.95)(Duration/1E6) as p95,
    quantile(0.99)(Duration/1E6) as p99,
    SpanName
FROM tbl
GROUP BY SpanName ORDER BY 1 DESC LIMIT 50 FORMAT Null;

-- 按 ResourceAttributes 查询速度较慢
-- [冷查询] 0 rows in set. Elapsed: 6.432 sec. Processed 1.00 million rows, 5.65 GB (155.46 thousand rows/s., 879.07 MB/s.)
-- [热查询] 0 rows in set. Elapsed: 5.935 sec. Processed 1.00 million rows, 5.65 GB (168.50 thousand rows/s., 952.81 MB/s.)
SELECT
    COUNT(*),
    avg(Duration/1E6) as average,
    quantile(0.95)(Duration/1E6) as p95,
    quantile(0.99)(Duration/1E6) as p99,
    ResourceAttributes['1'] as hostname
FROM tbl
GROUP BY hostname ORDER BY 1 DESC LIMIT 50 FORMAT Null;
