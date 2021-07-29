private section.

  constants GC_INCL_NAME type STRING value 'ZCL_ZGWSO_MPC=================CP' ##NO_TEXT.

  methods DEFINE_SALESORDER
    raising
      /IWBEP/CX_MGW_MED_EXCEPTION .
  methods DEFINE_SALESORDERITEM
    raising
      /IWBEP/CX_MGW_MED_EXCEPTION .