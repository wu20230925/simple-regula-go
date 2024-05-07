// Package model provides primitives to interact with the openapi HTTP API.
package model

import (
	"encoding/json"
	"time"
)

// Defines values for AuthenticityResultType.
const (
	AuthenticityResultTypeAXIALPROTECTION       AuthenticityResultType = 8
	AuthenticityResultTypeBARCODEFORMATCHECK    AuthenticityResultType = 65536
	AuthenticityResultTypeEXTENDEDMRZCHECK      AuthenticityResultType = 8.388608e+06
	AuthenticityResultTypeEXTENDEDOCRCHECK      AuthenticityResultType = 4.194304e+06
	AuthenticityResultTypeFINGERPRINTCOMPARISON AuthenticityResultType = 1.048576e+06
	AuthenticityResultTypeHOLOGRAMDETECTION     AuthenticityResultType = 524288
	AuthenticityResultTypeHOLOGRAMS             AuthenticityResultType = 4096
	AuthenticityResultTypeIMAGEPATTERN          AuthenticityResultType = 4
	AuthenticityResultTypeIPI                   AuthenticityResultType = 128
	AuthenticityResultTypeIRB900                AuthenticityResultType = 2
	AuthenticityResultTypeIRVISIBILITY          AuthenticityResultType = 32
	AuthenticityResultTypeKINEGRAM              AuthenticityResultType = 131072
	AuthenticityResultTypeLETTERSCREEN          AuthenticityResultType = 262144
	AuthenticityResultTypeLIVENESS              AuthenticityResultType = 2.097152e+06
	AuthenticityResultTypeOCRSECURITYTEXT       AuthenticityResultType = 64
	AuthenticityResultTypeOVI                   AuthenticityResultType = 1024
	AuthenticityResultTypePHOTOAREA             AuthenticityResultType = 8192
	AuthenticityResultTypePHOTOEMBEDTYPE        AuthenticityResultType = 512
	AuthenticityResultTypePORTRAITCOMPARISON    AuthenticityResultType = 32768
	AuthenticityResultTypeUVFIBERS              AuthenticityResultType = 16
	AuthenticityResultTypeUVLUMINESCENCE        AuthenticityResultType = 1
)

// Defines values for CheckDiagnose.
const (
	CheckDiagnoseBACKGROUNDCOMPARISONERROR               CheckDiagnose = 25
	CheckDiagnoseBADAREAINAXIAL                          CheckDiagnose = 60
	CheckDiagnoseBARCODEDATAFORMATERROR                  CheckDiagnose = 141
	CheckDiagnoseBARCODESIZEPARAMSERROR                  CheckDiagnose = 142
	CheckDiagnoseBARCODEWASREADWITHERRORS                CheckDiagnose = 140
	CheckDiagnoseCHDICAOIDBBASE32ERROR                   CheckDiagnose = 243
	CheckDiagnoseCHDICAOIDBCERTIFICATEMUSTNOTBEPRESENT   CheckDiagnose = 248
	CheckDiagnoseCHDICAOIDBMESSAGEZONEEMPTY              CheckDiagnose = 245
	CheckDiagnoseCHDICAOIDBSIGNATUREMUSTBEPRESENT        CheckDiagnose = 246
	CheckDiagnoseCHDICAOIDBSIGNATUREMUSTNOTBEPRESENT     CheckDiagnose = 247
	CheckDiagnoseCHDICAOIDBZIPPEDERROR                   CheckDiagnose = 244
	CheckDiagnoseCONTACTCHIPTYPEMISMATCH                 CheckDiagnose = 29
	CheckDiagnoseDOCLIVENESSELECTRONICDEVICEDETECTED     CheckDiagnose = 240
	CheckDiagnoseDOCLIVENESSINVALIDBARCODEBACKGROUND     CheckDiagnose = 241
	CheckDiagnoseDOCUMENTISCANCELLING                    CheckDiagnose = 122
	CheckDiagnoseELEMENTSHOULDBECOLORED                  CheckDiagnose = 42
	CheckDiagnoseELEMENTSHOULDBEGRAYSCALE                CheckDiagnose = 43
	CheckDiagnoseEXCEPTIONINMODULE                       CheckDiagnose = 4
	CheckDiagnoseFALSEIPPARAMETERS                       CheckDiagnose = 65
	CheckDiagnoseFALSELUMINESCENCEERROR                  CheckDiagnose = 21
	CheckDiagnoseFALSELUMINESCENCEINBLANK                CheckDiagnose = 55
	CheckDiagnoseFALSELUMINESCENCEINMRZ                  CheckDiagnose = 51
	CheckDiagnoseFIBERSNOTFOUND                          CheckDiagnose = 30
	CheckDiagnoseFIELDPOSCORRECTORFACEABSENCECHECKERROR  CheckDiagnose = 85
	CheckDiagnoseFIELDPOSCORRECTORFACEPRESENCECHECKERROR CheckDiagnose = 84
	CheckDiagnoseFIELDPOSCORRECTORGLARESINPHOTOAREA      CheckDiagnose = 81
	CheckDiagnoseFIELDPOSCORRECTORHIGHLIGHTIR            CheckDiagnose = 80
	CheckDiagnoseFIELDPOSCORRECTORLANDMARKSCHECKERROR    CheckDiagnose = 83
	CheckDiagnoseFIELDPOSCORRECTORPHOTOREPLACED          CheckDiagnose = 82
	CheckDiagnoseFIELDSCOMPARISONLOGICERROR              CheckDiagnose = 14
	CheckDiagnoseFINGERPRINTSCOMPARISONMISMATCH          CheckDiagnose = 170
	CheckDiagnoseFIXEDPATTERNERROR                       CheckDiagnose = 22
	CheckDiagnoseGLARESINBARCODEAREA                     CheckDiagnose = 144
	CheckDiagnoseHOLOGRAMELEMENTABSENT                   CheckDiagnose = 100
	CheckDiagnoseHOLOGRAMELEMENTPRESENT                  CheckDiagnose = 102
	CheckDiagnoseHOLOGRAMFRAMESISABSENT                  CheckDiagnose = 103
	CheckDiagnoseHOLOGRAMHOLOFIELDISABSENT               CheckDiagnose = 104
	CheckDiagnoseHOLOGRAMSIDETOPIMAGESABSENT             CheckDiagnose = 101
	CheckDiagnoseHOLOPHOTOALGORITHMSSTEPSERROR           CheckDiagnose = 184
	CheckDiagnoseHOLOPHOTODOCUMENTOUTSIDEFRAME           CheckDiagnose = 187
	CheckDiagnoseHOLOPHOTOFACECOMPARISONFAILED           CheckDiagnose = 181
	CheckDiagnoseHOLOPHOTOFACENOTDETECTED                CheckDiagnose = 180
	CheckDiagnoseHOLOPHOTOFINISHEDBYTIMEOUT              CheckDiagnose = 186
	CheckDiagnoseHOLOPHOTOGLAREINCENTERABSENT            CheckDiagnose = 182
	CheckDiagnoseHOLOPHOTOHOLOAREASNOTLOADED             CheckDiagnose = 185
	CheckDiagnoseHOLOPHOTOHOLOELEMENTSHAPEERROR          CheckDiagnose = 183
	CheckDiagnoseINCORRECTBACKGROUNDLIGHT                CheckDiagnose = 24
	CheckDiagnoseINCORRECTTEXTCOLOR                      CheckDiagnose = 26
	CheckDiagnoseINTERNALERROR                           CheckDiagnose = 3
	CheckDiagnoseINVALIDCHECKSUM                         CheckDiagnose = 10
	CheckDiagnoseINVALIDFIELDFORMAT                      CheckDiagnose = 15
	CheckDiagnoseINVALIDINPUTDATA                        CheckDiagnose = 2
	CheckDiagnoseINVISIBLEELEMENTPRESENT                 CheckDiagnose = 40
	CheckDiagnoseLASINKINVALIDLINESFREQUENCY             CheckDiagnose = 230
	CheckDiagnoseLASTDIAGNOSEVALUE                       CheckDiagnose = 250
	CheckDiagnoseLIVENESSDEPTHCHECKFAILED                CheckDiagnose = 190
	CheckDiagnoseLOGICERROR                              CheckDiagnose = 12
	CheckDiagnoseLOWCONTRASTINIRLIGHT                    CheckDiagnose = 23
	CheckDiagnoseMOBILEIMAGESUNSUITABLELIGHTCONDITIONS   CheckDiagnose = 160
	CheckDiagnoseMOBILEIMAGESWHITEUVNODIFFERENCE         CheckDiagnose = 161
	CheckDiagnoseMRZQUALITYWRONGBACKGROUND               CheckDiagnose = 201
	CheckDiagnoseMRZQUALITYWRONGFONTTYPE                 CheckDiagnose = 205
	CheckDiagnoseMRZQUALITYWRONGLINEPOSITION             CheckDiagnose = 204
	CheckDiagnoseMRZQUALITYWRONGMRZHEIGHT                CheckDiagnose = 203
	CheckDiagnoseMRZQUALITYWRONGMRZWIDTH                 CheckDiagnose = 202
	CheckDiagnoseMRZQUALITYWRONGSYMBOLPOSITION           CheckDiagnose = 200
	CheckDiagnoseNECESSARYIMAGENOTFOUND                  CheckDiagnose = 7
	CheckDiagnoseNOTALLBARCODESREAD                      CheckDiagnose = 143
	CheckDiagnoseOCRQUALITYINVALIDBACKGROUND             CheckDiagnose = 222
	CheckDiagnoseOCRQUALITYINVALIDFONT                   CheckDiagnose = 221
	CheckDiagnoseOCRQUALITYTEXTPOSITION                  CheckDiagnose = 220
	CheckDiagnoseOVIBADCOLORFRONT                        CheckDiagnose = 93
	CheckDiagnoseOVIBADCOLORPERCENT                      CheckDiagnose = 96
	CheckDiagnoseOVIBADCOLORSIDE                         CheckDiagnose = 94
	CheckDiagnoseOVICOLORINVARIABLE                      CheckDiagnose = 92
	CheckDiagnoseOVIINSUFFICIENTAREA                     CheckDiagnose = 91
	CheckDiagnoseOVIIRINVISIBLE                          CheckDiagnose = 90
	CheckDiagnoseOVIWIDECOLORSPREAD                      CheckDiagnose = 95
	CheckDiagnosePASS                                    CheckDiagnose = 1
	CheckDiagnosePHOTOCORNERSISWRONG                     CheckDiagnose = 121
	CheckDiagnosePHOTOFALSELUMINESCENCE                  CheckDiagnose = 27
	CheckDiagnosePHOTOISNOTRECTANGLE                     CheckDiagnose = 120
	CheckDiagnosePHOTOPATTERNDIFFERENTCOLORS             CheckDiagnose = 112
	CheckDiagnosePHOTOPATTERNDIFFERENTLINESTHICKNESS     CheckDiagnose = 119
	CheckDiagnosePHOTOPATTERNINTERRUPTED                 CheckDiagnose = 110
	CheckDiagnosePHOTOPATTERNINVALIDCOLOR                CheckDiagnose = 116
	CheckDiagnosePHOTOPATTERNIRVISIBLE                   CheckDiagnose = 113
	CheckDiagnosePHOTOPATTERNNOTINTERSECT                CheckDiagnose = 114
	CheckDiagnosePHOTOPATTERNPATTERNNOTFOUND             CheckDiagnose = 118
	CheckDiagnosePHOTOPATTERNSHIFTED                     CheckDiagnose = 111
	CheckDiagnosePHOTOPATTERNSHIFTEDVERT                 CheckDiagnose = 117
	CheckDiagnosePHOTOSIDESNOTFOUND                      CheckDiagnose = 8
	CheckDiagnosePHOTOSIZEISWRONG                        CheckDiagnose = 115
	CheckDiagnosePHOTOWHITEIRDONTMATCH                   CheckDiagnose = 44
	CheckDiagnosePORTRAITCOMPARISONNOLIVEPHOTO           CheckDiagnose = 154
	CheckDiagnosePORTRAITCOMPARISONNOPORTRAITDETECTED    CheckDiagnose = 156
	CheckDiagnosePORTRAITCOMPARISONNOSERVICELICENSE      CheckDiagnose = 155
	CheckDiagnosePORTRAITCOMPARISONNOSERVICEREPLY        CheckDiagnose = 151
	CheckDiagnosePORTRAITCOMPARISONNOTENOUGHIMAGES       CheckDiagnose = 153
	CheckDiagnosePORTRAITCOMPARISONPORTRAITSDIFFER       CheckDiagnose = 150
	CheckDiagnosePORTRAITCOMPARISONSERVICEERROR          CheckDiagnose = 152
	CheckDiagnoseSOURCESCOMPARISONERROR                  CheckDiagnose = 13
	CheckDiagnoseSPECKSINUV                              CheckDiagnose = 33
	CheckDiagnoseSYNTAXERROR                             CheckDiagnose = 11
	CheckDiagnoseTEXTCOLORSHOULDBEBLUE                   CheckDiagnose = 130
	CheckDiagnoseTEXTCOLORSHOULDBEGREEN                  CheckDiagnose = 131
	CheckDiagnoseTEXTCOLORSHOULDBERED                    CheckDiagnose = 132
	CheckDiagnoseTEXTSHOULDBEBLACK                       CheckDiagnose = 133
	CheckDiagnoseTOOLOWRESOLUTION                        CheckDiagnose = 34
	CheckDiagnoseTOOMANYOBJECTS                          CheckDiagnose = 31
	CheckDiagnoseTOOMUCHSHIFT                            CheckDiagnose = 28
	CheckDiagnoseTRUELUMINESCENCEERROR                   CheckDiagnose = 20
	CheckDiagnoseUNCERTAINVERIFICATION                   CheckDiagnose = 5
	CheckDiagnoseUNKNOWN                                 CheckDiagnose = 0
	CheckDiagnoseUVDULLPAPERBLANK                        CheckDiagnose = 53
	CheckDiagnoseUVDULLPAPERERROR                        CheckDiagnose = 54
	CheckDiagnoseUVDULLPAPERMRZ                          CheckDiagnose = 50
	CheckDiagnoseUVDULLPAPERPHOTO                        CheckDiagnose = 52
	CheckDiagnoseVISIBLEELEMENTABSENT                    CheckDiagnose = 41
)

// Defines values for CheckResult.
const (
	CheckResultERROR      CheckResult = 0
	CheckResultOK         CheckResult = 1
	CheckResultWASNOTDONE CheckResult = 2
)

// Defines values for Critical.
const (
	CRITICAL    Critical = 1
	NOTCRITICAL Critical = 0
)

// Defines values for DocumentFormat.
const (
	DocumentFormatA4       DocumentFormat = 4
	DocumentFormatCUSTOM   DocumentFormat = 1000
	DocumentFormatFLEXIBLE DocumentFormat = 1002
	DocumentFormatID1      DocumentFormat = 0
	DocumentFormatID1180   DocumentFormat = 11
	DocumentFormatID1270   DocumentFormat = 12
	DocumentFormatID190    DocumentFormat = 10
	DocumentFormatID2      DocumentFormat = 1
	DocumentFormatID2180   DocumentFormat = 13
	DocumentFormatID3      DocumentFormat = 2
	DocumentFormatID3180   DocumentFormat = 14
	DocumentFormatID3X2    DocumentFormat = 5
	DocumentFormatNON      DocumentFormat = 3
)

// Defines values for DocumentType.
const (
	DocumentTypeADDRESSCARD                                   DocumentType = 183
	DocumentTypeAIRPORTIMMIGRATIONCARD                        DocumentType = 184
	DocumentTypeALIENREGISTRATIONCARD                         DocumentType = 185
	DocumentTypeALIENSIDENTITYCARD                            DocumentType = 22
	DocumentTypeALIENSPASSPORT                                DocumentType = 27
	DocumentTypeALTERNATIVEIDENTITYCARD                       DocumentType = 28
	DocumentTypeAPEHCARD                                      DocumentType = 186
	DocumentTypeARMEDFORCESIDENTITYCARD                       DocumentType = 225
	DocumentTypeAUTHORIZATIONCARD                             DocumentType = 32
	DocumentTypeBEGINNERPERMIT                                DocumentType = 33
	DocumentTypeBORDERCROSSINGCARD                            DocumentType = 34
	DocumentTypeBORDERCROSSINGPERMIT                          DocumentType = 217
	DocumentTypeCERTIFICATEOFCITIZENSHIP                      DocumentType = 182
	DocumentTypeCERTIFICATEOFCOMPETENCY                       DocumentType = 237
	DocumentTypeCERTIFICATEOFPROFICIENCY                      DocumentType = 238
	DocumentTypeCHAUFFEURLICENSE                              DocumentType = 35
	DocumentTypeCHAUFFEURLICENSEUNDER18                       DocumentType = 36
	DocumentTypeCHAUFFEURLICENSEUNDER21                       DocumentType = 37
	DocumentTypeCOMMERCIALDRIVINGLICENSE                      DocumentType = 38
	DocumentTypeCOMMERCIALDRIVINGLICENSEINSTRUCTIONALPERMIT   DocumentType = 39
	DocumentTypeCOMMERCIALDRIVINGLICENSENOVICE                DocumentType = 169
	DocumentTypeCOMMERCIALDRIVINGLICENSENOVICEUNDER18         DocumentType = 170
	DocumentTypeCOMMERCIALDRIVINGLICENSENOVICEUNDER21         DocumentType = 171
	DocumentTypeCOMMERCIALDRIVINGLICENSEUNDER18               DocumentType = 40
	DocumentTypeCOMMERCIALDRIVINGLICENSEUNDER21               DocumentType = 41
	DocumentTypeCOMMERCIALINSTRUCTIONPERMIT                   DocumentType = 42
	DocumentTypeCOMMERCIALNEWPERMIT                           DocumentType = 43
	DocumentTypeCOMPANYCARD                                   DocumentType = 221
	DocumentTypeCONCEALEDCARRYLICENSE                         DocumentType = 44
	DocumentTypeCONCEALEDFIREARMPERMIT                        DocumentType = 45
	DocumentTypeCONDITIONALDRIVINGLICENSE                     DocumentType = 46
	DocumentTypeCONSULARIDENTITYCARD                          DocumentType = 213
	DocumentTypeCOUPONTODRIVINGLICENSE                        DocumentType = 187
	DocumentTypeCREWMEMBERCERTIFICATE                         DocumentType = 188
	DocumentTypeDEPARTMENTOFVETERANSAFFAIRSIDENTITYCARD       DocumentType = 47
	DocumentTypeDIPLOMATICDRIVINGLICENSE                      DocumentType = 48
	DocumentTypeDIPLOMATICIDENTITYCARD                        DocumentType = 212
	DocumentTypeDIPLOMATICPASSPORT                            DocumentType = 13
	DocumentTypeDOCUMENTFORRETURN                             DocumentType = 189
	DocumentTypeDOCUMENTOFIDENTITY                            DocumentType = 216
	DocumentTypeDOMESTICPASSPORT                              DocumentType = 222
	DocumentTypeDRIVERCARD                                    DocumentType = 228
	DocumentTypeDRIVERTRAININGCERTIFICATE                     DocumentType = 229
	DocumentTypeDRIVINGLICENSE                                DocumentType = 49
	DocumentTypeDRIVINGLICENSEINSTRUCTIONALPERMIT             DocumentType = 50
	DocumentTypeDRIVINGLICENSEINSTRUCTIONALPERMITUNDER18      DocumentType = 51
	DocumentTypeDRIVINGLICENSEINSTRUCTIONALPERMITUNDER21      DocumentType = 52
	DocumentTypeDRIVINGLICENSELEARNERSPERMIT                  DocumentType = 53
	DocumentTypeDRIVINGLICENSELEARNERSPERMITUNDER18           DocumentType = 54
	DocumentTypeDRIVINGLICENSELEARNERSPERMITUNDER21           DocumentType = 55
	DocumentTypeDRIVINGLICENSENOVICE                          DocumentType = 56
	DocumentTypeDRIVINGLICENSENOVICEUNDER18                   DocumentType = 57
	DocumentTypeDRIVINGLICENSENOVICEUNDER21                   DocumentType = 58
	DocumentTypeDRIVINGLICENSEREGISTEREDOFFENDER              DocumentType = 59
	DocumentTypeDRIVINGLICENSERESTRICTEDUNDER18               DocumentType = 60
	DocumentTypeDRIVINGLICENSERESTRICTEDUNDER21               DocumentType = 61
	DocumentTypeDRIVINGLICENSETEMPORARYVISITOR                DocumentType = 62
	DocumentTypeDRIVINGLICENSETEMPORARYVISITORUNDER18         DocumentType = 63
	DocumentTypeDRIVINGLICENSETEMPORARYVISITORUNDER21         DocumentType = 64
	DocumentTypeDRIVINGLICENSEUNDER18                         DocumentType = 65
	DocumentTypeDRIVINGLICENSEUNDER19                         DocumentType = 176
	DocumentTypeDRIVINGLICENSEUNDER21                         DocumentType = 66
	DocumentTypeECARD                                         DocumentType = 190
	DocumentTypeEMERGENCYPASSPORT                             DocumentType = 26
	DocumentTypeEMPLOYMENTCARD                                DocumentType = 191
	DocumentTypeEMPLOYMENTDRIVINGPERMIT                       DocumentType = 67
	DocumentTypeENHANCEDCHAUFFEURLICENSE                      DocumentType = 68
	DocumentTypeENHANCEDCHAUFFEURLICENSEUNDER18               DocumentType = 69
	DocumentTypeENHANCEDCHAUFFEURLICENSEUNDER21               DocumentType = 70
	DocumentTypeENHANCEDCOMMERCIALDRIVINGLICENSE              DocumentType = 71
	DocumentTypeENHANCEDDRIVINGLICENSE                        DocumentType = 72
	DocumentTypeENHANCEDDRIVINGLICENSEUNDER18                 DocumentType = 73
	DocumentTypeENHANCEDDRIVINGLICENSEUNDER21                 DocumentType = 74
	DocumentTypeENHANCEDIDENTITYCARD                          DocumentType = 75
	DocumentTypeENHANCEDIDENTITYCARDUNDER18                   DocumentType = 76
	DocumentTypeENHANCEDIDENTITYCARDUNDER21                   DocumentType = 77
	DocumentTypeENHANCEDOPERATORSLICENSE                      DocumentType = 78
	DocumentTypeFIREARMSPERMIT                                DocumentType = 79
	DocumentTypeFULLPROVISIONALLICENSE                        DocumentType = 80
	DocumentTypeFULLPROVISIONALLICENSEUNDER18                 DocumentType = 81
	DocumentTypeFULLPROVISIONALLICENSEUNDER21                 DocumentType = 82
	DocumentTypeGENEVACONVENTIONSIDENTITYCARD                 DocumentType = 83
	DocumentTypeGRADUATEDDRIVINGLICENSEUNDER18                DocumentType = 84
	DocumentTypeGRADUATEDDRIVINGLICENSEUNDER21                DocumentType = 85
	DocumentTypeGRADUATEDINSTRUCTIONPERMITUNDER18             DocumentType = 86
	DocumentTypeGRADUATEDINSTRUCTIONPERMITUNDER21             DocumentType = 87
	DocumentTypeGRADUATEDLICENSEUNDER18                       DocumentType = 88
	DocumentTypeGRADUATEDLICENSEUNDER21                       DocumentType = 89
	DocumentTypeHANDGUNCARRYPERMIT                            DocumentType = 90
	DocumentTypeHEALTHCARD                                    DocumentType = 181
	DocumentTypeHKSARIMMIGRATIONFORM                          DocumentType = 192
	DocumentTypeIDENTITYANDPRIVILEGECARD                      DocumentType = 91
	DocumentTypeIDENTITYCARD                                  DocumentType = 12
	DocumentTypeIDENTITYCARDFORRESIDENCE                      DocumentType = 16
	DocumentTypeIDENTITYCARDMOBILITYIMPAIRED                  DocumentType = 92
	DocumentTypeIDENTITYCARDREGISTEREDOFFENDER                DocumentType = 93
	DocumentTypeIDENTITYCARDTEMPORARYVISITOR                  DocumentType = 94
	DocumentTypeIDENTITYCARDTEMPORARYVISITORUNDER18           DocumentType = 95
	DocumentTypeIDENTITYCARDTEMPORARYVISITORUNDER21           DocumentType = 96
	DocumentTypeIDENTITYCARDUNDER18                           DocumentType = 97
	DocumentTypeIDENTITYCARDUNDER19                           DocumentType = 177
	DocumentTypeIDENTITYCARDUNDER21                           DocumentType = 98
	DocumentTypeIDENTITYCERTIFICATE                           DocumentType = 223
	DocumentTypeIGNITIONINTERLOCKPERMIT                       DocumentType = 100
	DocumentTypeIMMIGRANTCARD                                 DocumentType = 193
	DocumentTypeIMMIGRANTVISA                                 DocumentType = 101
	DocumentTypeINCOMETAXCARD                                 DocumentType = 214
	DocumentTypeINSTRUCTIONPERMIT                             DocumentType = 102
	DocumentTypeINSTRUCTIONPERMITUNDER18                      DocumentType = 103
	DocumentTypeINSTRUCTIONPERMITUNDER21                      DocumentType = 104
	DocumentTypeINTERIMDRIVINGLICENSE                         DocumentType = 105
	DocumentTypeINTERIMIDENTITYCARD                           DocumentType = 106
	DocumentTypeINTERIMINSTRUCTIONALPERMIT                    DocumentType = 236
	DocumentTypeINTERMEDIATEDRIVINGLICENSE                    DocumentType = 107
	DocumentTypeINTERMEDIATEDRIVINGLICENSEUNDER18             DocumentType = 108
	DocumentTypeINTERMEDIATEDRIVINGLICENSEUNDER21             DocumentType = 109
	DocumentTypeINVOICE                                       DocumentType = 241
	DocumentTypeJUNIORDRIVINGLICENSE                          DocumentType = 110
	DocumentTypeLABOURCARD                                    DocumentType = 194
	DocumentTypeLAISSEZPASSER                                 DocumentType = 195
	DocumentTypeLAWYERIDENTITYCERTIFICATE                     DocumentType = 196
	DocumentTypeLEARNERINSTRUCTIONALPERMIT                    DocumentType = 111
	DocumentTypeLEARNERLICENSE                                DocumentType = 112
	DocumentTypeLEARNERLICENSEUNDER18                         DocumentType = 113
	DocumentTypeLEARNERLICENSEUNDER21                         DocumentType = 114
	DocumentTypeLEARNERPERMIT                                 DocumentType = 115
	DocumentTypeLEARNERPERMITUNDER18                          DocumentType = 116
	DocumentTypeLEARNERPERMITUNDER21                          DocumentType = 117
	DocumentTypeLICENSECARD                                   DocumentType = 197
	DocumentTypeLIMITEDLICENSE                                DocumentType = 118
	DocumentTypeLIMITEDPERMIT                                 DocumentType = 119
	DocumentTypeLIMITEDTERMDRIVINGLICENSE                     DocumentType = 120
	DocumentTypeLIMITEDTERMIDENTITYCARD                       DocumentType = 121
	DocumentTypeLIQUORIDENTITYCARD                            DocumentType = 122
	DocumentTypeMARINELICENSE                                 DocumentType = 233
	DocumentTypeMEMBERSHIPCARD                                DocumentType = 231
	DocumentTypeNATIONALIDENTITYCARD                          DocumentType = 20
	DocumentTypeNEWPERMIT                                     DocumentType = 123
	DocumentTypeNEWPERMITUNDER18                              DocumentType = 124
	DocumentTypeNEWPERMITUNDER21                              DocumentType = 125
	DocumentTypeNONUSCITIZENDRIVINGLICENSE                    DocumentType = 126
	DocumentTypeNOTDEFINED                                    DocumentType = 0
	DocumentTypeOCCUPATIONALDRIVINGLICENSE                    DocumentType = 127
	DocumentTypeONEIDATRIBEOFINDIANSIDENTITYCARD              DocumentType = 128
	DocumentTypeOPERATORLICENSE                               DocumentType = 129
	DocumentTypeOPERATORLICENSEUNDER18                        DocumentType = 130
	DocumentTypeOPERATORLICENSEUNDER21                        DocumentType = 131
	DocumentTypeORIGINCARD                                    DocumentType = 25
	DocumentTypeOTHER                                         DocumentType = 99
	DocumentTypePASSENGERLOCATORFORM                          DocumentType = 242
	DocumentTypePASSPORT                                      DocumentType = 11
	DocumentTypePASSPORTCARD                                  DocumentType = 172
	DocumentTypePASSPORTCHILD                                 DocumentType = 199
	DocumentTypePASSPORTCONSULAR                              DocumentType = 200
	DocumentTypePASSPORTDIPLOMATICSERVICE                     DocumentType = 201
	DocumentTypePASSPORTLIMITEDVALIDITY                       DocumentType = 218
	DocumentTypePASSPORTOFFICIAL                              DocumentType = 202
	DocumentTypePASSPORTPAGE                                  DocumentType = 240
	DocumentTypePASSPORTPROVISIONAL                           DocumentType = 203
	DocumentTypePASSPORTRESIDENTCARD                          DocumentType = 173
	DocumentTypePASSPORTSPECIAL                               DocumentType = 204
	DocumentTypePASSPORTSTATELESS                             DocumentType = 198
	DocumentTypePERMANENTDRIVINGLICENSE                       DocumentType = 132
	DocumentTypePERMISSIONTOTHELOCALBORDERTRAFFIC             DocumentType = 205
	DocumentTypePERMITTOREENTER                               DocumentType = 133
	DocumentTypePERSONALIDENTIFICATIONVERIFICATION            DocumentType = 174
	DocumentTypePRIVILEGEDIDENTITYCARD                        DocumentType = 23
	DocumentTypePROBATIONARYAUTOLICENSE                       DocumentType = 134
	DocumentTypePROBATIONARYDRIVINGLICENSEUNDER18             DocumentType = 135
	DocumentTypePROBATIONARYDRIVINGLICENSEUNDER21             DocumentType = 136
	DocumentTypePROBATIONARYVEHICLESALESPERSONLICENSE         DocumentType = 137
	DocumentTypePROFESSIONALCARD                              DocumentType = 226
	DocumentTypePROVISIONALDRIVINGLICENSE                     DocumentType = 138
	DocumentTypePROVISIONALDRIVINGLICENSEUNDER18              DocumentType = 139
	DocumentTypePROVISIONALDRIVINGLICENSEUNDER21              DocumentType = 140
	DocumentTypePROVISIONALLICENSE                            DocumentType = 141
	DocumentTypePROVISIONALLICENSEUNDER18                     DocumentType = 142
	DocumentTypePROVISIONALLICENSEUNDER21                     DocumentType = 143
	DocumentTypePUBLICPASSENGERCHAUFFEURLICENSE               DocumentType = 144
	DocumentTypePUBLICVEHICLEDRIVERAUTHORITYCARD              DocumentType = 232
	DocumentTypeQUALIFICATIONDRIVINGLICENSE                   DocumentType = 230
	DocumentTypeRACINGANDGAMINGCOMISSIONCARD                  DocumentType = 145
	DocumentTypeREFUGEETRAVELDOCUMENT                         DocumentType = 146
	DocumentTypeREGISTRATIONCERTIFICATE                       DocumentType = 206
	DocumentTypeREGISTRATIONSTAMP                             DocumentType = 227
	DocumentTypeRENEWALPERMIT                                 DocumentType = 147
	DocumentTypeRESIDENCEPERMIT                               DocumentType = 215
	DocumentTypeRESIDENCEPERMITIDENTITYCARD                   DocumentType = 24
	DocumentTypeRESIDENTIDCARD                                DocumentType = 224
	DocumentTypeRESTRICTEDCOMMERCIALDRIVERLICENSE             DocumentType = 148
	DocumentTypeRESTRICTEDDRIVERLICENSE                       DocumentType = 149
	DocumentTypeRESTRICTEDPERMIT                              DocumentType = 150
	DocumentTypeSEAMANSIDENTITYDOCUMENT                       DocumentType = 15
	DocumentTypeSEASONALCITIZENIDENTITYCARD                   DocumentType = 153
	DocumentTypeSEASONALPERMIT                                DocumentType = 151
	DocumentTypeSEASONALRESIDENTIDENTITYCARD                  DocumentType = 152
	DocumentTypeSEDESOLCARD                                   DocumentType = 207
	DocumentTypeSERVICEPASSPORT                               DocumentType = 14
	DocumentTypeSEXOFFENDER                                   DocumentType = 154
	DocumentTypeSIMCARD                                       DocumentType = 219
	DocumentTypeSOCIALCARD                                    DocumentType = 208
	DocumentTypeSOCIALIDENTITYCARD                            DocumentType = 21
	DocumentTypeSOCIALSECURITYCARD                            DocumentType = 155
	DocumentTypeTAXCARD                                       DocumentType = 220
	DocumentTypeTBCARD                                        DocumentType = 209
	DocumentTypeTEMPORARYCOMMERCIALDRIVINGLICENSE             DocumentType = 235
	DocumentTypeTEMPORARYDRIVINGLICENSE                       DocumentType = 156
	DocumentTypeTEMPORARYDRIVINGLICENSEUNDER18                DocumentType = 157
	DocumentTypeTEMPORARYDRIVINGLICENSEUNDER21                DocumentType = 158
	DocumentTypeTEMPORARYIDENTITYCARD                         DocumentType = 159
	DocumentTypeTEMPORARYINSTRUCTIONPERMITIDENTITYCARD        DocumentType = 160
	DocumentTypeTEMPORARYINSTRUCTIONPERMITIDENTITYCARDUNDER18 DocumentType = 161
	DocumentTypeTEMPORARYINSTRUCTIONPERMITIDENTITYCARDUNDER21 DocumentType = 162
	DocumentTypeTEMPORARYLEARNERLICENSE                       DocumentType = 234
	DocumentTypeTEMPORARYOPERATORLICENSE                      DocumentType = 175
	DocumentTypeTEMPORARYPASSPORT                             DocumentType = 179
	DocumentTypeTEMPORARYVISITORDRIVINGLICENSE                DocumentType = 163
	DocumentTypeTEMPORARYVISITORDRIVINGLICENSEUNDER18         DocumentType = 164
	DocumentTypeTEMPORARYVISITORDRIVINGLICENSEUNDER21         DocumentType = 165
	DocumentTypeTRADELICENSE                                  DocumentType = 239
	DocumentTypeTRAVELDOCUMENT                                DocumentType = 17
	DocumentTypeUNIFORMEDSERVICESIDENTITYCARD                 DocumentType = 166
	DocumentTypeVEHICLEPASSPORT                               DocumentType = 210
	DocumentTypeVEHICLESALESPERSONLICENSE                     DocumentType = 167
	DocumentTypeVISA                                          DocumentType = 178
	DocumentTypeVISAID2                                       DocumentType = 29
	DocumentTypeVISAID3                                       DocumentType = 30
	DocumentTypeVOTINGCARD                                    DocumentType = 180
	DocumentTypeWDOCUMENT                                     DocumentType = 211
	DocumentTypeWORKERIDENTIFICATIONCREDENTIAL                DocumentType = 168
)

