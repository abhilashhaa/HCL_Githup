class ZCL_ZAS_EMP_MPC definition
  public
  inheriting from /IWBEP/CL_MGW_PUSH_ABS_MODEL
  create public .

public section.

  types:
         TS_ZAS_EMP type ZAS_EMP .
  types:
    TT_ZAS_EMP type standard table of TS_ZAS_EMP .
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
         TS_ZAS_EMPNOM type ZAS_EMPNOM .
  types:
    TT_ZAS_EMPNOM type standard table of TS_ZAS_EMPNOM .

  constants GC_ZAS_EMP type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'ZAS_EMP' ##NO_TEXT.
  constants GC_ZAS_EMPNOM type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'ZAS_EMPnom' ##NO_TEXT.

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