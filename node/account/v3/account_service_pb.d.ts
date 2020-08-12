// package: kin.agora.account.v3
// file: account/v3/account_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v3_model_pb from "../../common/v3/model_pb";

export class AccountInfo extends jspb.Message { 

    hasAccountId(): boolean;
    clearAccountId(): void;
    getAccountId(): common_v3_model_pb.StellarAccountId | undefined;
    setAccountId(value?: common_v3_model_pb.StellarAccountId): AccountInfo;

    getSequenceNumber(): string;
    setSequenceNumber(value: string): AccountInfo;

    getBalance(): string;
    setBalance(value: string): AccountInfo;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountInfo.AsObject;
    static toObject(includeInstance: boolean, msg: AccountInfo): AccountInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountInfo;
    static deserializeBinaryFromReader(message: AccountInfo, reader: jspb.BinaryReader): AccountInfo;
}

export namespace AccountInfo {
    export type AsObject = {
        accountId?: common_v3_model_pb.StellarAccountId.AsObject,
        sequenceNumber: string,
        balance: string,
    }
}

export class CreateAccountRequest extends jspb.Message { 

    hasAccountId(): boolean;
    clearAccountId(): void;
    getAccountId(): common_v3_model_pb.StellarAccountId | undefined;
    setAccountId(value?: common_v3_model_pb.StellarAccountId): CreateAccountRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateAccountRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateAccountRequest): CreateAccountRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateAccountRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateAccountRequest;
    static deserializeBinaryFromReader(message: CreateAccountRequest, reader: jspb.BinaryReader): CreateAccountRequest;
}

export namespace CreateAccountRequest {
    export type AsObject = {
        accountId?: common_v3_model_pb.StellarAccountId.AsObject,
    }
}

export class CreateAccountResponse extends jspb.Message { 
    getResult(): CreateAccountResponse.Result;
    setResult(value: CreateAccountResponse.Result): CreateAccountResponse;


    hasAccountInfo(): boolean;
    clearAccountInfo(): void;
    getAccountInfo(): AccountInfo | undefined;
    setAccountInfo(value?: AccountInfo): CreateAccountResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateAccountResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CreateAccountResponse): CreateAccountResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateAccountResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateAccountResponse;
    static deserializeBinaryFromReader(message: CreateAccountResponse, reader: jspb.BinaryReader): CreateAccountResponse;
}

export namespace CreateAccountResponse {
    export type AsObject = {
        result: CreateAccountResponse.Result,
        accountInfo?: AccountInfo.AsObject,
    }

    export enum Result {
    OK = 0,
    EXISTS = 1,
    }

}

export class GetAccountInfoRequest extends jspb.Message { 

    hasAccountId(): boolean;
    clearAccountId(): void;
    getAccountId(): common_v3_model_pb.StellarAccountId | undefined;
    setAccountId(value?: common_v3_model_pb.StellarAccountId): GetAccountInfoRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetAccountInfoRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetAccountInfoRequest): GetAccountInfoRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetAccountInfoRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetAccountInfoRequest;
    static deserializeBinaryFromReader(message: GetAccountInfoRequest, reader: jspb.BinaryReader): GetAccountInfoRequest;
}

export namespace GetAccountInfoRequest {
    export type AsObject = {
        accountId?: common_v3_model_pb.StellarAccountId.AsObject,
    }
}

export class GetAccountInfoResponse extends jspb.Message { 
    getResult(): GetAccountInfoResponse.Result;
    setResult(value: GetAccountInfoResponse.Result): GetAccountInfoResponse;


    hasAccountInfo(): boolean;
    clearAccountInfo(): void;
    getAccountInfo(): AccountInfo | undefined;
    setAccountInfo(value?: AccountInfo): GetAccountInfoResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetAccountInfoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetAccountInfoResponse): GetAccountInfoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetAccountInfoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetAccountInfoResponse;
    static deserializeBinaryFromReader(message: GetAccountInfoResponse, reader: jspb.BinaryReader): GetAccountInfoResponse;
}

