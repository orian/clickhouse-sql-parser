
       CREATE TABLE hits_buffer AS hits_10m_single 
       ENGINE = Buffer(default, hits_wide, 1, 0, 0, 10000, 10000, 0, 0)
    