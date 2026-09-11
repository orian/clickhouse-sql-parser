SET cast_keep_nullable = 0;
SELECT CAST(toNullable(toInt32(0)) AS Int32) as x, toTypeName(x);
