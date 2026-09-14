SELECT sum(doc.shipping.cost) AS shipping_cost FROM t1;
-- DB::Exception: 集計関数 sum の引数の型 Dynamic は使用できません。 (ILLEGAL_TYPE_OF_ARGUMENT)
