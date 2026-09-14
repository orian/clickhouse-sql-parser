  -- إنشاء جدول يحتوي على جميع أنواع البيانات المطلوبة
  CREATE TABLE all_types_test (
      `id` UInt32,
      
      -- نوع String
      `String_value` String,
      
      -- أنواع الأعداد الصحيحة غير الموقَّعة
      `UInt8_value` UInt8,
      `UInt16_value` UInt16,
      `UInt32_value` UInt32,
      `UInt64_value` UInt64,
      
      -- أنواع الأعداد الصحيحة الموقَّعة
      `Int8_value` Int8,
      `Int16_value` Int16,
      `Int32_value` Int32,
      `Int64_value` Int64,
      
      -- أنواع الأعداد ذات الفاصلة العائمة
      `Float32_value` Float32,
      `Float64_value` Float64,
      
      -- أنواع التاريخ/الوقت
      `Date_value` Date,
      `DateTime_value` DateTime,
      
      -- أنواع الشبكة
      `IPv4_value` IPv4,
      `IPv6_value` IPv6,
      
      -- نوع UUID
      `UUID_value` UUID
  ) ENGINE = MergeTree() 
  ORDER BY id;
