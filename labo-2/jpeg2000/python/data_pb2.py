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
  package='data',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\ndata.proto\x12\x04\x64\x61ta\"\x1a\n\x08ImageRow\x12\x0e\n\x06values\x18\x01 \x03(\x02\")\n\tImageData\x12\x1c\n\x04rows\x18\x01 \x03(\x0b\x32\x0e.data.ImageRow\"8\n\tPythonDWT\x12\x0c\n\x04mode\x18\x01 \x01(\t\x12\x1d\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x0f.data.ImageData\"\x1c\n\x0bWaveletHaar\x12\r\n\x05level\x18\x01 \x01(\r\"7\n\x11WaveletDaubechies\x12\r\n\x05level\x18\x01 \x01(\r\x12\x13\n\x0b\x63oefficient\x18\x02 \x01(\r\"\x0e\n\x0cWaveletDummy\"\x8e\x01\n\rWaveletConfig\x12!\n\x04haar\x18\x01 \x01(\x0b\x32\x11.data.WaveletHaarH\x00\x12-\n\ndaubechies\x18\x02 \x01(\x0b\x32\x17.data.WaveletDaubechiesH\x00\x12#\n\x05\x64ummy\x18\x03 \x01(\x0b\x32\x12.data.WaveletDummyH\x00\x42\x06\n\x04\x64\x61ta\"B\n\x12QuantifierDeadZone\x12\r\n\x05width\x18\x01 \x01(\r\x12\r\n\x05\x64\x65lta\x18\x02 \x01(\r\x12\x0e\n\x06offset\x18\x03 \x01(\x02\"$\n\x13QuantifierMidThread\x12\r\n\x05\x64\x65lta\x18\x01 \x01(\r\"z\n\x10QuantifierConfig\x12-\n\tdead_zone\x18\x01 \x01(\x0b\x32\x18.data.QuantifierDeadZoneH\x00\x12/\n\nmid_thread\x18\x02 \x01(\x0b\x32\x19.data.QuantifierMidThreadH\x00\x42\x06\n\x04\x64\x61ta\"\x1e\n\x0e\x46ileImageLayer\x12\x0c\n\x04rows\x18\x01 \x03(\x0c\"\xbe\x01\n\x0f\x46ileImageHeader\x12\r\n\x05width\x18\x01 \x01(\r\x12\x0e\n\x06height\x18\x02 \x01(\r\x12\x12\n\nconversion\x18\x03 \x01(\x08\x12&\n\x0bsubsampling\x18\x04 \x01(\x0e\x32\x11.data.Subsampling\x12$\n\x07wavelet\x18\x05 \x01(\x0b\x32\x13.data.WaveletConfig\x12*\n\nquantifier\x18\x06 \x01(\x0b\x32\x16.data.QuantifierConfig\"r\n\rFileImageData\x12\x1f\n\x01y\x18\x01 \x01(\x0b\x32\x14.data.FileImageLayer\x12\x1f\n\x01u\x18\x02 \x01(\x0b\x32\x14.data.FileImageLayer\x12\x1f\n\x01v\x18\x03 \x01(\x0b\x32\x14.data.FileImageLayer\"U\n\tFileImage\x12%\n\x06header\x18\x01 \x01(\x0b\x32\x15.data.FileImageHeader\x12!\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x13.data.FileImageData*a\n\x0bSubsampling\x12\x13\n\x0fSUBSAMPLING_410\x10\x00\x12\x13\n\x0fSUBSAMPLING_420\x10\x01\x12\x13\n\x0fSUBSAMPLING_422\x10\x02\x12\x13\n\x0fSUBSAMPLING_444\x10\x03\x62\x06proto3')
)

_SUBSAMPLING = _descriptor.EnumDescriptor(
  name='Subsampling',
  full_name='data.Subsampling',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SUBSAMPLING_410', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBSAMPLING_420', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBSAMPLING_422', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUBSAMPLING_444', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1055,
  serialized_end=1152,
)
_sym_db.RegisterEnumDescriptor(_SUBSAMPLING)

Subsampling = enum_type_wrapper.EnumTypeWrapper(_SUBSAMPLING)
SUBSAMPLING_410 = 0
SUBSAMPLING_420 = 1
SUBSAMPLING_422 = 2
SUBSAMPLING_444 = 3



