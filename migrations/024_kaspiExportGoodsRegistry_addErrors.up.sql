ALTER TABLE "kaspi_export_goods_registry"
  ADD COLUMN IF NOT EXISTS errors TEXT[];

  