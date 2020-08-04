// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var account_v3_account_service_pb = require('../../account/v3/account_service_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');
var common_v3_model_pb = require('../../common/v3/model_pb.js');

function serialize_kin_agora_account_v3_CreateAccountRequest(arg) {
  if (!(arg instanceof account_v3_account_service_pb.CreateAccountRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v3.CreateAccountRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_CreateAccountRequest(buffer_arg) {
  return account_v3_account_service_pb.CreateAccountRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v3_CreateAccountResponse(arg) {
  if (!(arg instanceof account_v3_account_service_pb.CreateAccountResponse)) {
    throw new Error('Expected argument of type kin.agora.account.v3.CreateAccountResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_CreateAccountResponse(buffer_arg) {
  return account_v3_account_service_pb.CreateAccountResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v3_Events(arg) {
  if (!(arg instanceof account_v3_account_service_pb.Events)) {
    throw new Error('Expected argument of type kin.agora.account.v3.Events');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_Events(buffer_arg) {
  return account_v3_account_service_pb.Events.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v3_GetAccountInfoRequest(arg) {
  if (!(arg instanceof account_v3_account_service_pb.GetAccountInfoRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v3.GetAccountInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_GetAccountInfoRequest(buffer_arg) {
  return account_v3_account_service_pb.GetAccountInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v3_GetAccountInfoResponse(arg) {
  if (!(arg instanceof account_v3_account_service_pb.GetAccountInfoResponse)) {
    throw new Error('Expected argument of type kin.agora.account.v3.GetAccountInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_GetAccountInfoResponse(buffer_arg) {
  return account_v3_account_service_pb.GetAccountInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_account_v3_GetEventsRequest(arg) {
  if (!(arg instanceof account_v3_account_service_pb.GetEventsRequest)) {
    throw new Error('Expected argument of type kin.agora.account.v3.GetEventsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_account_v3_GetEventsRequest(buffer_arg) {
  return account_v3_account_service_pb.GetEventsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var AccountService = exports.AccountService = {
  // CreateAccount creates an account using a seed account configured
// by the service.
createAccount: {
    path: '/kin.agora.account.v3.Account/CreateAccount',
    requestStream: false,
    responseStream: false,
    requestType: account_v3_account_service_pb.CreateAccountRequest,
    responseType: account_v3_account_service_pb.CreateAccountResponse,
    requestSerialize: serialize_kin_agora_account_v3_CreateAccountRequest,
    requestDeserialize: deserialize_kin_agora_account_v3_CreateAccountRequest,
    responseSerialize: serialize_kin_agora_account_v3_CreateAccountResponse,
    responseDeserialize: deserialize_kin_agora_account_v3_CreateAccountResponse,
  },
  // GetAccountInfo returns the information of a specified account.
getAccountInfo: {
    path: '/kin.agora.account.v3.Account/GetAccountInfo',
    requestStream: false,
    responseStream: false,
    requestType: account_v3_account_service_pb.GetAccountInfoRequest,
    responseType: account_v3_account_service_pb.GetAccountInfoResponse,
    requestSerialize: serialize_kin_agora_account_v3_GetAccountInfoRequest,
    requestDeserialize: deserialize_kin_agora_account_v3_GetAccountInfoRequest,
    responseSerialize: serialize_kin_agora_account_v3_GetAccountInfoResponse,
    responseDeserialize: deserialize_kin_agora_account_v3_GetAccountInfoResponse,
  },
  // GetEvents returns a stream of events related to the specified account.
//
// Note: Only events occurring after stream initiation will be returned.
getEvents: {
    path: '/kin.agora.account.v3.Account/GetEvents',
    requestStream: false,
    responseStream: true,
    requestType: account_v3_account_service_pb.GetEventsRequest,
    responseType: account_v3_account_service_pb.Events,
    requestSerialize: serialize_kin_agora_account_v3_GetEventsRequest,
    requestDeserialize: deserialize_kin_agora_account_v3_GetEventsRequest,
    responseSerialize: serialize_kin_agora_account_v3_Events,
    responseDeserialize: deserialize_kin_agora_account_v3_Events,
  },
};

exports.AccountClient = grpc.makeGenericClientConstructor(AccountService);