_IMAGEROW = _descriptor.Descriptor(
  name='ImageRow',
  full_name='data.ImageRow',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='values', full_name='data.ImageRow.values', index=0,
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
  serialized_end=46,
)


_IMAGEDATA = _descriptor.Descriptor(
  name='ImageData',
  full_name='data.ImageData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rows', full_name='data.ImageData.rows', index=0,
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
  serialized_start=48,
  serialized_end=89,
)


_PYTHONDWT = _descriptor.Descriptor(
  name='PythonDWT',
  full_name='data.PythonDWT',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='data.PythonDWT.mode', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='data.PythonDWT.data', index=1,
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
  serialized_start=91,
  serialized_end=147,
)


_WAVELETHAAR = _descriptor.Descriptor(
  name='WaveletHaar',
  full_name='data.WaveletHaar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='level', full_name='data.WaveletHaar.level', index=0,
      number=1, type=13, cpp_type=3, label=1,
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
  serialized_start=149,
  serialized_end=177,
)


_WAVELETDAUBECHIES = _descriptor.Descriptor(
  name='WaveletDaubechies',
  full_name='data.WaveletDaubechies',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='level', full_name='data.WaveletDaubechies.level', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='coefficient', full_name='data.WaveletDaubechies.coefficient', index=1,
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
  serialized_start=179,
  serialized_end=234,
)


_WAVELETDUMMY = _descriptor.Descriptor(
  name='WaveletDummy',
  full_name='data.WaveletDummy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=236,
  serialized_end=250,
)


