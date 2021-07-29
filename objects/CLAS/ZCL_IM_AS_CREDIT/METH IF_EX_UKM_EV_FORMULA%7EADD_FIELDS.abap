  method IF_EX_UKM_EV_FORMULA~ADD_FIELDS.

 CONSTANTS: lc_empty  TYPE sfbefsym VALUE '',
            lc_bp_gen TYPE sfbefsym VALUE 'UKMBP_VECTOR_IT'.
 DATA: wa_operands TYPE sfbeoprnd.

  CASE i_key.
    WHEN lc_empty.
      CLEAR wa_operands.
      wa_operands-tech_name = 'UKMBP_VECTOR_IT'.
      wa_operands-descriptn = 'SAP Credit Management: Credit Segment'.
      APPEND wa_operands TO ct_operands.
    WHEN lc_bp_gen.
      CLEAR wa_operands.
      wa_operands-tech_name = 'OVDUE_DAYS'.
      wa_operands-descriptn = 'Days in Arrears'.
      wa_operands-type = 'DEC3'.
      APPEND wa_operands TO ct_operands.

*      CLEAR wa_operands.
*      wa_operands-tech_name = 'ZTABLE-ZFIELD2'.
*      wa_operands-descriptn = 'Field2 Description'.
*      wa_operands-type = 'INT4'.
*      APPEND wa_operands TO ct_operands.
    WHEN OTHERS.
  ENDCASE.

  endmethod.