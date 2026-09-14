
        SELECT groupArray(id)
        FROM dict_src_const_path
        WHERE dictGetUInt32('dict_const_path', 'value', id) = 1
        FORMAT NULL;
    