_WAVELETCONFIG = _descriptor.Descriptor(
  name='WaveletConfig',
  full_name='data.WaveletConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='haar', full_name='data.WaveletConfig.haar', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='daubechies', full_name='data.WaveletConfig.daubechies', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dummy', full_name='data.WaveletConfig.dummy', index=2,
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
    _descriptor.OneofDescriptor(
      name='data', full_name='data.WaveletConfig.data',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=253,
  serialized_end=395,
)


_QUANTIFIERDEADZONE = _descriptor.Descriptor(
  name='QuantifierDeadZone',
  full_name='data.QuantifierDeadZone',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='data.QuantifierDeadZone.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='delta', full_name='data.QuantifierDeadZone.delta', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='data.QuantifierDeadZone.offset', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=397,
  serialized_end=463,
)


_QUANTIFIERMIDTHREAD = _descriptor.Descriptor(
  name='QuantifierMidThread',
  full_name='data.QuantifierMidThread',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='delta', full_name='data.QuantifierMidThread.delta', index=0,
      number=1, type=13, cpp_type=3, label=1,
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
  serialized_start=465,
  serialized_end=501,
)


_QUANTIFIERCONFIG = _descriptor.Descriptor(
  name='QuantifierConfig',
  full_name='data.QuantifierConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dead_zone', full_name='data.QuantifierConfig.dead_zone', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mid_thread', full_name='data.QuantifierConfig.mid_thread', index=1,
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
    _descriptor.OneofDescriptor(
      name='data', full_name='data.QuantifierConfig.data',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=503,
  serialized_end=625,
)


_FILEIMAGELAYER = _descriptor.Descriptor(
  name='FileImageLayer',
  full_name='data.FileImageLayer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rows', full_name='data.FileImageLayer.rows', index=0,
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
  serialized_start=627,
  serialized_end=657,
)


_FILEIMAGEHEADER = _descriptor.Descriptor(
  name='FileImageHeader',
  full_name='data.FileImageHeader',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='data.FileImageHeader.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='height', full_name='data.FileImageHeader.height', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='conversion', full_name='data.FileImageHeader.conversion', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subsampling', full_name='data.FileImageHeader.subsampling', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='wavelet', full_name='data.FileImageHeader.wavelet', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='quantifier', full_name='data.FileImageHeader.quantifier', index=5,
      number=6, type=11, cpp_type=10, label=1,
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
  serialized_start=660,
  serialized_end=850,
)


_FILEIMAGEDATA = _descriptor.Descriptor(
  name='FileImageData',
  full_name='data.FileImageData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='y', full_name='data.FileImageData.y', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='u', full_name='data.FileImageData.u', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='v', full_name='data.FileImageData.v', index=2,
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
  serialized_start=852,
  serialized_end=966,
)


_FILEIMAGE = _descriptor.Descriptor(
  name='FileImage',
  full_name='data.FileImage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='data.FileImage.header', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='data.FileImage.data', index=1,
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
  serialized_start=968,
  serialized_end=1053,
)

_IMAGEDATA.fields_by_name['rows'].message_type = _IMAGEROW
_PYTHONDWT.fields_by_name['data'].message_type = _IMAGEDATA
_WAVELETCONFIG.fields_by_name['haar'].message_type = _WAVELETHAAR
_WAVELETCONFIG.fields_by_name['daubechies'].message_type = _WAVELETDAUBECHIES
_WAVELETCONFIG.fields_by_name['dummy'].message_type = _WAVELETDUMMY
_WAVELETCONFIG.oneofs_by_name['data'].fields.append(
  _WAVELETCONFIG.fields_by_name['haar'])
_WAVELETCONFIG.fields_by_name['haar'].containing_oneof = _WAVELETCONFIG.oneofs_by_name['data']
_WAVELETCONFIG.oneofs_by_name['data'].fields.append(
  _WAVELETCONFIG.fields_by_name['daubechies'])
_WAVELETCONFIG.fields_by_name['daubechies'].containing_oneof = _WAVELETCONFIG.oneofs_by_name['data']
_WAVELETCONFIG.oneofs_by_name['data'].fields.append(
  _WAVELETCONFIG.fields_by_name['dummy'])
_WAVELETCONFIG.fields_by_name['dummy'].containing_oneof = _WAVELETCONFIG.oneofs_by_name['data']
_QUANTIFIERCONFIG.fields_by_name['dead_zone'].message_type = _QUANTIFIERDEADZONE
_QUANTIFIERCONFIG.fields_by_name['mid_thread'].message_type = _QUANTIFIERMIDTHREAD
_QUANTIFIERCONFIG.oneofs_by_name['data'].fields.append(
  _QUANTIFIERCONFIG.fields_by_name['dead_zone'])
_QUANTIFIERCONFIG.fields_by_name['dead_zone'].containing_oneof = _QUANTIFIERCONFIG.oneofs_by_name['data']
_QUANTIFIERCONFIG.oneofs_by_name['data'].fields.append(
  _QUANTIFIERCONFIG.fields_by_name['mid_thread'])
_QUANTIFIERCONFIG.fields_by_name['mid_thread'].containing_oneof = _QUANTIFIERCONFIG.oneofs_by_name['data']
_FILEIMAGEHEADER.fields_by_name['subsampling'].enum_type = _SUBSAMPLING
_FILEIMAGEHEADER.fields_by_name['wavelet'].message_type = _WAVELETCONFIG
_FILEIMAGEHEADER.fields_by_name['quantifier'].message_type = _QUANTIFIERCONFIG
_FILEIMAGEDATA.fields_by_name['y'].message_type = _FILEIMAGELAYER
_FILEIMAGEDATA.fields_by_name['u'].message_type = _FILEIMAGELAYER
_FILEIMAGEDATA.fields_by_name['v'].message_type = _FILEIMAGELAYER
_FILEIMAGE.fields_by_name['header'].message_type = _FILEIMAGEHEADER
_FILEIMAGE.fields_by_name['data'].message_type = _FILEIMAGEDATA
DESCRIPTOR.message_types_by_name['ImageRow'] = _IMAGEROW
DESCRIPTOR.message_types_by_name['ImageData'] = _IMAGEDATA
DESCRIPTOR.message_types_by_name['PythonDWT'] = _PYTHONDWT
DESCRIPTOR.message_types_by_name['WaveletHaar'] = _WAVELETHAAR
DESCRIPTOR.message_types_by_name['WaveletDaubechies'] = _WAVELETDAUBECHIES
DESCRIPTOR.message_types_by_name['WaveletDummy'] = _WAVELETDUMMY
DESCRIPTOR.message_types_by_name['WaveletConfig'] = _WAVELETCONFIG
DESCRIPTOR.message_types_by_name['QuantifierDeadZone'] = _QUANTIFIERDEADZONE
DESCRIPTOR.message_types_by_name['QuantifierMidThread'] = _QUANTIFIERMIDTHREAD
DESCRIPTOR.message_types_by_name['QuantifierConfig'] = _QUANTIFIERCONFIG
DESCRIPTOR.message_types_by_name['FileImageLayer'] = _FILEIMAGELAYER
DESCRIPTOR.message_types_by_name['FileImageHeader'] = _FILEIMAGEHEADER
DESCRIPTOR.message_types_by_name['FileImageData'] = _FILEIMAGEDATA
DESCRIPTOR.message_types_by_name['FileImage'] = _FILEIMAGE
DESCRIPTOR.enum_types_by_name['Subsampling'] = _SUBSAMPLING
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ImageRow = _reflection.GeneratedProtocolMessageType('ImageRow', (_message.Message,), dict(
  DESCRIPTOR = _IMAGEROW,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.ImageRow)
  ))
_sym_db.RegisterMessage(ImageRow)

ImageData = _reflection.GeneratedProtocolMessageType('ImageData', (_message.Message,), dict(
  DESCRIPTOR = _IMAGEDATA,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.ImageData)
  ))
