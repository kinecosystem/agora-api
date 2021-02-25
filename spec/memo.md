# Kin Binary Memo Format

This outlines the encoding required for a transaction memo to conform to the binary memo format expected by Kin.

```
 0               1               2               3
 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|MXB| Ver |TypeID   |            AppIdx             |    FK     |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                            FK Cont                            |
|                             ....                              |
|                             ....                              |
|                             ....                              |
|                             ....                              |
|                             ....                              |
|                             ....                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

**MXB**: Magic Byte Indicator (2 bits)

    This field is intended to add a hint for memo decoders to
    infer that the memo format is the one described in this
    document. If the MXB matches the defined constant, however,
    it does not gaurentee that the memo encoding is that of the
    one described in this document. The value of the magic byte
    is `0x1`.

**Ver**: Version (3 bits)

    This field indicates the memo encoding version. 3 bits
    allows for up to 7 versions before an extension version
    is required.

**TypeID**: Transaction Type (5 bits)

    This field indicates the 'meta' type of transaction(s)
    that the transaction the memo was embodied in was for.
    See Type Identifiers for mapping.

**AppIdx**: App Index (16 bits)

    This field refers to the App in which the transaction
    that the memo was embodied in was related to. The
    string representation for the corresponding App Index
    can be looked up via a central service. 16 bits allows
    for 65,536 apps before a memo revision is required.

**FK**: Foreign Key (230 bits)

    The identifier in an auxiliary transaction service that
    contains the metadata about what the transaction was for.
    For example, it may contain the advertised Kin price,
    a brief description (i.e. if a product sku, or an offer).

### Type Identifiers

| Value  | Type     |
| -----  |:--------:|
| 0      | NONE     |
| 1      | EARN     |
| 2      | SPEND    |
| 3      | P2P      |
| 4      | RESERVED |
| 5      | RESERVED |
| 6      | RESERVED |
| 7      | RESERVED |
| 8      | RESERVED |
| 9      | RESERVED |
| 10     | RESERVED |
| 11     | RESERVED |
| 12     | RESERVED |
| 13     | RESERVED |
| 14     | RESERVED |
| 15     | RESERVED |
| 16     | RESERVED |
| 17     | RESERVED |
| 18     | RESERVED |
| 19     | RESERVED |
| 20     | RESERVED |
| 21     | RESERVED |
| 22     | RESERVED |
| 23     | RESERVED |
| 24     | RESERVED |
| 25     | RESERVED |
| 26     | RESERVED |
| 27     | RESERVED |
| 28     | RESERVED |
| 29     | RESERVED |
| 30     | RESERVED |
| 31     | RESERVED |
