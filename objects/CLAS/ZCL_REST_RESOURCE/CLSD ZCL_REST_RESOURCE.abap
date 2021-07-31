class-pool .
*"* class pool for class ZCL_REST_RESOURCE

*"* local type definitions
include ZCL_REST_RESOURCE=============ccdef.

*"* class ZCL_REST_RESOURCE definition
*"* public declarations
  include ZCL_REST_RESOURCE=============cu.
*"* protected declarations
  include ZCL_REST_RESOURCE=============co.
*"* private declarations
  include ZCL_REST_RESOURCE=============ci.
endclass. "ZCL_REST_RESOURCE definition

*"* macro definitions
include ZCL_REST_RESOURCE=============ccmac.
*"* local class implementation
include ZCL_REST_RESOURCE=============ccimp.

class ZCL_REST_RESOURCE implementation.
*"* method's implementations
  include methods.
endclass. "ZCL_REST_RESOURCE implementation
