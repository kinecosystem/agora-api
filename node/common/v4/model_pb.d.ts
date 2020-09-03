// package: kin.agora.common.v4
// file: common/v4/model.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as validate_validate_pb from "../../validate/validate_pb";

export class SolanaAccountId extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): SolanaAccountId;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SolanaAccountId.AsObject;
    static toObject(includeInstance: boolean, msg: SolanaAccountId): SolanaAccountId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SolanaAccountId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SolanaAccountId;
    static deserializeBinaryFromReader(message: SolanaAccountId, reader: jspb.BinaryReader): SolanaAccountId;
}

export namespace SolanaAccountId {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class TransactionId extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): TransactionId;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionId.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionId): TransactionId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionId;
    static deserializeBinaryFromReader(message: TransactionId, reader: jspb.BinaryReader): TransactionId;
}

export namespace TransactionId {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class Blockhash extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): Blockhash;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Blockhash.AsObject;
    static toObject(includeInstance: boolean, msg: Blockhash): Blockhash.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Blockhash, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Blockhash;
    static deserializeBinaryFromReader(message: Blockhash, reader: jspb.BinaryReader): Blockhash;
}

export namespace Blockhash {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class TransactionSignature extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): TransactionSignature;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionSignature.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionSignature): TransactionSignature.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionSignature, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionSignature;
    static deserializeBinaryFromReader(message: TransactionSignature, reader: jspb.BinaryReader): TransactionSignature;
}

export namespace TransactionSignature {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class Transaction extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): Transaction;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Transaction.AsObject;
    static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Transaction;
    static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class TransactionError extends jspb.Message { 
    getReason(): TransactionError.Reason;
    setReason(value: TransactionError.Reason): TransactionError;

    getInstructionIndex(): number;
    setInstructionIndex(value: number): TransactionError;

    getRaw(): Uint8Array | string;
    getRaw_asU8(): Uint8Array;
    getRaw_asB64(): string;
    setRaw(value: Uint8Array | string): TransactionError;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionError.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionError): TransactionError.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionError, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionError;
    static deserializeBinaryFromReader(message: TransactionError, reader: jspb.BinaryReader): TransactionError;
}

export namespace TransactionError {
    export type AsObject = {
        reason: TransactionError.Reason,
        instructionIndex: number,
        raw: Uint8Array | string,
    }

    export enum Reason {
    NONE = 0,
    UNKNOWN = 1,
    UNAUTHORIZED = 2,
    BAD_NONCE = 3,
    INSUFFICIENT_FUNDS = 4,
    INVALID_ACCOUNT = 5,
    }

}

export class StellarTransaction extends jspb.Message { 
    getResultXdr(): Uint8Array | string;
    getResultXdr_asU8(): Uint8Array;
    getResultXdr_asB64(): string;
    setResultXdr(value: Uint8Array | string): StellarTransaction;

    getEnvelopeXdr(): Uint8Array | string;
    getEnvelopeXdr_asU8(): Uint8Array;
    getEnvelopeXdr_asB64(): string;
    setEnvelopeXdr(value: Uint8Array | string): StellarTransaction;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StellarTransaction.AsObject;
    static toObject(includeInstance: boolean, msg: StellarTransaction): StellarTransaction.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StellarTransaction, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StellarTransaction;
    static deserializeBinaryFromReader(message: StellarTransaction, reader: jspb.BinaryReader): StellarTransaction;
}

export namespace StellarTransaction {
    export type AsObject = {
        resultXdr: Uint8Array | string,
        envelopeXdr: Uint8Array | string,
    }
}

export enum Commitment {
    RECENT = 0,
    SINGLE = 1,
    ROOT = 2,
    MAX = 3,
}