export namespace GetAccountInfoResponse {
    export type AsObject = {
        result: GetAccountInfoResponse.Result,
        accountInfo?: AccountInfo.AsObject,
    }

    export enum Result {
    OK = 0,
    NOT_FOUND = 1,
    }

}

export class GetEventsRequest extends jspb.Message { 

    hasAccountId(): boolean;
    clearAccountId(): void;
    getAccountId(): common_v3_model_pb.StellarAccountId | undefined;
    setAccountId(value?: common_v3_model_pb.StellarAccountId): GetEventsRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetEventsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetEventsRequest): GetEventsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetEventsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetEventsRequest;
    static deserializeBinaryFromReader(message: GetEventsRequest, reader: jspb.BinaryReader): GetEventsRequest;
}

export namespace GetEventsRequest {
    export type AsObject = {
        accountId?: common_v3_model_pb.StellarAccountId.AsObject,
    }
}

export class Events extends jspb.Message { 
    getResult(): Events.Result;
    setResult(value: Events.Result): Events;

    clearEventsList(): void;
    getEventsList(): Array<Event>;
    setEventsList(value: Array<Event>): Events;
    addEvents(value?: Event, index?: number): Event;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Events.AsObject;
    static toObject(includeInstance: boolean, msg: Events): Events.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Events, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Events;
    static deserializeBinaryFromReader(message: Events, reader: jspb.BinaryReader): Events;
}

export namespace Events {
    export type AsObject = {
        result: Events.Result,
        eventsList: Array<Event.AsObject>,
    }

    export enum Result {
    OK = 0,
    NOT_FOUND = 1,
    }

}

export class Event extends jspb.Message { 

    hasAccountUpdateEvent(): boolean;
    clearAccountUpdateEvent(): void;
    getAccountUpdateEvent(): AccountUpdateEvent | undefined;
    setAccountUpdateEvent(value?: AccountUpdateEvent): Event;


    hasTransactionEvent(): boolean;
    clearTransactionEvent(): void;
    getTransactionEvent(): TransactionEvent | undefined;
    setTransactionEvent(value?: TransactionEvent): Event;


    getTypeCase(): Event.TypeCase;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Event.AsObject;
    static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Event;
    static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
}

export namespace Event {
    export type AsObject = {
        accountUpdateEvent?: AccountUpdateEvent.AsObject,
        transactionEvent?: TransactionEvent.AsObject,
    }

    export enum TypeCase {
        TYPE_NOT_SET = 0,
    
    ACCOUNT_UPDATE_EVENT = 1,

    TRANSACTION_EVENT = 2,

    }

}

export class AccountUpdateEvent extends jspb.Message { 

    hasAccountInfo(): boolean;
    clearAccountInfo(): void;
    getAccountInfo(): AccountInfo | undefined;
    setAccountInfo(value?: AccountInfo): AccountUpdateEvent;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountUpdateEvent.AsObject;
    static toObject(includeInstance: boolean, msg: AccountUpdateEvent): AccountUpdateEvent.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountUpdateEvent, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountUpdateEvent;
    static deserializeBinaryFromReader(message: AccountUpdateEvent, reader: jspb.BinaryReader): AccountUpdateEvent;
}

export namespace AccountUpdateEvent {
    export type AsObject = {
        accountInfo?: AccountInfo.AsObject,
    }
}

export class TransactionEvent extends jspb.Message { 
    getEnvelopeXdr(): Uint8Array | string;
    getEnvelopeXdr_asU8(): Uint8Array;
    getEnvelopeXdr_asB64(): string;
    setEnvelopeXdr(value: Uint8Array | string): TransactionEvent;

    getResultXdr(): Uint8Array | string;
    getResultXdr_asU8(): Uint8Array;
    getResultXdr_asB64(): string;
    setResultXdr(value: Uint8Array | string): TransactionEvent;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionEvent.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionEvent): TransactionEvent.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionEvent, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionEvent;
    static deserializeBinaryFromReader(message: TransactionEvent, reader: jspb.BinaryReader): TransactionEvent;
}

export namespace TransactionEvent {
    export type AsObject = {
        envelopeXdr: Uint8Array | string,
        resultXdr: Uint8Array | string,
    }
}
