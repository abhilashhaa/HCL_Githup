  method ZMONITOR_VALUESE_GET_ENTITYSET.
**TRY.
*CALL METHOD SUPER->MONITORVALUESET_GET_ENTITYSET
*  EXPORTING
*    IV_ENTITY_NAME           =
*    IV_ENTITY_SET_NAME       =
*    IV_SOURCE_NAME           =
*    IT_FILTER_SELECT_OPTIONS =
*    IS_PAGING                =
*    IT_KEY_TAB               =
*    IT_NAVIGATION_PATH       =
*    IT_ORDER                 =
*    IV_FILTER_STRING         =
*    IV_SEARCH_STRING         =
**    io_tech_request_context  =
**  IMPORTING
**    et_entityset             =
**    es_response_context      =
*    .
**  CATCH /iwbep/cx_mgw_busi_exception.
**  CATCH /iwbep/cx_mgw_tech_exception.
**ENDTRY.
    DATA :  er_entity TYPE zca_monitor_value,
          lv_sys_value TYPE csm_sys_cs,
          lt_contexts TYPE standard table of ALCONSEG, "csm_context_cache,
          ls_contexts TYPE csmcontx,
          ls_cpu_all TYPE cpu_all,
          lt_cpu_all TYPE TABLE OF cpu_all.

      DATA: session_list TYPE ssi_session_list,
        server_info  TYPE REF TO cl_server_info.


  FIELD-SYMBOLS: <fs_alconseg> TYPE alconseg.

  CLEAR : er_entity, ls_cpu_all, lt_cpu_all[],
          lv_sys_value, lt_contexts[].

* Get all servers
  CALL FUNCTION 'SCSM_CONTEXTS_FOR_SYSTEM'
   EXPORTING
     alsystem                       = syst-sysid
*     CSMSYSID                       =
*     ROUTES_BY_COMMTYPE             = AL_CSM_CTYPE_HIGHPRIO
*     USE_CACHE                      = 'X'
*     FILL_CACHE                     = 'X'
     do_not_ping_system             = 'X'
     mte_class_pattern              = 'R3ApplicationServer'
   IMPORTING
     system_ctx_and_segms           = lv_sys_value
   TABLES
*     UNREACHABLE_CONTEXTS           =
     all_contexts                   = lt_contexts
   EXCEPTIONS
     no_contexts_found              = 1
     no_contexts_reachable          = 2
     system_unknown                 = 3
     system_not_reachable           = 4
     only_local_sysid_allowed       = 5
     OTHERS                         = 6
            .
  IF sy-subrc <> 0.
* Implement suitable error handling here
  ENDIF.

 data : ls_dest type RFCDEST.

  LOOP AT lt_contexts assigning <fs_alconseg>.

   ls_dest = <fs_alconseg>-context.
CALL FUNCTION 'GET_CPU_ALL'
 EXPORTING
   local_remote                         = 'LOCAL'
   logical_destination                  = ls_dest "<fs_alconseg>-context
* IMPORTING
*   F_CPU_ALL_READ                       =
*   ACTIVEFLAG                           =
*   INTERVAL                             =
*   OP_SYSTEM                            =
*   DETAILSCOLL                          =
*   DETAILSREQI                          =
*   DETAILSMODE                          =
*   LASTCOLLWRT                          =
*   LASTCOLLINT                          =
*   NORMCOLLINT                          =
  TABLES
    tf_cpu_all                           = lt_cpu_all
* EXCEPTIONS
*   INTERNAL_ERROR_ADRESS_FAILED         = 1
*   INTERNAL_ERROR_DIFFERENT_FIELD       = 2
*   INTERNAL_ERROR_NO_NEW_LINE           = 3
*   COLLECTOR_NOT_RUNNING                = 4
*   SHARED_MEMORY_NOT_AVAILABLE          = 5
*   COLLECTOR_BUSY                       = 6
*   VERSION_CONFLICT                     = 7
*   NO_NETWORK_COLLECTOR_RUNNING         = 8
*   SYSTEM_FAILURE                       = 9
*   COMMUNICATION_FAILURE                = 10
*   OTHERS                               = 11
          .
IF sy-subrc <> 0.
* Implement suitable error handling here
ENDIF.

  READ TABLE lt_cpu_all INTO ls_cpu_all INDEX 1.
* Populate CPU usage to output

   CLEAR er_entity.
    er_entity-sysid = syst-sysid.  "<fs_alconseg>-sysid.
    er_entity-servername = ls_dest. "<fs_alconseg>-context.
    er_entity-idle_total = ls_cpu_all-idle_total.

* Find no of session for each server
 data : ls_servername type SSI_SERVERNAME.
TRY.
      ls_servername = <fs_alconseg>-context.
      CREATE OBJECT server_info exporting SERVER_NAME = ls_servername .
      session_list = server_info->get_session_list( with_application_info = '1' ).
     " delete adjacent duplicates from session_list comparing user_name.
      DESCRIBE TABLE session_list LINES er_entity-no_session.
endtry.
*    er_entity-NO_SESSION = 3.
* Populate logon group (SMLG)
   select single classname from RZLLITAB into er_entity-classname
          where APPLSERVER = ls_servername
          and   grouptype = ''.
   if sy-subrc ne 0.
      clear er_entity-classname.
   endif.
*    er_entity-CLASSNAME =  'TEST'.
   APPEND er_entity TO et_entityset.

   ENDLOOP.

  endmethod.