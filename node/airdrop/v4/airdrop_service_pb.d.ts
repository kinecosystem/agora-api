// package: kin.agora.airdrop.v4
// file: airdrop/v4/airdrop_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as validate_validate_pb from "../../validate/validate_pb";
import * as common_v4_model_pb from "../../common/v4/model_pb";

export class RequestAirdropRequest extends jspb.Message { 

    hasAccountId(): boolean;
    clearAccountId(): void;
    getAccountId(): common_v4_model_pb.SolanaAccountId | undefined;
    setAccountId(value?: common_v4_model_pb.SolanaAccountId): RequestAirdropRequest;

    getQuarks(): number;
    setQuarks(value: number): RequestAirdropRequest;

    getCommitment(): common_v4_model_pb.Commitment;
    setCommitment(value: common_v4_model_pb.Commitment): RequestAirdropRequest;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestAirdropRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RequestAirdropRequest): RequestAirdropRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestAirdropRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestAirdropRequest;
    static deserializeBinaryFromReader(message: RequestAirdropRequest, reader: jspb.BinaryReader): RequestAirdropRequest;
}

export namespace RequestAirdropRequest {
    export type AsObject = {
        accountId?: common_v4_model_pb.SolanaAccountId.AsObject,
        quarks: number,
        commitment: common_v4_model_pb.Commitment,
    }
}

export class RequestAirdropResponse extends jspb.Message { 
    getResult(): RequestAirdropResponse.Result;
    setResult(value: RequestAirdropResponse.Result): RequestAirdropResponse;


    hasSignature(): boolean;
    clearSignature(): void;
    getSignature(): common_v4_model_pb.TransactionSignature | undefined;
    setSignature(value?: common_v4_model_pb.TransactionSignature): RequestAirdropResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestAirdropResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RequestAirdropResponse): RequestAirdropResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestAirdropResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestAirdropResponse;
    static deserializeBinaryFromReader(message: RequestAirdropResponse, reader: jspb.BinaryReader): RequestAirdropResponse;
}

export namespace RequestAirdropResponse {
    export type AsObject = {
        result: RequestAirdropResponse.Result,
        signature?: common_v4_model_pb.TransactionSignature.AsObject,
    }

    export enum Result {
    OK = 0,
    NOT_FOUND = 1,
    INSUFFICIENT_KIN = 2,
    }

}
