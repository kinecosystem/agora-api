// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var airdrop_v4_airdrop_service_pb = require('../../airdrop/v4/airdrop_service_pb.js');
var validate_validate_pb = require('../../validate/validate_pb.js');
var common_v4_model_pb = require('../../common/v4/model_pb.js');

function serialize_kin_agora_airdrop_v4_RequestAirdropRequest(arg) {
  if (!(arg instanceof airdrop_v4_airdrop_service_pb.RequestAirdropRequest)) {
    throw new Error('Expected argument of type kin.agora.airdrop.v4.RequestAirdropRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_airdrop_v4_RequestAirdropRequest(buffer_arg) {
  return airdrop_v4_airdrop_service_pb.RequestAirdropRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_airdrop_v4_RequestAirdropResponse(arg) {
  if (!(arg instanceof airdrop_v4_airdrop_service_pb.RequestAirdropResponse)) {
    throw new Error('Expected argument of type kin.agora.airdrop.v4.RequestAirdropResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_airdrop_v4_RequestAirdropResponse(buffer_arg) {
  return airdrop_v4_airdrop_service_pb.RequestAirdropResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Airdrop service sends Kin to accounts for free. It is only available on development networks.
var AirdropService = exports.AirdropService = {
  // RequestAirdrop requests an air drop of kin to the target account.
requestAirdrop: {
    path: '/kin.agora.airdrop.v4.Airdrop/RequestAirdrop',
    requestStream: false,
    responseStream: false,
    requestType: airdrop_v4_airdrop_service_pb.RequestAirdropRequest,
    responseType: airdrop_v4_airdrop_service_pb.RequestAirdropResponse,
    requestSerialize: serialize_kin_agora_airdrop_v4_RequestAirdropRequest,
    requestDeserialize: deserialize_kin_agora_airdrop_v4_RequestAirdropRequest,
    responseSerialize: serialize_kin_agora_airdrop_v4_RequestAirdropResponse,
    responseDeserialize: deserialize_kin_agora_airdrop_v4_RequestAirdropResponse,
  },
};

exports.AirdropClient = grpc.makeGenericClientConstructor(AirdropService);