_sym_db.RegisterMessage(ImageData)

PythonDWT = _reflection.GeneratedProtocolMessageType('PythonDWT', (_message.Message,), dict(
  DESCRIPTOR = _PYTHONDWT,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.PythonDWT)
  ))
_sym_db.RegisterMessage(PythonDWT)

WaveletHaar = _reflection.GeneratedProtocolMessageType('WaveletHaar', (_message.Message,), dict(
  DESCRIPTOR = _WAVELETHAAR,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.WaveletHaar)
  ))
_sym_db.RegisterMessage(WaveletHaar)

WaveletDaubechies = _reflection.GeneratedProtocolMessageType('WaveletDaubechies', (_message.Message,), dict(
  DESCRIPTOR = _WAVELETDAUBECHIES,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.WaveletDaubechies)
  ))
_sym_db.RegisterMessage(WaveletDaubechies)

WaveletDummy = _reflection.GeneratedProtocolMessageType('WaveletDummy', (_message.Message,), dict(
  DESCRIPTOR = _WAVELETDUMMY,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.WaveletDummy)
  ))
_sym_db.RegisterMessage(WaveletDummy)

WaveletConfig = _reflection.GeneratedProtocolMessageType('WaveletConfig', (_message.Message,), dict(
  DESCRIPTOR = _WAVELETCONFIG,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.WaveletConfig)
  ))
_sym_db.RegisterMessage(WaveletConfig)

QuantifierDeadZone = _reflection.GeneratedProtocolMessageType('QuantifierDeadZone', (_message.Message,), dict(
  DESCRIPTOR = _QUANTIFIERDEADZONE,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.QuantifierDeadZone)
  ))
_sym_db.RegisterMessage(QuantifierDeadZone)

QuantifierMidThread = _reflection.GeneratedProtocolMessageType('QuantifierMidThread', (_message.Message,), dict(
  DESCRIPTOR = _QUANTIFIERMIDTHREAD,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.QuantifierMidThread)
  ))
_sym_db.RegisterMessage(QuantifierMidThread)

QuantifierConfig = _reflection.GeneratedProtocolMessageType('QuantifierConfig', (_message.Message,), dict(
  DESCRIPTOR = _QUANTIFIERCONFIG,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.QuantifierConfig)
  ))
_sym_db.RegisterMessage(QuantifierConfig)

FileImageLayer = _reflection.GeneratedProtocolMessageType('FileImageLayer', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGELAYER,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.FileImageLayer)
  ))
_sym_db.RegisterMessage(FileImageLayer)

FileImageHeader = _reflection.GeneratedProtocolMessageType('FileImageHeader', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGEHEADER,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.FileImageHeader)
  ))
_sym_db.RegisterMessage(FileImageHeader)

FileImageData = _reflection.GeneratedProtocolMessageType('FileImageData', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGEDATA,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.FileImageData)
  ))
_sym_db.RegisterMessage(FileImageData)

FileImage = _reflection.GeneratedProtocolMessageType('FileImage', (_message.Message,), dict(
  DESCRIPTOR = _FILEIMAGE,
  __module__ = 'data_pb2'
  # @@protoc_insertion_point(class_scope:data.FileImage)
  ))
_sym_db.RegisterMessage(FileImage)


# @@protoc_insertion_point(module_scope)
