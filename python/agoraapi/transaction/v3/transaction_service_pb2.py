# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v3/transaction_service.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from common.v3 import model_pb2 as common_dot_v3_dot_model__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='transaction/v3/transaction_service.proto',
  package='kin.agora.transaction.v3',
  syntax='proto3',
  serialized_options=b'\n org.kin.agora.gen.transaction.v3ZNgithub.com/kinecosystem/agora-api/genproto/transaction/v3;transaction\242\002\020APBTransactionV3',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n(transaction/v3/transaction_service.proto\x12\x18kin.agora.transaction.v3\x1a\x17validate/validate.proto\x1a\x15\x63ommon/v3/model.proto\"\xf4\x01\n\x11GetHistoryRequest\x12\x43\n\naccount_id\x18\x01 \x01(\x0b\x32%.kin.agora.common.v3.StellarAccountIdB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x30\n\x06\x63ursor\x18\x02 \x01(\x0b\x32 .kin.agora.transaction.v3.Cursor\x12H\n\tdirection\x18\x03 \x01(\x0e\x32\x35.kin.agora.transaction.v3.GetHistoryRequest.Direction\"\x1e\n\tDirection\x12\x07\n\x03\x41SC\x10\x00\x12\x08\n\x04\x44\x45SC\x10\x01\"\xbd\x01\n\x12GetHistoryResponse\x12\x43\n\x06result\x18\x01 \x01(\x0e\x32\x33.kin.agora.transaction.v3.GetHistoryResponse.Result\x12\x41\n\x05items\x18\x02 \x03(\x0b\x32%.kin.agora.transaction.v3.HistoryItemB\x0b\xfa\x42\x08\x92\x01\x05\x08\x00\x10\x80\x01\"\x1f\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\r\n\tNOT_FOUND\x10\x01\"t\n\x18SubmitTransactionRequest\x12 \n\x0c\x65nvelope_xdr\x18\x01 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\x12\x36\n\x0cinvoice_list\x18\x02 \x01(\x0b\x32 .kin.agora.common.v3.InvoiceList\"\xc9\x02\n\x19SubmitTransactionResponse\x12J\n\x06result\x18\x01 \x01(\x0e\x32:.kin.agora.transaction.v3.SubmitTransactionResponse.Result\x12\x39\n\x0einvoice_errors\x18\x02 \x03(\x0b\x32!.kin.agora.common.v3.InvoiceError\x12\x32\n\x04hash\x18\x03 \x01(\x0b\x32$.kin.agora.common.v3.TransactionHash\x12\x12\n\x06ledger\x18\x04 \x01(\x03\x42\x02\x30\x01\x12\x1e\n\nresult_xdr\x18\x05 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x00\x18\x80P\"=\n\x06Result\x12\x06\n\x02OK\x10\x00\x12\n\n\x06\x46\x41ILED\x10\x01\x12\x0c\n\x08REJECTED\x10\x02\x12\x11\n\rINVOICE_ERROR\x10\x03\"a\n\x15GetTransactionRequest\x12H\n\x10transaction_hash\x18\x01 \x01(\x0b\x32$.kin.agora.common.v3.TransactionHashB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\"\xcb\x01\n\x16GetTransactionResponse\x12\x45\n\x05state\x18\x01 \x01(\x0e\x32\x36.kin.agora.transaction.v3.GetTransactionResponse.State\x12\x12\n\x06ledger\x18\x02 \x01(\x03\x42\x02\x30\x01\x12\x33\n\x04item\x18\x03 \x01(\x0b\x32%.kin.agora.transaction.v3.HistoryItem\"!\n\x05State\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0b\n\x07SUCCESS\x10\x01\"\xf7\x01\n\x0bHistoryItem\x12<\n\x04hash\x18\x01 \x01(\x0b\x32$.kin.agora.common.v3.TransactionHashB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x1e\n\nresult_xdr\x18\x02 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\x12 \n\x0c\x65nvelope_xdr\x18\x03 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80P\x12\x30\n\x06\x63ursor\x18\x04 \x01(\x0b\x32 .kin.agora.transaction.v3.Cursor\x12\x36\n\x0cinvoice_list\x18\x05 \x01(\x0b\x32 .kin.agora.common.v3.InvoiceList\"#\n\x06\x43ursor\x12\x19\n\x05value\x18\x01 \x01(\x0c\x42\n\xfa\x42\x07z\x05\x10\x01\x18\x80\x01\x32\xe9\x02\n\x0bTransaction\x12g\n\nGetHistory\x12+.kin.agora.transaction.v3.GetHistoryRequest\x1a,.kin.agora.transaction.v3.GetHistoryResponse\x12|\n\x11SubmitTransaction\x12\x32.kin.agora.transaction.v3.SubmitTransactionRequest\x1a\x33.kin.agora.transaction.v3.SubmitTransactionResponse\x12s\n\x0eGetTransaction\x12/.kin.agora.transaction.v3.GetTransactionRequest\x1a\x30.kin.agora.transaction.v3.GetTransactionResponseB\x85\x01\n org.kin.agora.gen.transaction.v3ZNgithub.com/kinecosystem/agora-api/genproto/transaction/v3;transaction\xa2\x02\x10\x41PBTransactionV3b\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,common_dot_v3_dot_model__pb2.DESCRIPTOR,])



_GETHISTORYREQUEST_DIRECTION = _descriptor.EnumDescriptor(
  name='Direction',
  full_name='kin.agora.transaction.v3.GetHistoryRequest.Direction',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ASC', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DESC', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=333,
  serialized_end=363,
)
_sym_db.RegisterEnumDescriptor(_GETHISTORYREQUEST_DIRECTION)

_GETHISTORYRESPONSE_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='kin.agora.transaction.v3.GetHistoryResponse.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NOT_FOUND', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=524,
  serialized_end=555,
)
_sym_db.RegisterEnumDescriptor(_GETHISTORYRESPONSE_RESULT)

