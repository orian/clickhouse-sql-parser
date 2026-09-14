SELECT JSONExtract('{"a" : [1, 2, 3]}', 'a', 'Dynamic') AS dynamic, dynamicType(dynamic) AS dynamic_type;
