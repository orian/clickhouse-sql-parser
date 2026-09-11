SELECT sum(doc.shipping.cost) AS shipping_cost FROM t1;
-- DB::Exception: Недопустимый тип Dynamic аргумента для агрегатной функции sum. (ILLEGAL_TYPE_OF_ARGUMENT)