// Defines values for DocumentTypeRecognitionResult.
const (
	DocumentTypeRecognitionResultNEEDMOREIMAGES DocumentTypeRecognitionResult = 29
	DocumentTypeRecognitionResultOK             DocumentTypeRecognitionResult = 0
	DocumentTypeRecognitionResultUNKNOWN        DocumentTypeRecognitionResult = 15
)

// Defines values for GraphicFieldType.
const (
	GraphicFieldTypeBARCODE                GraphicFieldType = 205
	GraphicFieldTypeCOLORDYNAMIC           GraphicFieldType = 209
	GraphicFieldTypeCONTACTCHIP            GraphicFieldType = 213
	GraphicFieldTypeDOCUMENTFRONT          GraphicFieldType = 207
	GraphicFieldTypeDOCUMENTREAR           GraphicFieldType = 208
	GraphicFieldTypeEYE                    GraphicFieldType = 203
	GraphicFieldTypeFINGERLEFTFOURFINGERS  GraphicFieldType = 314
	GraphicFieldTypeFINGERLEFTINDEX        GraphicFieldType = 301
	GraphicFieldTypeFINGERLEFTLITTLE       GraphicFieldType = 304
	GraphicFieldTypeFINGERLEFTMIDDLE       GraphicFieldType = 302
	GraphicFieldTypeFINGERLEFTRING         GraphicFieldType = 303
	GraphicFieldTypeFINGERLEFTTHUMB        GraphicFieldType = 300
	GraphicFieldTypeFINGERPRINT            GraphicFieldType = 202
	GraphicFieldTypeFINGERRIGHTFOURFINGERS GraphicFieldType = 313
	GraphicFieldTypeFINGERRIGHTINDEX       GraphicFieldType = 306
	GraphicFieldTypeFINGERRIGHTLITTLE      GraphicFieldType = 309
	GraphicFieldTypeFINGERRIGHTMIDDLE      GraphicFieldType = 307
	GraphicFieldTypeFINGERRIGHTRING        GraphicFieldType = 308
	GraphicFieldTypeFINGERRIGHTTHUMB       GraphicFieldType = 305
	GraphicFieldTypeFINGERTWOTHUMBS        GraphicFieldType = 315
	GraphicFieldTypeGHOSTPORTRAIT          GraphicFieldType = 210
	GraphicFieldTypeOTHER                  GraphicFieldType = 250
	GraphicFieldTypePORTRAIT               GraphicFieldType = 201
	GraphicFieldTypePROOFOFCITIZENSHIP     GraphicFieldType = 206
	GraphicFieldTypeSIGNATURE              GraphicFieldType = 204
	GraphicFieldTypeSTAMP                  GraphicFieldType = 211
)

// Defines values for ImageQualityCheckType.
const (
	Bounds          ImageQualityCheckType = 5
	Brightness      ImageQualityCheckType = 9
	ImageColorness  ImageQualityCheckType = 3
	ImageFocus      ImageQualityCheckType = 1
	ImageGlares     ImageQualityCheckType = 0
	ImageResolution ImageQualityCheckType = 2
	Perspective     ImageQualityCheckType = 4
	Portrait        ImageQualityCheckType = 7
	ScreenCapture   ImageQualityCheckType = 6
)

// Defines values for LCID.
const (
	LCIDABKHAZIAN                LCID = 10011
	LCIDAFRIKAANS                LCID = 1078
	LCIDALBANIAN                 LCID = 1052
	LCIDAMHARIC                  LCID = 1118
	LCIDARABIC                   LCID = 4096
	LCIDARABICALGERIA            LCID = 5121
	LCIDARABICARMENIAN           LCID = 1067
	LCIDARABICBAHRAIN            LCID = 15361
	LCIDARABICEGYPT              LCID = 3073
	LCIDARABICIRAQ               LCID = 2049
	LCIDARABICJORDAN             LCID = 11265
	LCIDARABICKUWAIT             LCID = 13313
	LCIDARABICLEBANON            LCID = 12289
	LCIDARABICLIBYA              LCID = 4097
	LCIDARABICMOROCCO            LCID = 6145
	LCIDARABICOMAN               LCID = 8193
	LCIDARABICQATAR              LCID = 16385
	LCIDARABICSAUDIARABIA        LCID = 1025
	LCIDARABICSYRIA              LCID = 10241
	LCIDARABICTUNISIA            LCID = 7169
	LCIDARABICUAE                LCID = 14337
	LCIDARABICYEMEN              LCID = 9217
	LCIDAZERICYRILIC             LCID = 2092
	LCIDAZERILATIN               LCID = 1068
	LCIDAssamese                 LCID = 1101
	LCIDBANKCARD                 LCID = 10003
	LCIDBANKCARDCVV2             LCID = 10004
	LCIDBANKCARDEXPIRYDATE       LCID = 10001
	LCIDBANKCARDNAME             LCID = 10002
	LCIDBANKCARDNUMBER           LCID = 10000
	LCIDBASQUE                   LCID = 1069
	LCIDBELARUSIAN               LCID = 1059
	LCIDBULGARIAN                LCID = 1026
	LCIDBURMESE                  LCID = 1109
	LCIDBengaliBangladesh        LCID = 2117
	LCIDBengaliIndia             LCID = 1093
	LCIDCATALAN                  LCID = 1027
	LCIDCHINESE                  LCID = 2052
	LCIDCROATIAN                 LCID = 1050
	LCIDCTCSIMPLIFIED            LCID = 50001
	LCIDCTCTRADITIONAL           LCID = 50002
	LCIDCZECH                    LCID = 1029
	LCIDDANISH                   LCID = 1030
	LCIDDUTCHBELGIUM             LCID = 2067
	LCIDDUTCHNETHERLANDS         LCID = 1043
	LCIDENGLISHAUSTRALIA         LCID = 3081
	LCIDENGLISHBELIZE            LCID = 10249
	LCIDENGLISHCANADA            LCID = 4105
	LCIDENGLISHCARRIBEAN         LCID = 9225
	LCIDENGLISHIRELAND           LCID = 6153
	LCIDENGLISHJAMAICA           LCID = 8201
	LCIDENGLISHNEWZEALAND        LCID = 5129
	LCIDENGLISHPHILIPPINES       LCID = 13321
	LCIDENGLISHSOUTHAFRICA       LCID = 7177
	LCIDENGLISHTRINIDAD          LCID = 11273
	LCIDENGLISHUK                LCID = 2057
	LCIDENGLISHUS                LCID = 1033
	LCIDENGLISHZIMBABWE          LCID = 12297
	LCIDESTONIAN                 LCID = 1061
	LCIDFAEROESE                 LCID = 1080
	LCIDFARSI                    LCID = 1065
	LCIDFINNISH                  LCID = 1035
	LCIDFRENCHBELGIUM            LCID = 2060
	LCIDFRENCHCANADA             LCID = 3084
	LCIDFRENCHFRANCE             LCID = 1036
	LCIDFRENCHLUXEMBOURG         LCID = 5132
	LCIDFRENCHMONACO             LCID = 6156
	LCIDFRENCHSWITZERLAND        LCID = 4108
	LCIDFYROMACEDONIAN           LCID = 1071
	LCIDGALICIAN                 LCID = 1110
	LCIDGEORGIAN                 LCID = 1079
	LCIDGERMANAUSTRIA            LCID = 3079
	LCIDGERMANGERMANY            LCID = 1031
	LCIDGERMANLIECHTENSTEIN      LCID = 5127
	LCIDGERMANLUXEMBOURG         LCID = 4103
	LCIDGERMANSWITZERLAND        LCID = 2055
	LCIDGREEK                    LCID = 1032
	LCIDGUJARATI                 LCID = 1095
	LCIDHEBREW                   LCID = 1037
	LCIDHINDIINDIA               LCID = 1081
	LCIDHUNGARIAN                LCID = 1038
	LCIDICELANDIC                LCID = 1039
	LCIDINDONESIAN               LCID = 1057
	LCIDITALIANITALY             LCID = 1040
	LCIDITALIANSWITZERLAND       LCID = 2064
	LCIDJAPANESE                 LCID = 1041
	LCIDKANNADA                  LCID = 1099
	LCIDKARAKALPAK               LCID = 10012
	LCIDKASHMIRI                 LCID = 1120
	LCIDKAZAKH                   LCID = 1087
	LCIDKHMER                    LCID = 1107
	LCIDKONKANI                  LCID = 1111
	LCIDKOREAN                   LCID = 1042
	LCIDKYRGYZCYRILICK           LCID = 1088
	LCIDLAO                      LCID = 1108
	LCIDLATIN                    LCID = 0
	LCIDLATVIAN                  LCID = 1062
	LCIDLITHUANIAN               LCID = 1063
	LCIDMALAYALAM                LCID = 1100
	LCIDMALAYBRUNEIDARUSSALAM    LCID = 2110
	LCIDMALAYMALAYSIA            LCID = 1086
	LCIDMALTESE                  LCID = 1082
	LCIDMARATHI                  LCID = 1102
	LCIDMONGOLIANCYRILIC         LCID = 1104
	LCIDNEPALI                   LCID = 1121
	LCIDNORWEGIANBOKMAL          LCID = 1044
	LCIDNORWEGIANNYORSK          LCID = 2068
	LCIDORIYA                    LCID = 1096
	LCIDPASHTO                   LCID = 1123
	LCIDPOLISH                   LCID = 1045
	LCIDPORTUGUESEBRAZIL         LCID = 1046
	LCIDPORTUGUESEPORTUGAL       LCID = 2070
	LCIDPUNJABI                  LCID = 1094
	LCIDRHAETOROMANIC            LCID = 1047
	LCIDROMANIAN                 LCID = 1048
	LCIDRUSSIAN                  LCID = 1049
	LCIDSANSKRIT                 LCID = 1103
	LCIDSERBIANCYRILIC           LCID = 3098
	LCIDSERBIANLATIN             LCID = 2074
	LCIDSINDHI                   LCID = 2137
	LCIDSINDHIINDIA              LCID = 1113
	LCIDSINHALA                  LCID = 1115
	LCIDSLOVAK                   LCID = 1051
	LCIDSLOVENIAN                LCID = 1060
	LCIDSPANICHCOLOMBIA          LCID = 9226
	LCIDSPANISHARGENTINA         LCID = 11274
	LCIDSPANISHBOLIVIA           LCID = 16394
	LCIDSPANISHCHILE             LCID = 13322
	LCIDSPANISHCOSTARICA         LCID = 5130
	LCIDSPANISHDOMINICANREPUBLIC LCID = 7178
	LCIDSPANISHECUADOR           LCID = 12298
	LCIDSPANISHELSALVADOR        LCID = 17418
	LCIDSPANISHGUATEMALA         LCID = 4106
	LCIDSPANISHHONDURAS          LCID = 18442
	LCIDSPANISHINTERNATIONALSORT LCID = 3082
	LCIDSPANISHMEXICO            LCID = 2058
	LCIDSPANISHNICARAGUA         LCID = 19466
	LCIDSPANISHPANAMA            LCID = 6154
	LCIDSPANISHPARAGUAY          LCID = 15370
	LCIDSPANISHPERU              LCID = 10250
	LCIDSPANISHPUERTORICO        LCID = 20490
	LCIDSPANISHTRADITIONALSORT   LCID = 1034
	LCIDSPANISHURUGUAY           LCID = 14346
	LCIDSPANISHVENEZUELA         LCID = 8202
	LCIDSWAHILI                  LCID = 1089
	LCIDSWEDISH                  LCID = 1053
	LCIDSWEDISHFINLAND           LCID = 2077
	LCIDTAJIKCYRILLIC            LCID = 1064
	LCIDTAMIL                    LCID = 1097
	LCIDTATAR                    LCID = 1092
	LCIDTELUGU                   LCID = 1098
	LCIDTHAITHAILAND             LCID = 1054
	LCIDTURKISH                  LCID = 1055
	LCIDTURKMEN                  LCID = 1090
	LCIDUKRAINIAN                LCID = 1058
	LCIDURDU                     LCID = 1056
	LCIDURDUDETECTION            LCID = 10560
	LCIDUZBEKCYRILIC             LCID = 2115
	LCIDUZBEKLATIN               LCID = 1091
	LCIDVIETNAMESE               LCID = 1066
)

// Defines values for Light.
const (
	IR    Light = 24
	OFF   Light = 0
	UV    Light = 128
	WHITE Light = 6
)

// Defines values for LogLevel.
const (
	LogLevelDEBUG      LogLevel = "Debug"
	LogLevelERROR      LogLevel = "Error"
	LogLevelFATALERROR LogLevel = "FatalError"
	LogLevelINFO       LogLevel = "Info"
	LogLevelWARNING    LogLevel = "Warning"
)

// Defines values for MRZFormat.
const (
	MRZFormatCAN    MRZFormat = "1x6"
	MRZFormatID1    MRZFormat = "3x30"
	MRZFormatID1230 MRZFormat = "2x30"
	MRZFormatID2    MRZFormat = "2x36"
	MRZFormatID3    MRZFormat = "2x44"
	MRZFormatIDL    MRZFormat = "1x30"
)

// Defines values for MeasureSystem.
const (
	IMPERIAL MeasureSystem = 1
	METRIC   MeasureSystem = 0
)

