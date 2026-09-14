    CREATE FILE FORMAT my_parquet_format TYPE = parquet;

    -- أنشئ المرحلة الخارجية التي تحدد حاوية S3 التي سيتم النسخ إليها
    CREATE OR REPLACE STAGE external_stage
    URL='s3://mybucket/mydataset'
    CREDENTIALS=(AWS_KEY_ID='<key>' AWS_SECRET_KEY='<secret>')
    FILE_FORMAT = my_parquet_format;

    -- طبّق البادئة "mydataset" على جميع الملفات وحدد حجمًا أقصى للملف قدره 150mb
    -- المعلَمة `header=true` مطلوبة للحصول على أسماء الأعمدة
    COPY INTO @external_stage/mydataset from mydataset max_file_size=157286400 header=true;
