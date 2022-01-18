// package: kin.agora.airdrop.v4
// file: airdrop/v4/airdrop_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as airdrop_v4_airdrop_service_pb from "../../airdrop/v4/airdrop_service_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v4_model_pb from "../../common/v4/model_pb";

interface IAirdropService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    requestAirdrop: IAirdropService_IRequestAirdrop;
}

interface IAirdropService_IRequestAirdrop extends grpc.MethodDefinition<airdrop_v4_airdrop_service_pb.RequestAirdropRequest, airdrop_v4_airdrop_service_pb.RequestAirdropResponse> {
    path: "/kin.agora.airdrop.v4.Airdrop/RequestAirdrop";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<airdrop_v4_airdrop_service_pb.RequestAirdropRequest>;
    requestDeserialize: grpc.deserialize<airdrop_v4_airdrop_service_pb.RequestAirdropRequest>;
    responseSerialize: grpc.serialize<airdrop_v4_airdrop_service_pb.RequestAirdropResponse>;
    responseDeserialize: grpc.deserialize<airdrop_v4_airdrop_service_pb.RequestAirdropResponse>;
}

export const AirdropService: IAirdropService;

export interface IAirdropServer extends grpc.UntypedServiceImplementation {
    requestAirdrop: grpc.handleUnaryCall<airdrop_v4_airdrop_service_pb.RequestAirdropRequest, airdrop_v4_airdrop_service_pb.RequestAirdropResponse>;
}

export interface IAirdropClient {
    requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
    requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
    requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
}

export class AirdropClient extends grpc.Client implements IAirdropClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
    public requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
    public requestAirdrop(request: airdrop_v4_airdrop_service_pb.RequestAirdropRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: airdrop_v4_airdrop_service_pb.RequestAirdropResponse) => void): grpc.ClientUnaryCall;
}