_SUBMITTRANSACTIONRESPONSE_RESULT = _descriptor.EnumDescriptor(
  name='Result',
  full_name='kin.agora.transaction.v3.SubmitTransactionResponse.Result',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FAILED', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='REJECTED', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='INVOICE_ERROR', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=944,
  serialized_end=1005,
)
_sym_db.RegisterEnumDescriptor(_SUBMITTRANSACTIONRESPONSE_RESULT)

_GETTRANSACTIONRESPONSE_STATE = _descriptor.EnumDescriptor(
  name='State',
  full_name='kin.agora.transaction.v3.GetTransactionResponse.State',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SUCCESS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1277,
  serialized_end=1310,
)
_sym_db.RegisterEnumDescriptor(_GETTRANSACTIONRESPONSE_STATE)


_GETHISTORYREQUEST = _descriptor.Descriptor(
  name='GetHistoryRequest',
  full_name='kin.agora.transaction.v3.GetHistoryRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='kin.agora.transaction.v3.GetHistoryRequest.account_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cursor', full_name='kin.agora.transaction.v3.GetHistoryRequest.cursor', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='direction', full_name='kin.agora.transaction.v3.GetHistoryRequest.direction', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _GETHISTORYREQUEST_DIRECTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=119,
  serialized_end=363,
)


_GETHISTORYRESPONSE = _descriptor.Descriptor(
  name='GetHistoryResponse',
  full_name='kin.agora.transaction.v3.GetHistoryResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='kin.agora.transaction.v3.GetHistoryResponse.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='items', full_name='kin.agora.transaction.v3.GetHistoryResponse.items', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\010\222\001\005\010\000\020\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _GETHISTORYRESPONSE_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=366,
  serialized_end=555,
)


_SUBMITTRANSACTIONREQUEST = _descriptor.Descriptor(
  name='SubmitTransactionRequest',
  full_name='kin.agora.transaction.v3.SubmitTransactionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='envelope_xdr', full_name='kin.agora.transaction.v3.SubmitTransactionRequest.envelope_xdr', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='invoice_list', full_name='kin.agora.transaction.v3.SubmitTransactionRequest.invoice_list', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=557,
  serialized_end=673,
)


_SUBMITTRANSACTIONRESPONSE = _descriptor.Descriptor(
  name='SubmitTransactionResponse',
  full_name='kin.agora.transaction.v3.SubmitTransactionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='kin.agora.transaction.v3.SubmitTransactionResponse.result', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='invoice_errors', full_name='kin.agora.transaction.v3.SubmitTransactionResponse.invoice_errors', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hash', full_name='kin.agora.transaction.v3.SubmitTransactionResponse.hash', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ledger', full_name='kin.agora.transaction.v3.SubmitTransactionResponse.ledger', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'0\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result_xdr', full_name='kin.agora.transaction.v3.SubmitTransactionResponse.result_xdr', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\000\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SUBMITTRANSACTIONRESPONSE_RESULT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=676,
  serialized_end=1005,
)


_GETTRANSACTIONREQUEST = _descriptor.Descriptor(
  name='GetTransactionRequest',
  full_name='kin.agora.transaction.v3.GetTransactionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='transaction_hash', full_name='kin.agora.transaction.v3.GetTransactionRequest.transaction_hash', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1007,
  serialized_end=1104,
)


