// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var account_v4_account_service_pb = require('../../account/v4/account_service_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');
var common_v4_model_pb = require('../../common/v4/model_pb.js');

function serialize_kin_agora_account_v4_CreateAccountRequest(arg) {
  if (!(arg instanceof account_v4_account_service_pb.CreateAccountRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v4.CreateAccountRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_CreateAccountRequest(buffer_arg) {
  return account_v4_account_service_pb.CreateAccountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_CreateAccountResponse(arg) {
  if (!(arg instanceof account_v4_account_service_pb.CreateAccountResponse)) {
    throw new Error('Expected argument of type kin.agora.account.v4.CreateAccountResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_CreateAccountResponse(buffer_arg) {
  return account_v4_account_service_pb.CreateAccountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_Events(arg) {
  if (!(arg instanceof account_v4_account_service_pb.Events)) {
    throw new Error('Expected argument of type kin.agora.account.v4.Events');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_Events(buffer_arg) {
  return account_v4_account_service_pb.Events.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_GetAccountInfoRequest(arg) {
  if (!(arg instanceof account_v4_account_service_pb.GetAccountInfoRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v4.GetAccountInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_GetAccountInfoRequest(buffer_arg) {
  return account_v4_account_service_pb.GetAccountInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_GetAccountInfoResponse(arg) {
  if (!(arg instanceof account_v4_account_service_pb.GetAccountInfoResponse)) {
    throw new Error('Expected argument of type kin.agora.account.v4.GetAccountInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_GetAccountInfoResponse(buffer_arg) {
  return account_v4_account_service_pb.GetAccountInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_GetEventsRequest(arg) {
  if (!(arg instanceof account_v4_account_service_pb.GetEventsRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v4.GetEventsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_GetEventsRequest(buffer_arg) {
  return account_v4_account_service_pb.GetEventsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_ResolveTokenAccountsRequest(arg) {
  if (!(arg instanceof account_v4_account_service_pb.ResolveTokenAccountsRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v4.ResolveTokenAccountsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_ResolveTokenAccountsRequest(buffer_arg) {
  return account_v4_account_service_pb.ResolveTokenAccountsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v4_ResolveTokenAccountsResponse(arg) {
  if (!(arg instanceof account_v4_account_service_pb.ResolveTokenAccountsResponse)) {
    throw new Error('Expected argument of type kin.agora.account.v4.ResolveTokenAccountsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v4_ResolveTokenAccountsResponse(buffer_arg) {
  return account_v4_account_service_pb.ResolveTokenAccountsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AccountService = exports.AccountService = {
  // CreateAccount creates a kin token account.
createAccount: {
    path: '/kin.agora.account.v4.Account/CreateAccount',
    requestStream: false,
    responseStream: false,
    requestType: account_v4_account_service_pb.CreateAccountRequest,
    responseType: account_v4_account_service_pb.CreateAccountResponse,
    requestSerialize: serialize_kin_agora_account_v4_CreateAccountRequest,
    requestDeserialize: deserialize_kin_agora_account_v4_CreateAccountRequest,
    responseSerialize: serialize_kin_agora_account_v4_CreateAccountResponse,
    responseDeserialize: deserialize_kin_agora_account_v4_CreateAccountResponse,
  },
  // GetAccountInfo returns the information of a specified account.
getAccountInfo: {
    path: '/kin.agora.account.v4.Account/GetAccountInfo',
    requestStream: false,
    responseStream: false,
    requestType: account_v4_account_service_pb.GetAccountInfoRequest,
    responseType: account_v4_account_service_pb.GetAccountInfoResponse,
    requestSerialize: serialize_kin_agora_account_v4_GetAccountInfoRequest,
    requestDeserialize: deserialize_kin_agora_account_v4_GetAccountInfoRequest,
    responseSerialize: serialize_kin_agora_account_v4_GetAccountInfoResponse,
    responseDeserialize: deserialize_kin_agora_account_v4_GetAccountInfoResponse,
  },
  // ResolveTokenAccounts resolves a set of Token Accounts owned by the specified account ID.
resolveTokenAccounts: {
    path: '/kin.agora.account.v4.Account/ResolveTokenAccounts',
    requestStream: false,
    responseStream: false,
    requestType: account_v4_account_service_pb.ResolveTokenAccountsRequest,
    responseType: account_v4_account_service_pb.ResolveTokenAccountsResponse,
    requestSerialize: serialize_kin_agora_account_v4_ResolveTokenAccountsRequest,
    requestDeserialize: deserialize_kin_agora_account_v4_ResolveTokenAccountsRequest,
    responseSerialize: serialize_kin_agora_account_v4_ResolveTokenAccountsResponse,
    responseDeserialize: deserialize_kin_agora_account_v4_ResolveTokenAccountsResponse,
  },
  // GetEvents returns a stream of events related to the specified account.
//
// Note: Only events occurring after stream initiation will be returned.
getEvents: {
    path: '/kin.agora.account.v4.Account/GetEvents',
    requestStream: false,
    responseStream: true,
    requestType: account_v4_account_service_pb.GetEventsRequest,
    responseType: account_v4_account_service_pb.Events,
    requestSerialize: serialize_kin_agora_account_v4_GetEventsRequest,
    requestDeserialize: deserialize_kin_agora_account_v4_GetEventsRequest,
    responseSerialize: serialize_kin_agora_account_v4_Events,
    responseDeserialize: deserialize_kin_agora_account_v4_Events,
  },
};

exports.AccountClient = grpc.makeGenericClientConstructor(AccountService);
