/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-PDU-Contents"
 * 	found in "../../../../api/e2ap/v2/e2ap_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RICsubscription_List_withCause_H_
#define	_RICsubscription_List_withCause_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ProtocolIE_SingleContainer;

/* RICsubscription-List-withCause */
typedef struct RICsubscription_List_withCause {
	A_SEQUENCE_OF(struct ProtocolIE_SingleContainer) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RICsubscription_List_withCause_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RICsubscription_List_withCause;

#ifdef __cplusplus
}
#endif

#endif	/* _RICsubscription_List_withCause_H_ */
#include "asn_internal.h"