// Defines values for ParsingNotificationCodes.
const (
	NtfLDSASNCertificateDuplicatingExtensions            ParsingNotificationCodes = -1.879048169e+09
	NtfLDSASNCertificateEmptyIssuer                      ParsingNotificationCodes = -1.879048187e+09
	NtfLDSASNCertificateEmptySubject                     ParsingNotificationCodes = -1.879048186e+09
	NtfLDSASNCertificateForcedDefaultCSCARole            ParsingNotificationCodes = -1.879048178e+09
	NtfLDSASNCertificateForcedDefaultDSRole              ParsingNotificationCodes = -1.879048177e+09
	NtfLDSASNCertificateIncorrectIssuerSubjectDS         ParsingNotificationCodes = -1.879048176e+09
	NtfLDSASNCertificateIncorrectTimeCoding              ParsingNotificationCodes = -1.879048189e+09
	NtfLDSASNCertificateIncorrectUseOfGeneralizedTime    ParsingNotificationCodes = -1.879048188e+09
	NtfLDSASNCertificateIncorrectVersion                 ParsingNotificationCodes = -1.879048191e+09
	NtfLDSASNCertificateNonMatchingSignatureAlgorithm    ParsingNotificationCodes = -1.87904819e+09
	NtfLDSASNCertificateUnsupportedCriticalExtension     ParsingNotificationCodes = -1.879048184e+09
	NtfLDSASNSignedDataContentOIDIncorrect               ParsingNotificationCodes = -1.879047775e+09
	NtfLDSASNSignedDataOIDIncorrect                      ParsingNotificationCodes = -1.879047936e+09
	NtfLDSASNSignedDataVersionIncorrect                  ParsingNotificationCodes = -1.879047776e+09
	NtfLDSASNSignerInfoContentTypeAttrData               ParsingNotificationCodes = -1.879047919e+09
	NtfLDSASNSignerInfoContentTypeAttrMissing            ParsingNotificationCodes = -1.87904792e+09
	NtfLDSASNSignerInfoContentTypeAttrValue              ParsingNotificationCodes = -1.879047918e+09
	NtfLDSASNSignerInfoListContentDescriptionAttrData    ParsingNotificationCodes = -1.879047905e+09
	NtfLDSASNSignerInfoListContentDescriptionAttrMissing ParsingNotificationCodes = -1.879047906e+09
	NtfLDSASNSignerInfoMessageDigestAttrData             ParsingNotificationCodes = -1.879047922e+09
	NtfLDSASNSignerInfoMessageDigestAttrMissing          ParsingNotificationCodes = -1.879047923e+09
	NtfLDSASNSignerInfoMessageDigestAttrValue            ParsingNotificationCodes = -1.879047921e+09
	NtfLDSASNSignerInfoSIDDigestAlgorithmNotListed       ParsingNotificationCodes = -1.879047924e+09
	NtfLDSASNSignerInfoSIDIncorrectChoice                ParsingNotificationCodes = -1.879047925e+09
	NtfLDSASNSignerInfoSigningTimeAttrData               ParsingNotificationCodes = -1.879047908e+09
	NtfLDSASNSignerInfoSigningTimeAttrMissing            ParsingNotificationCodes = -1.879047909e+09
	NtfLDSASNSignerInfoSigningTimeAttrValue              ParsingNotificationCodes = -1.879047907e+09
	NtfLDSASNSignerInfoVersionIncorrect                  ParsingNotificationCodes = -1.879047926e+09
	NtfLDSAuthMLSignerInfoCertificateCantFindCSCA        ParsingNotificationCodes = -1.845493481e+09
	NtfLDSAuthMLSignerInfoCertificateRevoked             ParsingNotificationCodes = -1.84549348e+09
	NtfLDSAuthMLSignerInfoCertificateRootIsNotTrusted    ParsingNotificationCodes = -1.845493482e+09
	NtfLDSAuthMLSignerInfoCertificateSignatureInvalid    ParsingNotificationCodes = -1.845493479e+09
	NtfLDSAuthMLSignerInfoCertificateValidity            ParsingNotificationCodes = -1.845493483e+09
	NtfLDSAuthSignerInfoCertificateCantFindCSCA          ParsingNotificationCodes = -1.879047913e+09
	NtfLDSAuthSignerInfoCertificateRevoked               ParsingNotificationCodes = -1.879047912e+09
	NtfLDSAuthSignerInfoCertificateRootIsNotTrusted      ParsingNotificationCodes = -1.879047914e+09
	NtfLDSAuthSignerInfoCertificateSignatureInvalid      ParsingNotificationCodes = -1.879047911e+09
	NtfLDSAuthSignerInfoCertificateValidity              ParsingNotificationCodes = -1.879047915e+09
	NtfLDSBSIBlackListVersionIncorrect                   ParsingNotificationCodes = -1.87904772e+09
	NtfLDSBSIDefectListVersionIncorrect                  ParsingNotificationCodes = -1.879047728e+09
	NtfLDSBiometricsBDBDataEyeColor                      ParsingNotificationCodes = -1.87793408e+09
	NtfLDSBiometricsBDBDataFaceImageType                 ParsingNotificationCodes = -1.877409792e+09
	NtfLDSBiometricsBDBDataGender                        ParsingNotificationCodes = -1.877999616e+09
	NtfLDSBiometricsBDBDataHairColor                     ParsingNotificationCodes = -1.877868544e+09
	NtfLDSBiometricsBDBDataImageDataType                 ParsingNotificationCodes = -1.877344256e+09
	NtfLDSBiometricsBDBDataLengthIncorrect               ParsingNotificationCodes = -1.878327296e+09
	NtfLDSBiometricsBDBDataPoseAnglePitch                ParsingNotificationCodes = -1.877737472e+09
	NtfLDSBiometricsBDBDataPoseAngleRoll                 ParsingNotificationCodes = -1.877671936e+09
	NtfLDSBiometricsBDBDataPoseAngleUPitch               ParsingNotificationCodes = -1.877540864e+09
	NtfLDSBiometricsBDBDataPoseAngleURoll                ParsingNotificationCodes = -1.877475328e+09
	NtfLDSBiometricsBDBDataPoseAngleUYaw                 ParsingNotificationCodes = -1.8776064e+09
	NtfLDSBiometricsBDBDataPoseAngleYaw                  ParsingNotificationCodes = -1.877803008e+09
	NtfLDSBiometricsBDBFormatIDIncorrect                 ParsingNotificationCodes = -1.878458368e+09
	NtfLDSBiometricsBDBImageMissing                      ParsingNotificationCodes = -1.878523904e+09
	NtfLDSBiometricsBDBVersionIncorrect                  ParsingNotificationCodes = -1.878392832e+09
	NtfLDSBiometricsFormatOwnerIncorrect                 ParsingNotificationCodes = -1.87891712e+09
	NtfLDSBiometricsFormatOwnerMissing                   ParsingNotificationCodes = -1.878982656e+09
	NtfLDSBiometricsFormatTypeIncorrect                  ParsingNotificationCodes = -1.878786048e+09
	NtfLDSBiometricsFormatTypeMissing                    ParsingNotificationCodes = -1.878851584e+09
	NtfLDSBiometricsSubTypeIncorrect                     ParsingNotificationCodes = -1.87858944e+09
	NtfLDSBiometricsSubTypeMissing                       ParsingNotificationCodes = -1.878654976e+09
	NtfLDSBiometricsTypeIncorrect                        ParsingNotificationCodes = -1.878720512e+09
	NtfLDSCVCertificateNonCVCADomainParameters           ParsingNotificationCodes = -1.862270461e+09
	NtfLDSCVCertificatePrivateKeyIncorrectVersion        ParsingNotificationCodes = -1.86227046e+09
	NtfLDSCVCertificateProfileIncorrectVersion           ParsingNotificationCodes = -1.862270463e+09
	NtfLDSCVCertificateValidity                          ParsingNotificationCodes = -1.862270462e+09
	NtfLDSICAOApplicationLDSVersionInconsistent          ParsingNotificationCodes = -1.879048142e+09
	NtfLDSICAOApplicationLDSVersionUnsupported           ParsingNotificationCodes = -1.879048144e+09
	NtfLDSICAOApplicationUnicodeVersionInconsistent      ParsingNotificationCodes = -1.879048141e+09
	NtfLDSICAOApplicationUnicodeVersionUnsupported       ParsingNotificationCodes = -1.879048143e+09
	NtfLDSICAOCOMDGPMIncorrect                           ParsingNotificationCodes = -1.879048156e+09
	NtfLDSICAOCOMDGPMMissing                             ParsingNotificationCodes = -1.879048155e+09
	NtfLDSICAOCOMDGPMUnexpected                          ParsingNotificationCodes = -1.879048154e+09
	NtfLDSICAOCOMLDSVersionIncorrect                     ParsingNotificationCodes = -1.87904816e+09
	NtfLDSICAOCOMLDSVersionMissing                       ParsingNotificationCodes = -1.879048159e+09
	NtfLDSICAOCOMUnicodeVersionIncorrect                 ParsingNotificationCodes = -1.879048158e+09
	NtfLDSICAOCOMUnicodeVersionMissing                   ParsingNotificationCodes = -1.879048157e+09
	NtfLDSICAOCertificateChainCountryNonMatching         ParsingNotificationCodes = -1.8790476e+09
	NtfLDSICAOCertificateExtAuthKeyIDIncorrectData       ParsingNotificationCodes = -1.879047652e+09
	NtfLDSICAOCertificateExtAuthKeyIDKeyIDMissed         ParsingNotificationCodes = -1.879047651e+09
	NtfLDSICAOCertificateExtAuthKeyIDMissed              ParsingNotificationCodes = -1.879047653e+09
	NtfLDSICAOCertificateExtBasicCIncorrectData          ParsingNotificationCodes = -1.879047659e+09
	NtfLDSICAOCertificateExtBasicCIncorrectUsage1        ParsingNotificationCodes = -1.879047662e+09
	NtfLDSICAOCertificateExtBasicCIncorrectUsage2        ParsingNotificationCodes = -1.879047661e+09
	NtfLDSICAOCertificateExtBasicCMissed                 ParsingNotificationCodes = -1.879047663e+09
	NtfLDSICAOCertificateExtBasicCNotCritical            ParsingNotificationCodes = -1.87904766e+09
	NtfLDSICAOCertificateExtBasicCPathLenCIncorrect      ParsingNotificationCodes = -1.879047657e+09
	NtfLDSICAOCertificateExtBasicCPathLenCMissed         ParsingNotificationCodes = -1.879047658e+09
	NtfLDSICAOCertificateExtCRLDistPointEmpty            ParsingNotificationCodes = -1.879047617e+09
	NtfLDSICAOCertificateExtCRLDistPointIncorrectData    ParsingNotificationCodes = -1.879047618e+09
	NtfLDSICAOCertificateExtCRLDistPointMissed           ParsingNotificationCodes = -1.879047619e+09
	NtfLDSICAOCertificateExtCRLDistPointPointMissed      ParsingNotificationCodes = -1.879047616e+09
	NtfLDSICAOCertificateExtCSCAAltNamesNonMatching      ParsingNotificationCodes = -1.879047609e+09
	NtfLDSICAOCertificateExtCertPoliciesEmpty            ParsingNotificationCodes = -1.879047621e+09
	NtfLDSICAOCertificateExtCertPoliciesIncorrectData    ParsingNotificationCodes = -1.879047622e+09
	NtfLDSICAOCertificateExtCertPoliciesPolicyIDMissed   ParsingNotificationCodes = -1.87904762e+09
	NtfLDSICAOCertificateExtDocTypeListCritical          ParsingNotificationCodes = -1.879047604e+09
	NtfLDSICAOCertificateExtDocTypeListDocTypes          ParsingNotificationCodes = -1.879047624e+09
	NtfLDSICAOCertificateExtDocTypeListDocTypesEmpty     ParsingNotificationCodes = -1.879047623e+09
	NtfLDSICAOCertificateExtDocTypeListIncorrectData     ParsingNotificationCodes = -1.879047626e+09
	NtfLDSICAOCertificateExtDocTypeListMissed            ParsingNotificationCodes = -1.879047627e+09
	NtfLDSICAOCertificateExtDocTypeListNonCompliant      ParsingNotificationCodes = -1.879047605e+09
	NtfLDSICAOCertificateExtDocTypeListVersion           ParsingNotificationCodes = -1.879047625e+09
	NtfLDSICAOCertificateExtExtKeyUsageIncorrectData     ParsingNotificationCodes = -1.879047654e+09
	NtfLDSICAOCertificateExtExtKeyUsageIncorrectUsage    ParsingNotificationCodes = -1.879047655e+09
	NtfLDSICAOCertificateExtExtKeyUsageNotCritical       ParsingNotificationCodes = -1.879047656e+09
	NtfLDSICAOCertificateExtIssuerAltNameCritical        ParsingNotificationCodes = -1.879047631e+09
	NtfLDSICAOCertificateExtIssuerAltNameDNEmpty         ParsingNotificationCodes = -1.87904763e+09
	NtfLDSICAOCertificateExtIssuerAltNameDNIncorrect     ParsingNotificationCodes = -1.879047629e+09
	NtfLDSICAOCertificateExtIssuerAltNameDNNonCompliant  ParsingNotificationCodes = -1.879047628e+09
	NtfLDSICAOCertificateExtIssuerAltNameEmpty           ParsingNotificationCodes = -1.879047634e+09
	NtfLDSICAOCertificateExtIssuerAltNameIncorrectData   ParsingNotificationCodes = -1.879047635e+09
	NtfLDSICAOCertificateExtIssuerAltNameMissed          ParsingNotificationCodes = -1.879047636e+09
	NtfLDSICAOCertificateExtIssuerAltNameNonCompliant    ParsingNotificationCodes = -1.879047633e+09
	NtfLDSICAOCertificateExtKeyUsageIncorrectData        ParsingNotificationCodes = -1.879047664e+09
	NtfLDSICAOCertificateExtKeyUsageMissed               ParsingNotificationCodes = -1.879047666e+09
	NtfLDSICAOCertificateExtKeyUsageNotCritical          ParsingNotificationCodes = -1.879047665e+09
	NtfLDSICAOCertificateExtNameChangeCritical           ParsingNotificationCodes = -1.879047606e+09
	NtfLDSICAOCertificateExtNameChangeIncorrectData      ParsingNotificationCodes = -1.879047608e+09
	NtfLDSICAOCertificateExtNameChangeNonCompliant       ParsingNotificationCodes = -1.879047607e+09
	NtfLDSICAOCertificateExtOptionalCritical             ParsingNotificationCodes = -1.879047603e+09
	NtfLDSICAOCertificateExtPrivateKeyUPEmpty            ParsingNotificationCodes = -1.879047646e+09
	NtfLDSICAOCertificateExtPrivateKeyUPIncorrectData    ParsingNotificationCodes = -1.879047647e+09
	NtfLDSICAOCertificateExtPrivateKeyUPMissed           ParsingNotificationCodes = -1.879047648e+09
	NtfLDSICAOCertificateExtSubjectAltNameCritical       ParsingNotificationCodes = -1.87904764e+09
	NtfLDSICAOCertificateExtSubjectAltNameDNEmpty        ParsingNotificationCodes = -1.879047639e+09
	NtfLDSICAOCertificateExtSubjectAltNameDNIncorrect    ParsingNotificationCodes = -1.879047638e+09
	NtfLDSICAOCertificateExtSubjectAltNameDNNonCompliant ParsingNotificationCodes = -1.879047637e+09
	NtfLDSICAOCertificateExtSubjectAltNameEmpty          ParsingNotificationCodes = -1.879047643e+09
	NtfLDSICAOCertificateExtSubjectAltNameIncorrectData  ParsingNotificationCodes = -1.879047644e+09
	NtfLDSICAOCertificateExtSubjectAltNameMissed         ParsingNotificationCodes = -1.879047645e+09
	NtfLDSICAOCertificateExtSubjectAltNameNonCompliant   ParsingNotificationCodes = -1.879047642e+09
	NtfLDSICAOCertificateExtSubjectKeyIDIncorrectData    ParsingNotificationCodes = -1.879047649e+09
	NtfLDSICAOCertificateExtSubjectKeyIDMissed           ParsingNotificationCodes = -1.87904765e+09
	NtfLDSICAOCertificateExtUsingNonCompliantData        ParsingNotificationCodes = -1.879047667e+09
	NtfLDSICAOCertificateIssuerAttributeNonCompliant     ParsingNotificationCodes = -1.879047612e+09
	NtfLDSICAOCertificateIssuerCommonNameMissed          ParsingNotificationCodes = -1.879047677e+09
	NtfLDSICAOCertificateIssuerCountryMissed             ParsingNotificationCodes = -1.879047678e+09
	NtfLDSICAOCertificateIssuerCountryNonCompliant       ParsingNotificationCodes = -1.879047676e+09
	NtfLDSICAOCertificateIssuerSNNonCompliant            ParsingNotificationCodes = -1.879047614e+09
	NtfLDSICAOCertificateIssuerSubjectCountryNonMatching ParsingNotificationCodes = -1.87904761e+09
	NtfLDSICAOCertificateMRZCountryNonMatching           ParsingNotificationCodes = -1.879047598e+09
	NtfLDSICAOCertificateMissedExtensions                ParsingNotificationCodes = -1.879047669e+09
	NtfLDSICAOCertificateSNNonCompliant                  ParsingNotificationCodes = -1.879047615e+09
	NtfLDSICAOCertificateSubjectAttributeNonCompliant    ParsingNotificationCodes = -1.879047611e+09
	NtfLDSICAOCertificateSubjectCommonNameMissed         ParsingNotificationCodes = -1.879047674e+09
	NtfLDSICAOCertificateSubjectCommonNameNonCompliant   ParsingNotificationCodes = -1.879047601e+09
	NtfLDSICAOCertificateSubjectCountryMissed            ParsingNotificationCodes = -1.879047675e+09
	NtfLDSICAOCertificateSubjectCountryNonCompliant      ParsingNotificationCodes = -1.879047673e+09
	NtfLDSICAOCertificateSubjectNonCompliant             ParsingNotificationCodes = -1.879047602e+09
	NtfLDSICAOCertificateSubjectSNNonCompliant           ParsingNotificationCodes = -1.879047613e+09
	NtfLDSICAOCertificateUnsupportedPublicKeyAlgorithm   ParsingNotificationCodes = -1.87904767e+09
	NtfLDSICAOCertificateUnsupportedSignatureAlgorithm   ParsingNotificationCodes = -1.879047671e+09
	NtfLDSICAOCertificateUsingNonCompliantData           ParsingNotificationCodes = -1.879047672e+09
	NtfLDSICAOCertificateValidity                        ParsingNotificationCodes = -1.879047668e+09
	NtfLDSICAOCertificateVersionIncorrect                ParsingNotificationCodes = -1.879047679e+09
	NtfLDSICAOCertificateVersionMissed                   ParsingNotificationCodes = -1.87904768e+09
	NtfLDSICAOCertificateVisualMrzCountryNonMatching     ParsingNotificationCodes = -1.879047599e+09
	NtfLDSICAODeviationListVersionIncorrect              ParsingNotificationCodes = -1.879047736e+09
	NtfLDSICAOLDSObjectDGHashExtra                       ParsingNotificationCodes = -1.879047929e+09
	NtfLDSICAOLDSObjectDGHashMissing                     ParsingNotificationCodes = -1.87904793e+09
	NtfLDSICAOLDSObjectDGNumberIncorrect                 ParsingNotificationCodes = -1.879047931e+09
	NtfLDSICAOLDSObjectIncorrectContentOID               ParsingNotificationCodes = -1.879047932e+09
	NtfLDSICAOLDSObjectVersionIncorrect                  ParsingNotificationCodes = -1.879047928e+09
	NtfLDSICAOMasterListVersionIncorrect                 ParsingNotificationCodes = -1.879047744e+09
	NtfLDSICAOSignedDataCRLsIncorrectUsage               ParsingNotificationCodes = -1.879047758e+09
	NtfLDSICAOSignedDataCertificatesEmpty                ParsingNotificationCodes = -1.879047759e+09
	NtfLDSICAOSignedDataCertificatesMissed               ParsingNotificationCodes = -1.87904776e+09
	NtfLDSICAOSignedDataDigestAlgorithmsEmpty            ParsingNotificationCodes = -1.879047934e+09
	NtfLDSICAOSignedDataDigestAlgorithmsUnsupported      ParsingNotificationCodes = -1.879047933e+09
	NtfLDSICAOSignedDataSignerInfosMultipleEntries       ParsingNotificationCodes = -1.879047927e+09
	NtfLDSICAOSignedDataVersionIncorrect                 ParsingNotificationCodes = -1.879047935e+09
	NtfLDSMRZCountryCodeVisualMrzNonMatching             ParsingNotificationCodes = 139289
	NtfLDSMRZDOBError                                    ParsingNotificationCodes = 139280
	NtfLDSMRZDOBIncorrectChecksum                        ParsingNotificationCodes = 139281
	NtfLDSMRZDOBSyntaxError                              ParsingNotificationCodes = 139279
	NtfLDSMRZDOEError                                    ParsingNotificationCodes = 139284
	NtfLDSMRZDOEIncorrectChecksum                        ParsingNotificationCodes = 139285
	NtfLDSMRZDOESyntaxError                              ParsingNotificationCodes = 139283
	NtfLDSMRZDocumentTypeUnknown                         ParsingNotificationCodes = 139272
	NtfLDSMRZIncorrect                                   ParsingNotificationCodes = 139288
	NtfLDSMRZIncorrectChecksum                           ParsingNotificationCodes = 139287
	NtfLDSMRZIssuingStateSyntaxError                     ParsingNotificationCodes = 139273
	NtfLDSMRZNameIsVoid                                  ParsingNotificationCodes = 139274
	NtfLDSMRZNationalitySyntaxError                      ParsingNotificationCodes = 139278
	NtfLDSMRZNumberIncorrectChecksum                     ParsingNotificationCodes = 139277
	NtfLDSMRZOptionalDataIncorrectChecksum               ParsingNotificationCodes = 139286
	NtfLDSMRZSexIncorrect                                ParsingNotificationCodes = 139282
	NtfLDSSIAAInfoInconsistentAlgorithmReference         ParsingNotificationCodes = -1.862270962e+09
	NtfLDSSIAAInfoIncorrectVersion                       ParsingNotificationCodes = -1.862270964e+09
	NtfLDSSIAAInfoUnsupportedAlgorithm                   ParsingNotificationCodes = -1.862270963e+09
	NtfLDSSICADomainParamsUnsupportedAlgorithm           ParsingNotificationCodes = -1.86227097e+09
	NtfLDSSICAInfoIncorrectVersion                       ParsingNotificationCodes = -1.862270972e+09
	NtfLDSSICAPublicKeyUnsupportedAlgorithm              ParsingNotificationCodes = -1.862270971e+09
	NtfLDSSIEIDSecurityUnsupportedDigestAlgorithm        ParsingNotificationCodes = -1.862270967e+09
	NtfLDSSIPACEDomainParamsUnsupportedAlgorithm         ParsingNotificationCodes = -1.862270973e+09
	NtfLDSSIPACEDomainParamsUsingStdRef                  ParsingNotificationCodes = -1.862270974e+09
	NtfLDSSIPACEInfoDeprecatedVersion                    ParsingNotificationCodes = -1.862270975e+09
	NtfLDSSIPACEInfoUnsupportedStdParameters             ParsingNotificationCodes = -1.862270976e+09
	NtfLDSSIRIDomainParamsUnsupportedAlgorithm           ParsingNotificationCodes = -1.862270965e+09
	NtfLDSSIRIInfoIncorrectVersion                       ParsingNotificationCodes = -1.862270966e+09
	NtfLDSSIStorageCAAnonymousInfos                      ParsingNotificationCodes = -1.862270714e+09
	NtfLDSSIStorageCADomainParamsNoRequiredOption        ParsingNotificationCodes = -1.862270716e+09
	NtfLDSSIStorageCADomainParamsNotAvailable            ParsingNotificationCodes = -1.862270715e+09
	NtfLDSSIStorageCAIncorrectInfosQuantity              ParsingNotificationCodes = -1.862270711e+09
	NtfLDSSIStorageCAInfoNoMatchingDomainParams          ParsingNotificationCodes = -1.862270713e+09
	NtfLDSSIStorageCAInfoNoMatchingPublicKey             ParsingNotificationCodes = -1.862270712e+09
	NtfLDSSIStorageCAInfoNotAvailable                    ParsingNotificationCodes = -1.862270717e+09
	NtfLDSSIStorageCardInfoLocatorMultipleEntries        ParsingNotificationCodes = -1.862270709e+09
	NtfLDSSIStorageEIDSecurityInfoMultipleEntries        ParsingNotificationCodes = -1.862270708e+09
	NtfLDSSIStoragePACEInfoNoMatchingDomainParams        ParsingNotificationCodes = -1.862270718e+09
	NtfLDSSIStoragePACEInfoNoStdParameters               ParsingNotificationCodes = -1.862270719e+09
	NtfLDSSIStoragePACEInfoNotAvailable                  ParsingNotificationCodes = -1.86227072e+09
	NtfLDSSIStoragePACEInfosNonConsistant                ParsingNotificationCodes = -1.862270704e+09
	NtfLDSSIStoragePrivilegedTIIncorrectUsage            ParsingNotificationCodes = -1.862270706e+09
	NtfLDSSIStoragePrivilegedTIMultipleEntries           ParsingNotificationCodes = -1.862270707e+09
	NtfLDSSIStorageRIDomainParamsMultipleEntries         ParsingNotificationCodes = -1.862270705e+09
	NtfLDSSIStorageTAInfoNotAvailable                    ParsingNotificationCodes = -1.86227071e+09
	NtfLDSSITAInfoFileIDForVersion2                      ParsingNotificationCodes = -1.862270968e+09
	NtfLDSSITAInfoIncorrectVersion                       ParsingNotificationCodes = -1.862270969e+09
	NtfLDSTAPACEStaticBindingUsed                        ParsingNotificationCodes = -1.862270208e+09
	NtfLDSUnsupportedImageFormat                         ParsingNotificationCodes = -1.87904791e+09
)

// Defines values for ProcessingStatus.
const (
	FINISHED    ProcessingStatus = 1
	NOTFINISHED ProcessingStatus = 0
	TIMEOUT     ProcessingStatus = 2
)

// Defines values for Result.
const (
	ResultAUTHENTICITY           Result = 20
	ResultBARCODEGRAPHICS        Result = 19
	ResultBARCODES               Result = 5
	ResultBARCODETEXT            Result = 18
	ResultDOCUMENTIMAGE          Result = 1
	ResultDOCUMENTPOSITION       Result = 85
	ResultDOCUMENTTYPE           Result = 9
	ResultDOCUMENTTYPECANDIDATES Result = 8
	ResultENCRYPTEDRCL           Result = 49
	ResultFINGERPRINTCOMPARISON  Result = 39
	ResultIMAGEQUALITY           Result = 30
	ResultIMAGES                 Result = 37
	ResultLEXICALANALYSIS        Result = 15
	ResultLICENSE                Result = 50
	ResultMRZTEXT                Result = 3
	ResultPORTRAITCOMPARISON     Result = 34
	ResultRFIDBINARYDATA         Result = 104
	ResultRFIDGRAPHICS           Result = 103
	ResultRFIDORIGINALGRAPHICS   Result = 105
	ResultRFIDRAWDATA            Result = 101
	ResultRFIDTEXT               Result = 102
	ResultSTATUS                 Result = 33
	ResultTEXT                   Result = 36
	ResultVISUALGRAPHICS         Result = 6
	ResultVISUALTEXT             Result = 17
)

// Defines values for RfidLocation.
const (
	BACKPAGE RfidLocation = 2
	MAINPAGE RfidLocation = 1
	NONE     RfidLocation = 0
)

// Defines values for Scenario.
const (
	ScenarioBARCODE                 Scenario = "Barcode"
	ScenarioBARCODEANDLOCATE        Scenario = "BarcodeAndLocate"
	ScenarioCAPTURE                 Scenario = "Capture"
	ScenarioCREDITCARD              Scenario = "CreditCard"
	ScenarioDOCTYPE                 Scenario = "DocType"
	ScenarioFULLAUTH                Scenario = "FullAuth"
	ScenarioFULLPROCESS             Scenario = "FullProcess"
	ScenarioLOCATE                  Scenario = "Locate"
	ScenarioLOCATEVISUALANDMRZOROCR Scenario = "LocateVisual_And_MrzOrOcr"
	ScenarioMRZ                     Scenario = "Mrz"
	ScenarioMRZANDLOCATE            Scenario = "MrzAndLocate"
	ScenarioMRZORBARCODE            Scenario = "MrzOrBarcode"
	ScenarioMRZORBARCODEOROCR       Scenario = "MrzOrBarcodeOrOcr"
	ScenarioMRZORLOCATE             Scenario = "MrzOrLocate"
	ScenarioMRZOROCR                Scenario = "MrzOrOcr"
	ScenarioOCR                     Scenario = "Ocr"
	ScenarioOCRFREE                 Scenario = "OcrFree"
	ScenarioRUSSTAMP                Scenario = "RusStamp"
)

