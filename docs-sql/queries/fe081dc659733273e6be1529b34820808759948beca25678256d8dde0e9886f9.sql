coalesce(
  nullIf(ResourceAttributes['service.version'], ''),
  nullIf(ResourceAttributes['container.image.tag'], '')
)
