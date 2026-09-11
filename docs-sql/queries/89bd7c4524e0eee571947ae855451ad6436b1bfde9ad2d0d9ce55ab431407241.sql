SET input_format_json_infer_array_of_dynamic_from_array_of_different_types=0;
DESC format(JSONEachRow, $$
                              {"tuple" : [1, null, null]}
                              {"tuple" : [null, "Hello, World!", []]}
                              {"tuple" : [null, null, [1, 2, 3]]}
                         $$)
