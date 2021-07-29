  method DEFINE_ASSOCIATIONS.
*&---------------------------------------------------------------------*
*&           Generated code for the MODEL PROVIDER BASE CLASS         &*
*&                                                                     &*
*&  !!!NEVER MODIFY THIS CLASS. IN CASE YOU WANT TO CHANGE THE MODEL  &*
*&        DO THIS IN THE MODEL PROVIDER SUBCLASS!!!                   &*
*&                                                                     &*
*&---------------------------------------------------------------------*




data:
lo_annotation     type ref to /iwbep/if_mgw_odata_annotation,                   "#EC NEEDED
lo_entity_type    type ref to /iwbep/if_mgw_odata_entity_typ,                   "#EC NEEDED
lo_association    type ref to /iwbep/if_mgw_odata_assoc,                        "#EC NEEDED
lo_ref_constraint type ref to /iwbep/if_mgw_odata_ref_constr,                   "#EC NEEDED
lo_assoc_set      type ref to /iwbep/if_mgw_odata_assoc_set,                    "#EC NEEDED
lo_nav_property   type ref to /iwbep/if_mgw_odata_nav_prop.                     "#EC NEEDED

***********************************************************************************************************************************
*   ASSOCIATIONS
***********************************************************************************************************************************

 lo_association = model->create_association(
                            iv_association_name = 'zgw_emp' "#EC NOTEXT
                            iv_left_type        = 'ZAS_EMP' "#EC NOTEXT
                            iv_right_type       = 'ZAS_EMPnom' "#EC NOTEXT
                            iv_right_card       = 'M' "#EC NOTEXT
                            iv_left_card        = '1'  "#EC NOTEXT
                            iv_def_assoc_set    = abap_false ). "#EC NOTEXT
* Referential constraint for association - zgw_emp
lo_ref_constraint = lo_association->create_ref_constraint( ).
lo_ref_constraint->add_property( iv_principal_property = 'Id'   iv_dependent_property = 'Id' ). "#EC NOTEXT
lo_assoc_set = model->create_association_set( iv_association_set_name  = 'zgw_empSet'                         "#EC NOTEXT
                                              iv_left_entity_set_name  = 'ZAS_EMPSet'              "#EC NOTEXT
                                              iv_right_entity_set_name = 'ZAS_EMPnomSet'             "#EC NOTEXT
                                              iv_association_name      = 'zgw_emp' ).                                 "#EC NOTEXT


***********************************************************************************************************************************
*   NAVIGATION PROPERTIES
***********************************************************************************************************************************

* Navigation Properties for entity - ZAS_EMP
lo_entity_type = model->get_entity_type( iv_entity_name = 'ZAS_EMP' ). "#EC NOTEXT
lo_nav_property = lo_entity_type->create_navigation_property( iv_property_name  = 'zgw_empnav' "#EC NOTEXT
                                                              iv_abap_fieldname = 'ZGW_EMPNAV' "#EC NOTEXT
                                                              iv_association_name = 'zgw_emp' ). "#EC NOTEXT
  endmethod.