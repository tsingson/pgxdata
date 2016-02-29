package data
// This file is automatically generated by pgxdata.

import (
  "strings"

  "github.com/jackc/pgx"
)

type Blob struct {
  ID Int32
  Payload Bytes
}

func CountBlob(db Queryer) (int64, error) {
  var n int64
  sql := `select count(*) from "blob"`
  err := db.QueryRow(sql).Scan(&n)
  return n, err
}

func SelectAllBlob(db Queryer) ([]Blob, error) {
  sql := `select
  "id",
  "payload"
from "blob"`

  var rows []Blob

  dbRows, err := db.Query(sql)
  if err != nil {
    return nil, err
  }

  for dbRows.Next() {
    var row Blob
    dbRows.Scan(
&row.ID,
    &row.Payload,
    )
    rows = append(rows, row)
  }

  if dbRows.Err() != nil {
    return nil, dbRows.Err()
  }

  return rows, nil
}

func SelectBlobByPK(
  db Queryer,
  id int32,
) (*Blob, error) {
  sql := `select
  "id",
  "payload"
from "blob"
where "id"=$1`

  var row Blob
  err := db.QueryRow(sql , id).Scan(
&row.ID,
    &row.Payload,
    )
  if err == pgx.ErrNoRows {
    return nil, ErrNotFound
  } else if err != nil {
    return nil, err
  }

  return &row, nil
}

func InsertBlob(db Queryer, row *Blob) error {
  args := pgx.QueryArgs(make([]interface{}, 0, 2))

  var columns, values []string

  row.ID.addInsert(`id`, &columns, &values, &args)
  row.Payload.addInsert(`payload`, &columns, &values, &args)


  sql := `insert into "blob"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "id"
  `

  return db.QueryRow(sql, args...).Scan(&row.ID)
}

func UpdateBlob(db Queryer,
  id int32,
  row *Blob,
) error {
  sets := make([]string, 0, 2)
  args := pgx.QueryArgs(make([]interface{}, 0, 2))

  row.ID.addUpdate(`id`, &sets, &args)
  row.Payload.addUpdate(`payload`, &sets, &args)


  if len(sets) == 0 {
    return nil
  }

  sql := `update "blob" set ` + strings.Join(sets, ", ") + ` where `  + `"id"=` + args.Append(id)

  commandTag, err := db.Exec(sql, args...)
  if err != nil {
    return err
  }
  if commandTag.RowsAffected() != 1 {
    return ErrNotFound
  }
  return nil
}

func DeleteBlob(db Queryer,
  id int32,
) error {
  args := pgx.QueryArgs(make([]interface{}, 0, 1))

  sql := `delete from "blob" where `  + `"id"=` + args.Append(id)

  commandTag, err := db.Exec(sql, args...)
  if err != nil {
    return err
  }
  if commandTag.RowsAffected() != 1 {
    return ErrNotFound
  }
  return nil
}

