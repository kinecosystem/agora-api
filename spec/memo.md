# Memo Encoding


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
    one described in this document.

**Ver**: Version (3 bits)

    This field indicates the memo encoding version. 3 bits
    allows for up to 7 versions before an extension version
    is required.

**TypeID**: Transaction Type (5 bits)

    This field indicates the 'meta' type of transaction(s)
    that the transaction the memo was emboddied in was for.
    See Type Identifiers for mapping.

**AppIdx**: App Index (16 bits)

    This field refers to the App in which the transcation
    that the memo was emboddied in was related to. The
    string representation for the corresponding App Index
    can be looked up via a central service. 16 bits allows
    for 65,536 apps before a memo revision is required.

**FK**: Foreign Key (230 bits)

    The identifier in an auxiliary transaction service that
    contains the metadata about what the transaction was for.
    For example, it may contain the advertised Kin price,
    a brief description (i.e. if a product sku, or an offer).

### Type Identifiers

| Value  | Type  |
| -----  |:-----:|
| 0      | TODO  |
| 1      | TODO  |
| 2      | TODO  |
| 3      | TODO  |
| 4      | TODO  |
| 5      | TODO  |
| 6      | TODO  |
| 7      | TODO  |
| 8      | TODO  |
| 9      | TODO  |
| 10     | TODO  |
| 11     | TODO  |
| 12     | TODO  |
| 13     | TODO  |
| 14     | TODO  |
| 15     | TODO  |
| 16     | TODO  |
| 17     | TODO  |
| 18     | TODO  |
| 19     | TODO  |
| 20     | TODO  |
| 21     | TODO  |
| 22     | TODO  |
| 23     | TODO  |
| 24     | TODO  |
| 25     | TODO  |
| 26     | TODO  |
| 27     | TODO  |
| 28     | TODO  |
| 29     | TODO  |
| 30     | TODO  |
| 31     | TODO  |
| 32     | TODO  |
