
        SELECT tags['k0'], tags['k1'], tags['k2'], tags['k3'],
               tags['k4'], tags['k5'], tags['k6'], tags['k7'],
               tags['k8'], tags['k9'], tags['k10'], tags['k11'],
               tags['k12'], tags['k13'], tags['k14'], tags['k15']
        FROM map_subcol_pw_compact
        FORMAT Null
        SETTINGS optimize_functions_to_subcolumns=1
    