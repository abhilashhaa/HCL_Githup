class ZCL_ZGWSO_MPC definition
  public
  inheriting from /IWBEP/CL_MGW_PUSH_ABS_MODEL
  create public .

public section.

  types:
      begin of TS_SALESORDER,
     SO_ID type SNWD_SO_ID,
     NOTE type SNWD_DESC,
     BUYER_ID type SNWD_PARTNER_ID,
     BUYER_NAME type SNWD_COMPANY_NAME,
     CURRENCY_CODE type SNWD_CURR_CODE,
     GROSS_AMOUNT type SNWD_TTL_GROSS_AMOUNT,
     NET_AMOUNT type SNWD_TTL_NET_AMOUNT,
     TAX_AMOUNT type SNWD_TTL_TAX_AMOUNT,
  end of TS_SALESORDER .
  types:
    TT_SALESORDER type standard table of TS_SALESORDER .
  types:
       begin of ts_text_element,
      artifact_name  type c length 40,       " technical name
      artifact_type  type c length 4,
      parent_artifact_name type c length 40, " technical name
      parent_artifact_type type c length 4,
      text_symbol    type textpoolky,
   end of ts_text_element .
  types:
             tt_text_elements type standard table of ts_text_element with key text_symbol .
  types:
      begin of TS_SALESORDERITEM,
     SO_ID type SNWD_SO_ID,
     SO_ITEM_POS type SNWD_SO_ITEM_POS,
     PRODUCT_ID type SNWD_PRODUCT_ID,
     NOTE type SNWD_DESC,
     CURRENCY_CODE type SNWD_CURR_CODE,
     GROSS_AMOUNT type SNWD_TTL_GROSS_AMOUNT,
     NET_AMOUNT type SNWD_TTL_NET_AMOUNT,
     TAX_AMOUNT type SNWD_TTL_TAX_AMOUNT,
     QUANTITY type SNWD_QUANTITY,
     QUANTITY_UNIT type SNWD_QUANTITY_UNIT,
  end of TS_SALESORDERITEM .
  types:
    TT_SALESORDERITEM type standard table of TS_SALESORDERITEM .

  constants GC_SALESORDER type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'SalesOrder' ##NO_TEXT.
  constants GC_SALESORDERITEM type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'SalesOrderItem' ##NO_TEXT.

  methods LOAD_TEXT_ELEMENTS
  final
    returning
      value(RT_TEXT_ELEMENTS) type TT_TEXT_ELEMENTS
    raising
      /IWBEP/CX_MGW_MED_EXCEPTION .

  methods DEFINE
    redefinition .
  methods GET_LAST_MODIFIED
    redefinition .