// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var transaction_v4_transaction_service_pb = require('../../transaction/v4/transaction_service_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');
var common_v3_model_pb = require('../../common/v3/model_pb.js');
var common_v4_model_pb = require('../../common/v4/model_pb.js');

function serialize_kin_agora_transaction_v4_GetHistoryRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetHistoryRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetHistoryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetHistoryRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetHistoryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetHistoryResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetHistoryResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetHistoryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetHistoryResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetHistoryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetMinimumBalanceForRentExemptionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetMinimumBalanceForRentExemptionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetMinimumKinVersionRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetMinimumKinVersionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetMinimumKinVersionRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetMinimumKinVersionResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetMinimumKinVersionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetMinimumKinVersionResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetRecentBlockhashRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetRecentBlockhashRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetRecentBlockhashRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetRecentBlockhashRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetRecentBlockhashRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetRecentBlockhashResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetRecentBlockhashResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetRecentBlockhashResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetRecentBlockhashResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetRecentBlockhashResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetServiceConfigRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetServiceConfigRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetServiceConfigRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetServiceConfigRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetServiceConfigRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetServiceConfigResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetServiceConfigResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetServiceConfigResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetServiceConfigResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetServiceConfigResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetTransactionRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetTransactionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetTransactionRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_GetTransactionResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.GetTransactionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.GetTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_GetTransactionResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.GetTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_SignTransactionRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.SignTransactionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.SignTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_SignTransactionRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.SignTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_SignTransactionResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.SignTransactionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.SignTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_SignTransactionResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.SignTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_SubmitTransactionRequest(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.SubmitTransactionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.SubmitTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_SubmitTransactionRequest(buffer_arg) {
  return transaction_v4_transaction_service_pb.SubmitTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v4_SubmitTransactionResponse(arg) {
  if (!(arg instanceof transaction_v4_transaction_service_pb.SubmitTransactionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v4.SubmitTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v4_SubmitTransactionResponse(buffer_arg) {
  return transaction_v4_transaction_service_pb.SubmitTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TransactionService = exports.TransactionService = {
  // GetServiceConfig returns the service and token parameters for the token.
//
// The subsidizer key returned may vary based on the 'app-index' header.
getServiceConfig: {
    path: '/kin.agora.transaction.v4.Transaction/GetServiceConfig',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetServiceConfigRequest,
    responseType: transaction_v4_transaction_service_pb.GetServiceConfigResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetServiceConfigRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetServiceConfigRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetServiceConfigResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetServiceConfigResponse,
  },
  // GetMinimumKinVersion returns the minimum Kin version that is supported.
//
// This version will _never_ decrease in non-test scenarios, as it indicates
// a global migration has occured.
getMinimumKinVersion: {
    path: '/kin.agora.transaction.v4.Transaction/GetMinimumKinVersion',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest,
    responseType: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetMinimumKinVersionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetMinimumKinVersionRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetMinimumKinVersionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetMinimumKinVersionResponse,
  },
  // GetRecentBlockhash returns a recent block hash from the underlying network,
// which should be used when crafting transactions. If a transaction fails with
// DuplicateSignature or InvalidNonce, it is recommended that a new block hash
// is retrieved.
//
// Block hashes are expected to be valid for ~2 minutes.
getRecentBlockhash: {
    path: '/kin.agora.transaction.v4.Transaction/GetRecentBlockhash',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest,
    responseType: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetRecentBlockhashRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetRecentBlockhashRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetRecentBlockhashResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetRecentBlockhashResponse,
  },
  // GetMinimumBalanceForRentExemption returns the minimum amount of lamports that
// must be in an account for it not to be garbage collected.
getMinimumBalanceForRentExemption: {
    path: '/kin.agora.transaction.v4.Transaction/GetMinimumBalanceForRentExemption',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest,
    responseType: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetMinimumBalanceForRentExemptionResponse,
  },
  // GetHistory returns the transaction history for an account,
// with additional off-chain invoice data, if available.
getHistory: {
    path: '/kin.agora.transaction.v4.Transaction/GetHistory',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetHistoryRequest,
    responseType: transaction_v4_transaction_service_pb.GetHistoryResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetHistoryRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetHistoryRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetHistoryResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetHistoryResponse,
  },
  // SignTransaction signs the provided transaction, returning the signature to be used.
//
// The transaction may include the following types of instructions:
//   - SplAssociateTokenAccount::CreateAssociatedTokenAccount()
//   - SplToken::SetAuthority(CloseAuthority)
//   - SplToken::Transfer()
//   - SplToken::CloseAccount()
//   - Memo::Memo()
//
// The transaction can be divided into one or more 'regions', which are delineated by
// the memo instruction. Each instruction within a region is considered to be 'related to'
// the memo at the beginning of the region. The first (or only) region may not have a memo.
// For example, if there are instructions before the first memo instruction, or if there
// is no memo at all.
//
// If an invoice is applied, there must be a memo whose foreign key contains the SHA-226
// of the serialized memo. Additionally, the number of SplToken::Transfer instructions in
// the region _must_ match the number of invoices. Furthermore, the invoice cannot be
// referenced by more than one region.
//
// Examples:
//
// Basic Transfer (No Invoice)
//   1. SplToken::Transfer()
//
// Basic Transfer (Invoice)
//   1. Memo::Memo(Spend)
//   2. SplToken::Transfer()
//
// Transfer with Cleanup (Sender has token accounts A, B, sending to C)
//   1. Memo::Memo(GC) [Optional, 'memoless' region is ok]
//   2. SplToken::Transfer(B -> A)
//   3. SplToken::CloseAccount(B)
//   4. Memo::Memo(Spend)
//   5. SplToken::Transfer(A -> C)
//
// Transfer with Cleanup At End (Sender has token accounts A, B, sending to C)
//   1. Memo::Memo(Spend)
//   2. SplToken::Transfer(A -> C)
//   3. Memo::Memo(GC) [Required, delineate cleanup region from above]
//   4. SplToken::Transfer(B -> A)
//   5. SplToken::CloseAccount(B)
//
// Sender Creates Destination (No Invoice)
//   1. SplAssociateTokenAccount::CreateAssociatedTokenAccount()
//   2. SplToken::SetAuthority(CloseAuthority)
//   2. SplToken::Transfer()
//
// Sender Creates Destination (Invoice)
//   1. SplAssociateTokenAccount::CreateAssociatedTokenAccount()
//   2. SplToken::SetAuthority(CloseAuthority)
//   3. Memo::Memo(Earn)
//   4. SplToken::Transfer()
signTransaction: {
    path: '/kin.agora.transaction.v4.Transaction/SignTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.SignTransactionRequest,
    responseType: transaction_v4_transaction_service_pb.SignTransactionResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_SignTransactionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_SignTransactionRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_SignTransactionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_SignTransactionResponse,
  },
  // SubmitTransaction submits a transaction.
//
// If the transaction is already signed, the SignTransaction webhook will not
// be called.
submitTransaction: {
    path: '/kin.agora.transaction.v4.Transaction/SubmitTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.SubmitTransactionRequest,
    responseType: transaction_v4_transaction_service_pb.SubmitTransactionResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_SubmitTransactionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_SubmitTransactionRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_SubmitTransactionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_SubmitTransactionResponse,
  },
  // GetTransaction returns a transaction and additional off-chain invoice data, if available.
getTransaction: {
    path: '/kin.agora.transaction.v4.Transaction/GetTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v4_transaction_service_pb.GetTransactionRequest,
    responseType: transaction_v4_transaction_service_pb.GetTransactionResponse,
    requestSerialize: serialize_kin_agora_transaction_v4_GetTransactionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v4_GetTransactionRequest,
    responseSerialize: serialize_kin_agora_transaction_v4_GetTransactionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v4_GetTransactionResponse,
  },
};

exports.TransactionClient = grpc.makeGenericClientConstructor(TransactionService);
