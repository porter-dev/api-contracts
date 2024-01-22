// @generated by protoc-gen-es v1.4.2
// @generated from file porter/v1/compliance.proto (package porter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum porter.v1.EnumComplianceVendor
 */
export const EnumComplianceVendor = proto3.makeEnum(
  "porter.v1.EnumComplianceVendor",
  [
    {no: 0, name: "ENUM_COMPLIANCE_VENDOR_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_COMPLIANCE_VENDOR_VANTA", localName: "VANTA"},
  ],
);

/**
 * @generated from enum porter.v1.EnumComplianceCheckStatus
 */
export const EnumComplianceCheckStatus = proto3.makeEnum(
  "porter.v1.EnumComplianceCheckStatus",
  [
    {no: 0, name: "ENUM_COMPLIANCE_CHECK_STATUS_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ENUM_COMPLIANCE_CHECK_STATUS_PASSED", localName: "PASSED"},
    {no: 2, name: "ENUM_COMPLIANCE_CHECK_STATUS_FAILED", localName: "FAILED"},
    {no: 3, name: "ENUM_COMPLIANCE_CHECK_STATUS_NOT_APPLICABLE", localName: "NOT_APPLICABLE"},
  ],
);

/**
 * ContractComplianceCheckGroup representes a porter internal concept that represents some infrastructure level configuration
 * that is expected to be in place for a contract to be considered compliant.
 *
 * @generated from message porter.v1.ContractComplianceCheckGroup
 */
export const ContractComplianceCheckGroup = proto3.makeMessageType(
  "porter.v1.ContractComplianceCheckGroup",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(EnumComplianceCheckStatus) },
    { no: 3, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * VendorComplianceCheckGroup represents a vendor provided compliance test, which porter deems to be passing based on the status of the corresponding check group.
 *
 * @generated from message porter.v1.VendorComplianceCheck
 */
export const VendorComplianceCheck = proto3.makeMessageType(
  "porter.v1.VendorComplianceCheck",
  () => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "check_group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(EnumComplianceCheckStatus) },
    { no: 4, name: "reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