_GETTRANSACTIONRESPONSE = _descriptor.Descriptor(
  name='GetTransactionResponse',
  full_name='kin.agora.transaction.v3.GetTransactionResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='state', full_name='kin.agora.transaction.v3.GetTransactionResponse.state', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ledger', full_name='kin.agora.transaction.v3.GetTransactionResponse.ledger', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'0\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='item', full_name='kin.agora.transaction.v3.GetTransactionResponse.item', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _GETTRANSACTIONRESPONSE_STATE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1107,
  serialized_end=1310,
)


_HISTORYITEM = _descriptor.Descriptor(
  name='HistoryItem',
  full_name='kin.agora.transaction.v3.HistoryItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='hash', full_name='kin.agora.transaction.v3.HistoryItem.hash', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\005\212\001\002\020\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result_xdr', full_name='kin.agora.transaction.v3.HistoryItem.result_xdr', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='envelope_xdr', full_name='kin.agora.transaction.v3.HistoryItem.envelope_xdr', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200P', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cursor', full_name='kin.agora.transaction.v3.HistoryItem.cursor', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='invoice_list', full_name='kin.agora.transaction.v3.HistoryItem.invoice_list', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1313,
  serialized_end=1560,
)


_CURSOR = _descriptor.Descriptor(
  name='Cursor',
  full_name='kin.agora.transaction.v3.Cursor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='kin.agora.transaction.v3.Cursor.value', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\007z\005\020\001\030\200\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1562,
  serialized_end=1597,
)

_GETHISTORYREQUEST.fields_by_name['account_id'].message_type = common_dot_v3_dot_model__pb2._STELLARACCOUNTID
_GETHISTORYREQUEST.fields_by_name['cursor'].message_type = _CURSOR
_GETHISTORYREQUEST.fields_by_name['direction'].enum_type = _GETHISTORYREQUEST_DIRECTION
_GETHISTORYREQUEST_DIRECTION.containing_type = _GETHISTORYREQUEST
_GETHISTORYRESPONSE.fields_by_name['result'].enum_type = _GETHISTORYRESPONSE_RESULT
_GETHISTORYRESPONSE.fields_by_name['items'].message_type = _HISTORYITEM
_GETHISTORYRESPONSE_RESULT.containing_type = _GETHISTORYRESPONSE
_SUBMITTRANSACTIONREQUEST.fields_by_name['invoice_list'].message_type = common_dot_v3_dot_model__pb2._INVOICELIST
_SUBMITTRANSACTIONRESPONSE.fields_by_name['result'].enum_type = _SUBMITTRANSACTIONRESPONSE_RESULT
_SUBMITTRANSACTIONRESPONSE.fields_by_name['invoice_errors'].message_type = common_dot_v3_dot_model__pb2._INVOICEERROR
_SUBMITTRANSACTIONRESPONSE.fields_by_name['hash'].message_type = common_dot_v3_dot_model__pb2._TRANSACTIONHASH
_SUBMITTRANSACTIONRESPONSE_RESULT.containing_type = _SUBMITTRANSACTIONRESPONSE
_GETTRANSACTIONREQUEST.fields_by_name['transaction_hash'].message_type = common_dot_v3_dot_model__pb2._TRANSACTIONHASH
_GETTRANSACTIONRESPONSE.fields_by_name['state'].enum_type = _GETTRANSACTIONRESPONSE_STATE
_GETTRANSACTIONRESPONSE.fields_by_name['item'].message_type = _HISTORYITEM
_GETTRANSACTIONRESPONSE_STATE.containing_type = _GETTRANSACTIONRESPONSE
_HISTORYITEM.fields_by_name['hash'].message_type = common_dot_v3_dot_model__pb2._TRANSACTIONHASH
_HISTORYITEM.fields_by_name['cursor'].message_type = _CURSOR
_HISTORYITEM.fields_by_name['invoice_list'].message_type = common_dot_v3_dot_model__pb2._INVOICELIST
DESCRIPTOR.message_types_by_name['GetHistoryRequest'] = _GETHISTORYREQUEST
DESCRIPTOR.message_types_by_name['GetHistoryResponse'] = _GETHISTORYRESPONSE
DESCRIPTOR.message_types_by_name['SubmitTransactionRequest'] = _SUBMITTRANSACTIONREQUEST
DESCRIPTOR.message_types_by_name['SubmitTransactionResponse'] = _SUBMITTRANSACTIONRESPONSE
DESCRIPTOR.message_types_by_name['GetTransactionRequest'] = _GETTRANSACTIONREQUEST
DESCRIPTOR.message_types_by_name['GetTransactionResponse'] = _GETTRANSACTIONRESPONSE
DESCRIPTOR.message_types_by_name['HistoryItem'] = _HISTORYITEM
DESCRIPTOR.message_types_by_name['Cursor'] = _CURSOR
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetHistoryRequest = _reflection.GeneratedProtocolMessageType('GetHistoryRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETHISTORYREQUEST,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.GetHistoryRequest)
  })
