SELECT sum(doc.shipping.cost) AS shipping_cost FROM t1;
-- DB::Exception: 집계 함수 sum의 인수로 Dynamic 유형은 사용할 수 없습니다. (ILLEGAL_TYPE_OF_ARGUMENT)
