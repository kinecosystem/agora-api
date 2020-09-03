// package: kin.agora.common.v3
// file: common/v3/model.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as validate_validate_pb from "../../validate/validate_pb";

export class StellarAccountId extends jspb.Message { 
    getValue(): string;
    setValue(value: string): StellarAccountId;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StellarAccountId.AsObject;
    static toObject(includeInstance: boolean, msg: StellarAccountId): StellarAccountId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StellarAccountId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StellarAccountId;
    static deserializeBinaryFromReader(message: StellarAccountId, reader: jspb.BinaryReader): StellarAccountId;
}

export namespace StellarAccountId {
    export type AsObject = {
        value: string,
    }
}

export class TransactionHash extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): TransactionHash;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionHash.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionHash): TransactionHash.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionHash, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionHash;
    static deserializeBinaryFromReader(message: TransactionHash, reader: jspb.BinaryReader): TransactionHash;
}

export namespace TransactionHash {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class InvoiceHash extends jspb.Message { 
    getValue(): Uint8Array | string;
    getValue_asU8(): Uint8Array;
    getValue_asB64(): string;
    setValue(value: Uint8Array | string): InvoiceHash;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InvoiceHash.AsObject;
    static toObject(includeInstance: boolean, msg: InvoiceHash): InvoiceHash.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InvoiceHash, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InvoiceHash;
    static deserializeBinaryFromReader(message: InvoiceHash, reader: jspb.BinaryReader): InvoiceHash;
}

export namespace InvoiceHash {
    export type AsObject = {
        value: Uint8Array | string,
    }
}

export class Invoice extends jspb.Message { 
    clearItemsList(): void;
    getItemsList(): Array<Invoice.LineItem>;
    setItemsList(value: Array<Invoice.LineItem>): Invoice;
    addItems(value?: Invoice.LineItem, index?: number): Invoice.LineItem;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Invoice.AsObject;
    static toObject(includeInstance: boolean, msg: Invoice): Invoice.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Invoice, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Invoice;
    static deserializeBinaryFromReader(message: Invoice, reader: jspb.BinaryReader): Invoice;
}

export namespace Invoice {
    export type AsObject = {
        itemsList: Array<Invoice.LineItem.AsObject>,
    }


    export class LineItem extends jspb.Message { 
        getTitle(): string;
        setTitle(value: string): LineItem;

        getDescription(): string;
        setDescription(value: string): LineItem;

        getAmount(): string;
        setAmount(value: string): LineItem;

        getSku(): Uint8Array | string;
        getSku_asU8(): Uint8Array;
        getSku_asB64(): string;
        setSku(value: Uint8Array | string): LineItem;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): LineItem.AsObject;
        static toObject(includeInstance: boolean, msg: LineItem): LineItem.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: LineItem, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): LineItem;
        static deserializeBinaryFromReader(message: LineItem, reader: jspb.BinaryReader): LineItem;
    }

    export namespace LineItem {
        export type AsObject = {
            title: string,
            description: string,
            amount: string,
            sku: Uint8Array | string,
        }
    }

}

export class InvoiceList extends jspb.Message { 
    clearInvoicesList(): void;
    getInvoicesList(): Array<Invoice>;
    setInvoicesList(value: Array<Invoice>): InvoiceList;
    addInvoices(value?: Invoice, index?: number): Invoice;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InvoiceList.AsObject;
    static toObject(includeInstance: boolean, msg: InvoiceList): InvoiceList.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InvoiceList, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InvoiceList;
    static deserializeBinaryFromReader(message: InvoiceList, reader: jspb.BinaryReader): InvoiceList;
}

export namespace InvoiceList {
    export type AsObject = {
        invoicesList: Array<Invoice.AsObject>,
    }
}

export class InvoiceError extends jspb.Message { 
    getOpIndex(): number;
    setOpIndex(value: number): InvoiceError;


    hasInvoice(): boolean;
    clearInvoice(): void;
    getInvoice(): Invoice | undefined;
    setInvoice(value?: Invoice): InvoiceError;

    getReason(): InvoiceError.Reason;
    setReason(value: InvoiceError.Reason): InvoiceError;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InvoiceError.AsObject;
    static toObject(includeInstance: boolean, msg: InvoiceError): InvoiceError.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InvoiceError, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InvoiceError;
    static deserializeBinaryFromReader(message: InvoiceError, reader: jspb.BinaryReader): InvoiceError;
}

export namespace InvoiceError {
    export type AsObject = {
        opIndex: number,
        invoice?: Invoice.AsObject,
        reason: InvoiceError.Reason,
    }

    export enum Reason {
    UNKNOWN = 0,
    ALREADY_PAID = 1,
    WRONG_DESTINATION = 2,
    SKU_NOT_FOUND = 3,
    }

}
