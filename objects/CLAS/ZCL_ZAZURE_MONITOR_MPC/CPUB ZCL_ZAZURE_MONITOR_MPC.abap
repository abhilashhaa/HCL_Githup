class ZCL_ZAZURE_MONITOR_MPC definition
  public
  inheriting from /IWBEP/CL_MGW_PUSH_ABS_MODEL
  create public .

public section.

  types:
     TS_ZMONITOR_VALUE type ZCA_MONITOR_VALUE .
  types:
TT_ZMONITOR_VALUE type standard table of TS_ZMONITOR_VALUE .
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
     TS_MONITORVALUE type ZCA_MONITOR_VALUE .
  types:
TT_MONITORVALUE type standard table of TS_MONITORVALUE .
  types:
     TS_ZMONITOR_LOG type ZCA_AZURE_LOG .
  types:
TT_ZMONITOR_LOG type standard table of TS_ZMONITOR_LOG .

  constants GC_MONITORVALUE type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'MonitorValue' ##NO_TEXT.
  constants GC_ZMONITOR_LOG type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'ZMONITOR_LOG' ##NO_TEXT.
  constants GC_ZMONITOR_VALUE type /IWBEP/IF_MGW_MED_ODATA_TYPES=>TY_E_MED_ENTITY_NAME value 'ZMONITOR_VALUE' ##NO_TEXT.

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