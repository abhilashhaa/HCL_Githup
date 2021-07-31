  method /IWBEP/IF_MGW_APPL_SRV_RUNTIME~CREATE_ENTITY.
*&----------------------------------------------------------------------------------------------*
*&  Include           /IWBEP/DPC_TEMP_CRT_ENTITY_BASE
*&* This class has been generated on 27.03.2020 11:22:59 in client 200
*&*
*&*       WARNING--> NEVER MODIFY THIS CLASS <--WARNING
*&*   If you want to change the DPC implementation, use the
*&*   generated methods inside the DPC provider subclass - ZCL_ZAZURE_MONITOR_DPC_EXT
*&-----------------------------------------------------------------------------------------------*

 DATA zmonitor_valuese_create_entity TYPE zcl_zazure_monitor_mpc=>ts_zmonitor_value.
 DATA monitorvalueset_create_entity TYPE zcl_zazure_monitor_mpc=>ts_monitorvalue.
 DATA zmonitor_logset_create_entity TYPE zcl_zazure_monitor_mpc=>ts_zmonitor_log.
 DATA lv_entityset_name TYPE string.

lv_entityset_name = io_tech_request_context->get_entity_set_name( ).

CASE lv_entityset_name.
*-------------------------------------------------------------------------*
*             EntitySet -  ZMONITOR_VALUESet
*-------------------------------------------------------------------------*
     WHEN 'ZMONITOR_VALUESet'.
*     Call the entity set generated method
    zmonitor_valuese_create_entity(
         EXPORTING iv_entity_name     = iv_entity_name
                   iv_entity_set_name = iv_entity_set_name
                   iv_source_name     = iv_source_name
                   io_data_provider   = io_data_provider
                   it_key_tab         = it_key_tab
                   it_navigation_path = it_navigation_path
                   io_tech_request_context = io_tech_request_context
       	 IMPORTING er_entity          = zmonitor_valuese_create_entity
    ).
*     Send specific entity data to the caller interfaces
    copy_data_to_ref(
      EXPORTING
        is_data = zmonitor_valuese_create_entity
      CHANGING
        cr_data = er_entity
   ).

*-------------------------------------------------------------------------*
*             EntitySet -  MonitorValueSet
*-------------------------------------------------------------------------*
     WHEN 'MonitorValueSet'.
*     Call the entity set generated method
    monitorvalueset_create_entity(
         EXPORTING iv_entity_name     = iv_entity_name
                   iv_entity_set_name = iv_entity_set_name
                   iv_source_name     = iv_source_name
                   io_data_provider   = io_data_provider
                   it_key_tab         = it_key_tab
                   it_navigation_path = it_navigation_path
                   io_tech_request_context = io_tech_request_context
       	 IMPORTING er_entity          = monitorvalueset_create_entity
    ).
*     Send specific entity data to the caller interfaces
    copy_data_to_ref(
      EXPORTING
        is_data = monitorvalueset_create_entity
      CHANGING
        cr_data = er_entity
   ).

*-------------------------------------------------------------------------*
*             EntitySet -  ZMONITOR_LOGSet
*-------------------------------------------------------------------------*
     WHEN 'ZMONITOR_LOGSet'.
*     Call the entity set generated method
    zmonitor_logset_create_entity(
         EXPORTING iv_entity_name     = iv_entity_name
                   iv_entity_set_name = iv_entity_set_name
                   iv_source_name     = iv_source_name
                   io_data_provider   = io_data_provider
                   it_key_tab         = it_key_tab
                   it_navigation_path = it_navigation_path
                   io_tech_request_context = io_tech_request_context
       	 IMPORTING er_entity          = zmonitor_logset_create_entity
    ).
*     Send specific entity data to the caller interfaces
    copy_data_to_ref(
      EXPORTING
        is_data = zmonitor_logset_create_entity
      CHANGING
        cr_data = er_entity
   ).

  when others.
    super->/iwbep/if_mgw_appl_srv_runtime~create_entity(
       EXPORTING
         iv_entity_name = iv_entity_name
         iv_entity_set_name = iv_entity_set_name
         iv_source_name = iv_source_name
         io_data_provider   = io_data_provider
         it_key_tab = it_key_tab
         it_navigation_path = it_navigation_path
      IMPORTING
        er_entity = er_entity
  ).
ENDCASE.
  endmethod.