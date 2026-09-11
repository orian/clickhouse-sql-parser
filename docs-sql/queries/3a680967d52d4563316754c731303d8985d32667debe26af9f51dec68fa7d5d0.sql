{% for type in ['UInt8', 'UInt16', 'UInt32'] -%}
SELECT toTypeName(0::{{ type }});
{% endfor -%}
