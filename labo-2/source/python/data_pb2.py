# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: data.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='data.proto',
  package='main',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\ndata.proto\x12\x04main\"\x1f\n\rProtoImageRow\x12\x0e\n\x06values\x18\x01 \x03(\x02\"3\n\x0eProtoImageData\x12!\n\x04rows\x18\x01 \x03(\x0b\x32\x13.main.ProtoImageRow\"<\n\x08ProtoDWT\x12\x0c\n\x04mode\x18\x01 \x01(\t\x12\"\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x14.main.ProtoImageData\"\x1e\n\x0e\x46ileImageLayer\x12\x0c\n\x04rows\x18\x01 \x03(\x0c\"0\n\x0f\x46ileImageHeader\x12\r\n\x05width\x18\x01 \x01(\r\x12\x0e\n\x06height\x18\x02 \x01(\r\"r\n\rFileImageData\x12\x1f\n\x01y\x18\x01 \x01(\x0b\x32\x14.main.FileImageLayer\x12\x1f\n\x01u\x18\x02 \x01(\x0b\x32\x14.main.FileImageLayer\x12\x1f\n\x01v\x18\x03 \x01(\x0b\x32\x14.main.FileImageLayer\"U\n\tFileImage\x12%\n\x06header\x18\x01 \x01(\x0b\x32\x15.main.FileImageHeader\x12!\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x13.main.FileImageData*\'\n\x0bWaveletAlgo\x12\x08\n\x04HAAR\x10\x00\x12\x0e\n\nDAUBECHIES\x10\x01*\x1e\n\x0eQuantifierAlgo\x12\x0c\n\x08\x44\x45\x41\x44ZONE\x10\x00*\x19\n\x0e\x43ompressorAlgo\x12\x07\n\x03LZW\x10\x00\x62\x06proto3')
)

_WAVELETALGO = _descriptor.EnumDescriptor(
  name='WaveletAlgo',
  full_name='main.WaveletAlgo',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HAAR', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DAUBECHIES', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=453,
  serialized_end=492,
)
_sym_db.RegisterEnumDescriptor(_WAVELETALGO)

WaveletAlgo = enum_type_wrapper.EnumTypeWrapper(_WAVELETALGO)
_QUANTIFIERALGO = _descriptor.EnumDescriptor(
  name='QuantifierAlgo',
  full_name='main.QuantifierAlgo',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEADZONE', index=0, number=0,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=494,
  serialized_end=524,
)
_sym_db.RegisterEnumDescriptor(_QUANTIFIERALGO)

QuantifierAlgo = enum_type_wrapper.EnumTypeWrapper(_QUANTIFIERALGO)
_COMPRESSORALGO = _descriptor.EnumDescriptor(
  name='CompressorAlgo',
  full_name='main.CompressorAlgo',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='LZW', index=0, number=0,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=526,
  serialized_end=551,
)
_sym_db.RegisterEnumDescriptor(_COMPRESSORALGO)

CompressorAlgo = enum_type_wrapper.EnumTypeWrapper(_COMPRESSORALGO)
HAAR = 0
DAUBECHIES = 1
DEADZONE = 0
LZW = 0



_PROTOIMAGEROW = _descriptor.Descriptor(
  name='ProtoImageRow',
  full_name='main.ProtoImageRow',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='main.ProtoImageRow.values', index=0,
      number=1, type=2, cpp_type=6, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=20,
  serialized_end=51,
)


_PROTOIMAGEDATA = _descriptor.Descriptor(
  name='ProtoImageData',
  full_name='main.ProtoImageData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rows', full_name='main.ProtoImageData.rows', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=53,
  serialized_end=104,
)


_PROTODWT = _descriptor.Descriptor(
  name='ProtoDWT',
  full_name='main.ProtoDWT',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='main.ProtoDWT.mode', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='main.ProtoDWT.data', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=106,
  serialized_end=166,
)


_FILEIMAGELAYER = _descriptor.Descriptor(
  name='FileImageLayer',
  full_name='main.FileImageLayer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rows', full_name='main.FileImageLayer.rows', index=0,
      number=1, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=168,
  serialized_end=198,
)


