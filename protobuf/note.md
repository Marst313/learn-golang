# Protobuf

## Protobuf are a language-neutral, platform-neutral extensible mechanism for serializing structured [data](https://protobuf.dev/overview/).

## In proto3 all field are optional by default

### Data types

#### Numeric types

1. int32/64 _signed integer_
2. uint32/64 _unsigned integer_
3. sint32/64 _signed integer with zig-zag encoding, efficient for negative numbers_
4. fixed32/64 _fixed-size integer_
5. sfixed32/64 _signed fixed-size integer_

#### Floating-Point Types`

1. float _32bit floating-point number_
2. double _64bit floating-point number_

#### Boolean types

1. boolean _true | false_

#### String types

1. string _UTF-8 characters, textual data_

#### Bytes types

1. bytes _sequence of raw bytes , useful binary data_

### Best Practice for using proto

1. Choose the right type
2. Avoid using `float` for precise values
3. Use Enumerations for Categorical Data
4. Plan for changes

### Field Numbers

`<field_type> <field_name> = <field_number>;`

#### Rules for Field Numbers

1. Use positive integers
2. Avoid zero
3. Unique within a message

#### Best Practices for Field Numbers

1. Use small numbers for frequently used field
2. Use lareger number for less common fields
3. Reserve numbers for future use

#### Changing Field Numbers

1. Decommission fields
2. Add new fields

## Step by step creating an grpc server

1. Create a file.proto
2. Compile it to a specific programming language
3. install grpc on specific programming language (google.golang.org/grpc)
4. Create new connection from server -> grpc
5. Register a new connection -> grpc
6. Serve grpc with listener