// Defines values for SecurityFeatureType.
const (
	SecurityFeatureTypeBARCODE                           SecurityFeatureType = 17
	SecurityFeatureTypeBARCODESIZECHECK                  SecurityFeatureType = 42
	SecurityFeatureTypeBLANK                             SecurityFeatureType = 0
	SecurityFeatureTypeCHECKDIGITALSIGNATURE             SecurityFeatureType = 50
	SecurityFeatureTypeCLEARGHOSTPHOTO                   SecurityFeatureType = 22
	SecurityFeatureTypeCONTACTCHIPCLASSIFICATION         SecurityFeatureType = 51
	SecurityFeatureTypeFACEABSENCE                       SecurityFeatureType = 38
	SecurityFeatureTypeFACEPRESENCE                      SecurityFeatureType = 36
	SecurityFeatureTypeFALSELUMINESCENCE                 SecurityFeatureType = 4
	SecurityFeatureTypeFILL                              SecurityFeatureType = 1
	SecurityFeatureTypeFLUORESCENTOBJECT                 SecurityFeatureType = 34
	SecurityFeatureTypeGHOSTPHOTO                        SecurityFeatureType = 21
	SecurityFeatureTypeHOLOSIMPLE                        SecurityFeatureType = 5
	SecurityFeatureTypeHOLOVERIFYDYNAMIC                 SecurityFeatureType = 8
	SecurityFeatureTypeHOLOVERIFYMULTISTATIC             SecurityFeatureType = 7
	SecurityFeatureTypeHOLOVERIFYSTATIC                  SecurityFeatureType = 6
	SecurityFeatureTypeINVISIBLEOBJECT                   SecurityFeatureType = 23
	SecurityFeatureTypeLANDMARKCHECK                     SecurityFeatureType = 35
	SecurityFeatureTypeLASINK                            SecurityFeatureType = 43
	SecurityFeatureTypeLIVENESSBARCODEBACKGROUND         SecurityFeatureType = 45
	SecurityFeatureTypeLIVENESSDEPTH                     SecurityFeatureType = 32
	SecurityFeatureTypeLIVENESSELECTRONICDEVICE          SecurityFeatureType = 40
	SecurityFeatureTypeLIVENESSMLI                       SecurityFeatureType = 44
	SecurityFeatureTypeLIVENESSOVI                       SecurityFeatureType = 41
	SecurityFeatureTypeLIVENESSSCREENCAPTURE             SecurityFeatureType = 39
	SecurityFeatureTypeLOWCONTRASTOBJECT                 SecurityFeatureType = 24
	SecurityFeatureTypeMICROTEXT                         SecurityFeatureType = 33
	SecurityFeatureTypeMRZ                               SecurityFeatureType = 3
	SecurityFeatureTypeOCR                               SecurityFeatureType = 28
	SecurityFeatureTypePATTERNDIFFERENTLINESTHICKNESS    SecurityFeatureType = 18
	SecurityFeatureTypePATTERNIRINVISIBLE                SecurityFeatureType = 12
	SecurityFeatureTypePATTERNNOTINTERRUPTED             SecurityFeatureType = 9
	SecurityFeatureTypePATTERNNOTSHIFTED                 SecurityFeatureType = 10
	SecurityFeatureTypePATTERNSAMECOLORS                 SecurityFeatureType = 11
	SecurityFeatureTypePHOTO                             SecurityFeatureType = 2
	SecurityFeatureTypePHOTOCOLOR                        SecurityFeatureType = 25
	SecurityFeatureTypePHOTOCORNERS                      SecurityFeatureType = 27
	SecurityFeatureTypePHOTOSHAPE                        SecurityFeatureType = 26
	SecurityFeatureTypePHOTOSIZECHECK                    SecurityFeatureType = 13
	SecurityFeatureTypePORTRAITCOMPARISONBARCODEVSCAMERA SecurityFeatureType = 49
	SecurityFeatureTypePORTRAITCOMPARISONEXTVSBARCODE    SecurityFeatureType = 48
	SecurityFeatureTypePORTRAITCOMPARISONEXTVSCAMERA     SecurityFeatureType = 31
	SecurityFeatureTypePORTRAITCOMPARISONEXTVSRFID       SecurityFeatureType = 30
	SecurityFeatureTypePORTRAITCOMPARISONEXTVSVISUAL     SecurityFeatureType = 29
	SecurityFeatureTypePORTRAITCOMPARISONRFIDVSBARCODE   SecurityFeatureType = 47
	SecurityFeatureTypePORTRAITCOMPARISONRFIDVSCAMERA    SecurityFeatureType = 20
	SecurityFeatureTypePORTRAITCOMPARISONVSBARCODE       SecurityFeatureType = 46
	SecurityFeatureTypePORTRAITCOMPARISONVSCAMERA        SecurityFeatureType = 19
	SecurityFeatureTypePORTRAITCOMPARISONVSGHOST         SecurityFeatureType = 14
	SecurityFeatureTypePORTRAITCOMPARISONVSRFID          SecurityFeatureType = 15
	SecurityFeatureTypePORTRAITCOMPARISONVSVISUAL        SecurityFeatureType = 16
)

// Defines values for Source.
const (
	BARCODE  Source = "BARCODE"
	MAGNETIC Source = "MAGNETIC"
	MRZ      Source = "MRZ"
	RFID     Source = "RFID"
	VISUAL   Source = "VISUAL"
)