_FILEIMAGEHEADER = _descriptor.Descriptor(
  name='FileImageHeader',
  full_name='main.FileImageHeader',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='main.FileImageHeader.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='height', full_name='main.FileImageHeader.height', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=200,
  serialized_end=248,
)


_FILEIMAGEDATA = _descriptor.Descriptor(
  name='FileImageData',
  full_name='main.FileImageData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='y', full_name='main.FileImageData.y', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='u', full_name='main.FileImageData.u', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='v', full_name='main.FileImageData.v', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=250,
  serialized_end=364,
)


_FILEIMAGE = _descriptor.Descriptor(
  name='FileImage',
  full_name='main.FileImage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='main.FileImage.header', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='main.FileImage.data', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=366,
  serialized_end=451,
)

_PROTOIMAGEDATA.fields_by_name['rows'].message_type = _PROTOIMAGEROW
_PROTODWT.fields_by_name['data'].message_type = _PROTOIMAGEDATA
_FILEIMAGEDATA.fields_by_name['y'].message_type = _FILEIMAGELAYER
_FILEIMAGEDATA.fields_by_name['u'].message_type = _FILEIMAGELAYER
_FILEIMAGEDATA.fields_by_name['v'].message_type = _FILEIMAGELAYER
_FILEIMAGE.fields_by_name['header'].message_type = _FILEIMAGEHEADER
_FILEIMAGE.fields_by_name['data'].message_type = _FILEIMAGEDATA
DESCRIPTOR.message_types_by_name['ProtoImageRow'] = _PROTOIMAGEROW
DESCRIPTOR.message_types_by_name['ProtoImageData'] = _PROTOIMAGEDATA
DESCRIPTOR.message_types_by_name['ProtoDWT'] = _PROTODWT
DESCRIPTOR.message_types_by_name['FileImageLayer'] = _FILEIMAGELAYER
DESCRIPTOR.message_types_by_name['FileImageHeader'] = _FILEIMAGEHEADER
DESCRIPTOR.message_types_by_name['FileImageData'] = _FILEIMAGEDATA
DESCRIPTOR.message_types_by_name['FileImage'] = _FILEIMAGE
DESCRIPTOR.enum_types_by_name['WaveletAlgo'] = _WAVELETALGO
DESCRIPTOR.enum_types_by_name['QuantifierAlgo'] = _QUANTIFIERALGO
DESCRIPTOR.enum_types_by_name['CompressorAlgo'] = _COMPRESSORALGO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProtoImageRow = _reflection.GeneratedProtocolMessageType('ProtoImageRow', (_message.Message,), dict(
  DESCRIPTOR = _PROTOIMAGEROW,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.ProtoImageRow)
  ))
_sym_db.RegisterMessage(ProtoImageRow)

ProtoImageData = _reflection.GeneratedProtocolMessageType('ProtoImageData', (_message.Message,), dict(
  DESCRIPTOR = _PROTOIMAGEDATA,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.ProtoImageData)
  ))
_sym_db.RegisterMessage(ProtoImageData)

ProtoDWT = _reflection.GeneratedProtocolMessageType('ProtoDWT', (_message.Message,), dict(
  DESCRIPTOR = _PROTODWT,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.ProtoDWT)
  ))
_sym_db.RegisterMessage(ProtoDWT)

FileImageLayer = _reflection.GeneratedProtocolMessageType('FileImageLayer', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGELAYER,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.FileImageLayer)
  ))
_sym_db.RegisterMessage(FileImageLayer)

FileImageHeader = _reflection.GeneratedProtocolMessageType('FileImageHeader', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGEHEADER,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.FileImageHeader)
  ))
_sym_db.RegisterMessage(FileImageHeader)

FileImageData = _reflection.GeneratedProtocolMessageType('FileImageData', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGEDATA,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.FileImageData)
  ))
_sym_db.RegisterMessage(FileImageData)

FileImage = _reflection.GeneratedProtocolMessageType('FileImage', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGE,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:main.FileImage)
  ))
_sym_db.RegisterMessage(FileImage)


# @@protoc_insertion_point(module_scope)
