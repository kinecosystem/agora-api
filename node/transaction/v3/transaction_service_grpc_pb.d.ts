// package: kin.agora.transaction.v3
// file: transaction/v3/transaction_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as transaction_v3_transaction_service_pb from "../../transaction/v3/transaction_service_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v3_model_pb from "../../common/v3/model_pb";

interface ITransactionService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getHistory: ITransactionService_IGetHistory;
    submitTransaction: ITransactionService_ISubmitTransaction;
    getTransaction: ITransactionService_IGetTransaction;
}

interface ITransactionService_IGetHistory extends grpc.MethodDefinition<transaction_v3_transaction_service_pb.GetHistoryRequest, transaction_v3_transaction_service_pb.GetHistoryResponse> {
    path: string; // "/kin.agora.transaction.v3.Transaction/GetHistory"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v3_transaction_service_pb.GetHistoryRequest>;
    requestDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.GetHistoryRequest>;
    responseSerialize: grpc.serialize<transaction_v3_transaction_service_pb.GetHistoryResponse>;
    responseDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.GetHistoryResponse>;
}
interface ITransactionService_ISubmitTransaction extends grpc.MethodDefinition<transaction_v3_transaction_service_pb.SubmitTransactionRequest, transaction_v3_transaction_service_pb.SubmitTransactionResponse> {
    path: string; // "/kin.agora.transaction.v3.Transaction/SubmitTransaction"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v3_transaction_service_pb.SubmitTransactionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.SubmitTransactionRequest>;
    responseSerialize: grpc.serialize<transaction_v3_transaction_service_pb.SubmitTransactionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.SubmitTransactionResponse>;
}
interface ITransactionService_IGetTransaction extends grpc.MethodDefinition<transaction_v3_transaction_service_pb.GetTransactionRequest, transaction_v3_transaction_service_pb.GetTransactionResponse> {
    path: string; // "/kin.agora.transaction.v3.Transaction/GetTransaction"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<transaction_v3_transaction_service_pb.GetTransactionRequest>;
    requestDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.GetTransactionRequest>;
    responseSerialize: grpc.serialize<transaction_v3_transaction_service_pb.GetTransactionResponse>;
    responseDeserialize: grpc.deserialize<transaction_v3_transaction_service_pb.GetTransactionResponse>;
}

export const TransactionService: ITransactionService;

export interface ITransactionServer {
    getHistory: grpc.handleUnaryCall<transaction_v3_transaction_service_pb.GetHistoryRequest, transaction_v3_transaction_service_pb.GetHistoryResponse>;
    submitTransaction: grpc.handleUnaryCall<transaction_v3_transaction_service_pb.SubmitTransactionRequest, transaction_v3_transaction_service_pb.SubmitTransactionResponse>;
    getTransaction: grpc.handleUnaryCall<transaction_v3_transaction_service_pb.GetTransactionRequest, transaction_v3_transaction_service_pb.GetTransactionResponse>;
}

export interface ITransactionClient {
    getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
}

export class TransactionClient extends grpc.Client implements ITransactionClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public getHistory(request: transaction_v3_transaction_service_pb.GetHistoryRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetHistoryResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public submitTransaction(request: transaction_v3_transaction_service_pb.SubmitTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.SubmitTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
    public getTransaction(request: transaction_v3_transaction_service_pb.GetTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: transaction_v3_transaction_service_pb.GetTransactionResponse) => void): grpc.ClientUnaryCall;
}
