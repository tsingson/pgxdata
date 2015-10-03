package main

import (
  "encoding/json"
  "errors"
  "fmt"
  "strconv"
  "time"

  "github.com/jackc/pgx"
)

const (
  Undefined = iota
  Null      = iota
  Present   = iota
)


type Bool struct {
  Value  bool
  Status byte
}

func (attr Bool) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr Bool) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr Bool) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *Bool) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullBool
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.Bool
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr Bool) FormatCode() int16 {
  var nv pgx.NullBool
  return nv.FormatCode()
}

func (attr Bool) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullBool
  nv.Bool = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}

type Int16 struct {
  Value  int16
  Status byte
}

func (attr Int16) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr Int16) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr Int16) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *Int16) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullInt16
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.Int16
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr Int16) FormatCode() int16 {
  var nv pgx.NullInt16
  return nv.FormatCode()
}

func (attr Int16) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullInt16
  nv.Int16 = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}

type Int32 struct {
  Value  int32
  Status byte
}

func (attr Int32) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr Int32) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr Int32) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *Int32) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullInt32
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.Int32
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr Int32) FormatCode() int16 {
  var nv pgx.NullInt32
  return nv.FormatCode()
}

func (attr Int32) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullInt32
  nv.Int32 = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}

type Int64 struct {
  Value  int64
  Status byte
}

func (attr Int64) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr Int64) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr Int64) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *Int64) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullInt64
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.Int64
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr Int64) FormatCode() int16 {
  var nv pgx.NullInt64
  return nv.FormatCode()
}

func (attr Int64) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullInt64
  nv.Int64 = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}

type String struct {
  Value  string
  Status byte
}

func (attr String) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr String) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr String) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *String) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullString
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.String
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr String) FormatCode() int16 {
  var nv pgx.NullString
  return nv.FormatCode()
}

func (attr String) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullString
  nv.String = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}

type Time struct {
  Value  time.Time
  Status byte
}

func (attr Time) String() string {
  switch attr.Status {
  case Present:
    return fmt.Sprintf("%v", attr.Value)
  case Null:
    return "null"
  case Undefined:
    return "undefined"
  }

  panic("unreachable")
}

func (attr Time) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr Time) addInsert(columnName string, sets, values *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName)
      *values = append(*values, args.Append(nil))
  }
}

func (attr *Time) Scan(r *pgx.ValueReader) error {
  var nv pgx.NullTime
  err := nv.Scan(r)
  if err != nil {
    return err
  }

  attr.Value = nv.Time
  if nv.Valid {
    attr.Status = Present
  } else {
    attr.Status = Null
  }

  return nil
}


func (attr Time) FormatCode() int16 {
  var nv pgx.NullTime
  return nv.FormatCode()
}

func (attr Time) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  var nv pgx.NullTime
  nv.Time = attr.Value

  switch attr.Status {
  case Present:
    nv.Valid = true
  case Null:
    nv.Valid = false
  case Undefined:
    return errors.New("cannot encode undefined attr")
  }

  return nv.Encode(w, oid)
}


type Bytes struct {
  Value  []byte
  Status byte
}

func (attr *Bytes) addUpdate(columnName string, sets *[]string, args *pgx.QueryArgs) {
  switch attr.Status {
    case Present:
      *sets = append(*sets, columnName+"="+args.Append(attr.Value))
    case Null:
      *sets = append(*sets, columnName+"="+args.Append(nil))
  }
}

func (attr *Bytes) Scan(vr *pgx.ValueReader) error {
  if vr.Len() == -1 {
    attr.Value = nil
    attr.Status = Null
    return nil
  }

  attr.Value = vr.ReadBytes(vr.Len())
  attr.Status = Present
  return vr.Err()
}


func (attr Bytes) FormatCode() int16 {
  return pgx.BinaryFormatCode
}

func (attr Bytes) Encode(w *pgx.WriteBuf, oid pgx.Oid) error {
  if oid != pgx.ByteaOid {
    return pgx.SerializationError(fmt.Sprintf("Bytes.Encode cannot encode into OID %d", oid))
  }

  if attr.Status != Present {
    w.WriteInt32(-1)
    return nil
  }

  w.WriteBytes(attr.Value)
  return nil
}


func (attr Bool) MarshalJSON() ([]byte, error) {
  if attr.Status != Present {
    return []byte("null"), nil
  }
  if attr.Value {
    return []byte("true"), nil
  }
  return []byte("false"), nil
}

func (attr *Bool) UnmarshalJSON(bval []byte) error {
  sval := string(bval)

  switch sval {
  case "true":
    attr.Value = true
    attr.Status = Present
  case "false":
    attr.Value = false
    attr.Status = Present
  case "null":
    attr.Status = Null
  default:
    return errors.New("unknown Bool value")
  }

  return nil
}


func (attr Int16) MarshalJSON() ([]byte, error) {
  if attr.Status != Present {
    return []byte("null"), nil
  }
  return []byte(strconv.FormatInt(int64(attr.Value), 10)), nil
}

func (attr *Int16) UnmarshalJSON(bval []byte) error {
  sval := string(bval)

  if sval == "null" {
    attr.Status = Null
    return nil
  }

  nval, err := strconv.ParseInt(sval, 10, 16)
  if err != nil {
    return err
  }

  attr.Value = int16(nval)
  attr.Status = Present

  return nil
}

func (attr Int32) MarshalJSON() ([]byte, error) {
  if attr.Status != Present {
    return []byte("null"), nil
  }
  return []byte(strconv.FormatInt(int64(attr.Value), 10)), nil
}

func (attr *Int32) UnmarshalJSON(bval []byte) error {
  sval := string(bval)

  if sval == "null" {
    attr.Status = Null
    return nil
  }

  nval, err := strconv.ParseInt(sval, 10, 32)
  if err != nil {
    return err
  }

  attr.Value = int32(nval)
  attr.Status = Present

  return nil
}

func (attr Int64) MarshalJSON() ([]byte, error) {
  if attr.Status != Present {
    return []byte("null"), nil
  }
  return []byte(strconv.FormatInt(int64(attr.Value), 10)), nil
}

func (attr *Int64) UnmarshalJSON(bval []byte) error {
  sval := string(bval)

  if sval == "null" {
    attr.Status = Null
    return nil
  }

  nval, err := strconv.ParseInt(sval, 10, 64)
  if err != nil {
    return err
  }

  attr.Value = int64(nval)
  attr.Status = Present

  return nil
}


func (attr String) MarshalJSON() ([]byte, error) {
  if attr.Status != Present {
    return []byte("null"), nil
  }

  return json.Marshal(attr.Value)
}

func (attr *String) UnmarshalJSON(bval []byte) error {
  sval := string(bval)

  if sval == "null" {
    attr.Status = Null
    return nil
  }

  err := json.Unmarshal(bval, &attr.Value)
  if err != nil {
    return err
  }

  attr.Status = Present
  return nil
}