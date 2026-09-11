
        SELECT count() FROM map_subcol_pw_compact
        PREWHERE tags['k0'] != '' AND tags['k1'] != '' AND tags['k2'] != '' AND tags['k3'] != ''
            AND tags['k4'] != '' AND tags['k5'] != '' AND tags['k6'] != '' AND tags['k7'] != ''
            AND tags['k8'] != '' AND tags['k9'] != '' AND tags['k10'] != '' AND tags['k11'] != ''
            AND tags['k12'] != '' AND tags['k13'] != '' AND tags['k14'] != '' AND tags['k15'] != ''
        SETTINGS optimize_functions_to_subcolumns=1
    