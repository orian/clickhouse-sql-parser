SELECT sum(doc.shipping.cost) AS shipping_cost FROM t1;
-- DB::Exception: 聚合函数 sum 的参数类型 Dynamic 不合法。(ILLEGAL_TYPE_OF_ARGUMENT)