// Defines values for TextFieldType.
const (
	TextFieldTypeACADEMICTITLE                         TextFieldType = 255
	TextFieldTypeACCOMPANIEDBY                         TextFieldType = 444
	TextFieldTypeADDRESS                               TextFieldType = 17
	TextFieldTypeADDRESSAREA                           TextFieldType = 64
	TextFieldTypeADDRESSBLOCKNUMBER                    TextFieldType = 674
	TextFieldTypeADDRESSBUILDING                       TextFieldType = 66
	TextFieldTypeADDRESSBUILDINGTYPE                   TextFieldType = 680
	TextFieldTypeADDRESSCITY                           TextFieldType = 77
	TextFieldTypeADDRESSCITYSECTOR                     TextFieldType = 677
	TextFieldTypeADDRESSCITYTYPE                       TextFieldType = 679
	TextFieldTypeADDRESSCOUNTRY                        TextFieldType = 256
	TextFieldTypeADDRESSCOUNTYTYPE                     TextFieldType = 678
	TextFieldTypeADDRESSENTRANCE                       TextFieldType = 673
	TextFieldTypeADDRESSFLAT                           TextFieldType = 68
	TextFieldTypeADDRESSFLOORNUMBER                    TextFieldType = 672
	TextFieldTypeADDRESSHOUSE                          TextFieldType = 67
	TextFieldTypeADDRESSJURISDICTIONCODE               TextFieldType = 78
	TextFieldTypeADDRESSLOCATION                       TextFieldType = 189
	TextFieldTypeADDRESSMUNICIPALITY                   TextFieldType = 188
	TextFieldTypeADDRESSPOSTALCODE                     TextFieldType = 79
	TextFieldTypeADDRESSSTATE                          TextFieldType = 65
	TextFieldTypeADDRESSSTREET                         TextFieldType = 76
	TextFieldTypeADDRESSSTREETNUMBER                   TextFieldType = 675
	TextFieldTypeADDRESSSTREETTYPE                     TextFieldType = 676
	TextFieldTypeADDRESSZIPCODE                        TextFieldType = 257
	TextFieldTypeADMINISTRATIVENUMBER                  TextFieldType = 333
	TextFieldTypeAGE                                   TextFieldType = 185
	TextFieldTypeAGEATISSUE                            TextFieldType = 522
	TextFieldTypeAGY                                   TextFieldType = 351
	TextFieldTypeAIRLINEDESIGNATOROFBOARDINGPASSISSUER TextFieldType = 368
	TextFieldTypeAIRLINENAME                           TextFieldType = 457
	TextFieldTypeAIRLINENAMEFREQUENTFLYER              TextFieldType = 458
	TextFieldTypeAIRLINENUMERICCODE                    TextFieldType = 369
	TextFieldTypeAIRPORTFROM                           TextFieldType = 455
	TextFieldTypeAIRPORTTO                             TextFieldType = 456
	TextFieldTypeAKADATEOFBIRTH                        TextFieldType = 110
	TextFieldTypeAKAGIVENNAMES                         TextFieldType = 113
	TextFieldTypeAKANAMEPREFIX                         TextFieldType = 115
	TextFieldTypeAKANAMESUFFIX                         TextFieldType = 114
	TextFieldTypeAKASOCIALSECURITYNUMBER               TextFieldType = 111
	TextFieldTypeAKASURNAME                            TextFieldType = 112
	TextFieldTypeAKASURNAMEANDGIVENNAMES               TextFieldType = 170
	TextFieldTypeALLERGIES                             TextFieldType = 343
	TextFieldTypeALTCODE                               TextFieldType = 295
	TextFieldTypeALTDATEOFEXPIRY                       TextFieldType = 637
	TextFieldTypeAPPLICATIONNUMBER                     TextFieldType = 514
	TextFieldTypeARTISTICNAME                          TextFieldType = 254
	TextFieldTypeAUDITINFORMATION                      TextFieldType = 120
	TextFieldTypeAUTHORITY                             TextFieldType = 24
	TextFieldTypeAUTHORITYCODE                         TextFieldType = 73
	TextFieldTypeAUTHORITYRUS                          TextFieldType = 133
	TextFieldTypeAUTHORIZATIONNUMBER                   TextFieldType = 516
	TextFieldTypeBANKCARDNUMBER                        TextFieldType = 493
	TextFieldTypeBANKCARDVALIDTHRU                     TextFieldType = 494
	TextFieldTypeBANKNOTENUMBER                        TextFieldType = 252
	TextFieldTypeBDBTYPE                               TextFieldType = 305
	TextFieldTypeBENEFITSNUMBER                        TextFieldType = 280
	TextFieldTypeBINARYCODE                            TextFieldType = 296
	TextFieldTypeBIOMETRICFORMATOWNER                  TextFieldType = 309
	TextFieldTypeBIOMETRICFORMATTYPE                   TextFieldType = 310
	TextFieldTypeBIOMETRICPRODUCTID                    TextFieldType = 308
	TextFieldTypeBIOMETRICSUBTYPE                      TextFieldType = 307
	TextFieldTypeBIOMETRICTYPE                         TextFieldType = 306
	TextFieldTypeBLOODGROUP                            TextFieldType = 180
	TextFieldTypeBOOKLETNUMBER                         TextFieldType = 165
	TextFieldTypeCALIBER                               TextFieldType = 585
	TextFieldTypeCARDACCESSNUMBER                      TextFieldType = 159
	TextFieldTypeCATEGORY                              TextFieldType = 286
	TextFieldTypeCCWUNTIL                              TextFieldType = 359
	TextFieldTypeCDLCLASS                              TextFieldType = 265
	TextFieldTypeCENTURYDATEOFBIRTH                    TextFieldType = 468
	TextFieldTypeCHASSISNUMBER                         TextFieldType = 138
	TextFieldTypeCHECKINSEQUENCENUMBER                 TextFieldType = 367
	TextFieldTypeCHILDREN                              TextFieldType = 166
	TextFieldTypeCITIZENSHIPOFFIRSTPERSON              TextFieldType = 593
	TextFieldTypeCITIZENSHIPOFSECONDPERSON             TextFieldType = 594
	TextFieldTypeCITIZENSHIPSTATUS                     TextFieldType = 625
	TextFieldTypeCIVILSTATUS                           TextFieldType = 433
	TextFieldTypeCOMMERCIALVEHICLECODES                TextFieldType = 109
	TextFieldTypeCOMPANYNAME                           TextFieldType = 161
	TextFieldTypeCOMPARTMENTCODE                       TextFieldType = 366
	TextFieldTypeCOMPLEXION                            TextFieldType = 454
	TextFieldTypeCOMPLIANCETYPE                        TextFieldType = 271
	TextFieldTypeCONDITIONS                            TextFieldType = 287
	TextFieldTypeCONFIGURATION                         TextFieldType = 289
	TextFieldTypeCONTROLNUMBER                         TextFieldType = 143
	TextFieldTypeCOPY                                  TextFieldType = 167
	TextFieldTypeCOURTCODE                             TextFieldType = 345
	TextFieldTypeCSCCODE                               TextFieldType = 253
	TextFieldTypeCTY                                   TextFieldType = 346
	TextFieldTypeCURRENTDATE                           TextFieldType = 221
	TextFieldTypeCUSTODYINFO                           TextFieldType = 316
	TextFieldTypeCVV                                   TextFieldType = 595
	TextFieldTypeDATADISCRIMINATOR                     TextFieldType = 335
	TextFieldTypeDATEFIRSTRENEWAL                      TextFieldType = 511
	TextFieldTypeDATEOFARRIVAL                         TextFieldType = 646
	TextFieldTypeDATEOFBIRTH                           TextFieldType = 5
	TextFieldTypeDATEOFBIRTHCHECKDIGIT                 TextFieldType = 81
	TextFieldTypeDATEOFBIRTHCHECKSUM                   TextFieldType = 41
	TextFieldTypeDATEOFBIRTHOFHUSBAND                  TextFieldType = 592
	TextFieldTypeDATEOFBIRTHOFWIFE                     TextFieldType = 591
	TextFieldTypeDATEOFCREATION                        TextFieldType = 302
	TextFieldTypeDATEOFEXPIRY                          TextFieldType = 3
	TextFieldTypeDATEOFEXPIRYCHECKDIGIT                TextFieldType = 82
	TextFieldTypeDATEOFEXPIRYCHECKSUM                  TextFieldType = 42
	TextFieldTypeDATEOFFLIGHT                          TextFieldType = 356
	TextFieldTypeDATEOFINSURANCEEXPIRY                 TextFieldType = 596
	TextFieldTypeDATEOFISSUE                           TextFieldType = 4
	TextFieldTypeDATEOFISSUEBOARDINGPASS               TextFieldType = 358
	TextFieldTypeDATEOFISSUECHECKDIGIT                 TextFieldType = 55
	TextFieldTypeDATEOFISSUECHECKSUM                   TextFieldType = 54
	TextFieldTypeDATEOFPERSONALIZATION                 TextFieldType = 320
	TextFieldTypeDATEOFREGISTRATION                    TextFieldType = 70
	TextFieldTypeDATEOFRETIREMENT                      TextFieldType = 681
	TextFieldTypeDATESECONDRENEWAL                     TextFieldType = 512
	TextFieldTypeDAY                                   TextFieldType = 105
	TextFieldTypeDEPARTMENT                            TextFieldType = 277
	TextFieldTypeDEPTNUMBER                            TextFieldType = 341
	TextFieldTypeDESTINATION                           TextFieldType = 179
	TextFieldTypeDISCRETIONARYDATA                     TextFieldType = 290
	TextFieldTypeDLCDLRESTRICTIONCODE                  TextFieldType = 173
	TextFieldTypeDLCLASS                               TextFieldType = 20
	TextFieldTypeDLCLASSCODEA1FROM                     TextFieldType = 378
	TextFieldTypeDLCLASSCODEA1NOTES                    TextFieldType = 380
	TextFieldTypeDLCLASSCODEA1TO                       TextFieldType = 379
	TextFieldTypeDLCLASSCODEA2FROM                     TextFieldType = 426
	TextFieldTypeDLCLASSCODEA2NOTES                    TextFieldType = 428
	TextFieldTypeDLCLASSCODEA2TO                       TextFieldType = 427
	TextFieldTypeDLCLASSCODEA3FROM                     TextFieldType = 469
	TextFieldTypeDLCLASSCODEA3NOTES                    TextFieldType = 471
	TextFieldTypeDLCLASSCODEA3TO                       TextFieldType = 470
	TextFieldTypeDLCLASSCODEAFROM                      TextFieldType = 381
	TextFieldTypeDLCLASSCODEAMFROM                     TextFieldType = 423
	TextFieldTypeDLCLASSCODEAMNOTES                    TextFieldType = 425
	TextFieldTypeDLCLASSCODEAMTO                       TextFieldType = 424
	TextFieldTypeDLCLASSCODEANOTES                     TextFieldType = 383
	TextFieldTypeDLCLASSCODEATO                        TextFieldType = 382
	TextFieldTypeDLCLASSCODEB1FROM                     TextFieldType = 429
	TextFieldTypeDLCLASSCODEB1NOTES                    TextFieldType = 431
	TextFieldTypeDLCLASSCODEB1TO                       TextFieldType = 430
	TextFieldTypeDLCLASSCODEB2EFROM                    TextFieldType = 481
	TextFieldTypeDLCLASSCODEB2ENOTES                   TextFieldType = 483
	TextFieldTypeDLCLASSCODEB2ETO                      TextFieldType = 482
	TextFieldTypeDLCLASSCODEB2FROM                     TextFieldType = 475
	TextFieldTypeDLCLASSCODEB2NOTES                    TextFieldType = 477
	TextFieldTypeDLCLASSCODEB2TO                       TextFieldType = 476
	TextFieldTypeDLCLASSCODEBEFROM                     TextFieldType = 399
	TextFieldTypeDLCLASSCODEBENOTES                    TextFieldType = 401
	TextFieldTypeDLCLASSCODEBETO                       TextFieldType = 400
	TextFieldTypeDLCLASSCODEBFROM                      TextFieldType = 384
	TextFieldTypeDLCLASSCODEBNOTES                     TextFieldType = 386
	TextFieldTypeDLCLASSCODEBTO                        TextFieldType = 385
	TextFieldTypeDLCLASSCODEBTPFROM                    TextFieldType = 524
	TextFieldTypeDLCLASSCODEBTPNOTES                   TextFieldType = 525
	TextFieldTypeDLCLASSCODEBTPTO                      TextFieldType = 526
	TextFieldTypeDLCLASSCODEC1EFROM                    TextFieldType = 402
	TextFieldTypeDLCLASSCODEC1ENOTES                   TextFieldType = 404
	TextFieldTypeDLCLASSCODEC1ETO                      TextFieldType = 403
	TextFieldTypeDLCLASSCODEC1FROM                     TextFieldType = 387
	TextFieldTypeDLCLASSCODEC1NOTES                    TextFieldType = 389
	TextFieldTypeDLCLASSCODEC1TO                       TextFieldType = 388
	TextFieldTypeDLCLASSCODEC2FROM                     TextFieldType = 472
	TextFieldTypeDLCLASSCODEC2NOTES                    TextFieldType = 474
	TextFieldTypeDLCLASSCODEC2TO                       TextFieldType = 473
	TextFieldTypeDLCLASSCODEC3FROM                     TextFieldType = 527
	TextFieldTypeDLCLASSCODEC3NOTES                    TextFieldType = 528
	TextFieldTypeDLCLASSCODEC3TO                       TextFieldType = 529
	TextFieldTypeDLCLASSCODECAFROM                     TextFieldType = 622
	TextFieldTypeDLCLASSCODECANOTES                    TextFieldType = 624
	TextFieldTypeDLCLASSCODECATO                       TextFieldType = 623
	TextFieldTypeDLCLASSCODECDFROM                     TextFieldType = 638
	TextFieldTypeDLCLASSCODECDNOTES                    TextFieldType = 640
	TextFieldTypeDLCLASSCODECDTO                       TextFieldType = 639
	TextFieldTypeDLCLASSCODECEFROM                     TextFieldType = 405
	TextFieldTypeDLCLASSCODECENOTES                    TextFieldType = 407
	TextFieldTypeDLCLASSCODECETO                       TextFieldType = 406
	TextFieldTypeDLCLASSCODECFROM                      TextFieldType = 390
	TextFieldTypeDLCLASSCODECNOTES                     TextFieldType = 392
	TextFieldTypeDLCLASSCODECTO                        TextFieldType = 391
	TextFieldTypeDLCLASSCODED1EFROM                    TextFieldType = 408
	TextFieldTypeDLCLASSCODED1ENOTES                   TextFieldType = 410
	TextFieldTypeDLCLASSCODED1ETO                      TextFieldType = 409
	TextFieldTypeDLCLASSCODED1FROM                     TextFieldType = 393
	TextFieldTypeDLCLASSCODED1NOTES                    TextFieldType = 395
	TextFieldTypeDLCLASSCODED1TO                       TextFieldType = 394
	TextFieldTypeDLCLASSCODED2FROM                     TextFieldType = 478
	TextFieldTypeDLCLASSCODED2NOTES                    TextFieldType = 480
	TextFieldTypeDLCLASSCODED2TO                       TextFieldType = 479
	TextFieldTypeDLCLASSCODED3FROM                     TextFieldType = 634
	TextFieldTypeDLCLASSCODED3NOTES                    TextFieldType = 636
	TextFieldTypeDLCLASSCODED3TO                       TextFieldType = 635
	TextFieldTypeDLCLASSCODEDEFROM                     TextFieldType = 411
	TextFieldTypeDLCLASSCODEDENOTES                    TextFieldType = 413
	TextFieldTypeDLCLASSCODEDETO                       TextFieldType = 412
	TextFieldTypeDLCLASSCODEDFROM                      TextFieldType = 396
	TextFieldTypeDLCLASSCODEDNOTES                     TextFieldType = 398
	TextFieldTypeDLCLASSCODEDTO                        TextFieldType = 397
	TextFieldTypeDLCLASSCODEEBFROM                     TextFieldType = 657
	TextFieldTypeDLCLASSCODEEBNOTES                    TextFieldType = 658
	TextFieldTypeDLCLASSCODEEBTO                       TextFieldType = 659
	TextFieldTypeDLCLASSCODEEC1FROM                    TextFieldType = 663
	TextFieldTypeDLCLASSCODEEC1NOTES                   TextFieldType = 664
	TextFieldTypeDLCLASSCODEEC1TO                      TextFieldType = 665
	TextFieldTypeDLCLASSCODEECFROM                     TextFieldType = 660
	TextFieldTypeDLCLASSCODEECNOTES                    TextFieldType = 661
	TextFieldTypeDLCLASSCODEECTO                       TextFieldType = 662
	TextFieldTypeDLCLASSCODEEFROM                      TextFieldType = 530
	TextFieldTypeDLCLASSCODEENOTES                     TextFieldType = 531
	TextFieldTypeDLCLASSCODEETO                        TextFieldType = 532
	TextFieldTypeDLCLASSCODEFA1FROM                    TextFieldType = 539
	TextFieldTypeDLCLASSCODEFA1NOTES                   TextFieldType = 540
	TextFieldTypeDLCLASSCODEFA1TO                      TextFieldType = 541
	TextFieldTypeDLCLASSCODEFAFROM                     TextFieldType = 536
	TextFieldTypeDLCLASSCODEFANOTES                    TextFieldType = 537
	TextFieldTypeDLCLASSCODEFATO                       TextFieldType = 538
	TextFieldTypeDLCLASSCODEFBFROM                     TextFieldType = 542
	TextFieldTypeDLCLASSCODEFBNOTES                    TextFieldType = 543
	TextFieldTypeDLCLASSCODEFBTO                       TextFieldType = 544
	TextFieldTypeDLCLASSCODEFFROM                      TextFieldType = 533
	TextFieldTypeDLCLASSCODEFNOTES                     TextFieldType = 534
	TextFieldTypeDLCLASSCODEFTO                        TextFieldType = 535
	TextFieldTypeDLCLASSCODEG1FROM                     TextFieldType = 545
	TextFieldTypeDLCLASSCODEG1NOTES                    TextFieldType = 546
	TextFieldTypeDLCLASSCODEG1TO                       TextFieldType = 547
	TextFieldTypeDLCLASSCODEGFROM                      TextFieldType = 484
	TextFieldTypeDLCLASSCODEGNOTES                     TextFieldType = 486
	TextFieldTypeDLCLASSCODEGTO                        TextFieldType = 485
	TextFieldTypeDLCLASSCODEHCFROM                     TextFieldType = 610
	TextFieldTypeDLCLASSCODEHCNOTES                    TextFieldType = 612
	TextFieldTypeDLCLASSCODEHCTO                       TextFieldType = 611
	TextFieldTypeDLCLASSCODEHFROM                      TextFieldType = 548
	TextFieldTypeDLCLASSCODEHNOTES                     TextFieldType = 549
	TextFieldTypeDLCLASSCODEHRFROM                     TextFieldType = 607
	TextFieldTypeDLCLASSCODEHRNOTES                    TextFieldType = 609
	TextFieldTypeDLCLASSCODEHRTO                       TextFieldType = 608
	TextFieldTypeDLCLASSCODEHTO                        TextFieldType = 550
	TextFieldTypeDLCLASSCODEIFROM                      TextFieldType = 551
	TextFieldTypeDLCLASSCODEINOTES                     TextFieldType = 552
	TextFieldTypeDLCLASSCODEITO                        TextFieldType = 553
	TextFieldTypeDLCLASSCODEJFROM                      TextFieldType = 487
	TextFieldTypeDLCLASSCODEJNOTES                     TextFieldType = 489
	TextFieldTypeDLCLASSCODEJTO                        TextFieldType = 488
	TextFieldTypeDLCLASSCODEKFROM                      TextFieldType = 554
	TextFieldTypeDLCLASSCODEKNOTES                     TextFieldType = 555
	TextFieldTypeDLCLASSCODEKTO                        TextFieldType = 556
	TextFieldTypeDLCLASSCODELCFROM                     TextFieldType = 490
	TextFieldTypeDLCLASSCODELCNOTES                    TextFieldType = 492
	TextFieldTypeDLCLASSCODELCTO                       TextFieldType = 491
	TextFieldTypeDLCLASSCODELFROM                      TextFieldType = 417
	TextFieldTypeDLCLASSCODELKFROM                     TextFieldType = 557
	TextFieldTypeDLCLASSCODELKNOTES                    TextFieldType = 558
	TextFieldTypeDLCLASSCODELKTO                       TextFieldType = 559
	TextFieldTypeDLCLASSCODELNOTES                     TextFieldType = 419
	TextFieldTypeDLCLASSCODELRFROM                     TextFieldType = 601
	TextFieldTypeDLCLASSCODELRNOTES                    TextFieldType = 603
	TextFieldTypeDLCLASSCODELRTO                       TextFieldType = 602
	TextFieldTypeDLCLASSCODELTO                        TextFieldType = 418
	TextFieldTypeDLCLASSCODEMCFROM                     TextFieldType = 613
	TextFieldTypeDLCLASSCODEMCNOTES                    TextFieldType = 615
	TextFieldTypeDLCLASSCODEMCTO                       TextFieldType = 614
	TextFieldTypeDLCLASSCODEMFROM                      TextFieldType = 414
	TextFieldTypeDLCLASSCODEMNOTES                     TextFieldType = 416
	TextFieldTypeDLCLASSCODEMRFROM                     TextFieldType = 604
	TextFieldTypeDLCLASSCODEMRNOTES                    TextFieldType = 606
	TextFieldTypeDLCLASSCODEMRTO                       TextFieldType = 605
	TextFieldTypeDLCLASSCODEMTO                        TextFieldType = 415
	TextFieldTypeDLCLASSCODENFROM                      TextFieldType = 560
	TextFieldTypeDLCLASSCODENNOTES                     TextFieldType = 561
	TextFieldTypeDLCLASSCODENTFROM                     TextFieldType = 628
	TextFieldTypeDLCLASSCODENTNOTES                    TextFieldType = 630
	TextFieldTypeDLCLASSCODENTO                        TextFieldType = 562
	TextFieldTypeDLCLASSCODENTTO                       TextFieldType = 629
	TextFieldTypeDLCLASSCODEPWFROM                     TextFieldType = 654
	TextFieldTypeDLCLASSCODEPWNOTES                    TextFieldType = 655
	TextFieldTypeDLCLASSCODEPWTO                       TextFieldType = 656
	TextFieldTypeDLCLASSCODEREFROM                     TextFieldType = 616
	TextFieldTypeDLCLASSCODERENOTES                    TextFieldType = 618
	TextFieldTypeDLCLASSCODERETO                       TextFieldType = 617
	TextFieldTypeDLCLASSCODERFROM                      TextFieldType = 619
	TextFieldTypeDLCLASSCODERMFROM                     TextFieldType = 651
	TextFieldTypeDLCLASSCODERMNOTES                    TextFieldType = 652
	TextFieldTypeDLCLASSCODERMTO                       TextFieldType = 653
	TextFieldTypeDLCLASSCODERNOTES                     TextFieldType = 621
	TextFieldTypeDLCLASSCODERTO                        TextFieldType = 620
	TextFieldTypeDLCLASSCODESFROM                      TextFieldType = 563
	TextFieldTypeDLCLASSCODESNOTES                     TextFieldType = 564
	TextFieldTypeDLCLASSCODESTO                        TextFieldType = 565
	TextFieldTypeDLCLASSCODETBFROM                     TextFieldType = 566
	TextFieldTypeDLCLASSCODETBNOTES                    TextFieldType = 567
	TextFieldTypeDLCLASSCODETBTO                       TextFieldType = 568
	TextFieldTypeDLCLASSCODETFROM                      TextFieldType = 420
	TextFieldTypeDLCLASSCODETMFROM                     TextFieldType = 569
	TextFieldTypeDLCLASSCODETMNOTES                    TextFieldType = 570
	TextFieldTypeDLCLASSCODETMTO                       TextFieldType = 571
	TextFieldTypeDLCLASSCODETNFROM                     TextFieldType = 631
	TextFieldTypeDLCLASSCODETNNOTES                    TextFieldType = 633
	TextFieldTypeDLCLASSCODETNOTES                     TextFieldType = 422
	TextFieldTypeDLCLASSCODETNTO                       TextFieldType = 632
	TextFieldTypeDLCLASSCODETRFROM                     TextFieldType = 572
	TextFieldTypeDLCLASSCODETRNOTES                    TextFieldType = 573
	TextFieldTypeDLCLASSCODETRTO                       TextFieldType = 574
	TextFieldTypeDLCLASSCODETTO                        TextFieldType = 421
	TextFieldTypeDLCLASSCODETVFROM                     TextFieldType = 575
	TextFieldTypeDLCLASSCODETVNOTES                    TextFieldType = 576
	TextFieldTypeDLCLASSCODETVTO                       TextFieldType = 577
	TextFieldTypeDLCLASSCODEVFROM                      TextFieldType = 578
	TextFieldTypeDLCLASSCODEVNOTES                     TextFieldType = 579
	TextFieldTypeDLCLASSCODEVTO                        TextFieldType = 580
	TextFieldTypeDLCLASSCODEWFROM                      TextFieldType = 581
	TextFieldTypeDLCLASSCODEWNOTES                     TextFieldType = 582
	TextFieldTypeDLCLASSCODEWTO                        TextFieldType = 583
	TextFieldTypeDLDUPLICATEDATE                       TextFieldType = 176
	TextFieldTypeDLENDORSED                            TextFieldType = 21
	TextFieldTypeDLISSUETYPE                           TextFieldType = 177
	TextFieldTypeDLRECORDCREATED                       TextFieldType = 175
	TextFieldTypeDLRESTRICTIONCODE                     TextFieldType = 22
	TextFieldTypeDLUNDER18DATE                         TextFieldType = 174
	TextFieldTypeDLUNDER19DATE                         TextFieldType = 266
	TextFieldTypeDLUNDER21DATE                         TextFieldType = 23
	TextFieldTypeDNINUMBER                             TextFieldType = 519
	TextFieldTypeDOCUMENTCLASSCODE                     TextFieldType = 0
	TextFieldTypeDOCUMENTCLASSNAME                     TextFieldType = 37
	TextFieldTypeDOCUMENTDISCRIMINATOR                 TextFieldType = 334
	TextFieldTypeDOCUMENTNUMBER                        TextFieldType = 2
	TextFieldTypeDOCUMENTNUMBERCHECKDIGIT              TextFieldType = 80
	TextFieldTypeDOCUMENTNUMBERCHECKSUM                TextFieldType = 40
	TextFieldTypeDOCUMENTSERIES                        TextFieldType = 56
	TextFieldTypeDOCUMENTSTATUS                        TextFieldType = 682
	TextFieldTypeDODNUMBER                             TextFieldType = 348
	TextFieldTypeDONOR                                 TextFieldType = 18
	TextFieldTypeDOSSIERNUMBER                         TextFieldType = 169
	TextFieldTypeDSCERTIFICATEISSUER                   TextFieldType = 327
	TextFieldTypeDSCERTIFICATESUBJECT                  TextFieldType = 328
	TextFieldTypeDSCERTIFICATEVALIDFROM                TextFieldType = 329
	TextFieldTypeDSCERTIFICATEVALIDTO                  TextFieldType = 330
	TextFieldTypeDUFNUMBER                             TextFieldType = 350
	TextFieldTypeDURATIONOFSTAY                        TextFieldType = 103
	TextFieldTypeECENVIRONMENTALTYPE                   TextFieldType = 438
	TextFieldTypeEIDPLACEOFBIRTHCITY                   TextFieldType = 261
	TextFieldTypeEIDPLACEOFBIRTHCOUNTRY                TextFieldType = 263
	TextFieldTypeEIDPLACEOFBIRTHSTATE                  TextFieldType = 262
	TextFieldTypeEIDPLACEOFBIRTHSTREET                 TextFieldType = 260
	TextFieldTypeEIDPLACEOFBIRTHZIPCODE                TextFieldType = 264
	TextFieldTypeEIDRESIDENCEPERMIT1                   TextFieldType = 258
	TextFieldTypeEIDRESIDENCEPERMIT2                   TextFieldType = 259
	TextFieldTypeELECTRONICTICKETINDICATOR             TextFieldType = 365
	TextFieldTypeENDORSEMENTEXPIRATIONDATE             TextFieldType = 269
	TextFieldTypeENGINEMODEL                           TextFieldType = 140
	TextFieldTypeENGINENUMBER                          TextFieldType = 139
	TextFieldTypeENGINEPOWER                           TextFieldType = 136
	TextFieldTypeENGINEVOLUME                          TextFieldType = 137
	TextFieldTypeEQVCODE                               TextFieldType = 294
	TextFieldTypeEXAMDATE                              TextFieldType = 275
	TextFieldTypeEXCEPTINTANKS                         TextFieldType = 461
	TextFieldTypeEYESCOLOR                             TextFieldType = 15
	TextFieldTypeFACULTY                               TextFieldType = 517
	TextFieldTypeFAMILYNAME                            TextFieldType = 126
	TextFieldTypeFAMILYNAMETRUNCATION                  TextFieldType = 272
	TextFieldTypeFASTTRACK                             TextFieldType = 462
	TextFieldTypeFATHERCOUNTRYOFBIRTH                  TextFieldType = 510
	TextFieldTypeFATHERDATEOFBIRTH                     TextFieldType = 504
	TextFieldTypeFATHERGIVENNAME                       TextFieldType = 502
	TextFieldTypeFATHERPERSONALNUMBER                  TextFieldType = 506
	TextFieldTypeFATHERPLACEOFBIRTH                    TextFieldType = 508
	TextFieldTypeFATHERSNAME                           TextFieldType = 129
	TextFieldTypeFATHERSNAMERUS                        TextFieldType = 130
	TextFieldTypeFATHERSURNAME                         TextFieldType = 501
	TextFieldTypeFEDERALELECTIONS                      TextFieldType = 192
	TextFieldTypeFEE                                   TextFieldType = 298
	TextFieldTypeFIELDFROMMRZ                          TextFieldType = 220
	TextFieldTypeFINALCHECKDIGIT                       TextFieldType = 84
	TextFieldTypeFINALCHECKSUM                         TextFieldType = 44
	TextFieldTypeFIRSTISSUEDATE                        TextFieldType = 446
	TextFieldTypeFIRSTNAME                             TextFieldType = 645
	TextFieldTypeFIRSTNAMETRUNCATION                   TextFieldType = 273
	TextFieldTypeFIRSTSURNAME                          TextFieldType = 670
	TextFieldTypeFLIGHTNUMBER                          TextFieldType = 355
	TextFieldTypeFOLIONUMBER                           TextFieldType = 186
	TextFieldTypeFORMOFEDUCATION                       TextFieldType = 518
	TextFieldTypeFOURTHNAME                            TextFieldType = 649
	TextFieldTypeFREEBAGGAGEALLOWANCE                  TextFieldType = 373
	TextFieldTypeFREQUENTFLYERAIRLINEDESIGNATOR        TextFieldType = 371
	TextFieldTypeFREQUENTFLYERNUMBER                   TextFieldType = 372
	TextFieldTypeFROMAIRPORTCODE                       TextFieldType = 353
	TextFieldTypeFUELTYPE                              TextFieldType = 437
	TextFieldTypeGIVENNAMES                            TextFieldType = 9
	TextFieldTypeGIVENNAMESRUS                         TextFieldType = 127
	TextFieldTypeGNIBNUMBER                            TextFieldType = 340
	TextFieldTypeGRANDFATHERNAME                       TextFieldType = 497
	TextFieldTypeGRANDFATHERNAMEMATERNAL               TextFieldType = 669
	TextFieldTypeHAIRCOLOR                             TextFieldType = 16
	TextFieldTypeHEALTHNUMBER                          TextFieldType = 496
	TextFieldTypeHEIGHT                                TextFieldType = 13
	TextFieldTypeIDENTIFIER                            TextFieldType = 288
	TextFieldTypeIDENTITYCARDNUMBER                    TextFieldType = 142
	TextFieldTypeIDENTITYCARDNUMBERCHECKDIGIT          TextFieldType = 376
	TextFieldTypeIDENTITYCARDNUMBERCHECKSUM            TextFieldType = 375
	TextFieldTypeINTANKS                               TextFieldType = 460
	TextFieldTypeINVENTORYNUMBER                       TextFieldType = 121
	TextFieldTypeINVITATIONNUMBER                      TextFieldType = 28
	TextFieldTypeINVITATIONNUMBERCHECKDIGIT            TextFieldType = 86
	TextFieldTypeINVITATIONNUMBERCHECKSUM              TextFieldType = 46
	TextFieldTypeINVITEDBY                             TextFieldType = 451
	TextFieldTypeISOISSUERIDNUMBER                     TextFieldType = 336
	TextFieldTypeISSUERIDENTIFICATIONNUMBER            TextFieldType = 641
	TextFieldTypeISSUETIMESTAMP                        TextFieldType = 96
	TextFieldTypeISSUINGSTATECODE                      TextFieldType = 1
	TextFieldTypeISSUINGSTATECODENUMERIC               TextFieldType = 134
	TextFieldTypeISSUINGSTATENAME                      TextFieldType = 38
	TextFieldTypeJURISDICTIONENDORSEMENTCODE           TextFieldType = 124
	TextFieldTypeJURISDICTIONRESTRICTIONCODE           TextFieldType = 125
	TextFieldTypeJURISDICTIONVEHICLECLASS              TextFieldType = 123
	TextFieldTypeLASTNAME                              TextFieldType = 650
	TextFieldTypeLICENSENUMBER                         TextFieldType = 459
	TextFieldTypeLIMITEDDURATIONDOCUMENTINDICATOR      TextFieldType = 268
	TextFieldTypeLINE1CHECKDIGIT                       TextFieldType = 150
	TextFieldTypeLINE1CHECKSUM                         TextFieldType = 153
	TextFieldTypeLINE1OPTIONALDATA                     TextFieldType = 291
	TextFieldTypeLINE2CHECKDIGIT                       TextFieldType = 151
	TextFieldTypeLINE2CHECKSUM                         TextFieldType = 154
	TextFieldTypeLINE2OPTIONALDATA                     TextFieldType = 292
	TextFieldTypeLINE3CHECKDIGIT                       TextFieldType = 152
	TextFieldTypeLINE3CHECKSUM                         TextFieldType = 155
	TextFieldTypeLINE3OPTIONALDATA                     TextFieldType = 293
	TextFieldTypeMAILINGADDRESSCITY                    TextFieldType = 117
	TextFieldTypeMAILINGADDRESSJURISDICTIONCODE        TextFieldType = 118
	TextFieldTypeMAILINGADDRESSPOSTALCODE              TextFieldType = 119
	TextFieldTypeMAILINGADDRESSSTREET                  TextFieldType = 116
	TextFieldTypeMAKE                                  TextFieldType = 587
	TextFieldTypeMARITALSTATUS                         TextFieldType = 160
	TextFieldTypeMAXMASSOFTRAILERBRAKED                TextFieldType = 440
	TextFieldTypeMAXMASSOFTRAILERUNBRAKED              TextFieldType = 441
	TextFieldTypeMAXSPEED                              TextFieldType = 436
	TextFieldTypeMCNOVICEDATE                          TextFieldType = 349
	TextFieldTypeMEDICALINDICATORCODES                 TextFieldType = 98
	TextFieldTypeMIDDLENAME                            TextFieldType = 146
	TextFieldTypeMIDDLENAMETRUNCATION                  TextFieldType = 274
	TextFieldTypeMILITARYBOOKNUMBER                    TextFieldType = 178
	TextFieldTypeMILITARYSERVICEFROM                   TextFieldType = 626
	TextFieldTypeMILITARYSERVICETO                     TextFieldType = 627
	TextFieldTypeMODEL                                 TextFieldType = 586
	TextFieldTypeMONTH                                 TextFieldType = 106
	TextFieldTypeMONTHOFBIRTH                          TextFieldType = 671
	TextFieldTypeMORTGAGEBY                            TextFieldType = 597
	TextFieldTypeMOTHERCOUNTRYOFBIRTH                  TextFieldType = 509
	TextFieldTypeMOTHERDATEOFBIRTH                     TextFieldType = 503
	TextFieldTypeMOTHERGIVENNAME                       TextFieldType = 500
	TextFieldTypeMOTHERPERSONALNUMBER                  TextFieldType = 505
	TextFieldTypeMOTHERPLACEOFBIRTH                    TextFieldType = 507
	TextFieldTypeMOTHERSNAME                           TextFieldType = 10
	TextFieldTypeMOTHERSURNAME                         TextFieldType = 499
	TextFieldTypeMRZSTRINGS                            TextFieldType = 51
	TextFieldTypeMRZSTRINGSICAORFID                    TextFieldType = 464
	TextFieldTypeMRZSTRINGSWITHCORRECTCHECKSUMS        TextFieldType = 172
	TextFieldTypeMRZTYPE                               TextFieldType = 35
	TextFieldTypeNAMEPREFIX                            TextFieldType = 53
	TextFieldTypeNAMESUFFIX                            TextFieldType = 52
	TextFieldTypeNATIONALITY                           TextFieldType = 11
	TextFieldTypeNATIONALITYCODE                       TextFieldType = 26
	TextFieldTypeNATIONALITYCODENUMERIC                TextFieldType = 135
	TextFieldTypeNONRESIDENTINDICATOR                  TextFieldType = 99
	TextFieldTypeNUMBEROFAXLES                         TextFieldType = 448
	TextFieldTypeNUMBEROFCARDISSUANCE                  TextFieldType = 465
	TextFieldTypeNUMBEROFCARDISSUANCECHECKDIGIT        TextFieldType = 467
	TextFieldTypeNUMBEROFCARDISSUANCECHECKSUM          TextFieldType = 466
	TextFieldTypeNUMBEROFCYLINDERS                     TextFieldType = 588
	TextFieldTypeNUMBEROFDUPLICATES                    TextFieldType = 97
	TextFieldTypeNUMBEROFENTRIES                       TextFieldType = 104
	TextFieldTypeNUMBEROFSEATS                         TextFieldType = 434
	TextFieldTypeNUMBEROFSTANDINGPLACES                TextFieldType = 435
	TextFieldTypeOBSERVATIONS                          TextFieldType = 318
	TextFieldTypeOCRNUMBER                             TextFieldType = 191
	TextFieldTypeOLDDATEOFISSUE                        TextFieldType = 599
	TextFieldTypeOLDDOCUMENTNUMBER                     TextFieldType = 598
	TextFieldTypeOLDPLACEOFISSUE                       TextFieldType = 600
	TextFieldTypeOPTIONALDATA                          TextFieldType = 36
	TextFieldTypeOPTIONALDATACHECKDIGIT                TextFieldType = 195
	TextFieldTypeOPTIONALDATACHECKSUM                  TextFieldType = 194
	TextFieldTypeORGANIZATION                          TextFieldType = 276
	TextFieldTypeOTHER                                 TextFieldType = 50
	TextFieldTypeOTHERNAME                             TextFieldType = 317
	TextFieldTypeOTHERPERSONNAME                       TextFieldType = 322
	TextFieldTypeOTHERVALIDID                          TextFieldType = 315
	TextFieldTypeOWNER                                 TextFieldType = 463
	TextFieldTypePARENTSGIVENNAMES                     TextFieldType = 144
	TextFieldTypePASSPORTNUMBER                        TextFieldType = 27
	TextFieldTypePASSPORTNUMBERCHECKDIGIT              TextFieldType = 85
	TextFieldTypePASSPORTNUMBERCHECKSUM                TextFieldType = 45
	TextFieldTypePATRONHEADERVERSION                   TextFieldType = 304
	TextFieldTypePAYGRADE                              TextFieldType = 278
	TextFieldTypePAYLOADCAPACITY                       TextFieldType = 447
	TextFieldTypePAYMENTPERIODFROM                     TextFieldType = 642
	TextFieldTypePAYMENTPERIODTO                       TextFieldType = 643
	TextFieldTypePDF417CODEC                           TextFieldType = 374
	TextFieldTypePERMISSIBLEAXLELOAD                   TextFieldType = 449
	TextFieldTypePERMITDATEOFEXPIRY                    TextFieldType = 91
	TextFieldTypePERMITDATEOFISSUE                     TextFieldType = 93
	TextFieldTypePERMITDLCLASS                         TextFieldType = 90
	TextFieldTypePERMITENDORSED                        TextFieldType = 95
	TextFieldTypePERMITIDENTIFIER                      TextFieldType = 92
	TextFieldTypePERMITRESTRICTIONCODE                 TextFieldType = 94
	TextFieldTypePERSONALIZATIONSN                     TextFieldType = 321
	TextFieldTypePERSONALNUMBER                        TextFieldType = 7
	TextFieldTypePERSONALNUMBERCHECKDIGIT              TextFieldType = 83
	TextFieldTypePERSONALNUMBERCHECKSUM                TextFieldType = 43
	TextFieldTypePERSONALSUMMARY                       TextFieldType = 314
	TextFieldTypePERSONTONOTIFYADDRESS                 TextFieldType = 326
	TextFieldTypePERSONTONOTIFYDATEOFRECORD            TextFieldType = 323
	TextFieldTypePERSONTONOTIFYNAME                    TextFieldType = 324
	TextFieldTypePERSONTONOTIFYPHONE                   TextFieldType = 325
	TextFieldTypePHONE                                 TextFieldType = 311
	TextFieldTypePLACEOFBIRTH                          TextFieldType = 6
	TextFieldTypePLACEOFBIRTHAREA                      TextFieldType = 74
	TextFieldTypePLACEOFBIRTHCITY                      TextFieldType = 666
	TextFieldTypePLACEOFBIRTHRUS                       TextFieldType = 132
	TextFieldTypePLACEOFBIRTHSTATECODE                 TextFieldType = 75
	TextFieldTypePLACEOFEXAMINATION                    TextFieldType = 513
	TextFieldTypePLACEOFISSUE                          TextFieldType = 39
	TextFieldTypePLACEOFREGISTRATION                   TextFieldType = 69
	TextFieldTypePNRCODE                               TextFieldType = 352
	TextFieldTypePOLICEDISTRICT                        TextFieldType = 445
	TextFieldTypePOWERWEIGHTRATIO                      TextFieldType = 439
	TextFieldTypePRECINCT                              TextFieldType = 450
	TextFieldTypePREVIOUSTYPE                          TextFieldType = 200
	TextFieldTypePROFESSION                            TextFieldType = 312
	TextFieldTypePROFESSIONALIDNUMBER                  TextFieldType = 521
	TextFieldTypePSEUDOCODE                            TextFieldType = 297
	TextFieldTypePURPOSEOFENTRY                        TextFieldType = 452
	TextFieldTypeRACEETHNICITY                         TextFieldType = 122
	TextFieldTypeRANK                                  TextFieldType = 279
	TextFieldTypeREFERENCENUMBER                       TextFieldType = 193
	TextFieldTypeREFERENCENUMBERCHECKDIGIT             TextFieldType = 361
	TextFieldTypeREFERENCENUMBERCHECKSUM               TextFieldType = 360
	TextFieldTypeREGCERTBODYNUMBER                     TextFieldType = 60
	TextFieldTypeREGCERTBODYTYPE                       TextFieldType = 182
	TextFieldTypeREGCERTCARCOLOR                       TextFieldType = 59
	TextFieldTypeREGCERTCARMARK                        TextFieldType = 183
	TextFieldTypeREGCERTCARMODEL                       TextFieldType = 58
	TextFieldTypeREGCERTCARTYPE                        TextFieldType = 61
	TextFieldTypeREGCERTMAXWEIGHT                      TextFieldType = 62
	TextFieldTypeREGCERTREGNUMBER                      TextFieldType = 57
	TextFieldTypeREGCERTREGNUMBERCHECKDIGIT            TextFieldType = 156
	TextFieldTypeREGCERTREGNUMBERCHECKSUM              TextFieldType = 157
	TextFieldTypeREGCERTVEHICLEITSCODE                 TextFieldType = 158
	TextFieldTypeREGCERTVIN                            TextFieldType = 147
	TextFieldTypeREGCERTVINCHECKDIGIT                  TextFieldType = 148
	TextFieldTypeREGCERTVINCHECKSUM                    TextFieldType = 149
	TextFieldTypeREGCERTWEIGHT                         TextFieldType = 63
	TextFieldTypeRELATIONSHIP                          TextFieldType = 284
	TextFieldTypeRELIGION                              TextFieldType = 363
	TextFieldTypeREMAINDERTERM                         TextFieldType = 364
	TextFieldTypeRESIDENTFROM                          TextFieldType = 71
	TextFieldTypeRESIDENTUNTIL                         TextFieldType = 72
	TextFieldTypeRETIREMENTNUMBER                      TextFieldType = 520
	TextFieldTypeREVISIONDATE                          TextFieldType = 270
	TextFieldTypeROOMNUMBER                            TextFieldType = 362
	TextFieldTypeSBHINTEGRITYOPTIONS                   TextFieldType = 301
	TextFieldTypeSBHSECURITYOPTIONS                    TextFieldType = 300
	TextFieldTypeSEATNUMBER                            TextFieldType = 357
	TextFieldTypeSECONDNAME                            TextFieldType = 647
	TextFieldTypeSECONDSURNAME                         TextFieldType = 145
	TextFieldTypeSECTION                               TextFieldType = 190
	TextFieldTypeSELECTEEINDICATOR                     TextFieldType = 498
	TextFieldTypeSEQUENCENUMBER                        TextFieldType = 181
	TextFieldTypeSERIALNUMBER                          TextFieldType = 168
	TextFieldTypeSEX                                   TextFieldType = 12
	TextFieldTypeSIGNATURE                             TextFieldType = 683
	TextFieldTypeSKINCOLOR                             TextFieldType = 453
	TextFieldTypeSOCIALSECURITYNUMBER                  TextFieldType = 19
	TextFieldTypeSPCODE                                TextFieldType = 344
	TextFieldTypeSPECIALNOTES                          TextFieldType = 162
	TextFieldTypeSPONSOR                               TextFieldType = 283
	TextFieldTypeSPONSORSERVICE                        TextFieldType = 281
	TextFieldTypeSPONSORSSN                            TextFieldType = 347
	TextFieldTypeSPONSORSTATUS                         TextFieldType = 282
	TextFieldTypeSTAMPNUMBER                           TextFieldType = 299
	TextFieldTypeSTATUSDATEOFEXPIRY                    TextFieldType = 251
	TextFieldTypeSURNAME                               TextFieldType = 8
	TextFieldTypeSURNAMEANDGIVENNAMES                  TextFieldType = 25
	TextFieldTypeSURNAMEANDGIVENNAMESCHECKDIGIT        TextFieldType = 88
	TextFieldTypeSURNAMEANDGIVENNAMESCHECKSUM          TextFieldType = 48
	TextFieldTypeSURNAMEANDGIVENNAMESRUS               TextFieldType = 131
	TextFieldTypeSURNAMEATBIRTH                        TextFieldType = 432
	TextFieldTypeSURNAMEOFHUSBANDAFTERREGISTRATION     TextFieldType = 589
	TextFieldTypeSURNAMEOFSPOUSE                       TextFieldType = 163
	TextFieldTypeSURNAMEOFWIFEAFTERREGISTRATION        TextFieldType = 590
	TextFieldTypeTAX                                   TextFieldType = 319
	TextFieldTypeTAXNUMBER                             TextFieldType = 495
	TextFieldTypeTELEXCODE                             TextFieldType = 342
	TextFieldTypeTERRITORIALVALIDITY                   TextFieldType = 171
	TextFieldTypeTHIRDNAME                             TextFieldType = 648
	TextFieldTypeTICKETNUMBER                          TextFieldType = 370
	TextFieldTypeTITLE                                 TextFieldType = 313
	TextFieldTypeTOAIRPORTCODE                         TextFieldType = 354
	TextFieldTypeTRACKINGNUMBER                        TextFieldType = 164
	TextFieldTypeTRAILERHITCH                          TextFieldType = 443
	TextFieldTypeTRANSACTIONNUMBER                     TextFieldType = 184
	TextFieldTypeTRANSMISSIONTYPE                      TextFieldType = 442
	TextFieldTypeTYPEAPPROVALNUMBER                    TextFieldType = 332
	TextFieldTypeUNIQUECUSTOMERIDENTIFIER              TextFieldType = 108
	TextFieldTypeURL                                   TextFieldType = 584
	TextFieldTypeUSCIS                                 TextFieldType = 285
	TextFieldTypeVACCINATIONCERTIFICATEIDENTIFIER      TextFieldType = 644
	TextFieldTypeVALIDITYPERIOD                        TextFieldType = 303
	TextFieldTypeVEHICLECATEGORY                       TextFieldType = 141
	TextFieldTypeVETERAN                               TextFieldType = 377
	TextFieldTypeVISACLASS                             TextFieldType = 30
	TextFieldTypeVISAID                                TextFieldType = 29
	TextFieldTypeVISAIDCHECKDIGIT                      TextFieldType = 87
	TextFieldTypeVISAIDCHECKSUM                        TextFieldType = 47
	TextFieldTypeVISAIDRUS                             TextFieldType = 128
	TextFieldTypeVISANUMBER                            TextFieldType = 196
	TextFieldTypeVISANUMBERCHECKDIGIT                  TextFieldType = 198
	TextFieldTypeVISANUMBERCHECKSUM                    TextFieldType = 197
	TextFieldTypeVISASUBCLASS                          TextFieldType = 31
	TextFieldTypeVISATYPE                              TextFieldType = 100
	TextFieldTypeVISAVALIDFROM                         TextFieldType = 101
	TextFieldTypeVISAVALIDUNTIL                        TextFieldType = 102
	TextFieldTypeVISAVALIDUNTILCHECKDIGIT              TextFieldType = 89
	TextFieldTypeVISAVALIDUNTILCHECKSUM                TextFieldType = 49
	TextFieldTypeVOTER                                 TextFieldType = 199
	TextFieldTypeVOTERKEY                              TextFieldType = 187
	TextFieldTypeVOUCHERNUMBER                         TextFieldType = 515
	TextFieldTypeVRCDATAOBJECTENTRY                    TextFieldType = 331
	TextFieldTypeWEIGHT                                TextFieldType = 14
	TextFieldTypeWEIGHTPOUNDS                          TextFieldType = 267
	TextFieldTypeYEAR                                  TextFieldType = 107
	TextFieldTypeYEAROFBIRTH                           TextFieldType = 667
	TextFieldTypeYEAROFEXPIRY                          TextFieldType = 668
	TextFieldTypeYEARSSINCEISSUE                       TextFieldType = 523
)

