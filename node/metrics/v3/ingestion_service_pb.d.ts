// package: kin.agora.metrics.v3
// file: metrics/v3/ingestion_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class SubmitRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubmitRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SubmitRequest): SubmitRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SubmitRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubmitRequest;
    static deserializeBinaryFromReader(message: SubmitRequest, reader: jspb.BinaryReader): SubmitRequest;
}

export namespace SubmitRequest {
    export type AsObject = {
    }
}

export class SubmitResponse extends jspb.Message { 
    getResult(): SubmitResponse.Result;
    setResult(value: SubmitResponse.Result): SubmitResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubmitResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SubmitResponse): SubmitResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SubmitResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubmitResponse;
    static deserializeBinaryFromReader(message: SubmitResponse, reader: jspb.BinaryReader): SubmitResponse;
}

export namespace SubmitResponse {
    export type AsObject = {
        result: SubmitResponse.Result,
    }

    export enum Result {
    OK = 0,
    }

}
