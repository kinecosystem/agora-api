// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var transaction_v3_transaction_service_pb = require('../../transaction/v3/transaction_service_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');
var common_v3_model_pb = require('../../common/v3/model_pb.js');

function serialize_kin_agora_transaction_v3_GetHistoryRequest(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.GetHistoryRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.GetHistoryRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_GetHistoryRequest(buffer_arg) {
  return transaction_v3_transaction_service_pb.GetHistoryRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v3_GetHistoryResponse(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.GetHistoryResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.GetHistoryResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_GetHistoryResponse(buffer_arg) {
  return transaction_v3_transaction_service_pb.GetHistoryResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v3_GetTransactionRequest(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.GetTransactionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.GetTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_GetTransactionRequest(buffer_arg) {
  return transaction_v3_transaction_service_pb.GetTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v3_GetTransactionResponse(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.GetTransactionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.GetTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_GetTransactionResponse(buffer_arg) {
  return transaction_v3_transaction_service_pb.GetTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v3_SubmitTransactionRequest(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.SubmitTransactionRequest)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.SubmitTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_SubmitTransactionRequest(buffer_arg) {
  return transaction_v3_transaction_service_pb.SubmitTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_transaction_v3_SubmitTransactionResponse(arg) {
  if (!(arg instanceof transaction_v3_transaction_service_pb.SubmitTransactionResponse)) {
    throw new Error('Expected argument of type kin.agora.transaction.v3.SubmitTransactionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_transaction_v3_SubmitTransactionResponse(buffer_arg) {
  return transaction_v3_transaction_service_pb.SubmitTransactionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TransactionService = exports.TransactionService = {
  // GetHistory returns the transaction history for an account,
// with additional off-chain data if available.
getHistory: {
    path: '/kin.agora.transaction.v3.Transaction/GetHistory',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v3_transaction_service_pb.GetHistoryRequest,
    responseType: transaction_v3_transaction_service_pb.GetHistoryResponse,
    requestSerialize: serialize_kin_agora_transaction_v3_GetHistoryRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v3_GetHistoryRequest,
    responseSerialize: serialize_kin_agora_transaction_v3_GetHistoryResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v3_GetHistoryResponse,
  },
  // SubmitTransaction submits a transaction, which _may_ be whitelisted.
//
// If the memo does not conform to the 'memo standard' (todo: name),
// the transaction is not eligible for whitelisting.
submitTransaction: {
    path: '/kin.agora.transaction.v3.Transaction/SubmitTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v3_transaction_service_pb.SubmitTransactionRequest,
    responseType: transaction_v3_transaction_service_pb.SubmitTransactionResponse,
    requestSerialize: serialize_kin_agora_transaction_v3_SubmitTransactionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v3_SubmitTransactionRequest,
    responseSerialize: serialize_kin_agora_transaction_v3_SubmitTransactionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v3_SubmitTransactionResponse,
  },
  // GetTransaction blocks until the specified transaction is resolved,
// or until the RPC is cancelled by client / server, or fails.
//
// A transaction is resolved if it has succeeded for failed.
getTransaction: {
    path: '/kin.agora.transaction.v3.Transaction/GetTransaction',
    requestStream: false,
    responseStream: false,
    requestType: transaction_v3_transaction_service_pb.GetTransactionRequest,
    responseType: transaction_v3_transaction_service_pb.GetTransactionResponse,
    requestSerialize: serialize_kin_agora_transaction_v3_GetTransactionRequest,
    requestDeserialize: deserialize_kin_agora_transaction_v3_GetTransactionRequest,
    responseSerialize: serialize_kin_agora_transaction_v3_GetTransactionResponse,
    responseDeserialize: deserialize_kin_agora_transaction_v3_GetTransactionResponse,
  },
};

exports.TransactionClient = grpc.makeGenericClientConstructor(TransactionService);