_sym_db.RegisterMessage(GetHistoryRequest)

GetHistoryResponse = _reflection.GeneratedProtocolMessageType('GetHistoryResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETHISTORYRESPONSE,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.GetHistoryResponse)
  })
_sym_db.RegisterMessage(GetHistoryResponse)

SubmitTransactionRequest = _reflection.GeneratedProtocolMessageType('SubmitTransactionRequest', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITTRANSACTIONREQUEST,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.SubmitTransactionRequest)
  })
_sym_db.RegisterMessage(SubmitTransactionRequest)

SubmitTransactionResponse = _reflection.GeneratedProtocolMessageType('SubmitTransactionResponse', (_message.Message,), {
  'DESCRIPTOR' : _SUBMITTRANSACTIONRESPONSE,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.SubmitTransactionResponse)
  })
_sym_db.RegisterMessage(SubmitTransactionResponse)

GetTransactionRequest = _reflection.GeneratedProtocolMessageType('GetTransactionRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETTRANSACTIONREQUEST,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.GetTransactionRequest)
  })
_sym_db.RegisterMessage(GetTransactionRequest)

GetTransactionResponse = _reflection.GeneratedProtocolMessageType('GetTransactionResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETTRANSACTIONRESPONSE,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.GetTransactionResponse)
  })
_sym_db.RegisterMessage(GetTransactionResponse)

HistoryItem = _reflection.GeneratedProtocolMessageType('HistoryItem', (_message.Message,), {
  'DESCRIPTOR' : _HISTORYITEM,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.HistoryItem)
  })
_sym_db.RegisterMessage(HistoryItem)

Cursor = _reflection.GeneratedProtocolMessageType('Cursor', (_message.Message,), {
  'DESCRIPTOR' : _CURSOR,
  '__module__' : 'transaction.v3.transaction_service_pb2'
  # @@protoc_insertion_point(class_scope:kin.agora.transaction.v3.Cursor)
  })
_sym_db.RegisterMessage(Cursor)


DESCRIPTOR._options = None
_GETHISTORYREQUEST.fields_by_name['account_id']._options = None
_GETHISTORYRESPONSE.fields_by_name['items']._options = None
_SUBMITTRANSACTIONREQUEST.fields_by_name['envelope_xdr']._options = None
_SUBMITTRANSACTIONRESPONSE.fields_by_name['ledger']._options = None
_SUBMITTRANSACTIONRESPONSE.fields_by_name['result_xdr']._options = None
_GETTRANSACTIONREQUEST.fields_by_name['transaction_hash']._options = None
_GETTRANSACTIONRESPONSE.fields_by_name['ledger']._options = None
_HISTORYITEM.fields_by_name['hash']._options = None
_HISTORYITEM.fields_by_name['result_xdr']._options = None
_HISTORYITEM.fields_by_name['envelope_xdr']._options = None
_CURSOR.fields_by_name['value']._options = None

_TRANSACTION = _descriptor.ServiceDescriptor(
  name='Transaction',
  full_name='kin.agora.transaction.v3.Transaction',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1600,
  serialized_end=1961,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetHistory',
    full_name='kin.agora.transaction.v3.Transaction.GetHistory',
    index=0,
    containing_service=None,
    input_type=_GETHISTORYREQUEST,
    output_type=_GETHISTORYRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SubmitTransaction',
    full_name='kin.agora.transaction.v3.Transaction.SubmitTransaction',
    index=1,
    containing_service=None,
    input_type=_SUBMITTRANSACTIONREQUEST,
    output_type=_SUBMITTRANSACTIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetTransaction',
    full_name='kin.agora.transaction.v3.Transaction.GetTransaction',
    index=2,
    containing_service=None,
    input_type=_GETTRANSACTIONREQUEST,
    output_type=_GETTRANSACTIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_TRANSACTION)

DESCRIPTOR.services_by_name['Transaction'] = _TRANSACTION

# @@protoc_insertion_point(module_scope)
