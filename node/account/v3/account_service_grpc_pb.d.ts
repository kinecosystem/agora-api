// package: kin.agora.account.v3
// file: account/v3/account_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as account_v3_account_service_pb from "../../account/v3/account_service_pb";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v3_model_pb from "../../common/v3/model_pb";

interface IAccountService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createAccount: IAccountService_ICreateAccount;
    getAccountInfo: IAccountService_IGetAccountInfo;
    getEvents: IAccountService_IGetEvents;
}

interface IAccountService_ICreateAccount extends grpc.MethodDefinition<account_v3_account_service_pb.CreateAccountRequest, account_v3_account_service_pb.CreateAccountResponse> {
    path: string; // "/kin.agora.account.v3.Account/CreateAccount"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<account_v3_account_service_pb.CreateAccountRequest>;
    requestDeserialize: grpc.deserialize<account_v3_account_service_pb.CreateAccountRequest>;
    responseSerialize: grpc.serialize<account_v3_account_service_pb.CreateAccountResponse>;
    responseDeserialize: grpc.deserialize<account_v3_account_service_pb.CreateAccountResponse>;
}
interface IAccountService_IGetAccountInfo extends grpc.MethodDefinition<account_v3_account_service_pb.GetAccountInfoRequest, account_v3_account_service_pb.GetAccountInfoResponse> {
    path: string; // "/kin.agora.account.v3.Account/GetAccountInfo"
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<account_v3_account_service_pb.GetAccountInfoRequest>;
    requestDeserialize: grpc.deserialize<account_v3_account_service_pb.GetAccountInfoRequest>;
    responseSerialize: grpc.serialize<account_v3_account_service_pb.GetAccountInfoResponse>;
    responseDeserialize: grpc.deserialize<account_v3_account_service_pb.GetAccountInfoResponse>;
}
interface IAccountService_IGetEvents extends grpc.MethodDefinition<account_v3_account_service_pb.GetEventsRequest, account_v3_account_service_pb.Events> {
    path: string; // "/kin.agora.account.v3.Account/GetEvents"
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<account_v3_account_service_pb.GetEventsRequest>;
    requestDeserialize: grpc.deserialize<account_v3_account_service_pb.GetEventsRequest>;
    responseSerialize: grpc.serialize<account_v3_account_service_pb.Events>;
    responseDeserialize: grpc.deserialize<account_v3_account_service_pb.Events>;
}

export const AccountService: IAccountService;

export interface IAccountServer {
    createAccount: grpc.handleUnaryCall<account_v3_account_service_pb.CreateAccountRequest, account_v3_account_service_pb.CreateAccountResponse>;
    getAccountInfo: grpc.handleUnaryCall<account_v3_account_service_pb.GetAccountInfoRequest, account_v3_account_service_pb.GetAccountInfoResponse>;
    getEvents: grpc.handleServerStreamingCall<account_v3_account_service_pb.GetEventsRequest, account_v3_account_service_pb.Events>;
}

export interface IAccountClient {
    createAccount(request: account_v3_account_service_pb.CreateAccountRequest, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    createAccount(request: account_v3_account_service_pb.CreateAccountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    createAccount(request: account_v3_account_service_pb.CreateAccountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    getEvents(request: account_v3_account_service_pb.GetEventsRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<account_v3_account_service_pb.Events>;
    getEvents(request: account_v3_account_service_pb.GetEventsRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<account_v3_account_service_pb.Events>;
}

export class AccountClient extends grpc.Client implements IAccountClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public createAccount(request: account_v3_account_service_pb.CreateAccountRequest, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    public createAccount(request: account_v3_account_service_pb.CreateAccountRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    public createAccount(request: account_v3_account_service_pb.CreateAccountRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.CreateAccountResponse) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: account_v3_account_service_pb.GetAccountInfoRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: account_v3_account_service_pb.GetAccountInfoResponse) => void): grpc.ClientUnaryCall;
    public getEvents(request: account_v3_account_service_pb.GetEventsRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<account_v3_account_service_pb.Events>;
    public getEvents(request: account_v3_account_service_pb.GetEventsRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<account_v3_account_service_pb.Events>;
}
