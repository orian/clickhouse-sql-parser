
        SELECT groupArray(id)
        FROM dict_src_const_path
        WHERE dictGetString('dict_const_path', 'value_str', id) = 'value_1'
        FORMAT NULL;
    