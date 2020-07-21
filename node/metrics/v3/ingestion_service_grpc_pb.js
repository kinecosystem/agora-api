// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var metrics_v3_ingestion_service_pb = require('../../metrics/v3/ingestion_service_pb.js');

function serialize_kin_agora_metrics_v3_SubmitRequest(arg) {
  if (!(arg instanceof metrics_v3_ingestion_service_pb.SubmitRequest)) {
    throw new Error('Expected argument of type kin.agora.metrics.v3.SubmitRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_metrics_v3_SubmitRequest(buffer_arg) {
  return metrics_v3_ingestion_service_pb.SubmitRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_kin_agora_metrics_v3_SubmitResponse(arg) {
  if (!(arg instanceof metrics_v3_ingestion_service_pb.SubmitResponse)) {
    throw new Error('Expected argument of type kin.agora.metrics.v3.SubmitResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_kin_agora_metrics_v3_SubmitResponse(buffer_arg) {
  return metrics_v3_ingestion_service_pb.SubmitResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var IngestionService = exports.IngestionService = {
  submit: {
    path: '/kin.agora.metrics.v3.Ingestion/Submit',
    requestStream: false,
    responseStream: false,
    requestType: metrics_v3_ingestion_service_pb.SubmitRequest,
    responseType: metrics_v3_ingestion_service_pb.SubmitResponse,
    requestSerialize: serialize_kin_agora_metrics_v3_SubmitRequest,
    requestDeserialize: deserialize_kin_agora_metrics_v3_SubmitRequest,
    responseSerialize: serialize_kin_agora_metrics_v3_SubmitResponse,
    responseDeserialize: deserialize_kin_agora_metrics_v3_SubmitResponse,
  },
};

exports.IngestionClient = grpc.makeGenericClientConstructor(IngestionService);
