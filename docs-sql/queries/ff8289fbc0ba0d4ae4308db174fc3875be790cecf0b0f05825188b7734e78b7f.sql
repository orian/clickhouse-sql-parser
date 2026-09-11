SET input_format_try_infer_datetimes = 0, input_format_try_infer_dates = 0
DESC format(JSONEachRow, $$
                                {"date" : "2021-01-01"}
                                {"date" : "2022-01-01"}
                         $$)