// Defines values for TextPostProcessing.
const (
	CAPITAL   TextPostProcessing = 3
	LOWERCASE TextPostProcessing = 2
	NOCHANGE  TextPostProcessing = 0
	UPPERCASE TextPostProcessing = 1
)

// Defines values for VerificationResult.
const (
	COMPAREMATCH    VerificationResult = 3
	COMPARENOTMATCH VerificationResult = 4
	DISABLED        VerificationResult = 0
	NOTVERIFIED     VerificationResult = 2
	VERIFIED        VerificationResult = 1
)

// Defines values for Visibility.
const (
	COLORED         Visibility = 2
	GRAYSCALE       Visibility = 4
	INVISIBLE       Visibility = 0
	VISIBLE         Visibility = 1
	WHITEIRMATCHING Visibility = 8
)

// Defines values for MrzDetectModeEnum.
const (
	EMDMBlurBeforeBinarization MrzDetectModeEnum = 2
	EMDMDefault                MrzDetectModeEnum = 0
	EMDMResizeBinarizeWindow   MrzDetectModeEnum = 1
)

// Defines values for GetApiV2TransactionTransactionIdResultsParamsWithImages.
const (
	False GetApiV2TransactionTransactionIdResultsParamsWithImages = false
	True  GetApiV2TransactionTransactionIdResultsParamsWithImages = true
)

// AreaArray defines model for AreaArray.
type AreaArray struct {
	List   *[]RectangleCoordinates `json:"List,omitempty"`
	Points *[]PointArray           `json:"Points,omitempty"`
}

// AreaContainer Checked fragment coordinates
type AreaContainer struct {
	Count  *int                    `json:"Count,omitempty"`
	List   *[]RectangleCoordinates `json:"List,omitempty"`
	Points *[]PointsContainer      `json:"Points,omitempty"`
}

// AuthParams defines model for AuthParams.
type AuthParams struct {
	// CheckAxial This parameter is used to enable laminate integrity check in axial light
	CheckAxial *bool `json:"checkAxial,omitempty"`

	// CheckBarcodeFormat This parameter is used to enable Barcode format check (code metadata, data format, contents format, etc.)
	CheckBarcodeFormat *bool `json:"checkBarcodeFormat,omitempty"`

	// CheckExtMRZ This parameter is used to enable Extended MRZ Check
	CheckExtMRZ *bool `json:"checkExtMRZ,omitempty"`

	// CheckExtOCR This parameter is used to enable Extended OCR Check
	CheckExtOCR *bool `json:"checkExtOCR,omitempty"`

	// CheckFibers This parameter is used to enable Fibers detection
	CheckFibers *bool `json:"checkFibers,omitempty"`

	// CheckIPI This parameter is used to enable Invisible Personal Information (IPI) check
	CheckIPI *bool `json:"checkIPI,omitempty"`

	// CheckIRB900 This parameter is used to enable B900 ink MRZ contrast check in IR light
	CheckIRB900 *bool `json:"checkIRB900,omitempty"`

	// CheckIRVisibility This parameter is used to enable Document elements visibility check in IR light
	CheckIRVisibility *bool `json:"checkIRVisibility,omitempty"`

	// CheckImagePatterns This parameter is used to enable Image patterns presence/absence check (position, shape, color)
	CheckImagePatterns *bool `json:"checkImagePatterns,omitempty"`

	// CheckLetterScreen This parameter is used to enable LetterScreen check
	CheckLetterScreen *bool `json:"checkLetterScreen,omitempty"`

	// CheckLiveness This parameter is used to enable document liveness check
	CheckLiveness *bool `json:"checkLiveness,omitempty"`

	// CheckPhotoComparison This parameter is used to enable Portrait comparison check
	CheckPhotoComparison *bool `json:"checkPhotoComparison,omitempty"`

	// CheckPhotoEmbedding This parameter is used to enable Owner's photo embedding check (is photo printed or sticked)
	CheckPhotoEmbedding *bool `json:"checkPhotoEmbedding,omitempty"`

	// CheckUVLuminiscence This parameter is used to enable Document luminescence check in UV light
	CheckUVLuminiscence *bool           `json:"checkUVLuminiscence,omitempty"`
	LivenessParams      *LivenessParams `json:"livenessParams,omitempty"`
}

// AuthenticityCheckList defines model for AuthenticityCheckList.
type AuthenticityCheckList struct {
	// Count Count of items in List
	Count *int `json:"Count,omitempty"`

	// List Authenticity Check
	List []SchemasAuthenticityCheckResult `json:"List"`
}

// AuthenticityCheckResult defines model for AuthenticityCheckResult.
type AuthenticityCheckResult struct {
	List []AuthenticityCheckResult_List_Item `json:"List"`

	// Result 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Result CheckResult `json:"Result"`

	// Type Enumeration describes available authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/.
	Type AuthenticityResultType `json:"Type"`
}

// AuthenticityCheckResult_List_Item defines model for AuthenticityCheckResult.List.Item.
type AuthenticityCheckResult_List_Item struct {
	union json.RawMessage
}

// AuthenticityCheckResultItem Common fields for all authenticity result objects
type AuthenticityCheckResultItem struct {
	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult *CheckResult `json:"ElementResult,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`
}

