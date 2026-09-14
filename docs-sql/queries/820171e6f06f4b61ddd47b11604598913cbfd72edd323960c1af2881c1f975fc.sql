SELECT notHas('{"a" : 1, "b" : {"c" : 2}}'::JSON, 'b.c')
