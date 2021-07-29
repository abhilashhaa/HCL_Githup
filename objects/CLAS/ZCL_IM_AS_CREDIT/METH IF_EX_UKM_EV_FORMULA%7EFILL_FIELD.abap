  METHOD if_ex_ukm_ev_formula~fill_field.

    DATA: ls_but000   TYPE but000,
          lv_ovdueday TYPE ukm_ovdue_days,
          dref1       TYPE REF TO data.

    FIELD-SYMBOLS: <fs_field1> TYPE any.

    CLEAR ls_but000.
* read partner from BP screen
    CALL FUNCTION 'BUP_BUPA_BUT000_GET'
      IMPORTING
        e_but000 = ls_but000.

    CASE i_fieldname.
      WHEN 'OVDUE_DAYS'.
        CREATE DATA dref1 TYPE ukm_ovdue_days.
        ASSIGN dref1->* TO <fs_field1>.
        SELECT SINGLE ovdue_days
                INTO lv_ovdueday
               FROM ukmbp_vector_it
          WHERE partner = ls_but000-partner.
        IF sy-subrc = 0.
          <fs_field1> = lv_ovdueday.
          GET REFERENCE OF <fs_field1> INTO rd_result.
        ENDIF.
      WHEN OTHERS.
    ENDCASE.

  ENDMETHOD.