// AuthenticityResult defines model for AuthenticityResult.
type AuthenticityResult struct {
	AuthenticityCheckList AuthenticityCheckList `json:"AuthenticityCheckList"`
	BufLength             *int                  `json:"buf_length,omitempty"`
	Light                 *int                  `json:"light,omitempty"`
	ListIdx               *int                  `json:"list_idx,omitempty"`
	PageIdx               *int                  `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// AuthenticityResultType Enumeration describes available authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/.
type AuthenticityResultType int

// AuthenticityType Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
type AuthenticityType = int

// CheckDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
type CheckDiagnose int

// CheckResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
type CheckResult int

// ChosenDocumentType Contains information about one document type candidate
type ChosenDocumentType = OneCandidate

// ChosenDocumentTypeResult defines model for ChosenDocumentTypeResult.
type ChosenDocumentTypeResult struct {
	// OneCandidate Contains information about one document type candidate
	OneCandidate *OneCandidate `json:"OneCandidate,omitempty"`
	XMLBuffer    *string       `json:"XML_buffer,omitempty"`
	BufLength    *int          `json:"buf_length,omitempty"`
	Light        *int          `json:"light,omitempty"`
	ListIdx      *int          `json:"list_idx,omitempty"`
	PageIdx      *int          `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// ComparisonMatrix results comparison matrix. Elements of the matrix with indices 0, 1, 2, 3 take one of the values Disabled(0), Verified(1) or Not_Verified(2), elements with indices 4, 5, 6, 7, 8 are one of the values Disabled(0), Compare_Match(3) or Compare_Not_Match(4). Elements of the Matrix matrix have the following semantic meaning:
// - element with index 0 –– the result of verification of data from the MRZ; - 1 –– the result of verification of data from the RFID microcircuit; - 2 –– the result of verification of data from text areas of the document; - 3 –– the result of verification data from barcodes; - 4 - the result of comparing MRZ data and RFID microcircuits; - 5 - the result of comparing MRZ data and text areas of document filling; - 6 - the result of comparing MRZ data and bar codes; - 7 - the result of comparing the data of text areas of the document and the RFID chip; - 8 - the result of comparing the data of the text areas of the document and barcodes; - 9 - the result of comparing the data of the RFID chip and barcodes.
type ComparisonMatrix = []VerificationResult

// ContainerList List with various objects, containing processing results
type ContainerList struct {
	// Count Length of list (Count for items)
	Count *int                      `json:"Count,omitempty"`
	List  []ContainerList_List_Item `json:"List"`
}

// ContainerList_List_Item defines model for ContainerList.List.Item.
type ContainerList_List_Item struct {
	union json.RawMessage
}

// ContainerType Same as Result type, but used for safe parsing of not-described values. See Result type.
type ContainerType = int

// Critical Enumeration contains identifiers determining the criticality of the security element
type Critical int

// CrossSourceValueComparison defines model for CrossSourceValueComparison.
type CrossSourceValueComparison struct {
	// SourceLeft Document data sources
	SourceLeft Source `json:"sourceLeft"`

	// SourceRight Document data sources
	SourceRight Source `json:"sourceRight"`

	// Status 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Status CheckResult `json:"status"`
}

// DataModule defines model for DataModule.
type DataModule struct {
	MData      *string `json:"mData,omitempty"`
	MLength    *int    `json:"mLength,omitempty"`
	MReserved1 *int    `json:"mReserved1,omitempty"`
	MReserver2 *int    `json:"mReserver2,omitempty"`
	MType      *int    `json:"mType,omitempty"`
}

// DeviceInfo defines model for DeviceInfo.
type DeviceInfo struct {
	AppName       *string    `json:"app-name,omitempty"`
	LicenseId     *string    `json:"license-id,omitempty"`
	LicenseSerial *string    `json:"license-serial,omitempty"`
	ServerTime    *time.Time `json:"server-time,omitempty"`
	ValidUntil    *time.Time `json:"valid-until,omitempty"`
	Version       *string    `json:"version,omitempty"`
}

// DocBarCodeInfo defines model for DocBarCodeInfo.
type DocBarCodeInfo struct {
	DocBarCodeInfo *DocBarCodeInfoFieldsList `json:"DocBarCodeInfo,omitempty"`
	BufLength      *int                      `json:"buf_length,omitempty"`
	Light          *int                      `json:"light,omitempty"`
	ListIdx        *int                      `json:"list_idx,omitempty"`
	PageIdx        *int                      `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// DocBarCodeInfoFieldsList defines model for DocBarCodeInfoFieldsList.
type DocBarCodeInfoFieldsList struct {
	// ArrayFields Data from barcode
	ArrayFields *[]PArrayField `json:"ArrayFields,omitempty"`

	// Fields Count of array fields
	Fields *int `json:"Fields,omitempty"`
}

// DocVisualExtendedField defines model for DocVisualExtendedField.
type DocVisualExtendedField struct {
	// BufText Text field data in UTF8 format. Results of reading different lines of a multi-line field are separated by '^'
	BufText *string `json:"Buf_Text,omitempty"`

	// FieldName Field name. Only use to search values for fields with fieldType=50(other). In general, use wFieldType for lookup.
	FieldName string `json:"FieldName"`

	// FieldRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	FieldRect *RectangleCoordinates `json:"FieldRect,omitempty"`

	// RFIDOriginDG Origin data group information. Only for Result.RFID_TEXT results.
	RFIDOriginDG *int `json:"RFID_OriginDG,omitempty"`

	// RFIDOriginTagEntry Index of the text field record in origin data group. Only for Result.RFID_TEXT results.
	RFIDOriginTagEntry *int `json:"RFID_OriginTagEntry,omitempty"`

	// StringsResult Array of recognizing probabilities for a each line of text field. Only for Result.VISUAL_TEXT and Result.MRZ_TEXT results.
	StringsResult *[]StringRecognitionResult `json:"StringsResult,omitempty"`
	WFieldType    TextFieldType              `json:"wFieldType"`

	// WLCID Locale id. Used to tag same typed fields declared in several languages.
	// For example: name can be provided in both native and latin variants.
	// Based on Microsoft locale id (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/70feba9f-294e-491e-b6eb-56532684c37f).
	WLCID LCID `json:"wLCID"`
}

// DocVisualExtendedInfo Container for extracted text fields. Fields are identified by type and language
type DocVisualExtendedInfo struct {
	ArrayFields []DocVisualExtendedField `json:"ArrayFields"`
}

// DocumentFormat Defining the geometric format of documents in accordance with ISO / IEC 7810
type DocumentFormat int

// DocumentImage defines model for DocumentImage.
type DocumentImage = ImageData

// DocumentImageResult defines model for DocumentImageResult.
type DocumentImageResult struct {
	RawImageContainer ImageData `json:"RawImageContainer"`
	BufLength         *int      `json:"buf_length,omitempty"`
	Light             *int      `json:"light,omitempty"`
	ListIdx           *int      `json:"list_idx,omitempty"`
	PageIdx           *int      `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// DocumentPosition defines model for DocumentPosition.
type DocumentPosition struct {
	Angle       *float32 `json:"Angle,omitempty"`
	Center      *Point   `json:"Center,omitempty"`
	Dpi         *int     `json:"Dpi,omitempty"`
	Height      *int     `json:"Height,omitempty"`
	LeftBottom  *Point   `json:"LeftBottom,omitempty"`
	LeftTop     *Point   `json:"LeftTop,omitempty"`
	RightBottom *Point   `json:"RightBottom,omitempty"`
	RightTop    *Point   `json:"RightTop,omitempty"`
	Width       *int     `json:"Width,omitempty"`

	// DocFormat Defining the geometric format of documents in accordance with ISO / IEC 7810
	DocFormat *DocumentFormat `json:"docFormat,omitempty"`
}

// DocumentPositionResult defines model for DocumentPositionResult.
type DocumentPositionResult struct {
	DocumentPosition *DocumentPosition `json:"DocumentPosition,omitempty"`
	BufLength        *int              `json:"buf_length,omitempty"`
	Light            *int              `json:"light,omitempty"`
	ListIdx          *int              `json:"list_idx,omitempty"`
	PageIdx          *int              `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// DocumentType Possible values for document types
type DocumentType int

// DocumentTypeRecognitionResult defines model for DocumentTypeRecognitionResult.
type DocumentTypeRecognitionResult int

// DocumentTypesCandidates defines model for DocumentTypesCandidates.
type DocumentTypesCandidates = DocumentTypesCandidatesList

// DocumentTypesCandidatesList defines model for DocumentTypesCandidatesList.
type DocumentTypesCandidatesList struct {
	Candidates *[]OneCandidate                `json:"Candidates,omitempty"`
	RecResult  *DocumentTypeRecognitionResult `json:"RecResult,omitempty"`
}

// DocumentTypesCandidatesResult defines model for DocumentTypesCandidatesResult.
type DocumentTypesCandidatesResult struct {
	CandidatesList *DocumentTypesCandidatesList `json:"CandidatesList,omitempty"`
	BufLength      *int                         `json:"buf_length,omitempty"`
	Light          *int                         `json:"light,omitempty"`
	ListIdx        *int                         `json:"list_idx,omitempty"`
	PageIdx        *int                         `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// EncryptedRCL Base64 encoded data
type EncryptedRCL = []byte

// EncryptedRCLResult defines model for EncryptedRCLResult.
type EncryptedRCLResult struct {
	// EncryptedRCL Base64 encoded data
	EncryptedRCL EncryptedRCL `json:"EncryptedRCL"`
	BufLength    *int         `json:"buf_length,omitempty"`
	Light        *int         `json:"light,omitempty"`
	ListIdx      *int         `json:"list_idx,omitempty"`
	PageIdx      *int         `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// FDSIDList Extended document type info and Regula's 'Information Reference Systems' links
type FDSIDList struct {
	// ICAOCode ICAO code of the issuing country
	ICAOCode *string `json:"ICAOCode,omitempty"`

	// List Document identifiers in 'Information Reference Systems'
	List *[]int `json:"List,omitempty"`

	// DCountryName Issuing country name
	DCountryName *string `json:"dCountryName,omitempty"`

	// DDescription Document description
	DDescription *string `json:"dDescription,omitempty"`

	// DFormat Defining the geometric format of documents in accordance with ISO / IEC 7810
	DFormat *DocumentFormat `json:"dFormat,omitempty"`

	// DMRZ Flag indicating the presence of MRZ on the document
	DMRZ *bool `json:"dMRZ,omitempty"`

	// DStateCode Issuing state code
	DStateCode *string `json:"dStateCode,omitempty"`

	// DStateName Issuing state name
	DStateName *string `json:"dStateName,omitempty"`

	// DType Possible values for document types
	DType *DocumentType `json:"dType,omitempty"`

	// DYear Year of publication of the document
	DYear *string `json:"dYear,omitempty"`

	// IsDeprecated Whether the document is deprecated
	IsDeprecated *bool `json:"isDeprecated,omitempty"`
}

// FaceApi defines model for FaceApi.
type FaceApi struct {
	// Mode The processing mode: "match" or "match+search".
	Mode *string `json:"mode,omitempty"`

	// Proxy Proxy to use, should be set according to the <a href="https://curl.se/libcurl/c/CURLOPT_PROXY.html" target="_blank">cURL standard</a>.
	Proxy *string `json:"proxy,omitempty"`

	// ProxyType Proxy protocol type, should be set according to the <a href="https://curl.se/libcurl/c/CURLOPT_PROXYTYPE.html" target="_blank">cURL standard</a>.
	ProxyType *int `json:"proxy_type,omitempty"`

	// ProxyUserpwd Username and password to use for proxy authentication, should be set according to the <a href="https://curl.se/libcurl/c/CURLOPT_PROXYUSERPWD.html" target="_blank">cURL standard</a>.
	ProxyUserpwd *string `json:"proxy_userpwd,omitempty"`

	// Search A search filter that can be applied if the "match+search" mode is enabled. May include limit, threshold, group_ids. If the group_ids are specified, the search is performed only in these groups. Find more information in the <a href="https://dev.regulaforensics.com/FaceSDK-web-openapi/#tag/search/operation/search" target="_blank">OpenAPI documentation</a>.
	Search *struct {
		// GroupIds  The groups where to conduct the search.
		GroupIds *[]int `json:"group_ids,omitempty"`

		// Limit The maximum number of results to be returned.
		Limit *int `json:"limit,omitempty"`

		// Threshold The similarity threshold.
		Threshold *float32 `json:"threshold,omitempty"`
	} `json:"search,omitempty"`

	// ServiceTimeout The Regula Face Web service requests timeout, ms.
	ServiceTimeout *int `json:"serviceTimeout,omitempty"`

	// Threshold The similarity threshold, 0-100. Above 75 means that the faces' similarity is verified, below 75 is not.
	Threshold *int `json:"threshold,omitempty"`

	// Url The URL of the Regula Face Web service to be used.
	Url *string `json:"url,omitempty"`
}

// FiberResult defines model for FiberResult.
type FiberResult struct {
	// Area Fibers value for areas (in pixels)
	Area *[]int `json:"Area,omitempty"`

	// ColorValues Fibers color value
	ColorValues *[]int `json:"ColorValues,omitempty"`

	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult *CheckResult `json:"ElementResult,omitempty"`

	// ExpectedCount Expected fibers number. For UV_Fibers authentication result type
	ExpectedCount *int `json:"ExpectedCount,omitempty"`

	// Length Fibers length value for located areas (in pixels)
	Length *[]int `json:"Length,omitempty"`

	// LightDisp For UV_Background authentication result type
	LightDisp *int `json:"LightDisp,omitempty"`

	// LightValue Image light index
	LightValue *Light `json:"LightValue,omitempty"`

	// RectArray Coordinates of located areas for defined fibers type
	RectArray *[]RectangleCoordinates `json:"RectArray,omitempty"`

	// RectCount For UV_Fibers authenticity result type
	RectCount *int `json:"RectCount,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`

	// Width Fibers width value for located areas (in pixels)
	Width *[]int `json:"Width,omitempty"`
}

// GraphicField defines model for GraphicField.
type GraphicField struct {
	// FieldRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	FieldRect *RectangleCoordinates `json:"FieldRect,omitempty"`
	FieldType GraphicFieldType      `json:"FieldType"`

	// RFIDOriginDG Source data group file. Only for Result.RFID_GRAPHICS result.
	RFIDOriginDG *int `json:"RFID_OriginDG,omitempty"`

	// RFIDOriginDGTag Index of the source record of the image with biometric information in the information data group. Only for Result.RFID_GRAPHICS result.
	RFIDOriginDGTag *int `json:"RFID_OriginDGTag,omitempty"`

	// RFIDOriginEntryView Index of the variant of the biometric data template. Only for Result.RFID_GRAPHICS result.
	RFIDOriginEntryView *int `json:"RFID_OriginEntryView,omitempty"`

	// RFIDOriginTagEntry Index of the template in the record with biometric data. Only for Result.RFID_GRAPHICS result.
	RFIDOriginTagEntry *int      `json:"RFID_OriginTagEntry,omitempty"`
	Image              ImageData `json:"image"`
}

// GraphicFieldName Human readable field name. Do not bind to this name - use GraphicFieldType instead.
type GraphicFieldName = string

// GraphicFieldType defines model for GraphicFieldType.
type GraphicFieldType int

// GraphicFieldsList defines model for GraphicFieldsList.
type GraphicFieldsList struct {
	ArrayFields []GraphicField `json:"ArrayFields"`
}

// GraphicsResult defines model for GraphicsResult.
type GraphicsResult struct {
	DocGraphicsInfo *GraphicFieldsList `json:"DocGraphicsInfo,omitempty"`
	BufLength       *int               `json:"buf_length,omitempty"`
	Light           *int               `json:"light,omitempty"`
	ListIdx         *int               `json:"list_idx,omitempty"`
	PageIdx         *int               `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// IdentResult defines model for IdentResult.
type IdentResult struct {
	// Area Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	Area *RectangleCoordinates `json:"Area,omitempty"`

	// AreaList Checked fragment coordinates
	AreaList *AreaContainer `json:"AreaList,omitempty"`

	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult *CheckResult `json:"ElementResult,omitempty"`

	// ElementType Enumeration contains identifiers determining type of features for a document authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/security-feature-type/
	ElementType *SecurityFeatureType `json:"ElementType,omitempty"`
	EtalonImage *ImageData           `json:"EtalonImage,omitempty"`
	Image       *ImageData           `json:"Image,omitempty"`

	// LightIndex Image light index
	LightIndex *Light `json:"LightIndex,omitempty"`

	// PercentValue Probability percent for IMAGE_PATTERN check or element's visibility for IR_VISIBILITY
	PercentValue *int `json:"PercentValue,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`
}

// ImageBase64 Base64 encoded image
type ImageBase64 = string

// ImageData defines model for ImageData.
type ImageData struct {
	// Image Base64 encoded image
	Image ImageBase64 `json:"image"`
}

// ImageQA defines model for ImageQA.
type ImageQA struct {
	// AngleThreshold This parameter sets threshold for Image QA check of the presented document perspective angle in degrees. If actual document perspective angle is above this threshold, check will fail.
	AngleThreshold *int `json:"angleThreshold,omitempty"`

	// BrightnessThreshold Set the threshold for an actual document brightness below which the check fails
	BrightnessThreshold *float64 `json:"brightnessThreshold,omitempty"`

	// ColornessCheck This option enables colorness check while performing image quality validation.
	ColornessCheck *bool `json:"colornessCheck,omitempty"`

	// DocumentPositionIndent This parameter specifies the necessary margin. Default 0.
	DocumentPositionIndent *int `json:"documentPositionIndent,omitempty"`

	// DpiThreshold This parameter sets threshold for Image QA check of the presented document physical dpi. If actual document dpi is below this threshold, check will fail.
	DpiThreshold *int `json:"dpiThreshold,omitempty"`

	// FocusCheck This option enables focus check while performing image quality validation.
	FocusCheck *bool `json:"focusCheck,omitempty"`

	// GlaresCheck This option enables glares check while performing image quality validation.
	GlaresCheck *bool `json:"glaresCheck,omitempty"`

	// MoireCheck This option enables screen capture (moire patterns) check while performing image quality validation.
	MoireCheck *bool `json:"moireCheck,omitempty"`
}

// ImageQualityCheck defines model for ImageQualityCheck.
type ImageQualityCheck struct {
	Areas *AreaArray `json:"areas,omitempty"`

	// FeatureType Enumeration contains identifiers determining type of features for a document authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/security-feature-type/
	FeatureType *SecurityFeatureType `json:"featureType,omitempty"`
	Mean        *float32             `json:"mean,omitempty"`
	Probability *int                 `json:"probability,omitempty"`

	// Result 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Result *CheckResult `json:"result,omitempty"`
	StdDev *float32     `json:"std_dev,omitempty"`

	// Type Image quality check type
	Type *ImageQualityCheckType `json:"type,omitempty"`
}

// ImageQualityCheckList defines model for ImageQualityCheckList.
type ImageQualityCheckList struct {
	List []ImageQualityCheck `json:"List"`

	// Result 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Result CheckResult `json:"result"`
}

// ImageQualityCheckType Image quality check type
type ImageQualityCheckType int

// ImageQualityResult defines model for ImageQualityResult.
type ImageQualityResult struct {
	ImageQualityCheckList ImageQualityCheckList `json:"ImageQualityCheckList"`
	BufLength             *int                  `json:"buf_length,omitempty"`
	Light                 *int                  `json:"light,omitempty"`
	ListIdx               *int                  `json:"list_idx,omitempty"`
	PageIdx               *int                  `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// ImageTransactionData defines model for ImageTransactionData.
type ImageTransactionData struct {
	Image *ImagesFieldValue `json:"image,omitempty"`
}

// Images defines model for Images.
type Images struct {
	AvailableSourceList []ImagesAvailableSource `json:"availableSourceList"`
	FieldList           []ImagesField           `json:"fieldList"`
}

// ImagesAvailableSource defines model for ImagesAvailableSource.
type ImagesAvailableSource struct {
	// ContainerType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ContainerType *ContainerType `json:"containerType,omitempty"`

	// Source Document data sources
	Source Source `json:"source"`
}

// ImagesField defines model for ImagesField.
type ImagesField struct {
	// FieldName Human readable field name. Do not bind to this name - use GraphicFieldType instead.
	FieldName GraphicFieldName   `json:"fieldName"`
	FieldType GraphicFieldType   `json:"fieldType"`
	ValueList []ImagesFieldValue `json:"valueList"`
}

// ImagesFieldValue defines model for ImagesFieldValue.
type ImagesFieldValue struct {
	// ContainerType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ContainerType ContainerType `json:"containerType"`

	// FieldRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	FieldRect *RectangleCoordinates `json:"fieldRect,omitempty"`

	// LightIndex Image light index
	LightIndex Light `json:"lightIndex"`

	// OriginalPageIndex Original page index
	OriginalPageIndex *int `json:"originalPageIndex,omitempty"`

	// OriginalValue Base64 encoded image
	OriginalValue *ImageBase64 `json:"originalValue,omitempty"`

	// PageIndex Page index of the image from input list
	PageIndex PageIndex `json:"pageIndex"`

	// RfidOrigin Location of data in RFID chip
	RfidOrigin *RfidOrigin `json:"rfidOrigin,omitempty"`

	// Source Document data sources
	Source Source `json:"source"`

	// Value Base64 encoded image
	Value ImageBase64 `json:"value"`
}

// ImagesResult defines model for ImagesResult.
type ImagesResult struct {
	Images    SchemasImages `json:"Images"`
	BufLength *int          `json:"buf_length,omitempty"`
	Light     *int          `json:"light,omitempty"`
	ListIdx   *int          `json:"list_idx,omitempty"`
	PageIdx   *int          `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// InData defines model for InData.
type InData struct {
	Images *[]InDataTransactionImagesFieldValue `json:"images,omitempty"`

	// Video Video
	Video *struct {
		// Metadata A free-form object containing video's extended attributes.
		Metadata *map[string]interface{} `json:"metadata,omitempty"`

		// Url Video url
		Url *string `json:"url,omitempty"`
	} `json:"video,omitempty"`
}

// InDataTransactionImagesFieldValue defines model for InDataTransactionImagesFieldValue.
type InDataTransactionImagesFieldValue struct {
	// Light Image light index
	Light   *Light `json:"light,omitempty"`
	ListIdx *int   `json:"listIdx,omitempty"`

	// PageIdx Page index of the image from input list
	PageIdx *PageIndex `json:"pageIdx,omitempty"`

	// Url Image url
	Url *string `json:"url,omitempty"`
}

// LCID Locale id. Used to tag same typed fields declared in several languages.
// For example: name can be provided in both native and latin variants.
// Based on Microsoft locale id (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/70feba9f-294e-491e-b6eb-56532684c37f).
type LCID int

// LexicalAnalysisResult defines model for LexicalAnalysisResult.
type LexicalAnalysisResult struct {
	ListVerifiedFields *ListVerifiedFields `json:"ListVerifiedFields,omitempty"`
	BufLength          *int                `json:"buf_length,omitempty"`
	Light              *int                `json:"light,omitempty"`
	ListIdx            *int                `json:"list_idx,omitempty"`
	PageIdx            *int                `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// License Base64 encoded data
type License = []byte

// LicenseResult defines model for LicenseResult.
type LicenseResult struct {
	// License Base64 encoded data
	License   License `json:"License"`
	BufLength *int    `json:"buf_length,omitempty"`
	Light     *int    `json:"light,omitempty"`
	ListIdx   *int    `json:"list_idx,omitempty"`
	PageIdx   *int    `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// Light Image light index
type Light int

// ListVerifiedFields defines model for ListVerifiedFields.
type ListVerifiedFields struct {
	PFieldMaps *[]VerifiedFieldMap `json:"pFieldMaps,omitempty"`
}

// LivenessParams defines model for LivenessParams.
type LivenessParams struct {
	// CheckED This parameter is used to enable Electronic device detection
	CheckED *bool `json:"checkED,omitempty"`

	// CheckHolo This parameter is used to enable Hologram detection
	CheckHolo *bool `json:"checkHolo,omitempty"`

	// CheckMLI This parameter is used to enable MLI check
	CheckMLI *bool `json:"checkMLI,omitempty"`

	// CheckOVI This parameter is used to enable OVI check
	CheckOVI *bool `json:"checkOVI,omitempty"`
}

// LogLevel defines model for LogLevel.
type LogLevel string

// MRZFormat defines model for MRZFormat.
type MRZFormat string

// MeasureSystem defines model for MeasureSystem.
type MeasureSystem int

// OCRSecurityTextResult defines model for OCRSecurityTextResult.
type OCRSecurityTextResult struct {
	// CriticalFlag Enumeration contains identifiers determining the criticality of the security element
	CriticalFlag *Critical `json:"CriticalFlag,omitempty"`

	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult    *CheckResult `json:"ElementResult,omitempty"`
	EtalonFieldType  *int         `json:"EtalonFieldType,omitempty"`
	EtalonLightType  *int         `json:"EtalonLightType,omitempty"`
	EtalonResultOCR  *string      `json:"EtalonResultOCR,omitempty"`
	EtalonResultType *int         `json:"EtalonResultType,omitempty"`

	// FieldRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	FieldRect *RectangleCoordinates `json:"FieldRect,omitempty"`

	// LightType Image light index
	LightType             *Light  `json:"LightType,omitempty"`
	Reserved1             *int    `json:"Reserved1,omitempty"`
	Reserved2             *int    `json:"Reserved2,omitempty"`
	SecurityTextResultOCR *string `json:"SecurityTextResultOCR,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`
}

// OneCandidate Contains information about one document type candidate
type OneCandidate struct {
	// AuthenticityNecessaryLights Combination of lighting scheme identifiers (combination of Light enum) needed to perform all authenticity checks specified in CheckAuthenticity
	AuthenticityNecessaryLights *int `json:"AuthenticityNecessaryLights,omitempty"`

	// CheckAuthenticity Set of authentication options provided for this type of document (combination of Authenticity enum)
	CheckAuthenticity *int `json:"CheckAuthenticity,omitempty"`

	// DocumentName Document name
	DocumentName *string `json:"DocumentName,omitempty"`

	// FDSIDList Extended document type info and Regula's 'Information Reference Systems' links
	FDSIDList *FDSIDList `json:"FDSIDList,omitempty"`

	// ID Unique document type template identifier (Regula's internal numeric code)
	ID *int `json:"ID,omitempty"`

	// NecessaryLights Combination of lighting scheme identifiers (Light enum) required to conduct OCR for this type of document
	NecessaryLights *int `json:"NecessaryLights,omitempty"`

	// P A measure of the likelihood of correct recognition in the analysis of this type of document
	P *float32 `json:"P,omitempty"`

	// RFIDPresence Determines the presence and location of an RFID chip in a document. 0 - no rfid chip; 1 - chip is located in the document data page; 2 - chip is located in the back page or inlay of the document
	RFIDPresence *RfidLocation `json:"RFID_Presence,omitempty"`

	// UVExp The required exposure value of the camera when receiving images of a document of this type for a UV lighting scheme
	UVExp *int `json:"UVExp,omitempty"`
}

// OriginalSymbol defines model for OriginalSymbol.
type OriginalSymbol struct {
	// Code Unicode symbol code
	Code *int64 `json:"code,omitempty"`

	// Probability Probability of correctness reading of a single character
	Probability *int `json:"probability,omitempty"`

	// Rect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	Rect *RectangleCoordinates `json:"rect,omitempty"`
}

// OutData defines model for OutData.
type OutData struct {
	Images *[]OutDataTransactionImagesFieldValue `json:"images,omitempty"`

	// Url Response url
	Url *string `json:"url,omitempty"`
}

// OutDataTransactionImagesFieldValue defines model for OutDataTransactionImagesFieldValue.
type OutDataTransactionImagesFieldValue struct {
	FieldType *GraphicFieldType `json:"fieldType,omitempty"`

	// Light Image light index
	Light   *Light `json:"light,omitempty"`
	ListIdx *int   `json:"listIdx,omitempty"`

	// PageIdx Page index of the image from input list
	PageIdx *PageIndex `json:"pageIdx,omitempty"`

	// Url Image url
	Url *string `json:"url,omitempty"`
}

// PageIndex Page index of the image from input list
type PageIndex = int

// ParsingNotificationCodes The enumeration contains possible values of notification codes returned during the RFID chip processing.
type ParsingNotificationCodes int

// PerDocumentConfig defines model for PerDocumentConfig.
type PerDocumentConfig struct {
	// DocID Specific template IDs, for which apply current custom configuration
	DocID *[]int `json:"docID,omitempty"`

	// ExcludeAuthChecks Contains items from AuthenticityResultType as sum via OR operation
	ExcludeAuthChecks *int `json:"excludeAuthChecks,omitempty"`
}

// PhotoIdentResult defines model for PhotoIdentResult.
type PhotoIdentResult struct {
	Angle *int `json:"Angle,omitempty"`

	// Area Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	Area *RectangleCoordinates `json:"Area,omitempty"`

	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult   *CheckResult `json:"ElementResult,omitempty"`
	FieldTypesCount *int         `json:"FieldTypesCount,omitempty"`
	FieldTypesList  *[]int       `json:"FieldTypesList,omitempty"`

	// LightIndex Image light index
	LightIndex   *Light                 `json:"LightIndex,omitempty"`
	Reserved3    *int                   `json:"Reserved3,omitempty"`
	ResultImages *RawImageContainerList `json:"ResultImages,omitempty"`
	SourceImage  *ImageData             `json:"SourceImage,omitempty"`
	Step         *int                   `json:"Step,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`
}

// Point defines model for Point.
type Point struct {
	X *int `json:"x,omitempty"`
	Y *int `json:"y,omitempty"`
}

// PointArray defines model for PointArray.
type PointArray struct {
	PointsList *[]Point `json:"PointsList,omitempty"`
}

// PointsContainer defines model for PointsContainer.
type PointsContainer struct {
	PointCount *int     `json:"PointCount,omitempty"`
	PointsList *[]Point `json:"PointsList,omitempty"`
}

// ProcessParams defines model for ProcessParams.
type ProcessParams struct {
	// AlreadyCropped This option can be enabled if you know for sure that the image you provide contains already cropped document by its edges. This was designed to process on the server side images captured and cropped on mobile. Disabled by default.
	AlreadyCropped *bool       `json:"alreadyCropped,omitempty"`
	AuthParams     *AuthParams `json:"authParams,omitempty"`

	// CheckAuth This parameter is used to enable authenticity checks
	CheckAuth *bool `json:"checkAuth,omitempty"`

	// CheckRequiredTextFields When enabled, each field in template will be checked for value presence and if the field is marked as required, but has no value, it will have 'error' in validity status. Disabled by default.
	CheckRequiredTextFields *bool `json:"checkRequiredTextFields,omitempty"`

	// Config This option allows setting additional custom configuration per document type. If recognized document has ID specified in config, processing adjusts according to designated configuration.
	Config      *[]PerDocumentConfig `json:"config,omitempty"`
	ConvertCase *TextPostProcessing  `json:"convertCase,omitempty"`

	// CustomParams This option allows passing custom processing parameters that can be implemented in future without changing API.
	CustomParams *map[string]interface{} `json:"customParams,omitempty"`

	// DateFormat This option allows you to set dates format so that solution will return dates in this format. For example, if you supply 'MM/dd/yyyy', and document have printed date '09 JUL 2020' for the date os issue, you will get '07/09/2020' as a result. By default it is set to system locale default (where the service is running).
	DateFormat *string `json:"dateFormat,omitempty"`

	// DepersonalizeLog When enabled, all personal data will be forcibly removed from the logs. Disabled by default.
	DepersonalizeLog *bool `json:"depersonalizeLog,omitempty"`

	// DeviceId This parameter is used to specify the document reader device type from which input images were captured. Default 0.
	DeviceId *int `json:"deviceId,omitempty"`

	// DeviceType This parameter is used to specify the document reader device type from which input images were captured. Default 0.
	DeviceType *int `json:"deviceType,omitempty"`

	// DeviceTypeHex This parameter is used to specify the document reader device type from which input images were captured
	DeviceTypeHex *string `json:"deviceTypeHex,omitempty"`

	// DisablePerforationOCR When enabled, OCR of perforated fields in the document template will not be performed. Disabled by default.
	DisablePerforationOCR *bool `json:"disablePerforationOCR,omitempty"`

	// DoDetectCan This parameter allows enabling the CAN (Card Access Number) detection and recognition when using scenarios with document location and MRZ reading, such as the MrzAndLocate scenario.
	DoDetectCan *bool `json:"doDetectCan,omitempty"`

	// DocumentGroupFilter List of specific eligible document types from DocumentType enum to recognize from.  You may, for example, specify only passports to be recognized by setting this property. Empty by default.
	DocumentGroupFilter *[]DocumentType `json:"documentGroupFilter,omitempty"`

	// DocumentIdList List of the document ID's to process. All documents will be processed, if empty.
	DocumentIdList *[]int `json:"documentIdList,omitempty"`

	// DoublePageSpread Enable this option if the image you provide contains double page spread of the passport and you want to process both pages in one go. It makes sense to use it for documents that have meaningful information on both pages, like Russian domestic passport, or some others. Disabled by default.
	// Deprecated:
	DoublePageSpread *bool    `json:"doublePageSpread,omitempty"`
	FaceApi          *FaceApi `json:"faceApi,omitempty"`

	// FastDocDetect When enabled, shorten the list of candidates to process during document detection in a single image process mode. Reduces processing time for specific backgrounds. Enabled by default.
	// Deprecated:
	FastDocDetect *bool `json:"fastDocDetect,omitempty"`

	// FieldTypesFilter List of text field types to extract. If empty, all text fields from template will be extracted. Narrowing the list can shorten processing time. Empty by default.
	FieldTypesFilter *[]TextFieldType `json:"fieldTypesFilter,omitempty"`

	// ForceDocFormat Defining the geometric format of documents in accordance with ISO / IEC 7810
	ForceDocFormat *DocumentFormat `json:"forceDocFormat,omitempty"`

	// ForceDocID Force use of specific template ID and skip document type identification step.
	ForceDocID *int `json:"forceDocID,omitempty"`

	// ForceReadMrzBeforeLocate When enabled, make sure that in series processing MRZ is located fully inside the result document image,  if present on the document. Enabling this option may add extra processing time, by disabling optimizations, but allows more stability in output image quality. Disabled by default.
	ForceReadMrzBeforeLocate *bool `json:"forceReadMrzBeforeLocate,omitempty"`

	// GenerateDoublePageSpreadImage When enabled together with "doublePageSpread" and there is a passport with two pages spread in the image, pages will be cropped, straightened and aligned together, as if the document was captured on a flatbed scanner. Disabled by default.
	GenerateDoublePageSpreadImage *bool `json:"generateDoublePageSpreadImage,omitempty"`

	// IgnoreDeviceIdFromImage This parameter is used to tell the processing engine to ignore any parameters saved in the image when scanned from the document reader device. Default false
	IgnoreDeviceIdFromImage *bool `json:"ignoreDeviceIdFromImage,omitempty"`

	// ImageDpiOutMax This parameter controls maximum resolution in dpi of output images. Resolution will remain original in case 0 is supplied. By default is set to return images in response with resolution not greater than 300 dpi for all scenarios except FullAuth. In FullAuth scenario this limit is 1000 dpi by default.
	ImageDpiOutMax *int `json:"imageDpiOutMax,omitempty"`

	// ImageOutputMaxHeight This parameter allows setting maximum height in pixels of output images and thus reducing image size to desired. Does not change the aspect ratio. Changes disabled if equals to 0. Default 0.
	ImageOutputMaxHeight *int `json:"imageOutputMaxHeight,omitempty"`

	// ImageOutputMaxWidth This parameter allows setting maximum width in pixels of output images and thus reducing image size to desired. Does not change the aspect ratio. Changes disabled if equals to 0. Default 0.
	ImageOutputMaxWidth *int     `json:"imageOutputMaxWidth,omitempty"`
	ImageQa             *ImageQA `json:"imageQa,omitempty"`

	// LcidFilter The list of LCID types to recognize. If empty, values with all LCID types will be extracted. Empty by default.
	LcidFilter *[]int `json:"lcidFilter,omitempty"`

	// LcidIgnoreFilter The list of LCID types to ignore during the recognition. If empty, values with all LCID types will be extracted. Narrowing down the list can reduce processing time. Empty by default.
	LcidIgnoreFilter *[]int `json:"lcidIgnoreFilter,omitempty"`

	// Log When enabled, results will contain transaction processing log. Disabled by default
	Log      *bool     `json:"log,omitempty"`
	LogLevel *LogLevel `json:"logLevel,omitempty"`

	// MatchTextFieldMask When disabled, text field OCR will be done as is and then the recognized value will be matched to the field mask for validity. If enabled, we are trying to read a field value with maximum efforts to match the mask and provide a correctly formatted value, making assumptions based on the provided field mask in the template. Enabled by default.
	MatchTextFieldMask *bool          `json:"matchTextFieldMask,omitempty"`
	MeasureSystem      *MeasureSystem `json:"measureSystem,omitempty"`

	// MinimalHolderAge This options allows specifying the minimal age in years of the document holder for the document to be considered valid.
	MinimalHolderAge *int `json:"minimalHolderAge,omitempty"`

	// MrzDetectMode Make better MRZ detection on complex noisy backgrounds, like BW photocopy of some documents.
	MrzDetectMode *MrzDetectModeEnum `json:"mrzDetectMode,omitempty"`

	// MrzFormatsFilter This option allows limiting MRZ formats to be recognized by specifying them in array.
	MrzFormatsFilter *[]MRZFormat `json:"mrzFormatsFilter,omitempty"`

	// MultiDocOnImage This option allows locating and cropping multiple documents from one image if enabled. Disabled by default.
	MultiDocOnImage *bool `json:"multiDocOnImage,omitempty"`

	// NoGraphics When enabled, no graphic fields will be cropped from document image. Disabled by default.
	NoGraphics *bool `json:"noGraphics,omitempty"`

	// OneShotIdentification This parameter allows processing an image that contains a person and a document and compare the portrait photo from the document with the person's face
	OneShotIdentification *bool `json:"oneShotIdentification,omitempty"`

	// ParseBarcodes This option can be disabled to stop parsing after barcode is read. Enabled by default.
	ParseBarcodes *bool `json:"parseBarcodes,omitempty"`

	// ProcessAuth Authenticity checks that should be performed regardless of the document type. The available checks are listed in the eRPRM_Authenticity enum. Note that only supported by your license checks can be added.
	ProcessAuth *int64 `json:"processAuth,omitempty"`

	// RespectImageQuality When enabled, image quality checks status affects document optical and overall status. Disabled by default.
	RespectImageQuality *bool `json:"respectImageQuality,omitempty"`

	// ResultTypeOutput Types of results to return in response. See 'Result' enum for available options
	ResultTypeOutput *[]Result `json:"resultTypeOutput,omitempty"`

	// ReturnCroppedBarcode When enabled, returns cropped barcode images for unknown documents. Disabled by default.
	ReturnCroppedBarcode *bool `json:"returnCroppedBarcode,omitempty"`

	// ReturnUncroppedImage When enabled, returns input images in output. Disabled by default.
	ReturnUncroppedImage *bool `json:"returnUncroppedImage,omitempty"`

	// Rfid Params for the RFID chip data reprocessing
	Rfid *struct {
		// PaSensitiveCodesDisable A list of notification codes that should be ignored during passive authentication (PA)
		PaSensitiveCodesDisable *[]ParsingNotificationCodes `json:"paSensitiveCodesDisable,omitempty"`
	} `json:"rfid,omitempty"`

	// Scenario Document processing scenario
	Scenario Scenario `json:"scenario"`

	// ShiftExpiryDate This option allows shifting the date of expiry into the future or past for number of months specified. This is useful, for example, in some cases when document might be still valid for some period after original expiration date to prevent negative validity status for such documents. Or by shifting the date to the past will set negative validity for the documents that is about to expire in a specified number of months. 0 by default
	ShiftExpiryDate *int `json:"shiftExpiryDate,omitempty"`

	// SplitNames When enabled, the Surname and GivenNames fields from MRZ will be divided into ft_First_Name, ft_Second_Name, ft_Third_Name, ft_Fourth_Name, ft_Last_Name fields. Disabled by default.
	SplitNames *bool `json:"splitNames,omitempty"`

	// UpdateOCRValidityByGlare When enabled, fail OCR field validity, if there is a glare over the text field on the image. Disabled by default.
	UpdateOCRValidityByGlare *bool `json:"updateOCRValidityByGlare,omitempty"`

	// UseFaceApi This parameter allows comparing faces on Regula Face Web Service
	UseFaceApi *bool `json:"useFaceApi,omitempty"`
}

// ProcessRequest defines model for ProcessRequest.
type ProcessRequest struct {
	// ContainerList List with various objects, containing processing results
	ContainerList *ContainerList         `json:"ContainerList,omitempty"`
	List          *[]ProcessRequestImage `json:"List,omitempty"`

	// ExtPortrait Portrait photo from an external source
	ExtPortrait *string `json:"extPortrait,omitempty"`

	// LivePortrait Live portrait photo
	LivePortrait *string `json:"livePortrait,omitempty"`

	// PassBackObject Free-form object to be included in response. Must be object, not list or simple value. Do not affect document processing. Use it freely to pass your app params. Stored in process logs.
	PassBackObject *map[string]interface{} `json:"passBackObject,omitempty"`
	ProcessParam   ProcessParams           `json:"processParam"`
	SystemInfo     *ProcessSystemInfo      `json:"systemInfo,omitempty"`

	// Tag session id
	Tag *string `json:"tag,omitempty"`
}

// ProcessRequestImage defines model for ProcessRequestImage.
type ProcessRequestImage struct {
	ImageData ImageData `json:"ImageData"`

	// Light Image light index
	Light *Light `json:"light,omitempty"`

	// PageIdx page/image number
	PageIdx *int `json:"page_idx,omitempty"`
}

// ProcessResponse defines model for ProcessResponse.
type ProcessResponse struct {
	// ChipPage Determines the presence and location of an RFID chip in a document. 0 - no rfid chip; 1 - chip is located in the document data page; 2 - chip is located in the back page or inlay of the document
	ChipPage *RfidLocation `json:"ChipPage,omitempty"`

	// ContainerList List with various objects, containing processing results
	ContainerList      ContainerList    `json:"ContainerList"`
	ProcessingFinished ProcessingStatus `json:"ProcessingFinished"`
	TransactionInfo    TransactionInfo  `json:"TransactionInfo"`

	// ElapsedTime Time the document processing has taken, ms.
	ElapsedTime *int `json:"elapsedTime,omitempty"`

	// Log Base64 encoded transaction processing log
	Log                *string `json:"log,omitempty"`
	MorePagesAvailable *int    `json:"morePagesAvailable,omitempty"`

	// PassBackObject Free-form object provided in request. See passBackObject property of ProcessRequest.
	PassBackObject *map[string]interface{} `json:"passBackObject,omitempty"`
}

// ProcessSystemInfo defines model for ProcessSystemInfo.
type ProcessSystemInfo struct {
	// License Base64 encoded license file
	License *string `json:"license,omitempty"`

	// RecaptchaToken For internal use. Demo-sites recaptcha token.
	RecaptchaToken *string `json:"recaptcha_token,omitempty"`
}

// ProcessingStatus defines model for ProcessingStatus.
type ProcessingStatus int

// RawImageContainerList defines model for RawImageContainerList.
type RawImageContainerList struct {
	Count  *int         `json:"Count,omitempty"`
	Images *[]ImageData `json:"Images,omitempty"`
}

// RectangleCoordinates Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
type RectangleCoordinates struct {
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
	Top    int `json:"top"`
}

// Result defines model for Result.
type Result int

// ResultItem Common fields for all result objects
type ResultItem struct {
	BufLength *int `json:"buf_length,omitempty"`
	Light     *int `json:"light,omitempty"`
	ListIdx   *int `json:"list_idx,omitempty"`
	PageIdx   *int `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// RfidLocation Determines the presence and location of an RFID chip in a document. 0 - no rfid chip; 1 - chip is located in the document data page; 2 - chip is located in the back page or inlay of the document
type RfidLocation int

// RfidOrigin Location of data in RFID chip
type RfidOrigin struct {
	// Dg Source data group file
	Dg int `json:"dg"`

	// DgTag Index of the source record of the image with biometric information in the information data group
	DgTag *int `json:"dgTag,omitempty"`

	// EntryView Index of the variant of the biometric data template
	EntryView *int `json:"entryView,omitempty"`

	// TagEntry Index of the template in the record with biometric data
	TagEntry *int `json:"tagEntry,omitempty"`
}

// Scenario Document processing scenario
type Scenario string

// SecurityFeatureResult defines model for SecurityFeatureResult.
type SecurityFeatureResult struct {
	// AreaList Checked fragment coordinates
	AreaList *AreaContainer `json:"AreaList,omitempty"`

	// CriticalFlag Enumeration contains identifiers determining the criticality of the security element
	CriticalFlag *Critical `json:"CriticalFlag,omitempty"`

	// ElementDiagnose Enumeration contains identifiers which determinate the single document element authenticity check outcome reason: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/check-diagnose/
	ElementDiagnose *CheckDiagnose `json:"ElementDiagnose,omitempty"`

	// ElementRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	ElementRect *RectangleCoordinates `json:"ElementRect,omitempty"`

	// ElementResult 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ElementResult *CheckResult `json:"ElementResult,omitempty"`

	// ElementType Enumeration contains identifiers determining type of features for a document authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/security-feature-type/
	ElementType *SecurityFeatureType `json:"ElementType,omitempty"`
	Reserved2   *int                 `json:"Reserved2,omitempty"`

	// Type Same as authenticity result type, but used for safe parsing of not-described values: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/
	Type AuthenticityType `json:"Type"`

	// Visibility Enumeration contains visibility status of the security element
	Visibility *Visibility `json:"Visibility,omitempty"`
}

// SecurityFeatureType Enumeration contains identifiers determining type of features for a document authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/security-feature-type/
type SecurityFeatureType int

// Source Document data sources
type Source string

// SourceValidity defines model for SourceValidity.
type SourceValidity struct {
	// Source Document data sources
	Source Source `json:"source"`

	// Status 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Status CheckResult `json:"status"`
}

// Status defines model for Status.
type Status struct {
	DetailsOptical DetailsOptical `json:"detailsOptical"`
	DetailsRFID    *DetailsRFID   `json:"detailsRFID,omitempty"`

	// Optical 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Optical CheckResult `json:"optical"`

	// OverallStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	OverallStatus CheckResult `json:"overallStatus"`

	// Portrait 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Portrait *CheckResult `json:"portrait,omitempty"`

	// Rfid 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Rfid *CheckResult `json:"rfid,omitempty"`

	// StopList 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	StopList *CheckResult `json:"stopList,omitempty"`
}

// StatusResult defines model for StatusResult.
type StatusResult struct {
	Status    SchemasStatus `json:"Status"`
	BufLength *int          `json:"buf_length,omitempty"`
	Light     *int          `json:"light,omitempty"`
	ListIdx   *int          `json:"list_idx,omitempty"`
	PageIdx   *int          `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// StringRecognitionResult Describes single row recognition results in multi-line text field of a document
type StringRecognitionResult struct {
	// StringResult Array of recognition results for individual characters of a string
	StringResult []SymbolRecognitionResult `json:"StringResult"`
}

// SymbolCandidate Describes an individual character recognition candidate
type SymbolCandidate struct {
	// SymbolCode Unicode symbol code
	SymbolCode int `json:"SymbolCode"`

	// SymbolProbability character recognition probability (0–100,%)
	SymbolProbability int `json:"SymbolProbability"`
}

// SymbolRecognitionResult Describes a single character recognition results in the text field line
type SymbolRecognitionResult struct {
	// ListOfCandidates Array of candidate characters. Sorted in descending order of recognition probabilities (the first element has highest probability)
	ListOfCandidates []SymbolCandidate `json:"ListOfCandidates"`

	// SymbolRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	SymbolRect RectangleCoordinates `json:"SymbolRect"`
}

// Text Contains all document text fields data with validity and cross-source compare checks
type Text struct {
	AvailableSourceList []TextAvailableSource `json:"availableSourceList"`

	// ComparisonStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ComparisonStatus CheckResult `json:"comparisonStatus"`
	FieldList        []TextField `json:"fieldList"`

	// Status 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Status CheckResult `json:"status"`

	// ValidityStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ValidityStatus CheckResult `json:"validityStatus"`
}

// TextAvailableSource defines model for TextAvailableSource.
type TextAvailableSource struct {
	// ContainerType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ContainerType *ContainerType `json:"containerType,omitempty"`

	// Source Document data sources
	Source Source `json:"source"`

	// ValidityStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ValidityStatus CheckResult `json:"validityStatus"`
}

// TextDataResult defines model for TextDataResult.
type TextDataResult struct {
	// DocVisualExtendedInfo Container for extracted text fields. Fields are identified by type and language
	DocVisualExtendedInfo *DocVisualExtendedInfo `json:"DocVisualExtendedInfo,omitempty"`
	BufLength             *int                   `json:"buf_length,omitempty"`
	Light                 *int                   `json:"light,omitempty"`
	ListIdx               *int                   `json:"list_idx,omitempty"`
	PageIdx               *int                   `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// TextField defines model for TextField.
type TextField struct {
	ComparisonList []CrossSourceValueComparison `json:"comparisonList"`

	// ComparisonStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ComparisonStatus CheckResult `json:"comparisonStatus"`

	// FieldName Field name. Only use to search values for fields with fieldType=50(other). In general, use fieldType for lookup.
	FieldName string        `json:"fieldName"`
	FieldType TextFieldType `json:"fieldType"`

	// Lcid Locale id. Used to tag same typed fields declared in several languages.
	// For example: name can be provided in both native and latin variants.
	// Based on Microsoft locale id (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/70feba9f-294e-491e-b6eb-56532684c37f).
	Lcid *LCID `json:"lcid,omitempty"`

	// Status 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Status CheckResult `json:"status"`

	// ValidityList Validity of all field values for given source. If there are two values on different pages for one field-source pair, then validity also will include logical match checking. If such values do not match, validity will return error.
	ValidityList []SourceValidity `json:"validityList"`

	// ValidityStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ValidityStatus CheckResult `json:"validityStatus"`

	// Value The most confidence value, selected from valueList
	Value     string           `json:"value"`
	ValueList []TextFieldValue `json:"valueList"`
}

// TextFieldType defines model for TextFieldType.
type TextFieldType int

// TextFieldValue defines model for TextFieldValue.
type TextFieldValue struct {
	// FieldRect Coordinates of the rectangle region on a document image(result type 1). Represented by two points - (left, top) + (right, bottom)
	FieldRect       *RectangleCoordinates `json:"fieldRect,omitempty"`
	OriginalSymbols *[]OriginalSymbol     `json:"originalSymbols,omitempty"`

	// OriginalValue Original value as seen in the document
	OriginalValue *string `json:"originalValue,omitempty"`

	// PageIndex Page index of the image from input list
	PageIndex PageIndex `json:"pageIndex"`

	// Probability Min recognition probability. Combined minimum probability from single characters probabilities
	Probability *int `json:"probability,omitempty"`

	// RfidOrigin Location of data in RFID chip
	RfidOrigin *RfidOrigin `json:"rfidOrigin,omitempty"`

	// Source Document data sources
	Source Source `json:"source"`

	// Value Parsed/processed value. Date format converted for output, delimiters removed
	Value string `json:"value"`
}

// TextPostProcessing defines model for TextPostProcessing.
type TextPostProcessing int

// TextResult defines model for TextResult.
type TextResult struct {
	// Text Contains all document text fields data with validity and cross-source compare checks
	Text      SchemasText `json:"Text"`
	BufLength *int        `json:"buf_length,omitempty"`
	Light     *int        `json:"light,omitempty"`
	ListIdx   *int        `json:"list_idx,omitempty"`
	PageIdx   *int        `json:"page_idx,omitempty"`

	// ResultType Same as Result type, but used for safe parsing of not-described values. See Result type.
	ResultType ContainerType `json:"result_type"`
}

// TransactionImage defines model for TransactionImage.
type TransactionImage = ImageTransactionData

// TransactionInfo defines model for TransactionInfo.
type TransactionInfo struct {
	ComputerName  *string `json:"ComputerName,omitempty"`
	DateTime      *string `json:"DateTime,omitempty"`
	TransactionID *string `json:"TransactionID,omitempty"`
	UserName      *string `json:"UserName,omitempty"`
}

// TransactionProcessGetResponse defines model for TransactionProcessGetResponse.
type TransactionProcessGetResponse struct {
	InData        *InData  `json:"inData,omitempty"`
	OutData       *OutData `json:"outData,omitempty"`
	Tag           *string  `json:"tag,omitempty"`
	TransactionId *int     `json:"transactionId,omitempty"`
}

// TransactionProcessRequest defines model for TransactionProcessRequest.
type TransactionProcessRequest struct {
	// ContainerList List with various objects, containing processing results
	ContainerList *ContainerList         `json:"ContainerList,omitempty"`
	List          *[]ProcessRequestImage `json:"List,omitempty"`

	// ExtPortrait Portrait photo from an external source
	ExtPortrait *string `json:"extPortrait,omitempty"`

	// LivePortrait Live portrait photo
	LivePortrait *string `json:"livePortrait,omitempty"`

	// PassBackObject Free-form object to be included in response. Must be object, not list or simple value. Do not affect document processing. Use it freely to pass your app params. Stored in process logs.
	PassBackObject *map[string]interface{} `json:"passBackObject,omitempty"`
	ProcessParam   ProcessParams           `json:"processParam"`
	SystemInfo     *ProcessSystemInfo      `json:"systemInfo,omitempty"`

	// Tag session id
	Tag *string `json:"tag,omitempty"`
}

// VerificationResult defines model for VerificationResult.
type VerificationResult int

// VerifiedFieldMap defines model for VerifiedFieldMap.
type VerifiedFieldMap struct {
	FieldType *TextFieldType `json:"FieldType,omitempty"`

	// FieldBarcode Field data extracted from barcode
	FieldBarcode *string `json:"Field_Barcode,omitempty"`

	// FieldMRZ Field data extracted from mrz(machine readable zone)
	FieldMRZ *string `json:"Field_MRZ,omitempty"`

	// FieldRFID Field data extracted from rfid chip
	FieldRFID *string `json:"Field_RFID,omitempty"`

	// FieldVisual Field data extracted from visual zone
	FieldVisual *string `json:"Field_Visual,omitempty"`

	// LCID Locale id. Used to tag same typed fields declared in several languages.
	// For example: name can be provided in both native and latin variants.
	// Based on Microsoft locale id (https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-lcid/70feba9f-294e-491e-b6eb-56532684c37f).
	LCID *LCID `json:"LCID,omitempty"`

	// Matrix results comparison matrix. Elements of the matrix with indices 0, 1, 2, 3 take one of the values Disabled(0), Verified(1) or Not_Verified(2), elements with indices 4, 5, 6, 7, 8 are one of the values Disabled(0), Compare_Match(3) or Compare_Not_Match(4). Elements of the Matrix matrix have the following semantic meaning:
	// - element with index 0 –– the result of verification of data from the MRZ; - 1 –– the result of verification of data from the RFID microcircuit; - 2 –– the result of verification of data from text areas of the document; - 3 –– the result of verification data from barcodes; - 4 - the result of comparing MRZ data and RFID microcircuits; - 5 - the result of comparing MRZ data and text areas of document filling; - 6 - the result of comparing MRZ data and bar codes; - 7 - the result of comparing the data of text areas of the document and the RFID chip; - 8 - the result of comparing the data of the text areas of the document and barcodes; - 9 - the result of comparing the data of the RFID chip and barcodes.
	Matrix *ComparisonMatrix `json:"Matrix,omitempty"`
}

// Visibility Enumeration contains visibility status of the security element
type Visibility int

// BcPDF417INFO defines model for bcPDF417INFO.
type BcPDF417INFO struct {
	Angle        *float32 `json:"Angle,omitempty"`
	BcColumn     *int     `json:"bcColumn,omitempty"`
	BcErrorLevel *int     `json:"bcErrorLevel,omitempty"`
	BcRow        *int     `json:"bcRow,omitempty"`
	MinX         *float32 `json:"minX,omitempty"`
	MinY         *float32 `json:"minY,omitempty"`
}

// BcROIDETECT defines model for bcROI_DETECT.
type BcROIDETECT struct {
	Bottom *int `json:"bottom,omitempty"`
	Left   *int `json:"left,omitempty"`
	Right  *int `json:"right,omitempty"`
	Top    *int `json:"top,omitempty"`
}

// DetailsOptical defines model for detailsOptical.
type DetailsOptical struct {
	// DocType 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	DocType CheckResult `json:"docType"`

	// Expiry 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Expiry CheckResult `json:"expiry"`

	// ImageQA 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ImageQA CheckResult `json:"imageQA"`

	// Mrz 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Mrz CheckResult `json:"mrz"`

	// OverallStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	OverallStatus CheckResult `json:"overallStatus"`

	// PagesCount Number of processed pages in the document
	PagesCount int `json:"pagesCount"`

	// Security 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Security CheckResult `json:"security"`

	// Text 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Text CheckResult `json:"text"`
}

// DetailsRFID defines model for detailsRFID.
type DetailsRFID struct {
	// AA 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	AA CheckResult `json:"AA"`

	// BAC 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	BAC CheckResult `json:"BAC"`

	// CA 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	CA CheckResult `json:"CA"`

	// PA 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	PA CheckResult `json:"PA"`

	// PACE 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	PACE CheckResult `json:"PACE"`

	// TA 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	TA CheckResult `json:"TA"`

	// OverallStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	OverallStatus CheckResult `json:"overallStatus"`
}

// MrzDetectModeEnum Make better MRZ detection on complex noisy backgrounds, like BW photocopy of some documents.
type MrzDetectModeEnum int

// PArrayField defines model for pArrayField.
type PArrayField struct {
	BcAngleDETECT      *float32      `json:"bcAngle_DETECT,omitempty"`
	BcCodeResult       *int          `json:"bcCodeResult,omitempty"`
	BcCountModule      *int          `json:"bcCountModule,omitempty"`
	BcDataModule       *[]DataModule `json:"bcDataModule,omitempty"`
	BcPDF417INFO       *BcPDF417INFO `json:"bcPDF417INFO,omitempty"`
	BcROIDETECT        *BcROIDETECT  `json:"bcROI_DETECT,omitempty"`
	BcTextDecoderTypes *int          `json:"bcTextDecoderTypes,omitempty"`
	BcTextFieldType    *int          `json:"bcTextFieldType,omitempty"`
	BcTypeDECODE       *int          `json:"bcType_DECODE,omitempty"`
	BcTypeDETECT       *int          `json:"bcType_DETECT,omitempty"`
}

// SchemasAuthenticityCheckResult defines model for schemas-AuthenticityCheckResult.
type SchemasAuthenticityCheckResult struct {
	List []SchemasAuthenticityCheckResult_List_Item `json:"List"`

	// Result 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Result CheckResult `json:"Result"`

	// Type Enumeration describes available authenticity checks: https://docs.regulaforensics.com/develop/doc-reader-sdk/web-service/development/enums/authenticity-result-type/.
	Type AuthenticityResultType `json:"Type"`
}

// SchemasAuthenticityCheckResult_List_Item defines model for schemas-AuthenticityCheckResult.List.Item.
type SchemasAuthenticityCheckResult_List_Item struct {
	union json.RawMessage
}

// SchemasImages defines model for schemas-Images.
type SchemasImages struct {
	AvailableSourceList []ImagesAvailableSource `json:"availableSourceList"`
	FieldList           []ImagesField           `json:"fieldList"`
}

// SchemasStatus defines model for schemas-Status.
type SchemasStatus struct {
	DetailsOptical DetailsOptical `json:"detailsOptical"`
	DetailsRFID    *DetailsRFID   `json:"detailsRFID,omitempty"`

	// Optical 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Optical CheckResult `json:"optical"`

	// OverallStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	OverallStatus CheckResult `json:"overallStatus"`

	// Portrait 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Portrait *CheckResult `json:"portrait,omitempty"`

	// Rfid 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Rfid *CheckResult `json:"rfid,omitempty"`

	// StopList 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	StopList *CheckResult `json:"stopList,omitempty"`
}

// SchemasText Contains all document text fields data with validity and cross-source compare checks
type SchemasText struct {
	AvailableSourceList []TextAvailableSource `json:"availableSourceList"`

	// ComparisonStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ComparisonStatus CheckResult `json:"comparisonStatus"`
	FieldList        []TextField `json:"fieldList"`

	// Status 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	Status CheckResult `json:"status"`

	// ValidityStatus 0 - result is negative; 1 - result is positive; 2 - сheck was not performed
	ValidityStatus CheckResult `json:"validityStatus"`
}

// XRequest defines model for x-request.
type XRequest = string

// PingParams defines parameters for Ping.
type PingParams struct {
	XRequestID *XRequest `json:"X-RequestID,omitempty"`
}

// ApiProcessParams defines parameters for ApiProcess.
type ApiProcessParams struct {
	XRequestID *XRequest `json:"X-RequestID,omitempty"`
}

// GetApiV2TransactionTransactionIdFileParams defines parameters for GetApiV2TransactionTransactionIdFile.
type GetApiV2TransactionTransactionIdFileParams struct {
	// Name File name
	Name string `form:"name" json:"name"`
}

// GetApiV2TransactionTransactionIdResultsParams defines parameters for GetApiV2TransactionTransactionIdResults.
type GetApiV2TransactionTransactionIdResultsParams struct {
	// WithImages With base64 images or url
	WithImages *GetApiV2TransactionTransactionIdResultsParamsWithImages `form:"withImages,omitempty" json:"withImages,omitempty"`
}

// GetApiV2TransactionTransactionIdResultsParamsWithImages defines parameters for GetApiV2TransactionTransactionIdResults.
type GetApiV2TransactionTransactionIdResultsParamsWithImages bool

// ApiProcessJSONRequestBody defines body for ApiProcess for application/json ContentType.
type ApiProcessJSONRequestBody = ProcessRequest

// PostApiV2TransactionTransactionIdProcessJSONRequestBody defines body for PostApiV2TransactionTransactionIdProcess for application/json ContentType.
type PostApiV2TransactionTransactionIdProcessJSONRequestBody = TransactionProcessRequest

// AsSecurityFeatureResult returns the union data inside the AuthenticityCheckResult_List_Item as a SecurityFeatureResult
func (t AuthenticityCheckResult_List_Item) AsSecurityFeatureResult() (SecurityFeatureResult, error) {
	var body SecurityFeatureResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSecurityFeatureResult overwrites any union data inside the AuthenticityCheckResult_List_Item as the provided SecurityFeatureResult
func (t *AuthenticityCheckResult_List_Item) FromSecurityFeatureResult(v SecurityFeatureResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsIdentResult returns the union data inside the AuthenticityCheckResult_List_Item as a IdentResult
func (t AuthenticityCheckResult_List_Item) AsIdentResult() (IdentResult, error) {
	var body IdentResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromIdentResult overwrites any union data inside the AuthenticityCheckResult_List_Item as the provided IdentResult
func (t *AuthenticityCheckResult_List_Item) FromIdentResult(v IdentResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsFiberResult returns the union data inside the AuthenticityCheckResult_List_Item as a FiberResult
func (t AuthenticityCheckResult_List_Item) AsFiberResult() (FiberResult, error) {
	var body FiberResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFiberResult overwrites any union data inside the AuthenticityCheckResult_List_Item as the provided FiberResult
func (t *AuthenticityCheckResult_List_Item) FromFiberResult(v FiberResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsOCRSecurityTextResult returns the union data inside the AuthenticityCheckResult_List_Item as a OCRSecurityTextResult
func (t AuthenticityCheckResult_List_Item) AsOCRSecurityTextResult() (OCRSecurityTextResult, error) {
	var body OCRSecurityTextResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOCRSecurityTextResult overwrites any union data inside the AuthenticityCheckResult_List_Item as the provided OCRSecurityTextResult
func (t *AuthenticityCheckResult_List_Item) FromOCRSecurityTextResult(v OCRSecurityTextResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsPhotoIdentResult returns the union data inside the AuthenticityCheckResult_List_Item as a PhotoIdentResult
func (t AuthenticityCheckResult_List_Item) AsPhotoIdentResult() (PhotoIdentResult, error) {
	var body PhotoIdentResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPhotoIdentResult overwrites any union data inside the AuthenticityCheckResult_List_Item as the provided PhotoIdentResult
func (t *AuthenticityCheckResult_List_Item) FromPhotoIdentResult(v PhotoIdentResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

func (t AuthenticityCheckResult_List_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthenticityCheckResult_List_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsStatusResult returns the union data inside the ContainerList_List_Item as a StatusResult
func (t ContainerList_List_Item) AsStatusResult() (StatusResult, error) {
	var body StatusResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromStatusResult overwrites any union data inside the ContainerList_List_Item as the provided StatusResult
func (t *ContainerList_List_Item) FromStatusResult(v StatusResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsTextResult returns the union data inside the ContainerList_List_Item as a TextResult
func (t ContainerList_List_Item) AsTextResult() (TextResult, error) {
	var body TextResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextResult overwrites any union data inside the ContainerList_List_Item as the provided TextResult
func (t *ContainerList_List_Item) FromTextResult(v TextResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsDocumentImageResult returns the union data inside the ContainerList_List_Item as a DocumentImageResult
func (t ContainerList_List_Item) AsDocumentImageResult() (DocumentImageResult, error) {
	var body DocumentImageResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDocumentImageResult overwrites any union data inside the ContainerList_List_Item as the provided DocumentImageResult
func (t *ContainerList_List_Item) FromDocumentImageResult(v DocumentImageResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsImagesResult returns the union data inside the ContainerList_List_Item as a ImagesResult
func (t ContainerList_List_Item) AsImagesResult() (ImagesResult, error) {
	var body ImagesResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromImagesResult overwrites any union data inside the ContainerList_List_Item as the provided ImagesResult
func (t *ContainerList_List_Item) FromImagesResult(v ImagesResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsChosenDocumentTypeResult returns the union data inside the ContainerList_List_Item as a ChosenDocumentTypeResult
func (t ContainerList_List_Item) AsChosenDocumentTypeResult() (ChosenDocumentTypeResult, error) {
	var body ChosenDocumentTypeResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromChosenDocumentTypeResult overwrites any union data inside the ContainerList_List_Item as the provided ChosenDocumentTypeResult
func (t *ContainerList_List_Item) FromChosenDocumentTypeResult(v ChosenDocumentTypeResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsDocumentTypesCandidatesResult returns the union data inside the ContainerList_List_Item as a DocumentTypesCandidatesResult
func (t ContainerList_List_Item) AsDocumentTypesCandidatesResult() (DocumentTypesCandidatesResult, error) {
	var body DocumentTypesCandidatesResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDocumentTypesCandidatesResult overwrites any union data inside the ContainerList_List_Item as the provided DocumentTypesCandidatesResult
func (t *ContainerList_List_Item) FromDocumentTypesCandidatesResult(v DocumentTypesCandidatesResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsTextDataResult returns the union data inside the ContainerList_List_Item as a TextDataResult
func (t ContainerList_List_Item) AsTextDataResult() (TextDataResult, error) {
	var body TextDataResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextDataResult overwrites any union data inside the ContainerList_List_Item as the provided TextDataResult
func (t *ContainerList_List_Item) FromTextDataResult(v TextDataResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsGraphicsResult returns the union data inside the ContainerList_List_Item as a GraphicsResult
func (t ContainerList_List_Item) AsGraphicsResult() (GraphicsResult, error) {
	var body GraphicsResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGraphicsResult overwrites any union data inside the ContainerList_List_Item as the provided GraphicsResult
func (t *ContainerList_List_Item) FromGraphicsResult(v GraphicsResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsLexicalAnalysisResult returns the union data inside the ContainerList_List_Item as a LexicalAnalysisResult
func (t ContainerList_List_Item) AsLexicalAnalysisResult() (LexicalAnalysisResult, error) {
	var body LexicalAnalysisResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLexicalAnalysisResult overwrites any union data inside the ContainerList_List_Item as the provided LexicalAnalysisResult
func (t *ContainerList_List_Item) FromLexicalAnalysisResult(v LexicalAnalysisResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsAuthenticityResult returns the union data inside the ContainerList_List_Item as a AuthenticityResult
func (t ContainerList_List_Item) AsAuthenticityResult() (AuthenticityResult, error) {
	var body AuthenticityResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthenticityResult overwrites any union data inside the ContainerList_List_Item as the provided AuthenticityResult
func (t *ContainerList_List_Item) FromAuthenticityResult(v AuthenticityResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsImageQualityResult returns the union data inside the ContainerList_List_Item as a ImageQualityResult
func (t ContainerList_List_Item) AsImageQualityResult() (ImageQualityResult, error) {
	var body ImageQualityResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromImageQualityResult overwrites any union data inside the ContainerList_List_Item as the provided ImageQualityResult
func (t *ContainerList_List_Item) FromImageQualityResult(v ImageQualityResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsDocumentPositionResult returns the union data inside the ContainerList_List_Item as a DocumentPositionResult
func (t ContainerList_List_Item) AsDocumentPositionResult() (DocumentPositionResult, error) {
	var body DocumentPositionResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDocumentPositionResult overwrites any union data inside the ContainerList_List_Item as the provided DocumentPositionResult
func (t *ContainerList_List_Item) FromDocumentPositionResult(v DocumentPositionResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsDocBarCodeInfo returns the union data inside the ContainerList_List_Item as a DocBarCodeInfo
func (t ContainerList_List_Item) AsDocBarCodeInfo() (DocBarCodeInfo, error) {
	var body DocBarCodeInfo
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromDocBarCodeInfo overwrites any union data inside the ContainerList_List_Item as the provided DocBarCodeInfo
func (t *ContainerList_List_Item) FromDocBarCodeInfo(v DocBarCodeInfo) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsLicenseResult returns the union data inside the ContainerList_List_Item as a LicenseResult
func (t ContainerList_List_Item) AsLicenseResult() (LicenseResult, error) {
	var body LicenseResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLicenseResult overwrites any union data inside the ContainerList_List_Item as the provided LicenseResult
func (t *ContainerList_List_Item) FromLicenseResult(v LicenseResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsEncryptedRCLResult returns the union data inside the ContainerList_List_Item as a EncryptedRCLResult
func (t ContainerList_List_Item) AsEncryptedRCLResult() (EncryptedRCLResult, error) {
	var body EncryptedRCLResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEncryptedRCLResult overwrites any union data inside the ContainerList_List_Item as the provided EncryptedRCLResult
func (t *ContainerList_List_Item) FromEncryptedRCLResult(v EncryptedRCLResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

func (t ContainerList_List_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContainerList_List_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSecurityFeatureResult returns the union data inside the SchemasAuthenticityCheckResult_List_Item as a SecurityFeatureResult
func (t SchemasAuthenticityCheckResult_List_Item) AsSecurityFeatureResult() (SecurityFeatureResult, error) {
	var body SecurityFeatureResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSecurityFeatureResult overwrites any union data inside the SchemasAuthenticityCheckResult_List_Item as the provided SecurityFeatureResult
func (t *SchemasAuthenticityCheckResult_List_Item) FromSecurityFeatureResult(v SecurityFeatureResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsIdentResult returns the union data inside the SchemasAuthenticityCheckResult_List_Item as a IdentResult
func (t SchemasAuthenticityCheckResult_List_Item) AsIdentResult() (IdentResult, error) {
	var body IdentResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromIdentResult overwrites any union data inside the SchemasAuthenticityCheckResult_List_Item as the provided IdentResult
func (t *SchemasAuthenticityCheckResult_List_Item) FromIdentResult(v IdentResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsFiberResult returns the union data inside the SchemasAuthenticityCheckResult_List_Item as a FiberResult
func (t SchemasAuthenticityCheckResult_List_Item) AsFiberResult() (FiberResult, error) {
	var body FiberResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFiberResult overwrites any union data inside the SchemasAuthenticityCheckResult_List_Item as the provided FiberResult
func (t *SchemasAuthenticityCheckResult_List_Item) FromFiberResult(v FiberResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsOCRSecurityTextResult returns the union data inside the SchemasAuthenticityCheckResult_List_Item as a OCRSecurityTextResult
func (t SchemasAuthenticityCheckResult_List_Item) AsOCRSecurityTextResult() (OCRSecurityTextResult, error) {
	var body OCRSecurityTextResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOCRSecurityTextResult overwrites any union data inside the SchemasAuthenticityCheckResult_List_Item as the provided OCRSecurityTextResult
func (t *SchemasAuthenticityCheckResult_List_Item) FromOCRSecurityTextResult(v OCRSecurityTextResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// AsPhotoIdentResult returns the union data inside the SchemasAuthenticityCheckResult_List_Item as a PhotoIdentResult
func (t SchemasAuthenticityCheckResult_List_Item) AsPhotoIdentResult() (PhotoIdentResult, error) {
	var body PhotoIdentResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPhotoIdentResult overwrites any union data inside the SchemasAuthenticityCheckResult_List_Item as the provided PhotoIdentResult
func (t *SchemasAuthenticityCheckResult_List_Item) FromPhotoIdentResult(v PhotoIdentResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

func (t SchemasAuthenticityCheckResult_List_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SchemasAuthenticityCheckResult_List_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
