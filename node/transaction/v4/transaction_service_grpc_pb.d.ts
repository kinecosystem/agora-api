// package: kin.agora.transaction.v4
// file: transaction/v4/transaction_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as transaction_v4_transaction_service_pb from "../../transaction/v4/transaction_service_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v3_model_pb from "../../common/v3/model_pb";
import * as common_v4_model_pb from "../../common/v4/model_pb";

interface ITransactionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getServiceConfig: ITransactionService_IGetServiceConfig;
    getMinimumKinVersion: ITransactionService_IGetMinimumKinVersion;
    getRecentBlockhash: ITransactionService_IGetRecentBlockhash;
    getMinimumBalanceForRentExemption: ITransactionService_IGetMinimumBalanceForRentExemption;
    getHistory: ITransactionService_IGetHistory;
    signTransaction: ITransactionService_ISignTransaction;
    submitTransaction: ITransactionService_ISubmitTransaction;
    getTransaction: ITransactionService_IGetTransaction;
}

interface ITransactionService_IGetServiceConfig extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetServiceConfigRequest, transaction_v4_transaction_service_pb.GetServiceConfigResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetServiceConfig";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetServiceConfigRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetServiceConfigRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetServiceConfigResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetServiceConfigResponse>;
}
interface ITransactionService_IGetMinimumKinVersion extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetMinimumKinVersion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse>;
}
interface ITransactionService_IGetRecentBlockhash extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, transaction_v4_transaction_service_pb.GetRecentBlockhashResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetRecentBlockhash";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetRecentBlockhashRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetRecentBlockhashRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetRecentBlockhashResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetRecentBlockhashResponse>;
}
interface ITransactionService_IGetMinimumBalanceForRentExemption extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetMinimumBalanceForRentExemption";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse>;
}
interface ITransactionService_IGetHistory extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetHistoryRequest, transaction_v4_transaction_service_pb.GetHistoryResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetHistory";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetHistoryRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetHistoryRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetHistoryResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetHistoryResponse>;
}
interface ITransactionService_ISignTransaction extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.SignTransactionRequest, transaction_v4_transaction_service_pb.SignTransactionResponse> {
    path: "/kin.agora.transaction.v4.Transaction/SignTransaction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.SignTransactionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.SignTransactionRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.SignTransactionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.SignTransactionResponse>;
}
interface ITransactionService_ISubmitTransaction extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.SubmitTransactionRequest, transaction_v4_transaction_service_pb.SubmitTransactionResponse> {
    path: "/kin.agora.transaction.v4.Transaction/SubmitTransaction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.SubmitTransactionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.SubmitTransactionRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.SubmitTransactionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.SubmitTransactionResponse>;
}
interface ITransactionService_IGetTransaction extends grpc.MethodDefinition<transaction_v4_transaction_service_pb.GetTransactionRequest, transaction_v4_transaction_service_pb.GetTransactionResponse> {
    path: "/kin.agora.transaction.v4.Transaction/GetTransaction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetTransactionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetTransactionRequest>;
    responseSerialize: grpc.serialize<transaction_v4_transaction_service_pb.GetTransactionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v4_transaction_service_pb.GetTransactionResponse>;
}

export const TransactionService: ITransactionService;

export interface ITransactionServer extends grpc.UntypedServiceImplementation {
    getServiceConfig: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetServiceConfigRequest, transaction_v4_transaction_service_pb.GetServiceConfigResponse>;
    getMinimumKinVersion: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse>;
    getRecentBlockhash: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, transaction_v4_transaction_service_pb.GetRecentBlockhashResponse>;
    getMinimumBalanceForRentExemption: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse>;
    getHistory: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetHistoryRequest, transaction_v4_transaction_service_pb.GetHistoryResponse>;
    signTransaction: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.SignTransactionRequest, transaction_v4_transaction_service_pb.SignTransactionResponse>;
    submitTransaction: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.SubmitTransactionRequest, transaction_v4_transaction_service_pb.SubmitTransactionResponse>;
    getTransaction: grpc.handleUnaryCall<transaction_v4_transaction_service_pb.GetTransactionRequest, transaction_v4_transaction_service_pb.GetTransactionResponse>;
}

export interface ITransactionClient {
    getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
}

export class TransactionClient extends grpc.Client implements ITransactionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    public getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    public getServiceConfig(request: transaction_v4_transaction_service_pb.GetServiceConfigRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetServiceConfigResponse) => void): grpc.ClientUnaryCall;
    public getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    public getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    public getMinimumKinVersion(request: transaction_v4_transaction_service_pb.GetMinimumKinVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumKinVersionResponse) => void): grpc.ClientUnaryCall;
    public getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    public getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    public getRecentBlockhash(request: transaction_v4_transaction_service_pb.GetRecentBlockhashRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetRecentBlockhashResponse) => void): grpc.ClientUnaryCall;
    public getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    public getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    public getMinimumBalanceForRentExemption(request: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetMinimumBalanceForRentExemptionResponse) => void): grpc.ClientUnaryCall;
    public getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public getHistory(request: transaction_v4_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    public signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    public signTransaction(request: transaction_v4_transaction_service_pb.SignTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SignTransactionResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v4_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v4_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v4_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
}
