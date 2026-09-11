SET input_format_json_read_numbers_as_strings = 1;
DESC format(JSONEachRow, $$
                                {"value" : 1055}
                                {"value" : "unknown"}
                         $$)
