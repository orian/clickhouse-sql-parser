CREATE OR REPLACE MASKING POLICY mask_pii_data ON orders
    UPDATE name = '[REDACTED]'
    TO masked_data_viewer;
