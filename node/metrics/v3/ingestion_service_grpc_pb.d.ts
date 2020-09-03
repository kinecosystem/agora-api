// package: kin.agora.metrics.v3
// file: metrics/v3/ingestion_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as metrics_v3_ingestion_service_pb from "../../metrics/v3/ingestion_service_pb";

interface IIngestionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    submit: IIngestionService_ISubmit;
}

interface IIngestionService_ISubmit extends grpc.MethodDefinition<metrics_v3_ingestion_service_pb.SubmitRequest, metrics_v3_ingestion_service_pb.SubmitResponse> {
    path: string; // "/kin.agora.metrics.v3.Ingestion/Submit"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<metrics_v3_ingestion_service_pb.SubmitRequest>;
    requestDeserialize: grpc.deserialize<metrics_v3_ingestion_service_pb.SubmitRequest>;
    responseSerialize: grpc.serialize<metrics_v3_ingestion_service_pb.SubmitResponse>;
    responseDeserialize: grpc.deserialize<metrics_v3_ingestion_service_pb.SubmitResponse>;
}

export const IngestionService: IIngestionService;

export interface IIngestionServer {
    submit: grpc.handleUnaryCall<metrics_v3_ingestion_service_pb.SubmitRequest, metrics_v3_ingestion_service_pb.SubmitResponse>;
}

export interface IIngestionClient {
    submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
    submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
    submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
}

export class IngestionClient extends grpc.Client implements IIngestionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
    public submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
    public submit(request: metrics_v3_ingestion_service_pb.SubmitRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: metrics_v3_ingestion_service_pb.SubmitResponse) => void): grpc.ClientUnaryCall;
}
