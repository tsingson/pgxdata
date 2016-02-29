package data
// This file is automatically generated by pgxdata.

import (
  "strings"

  "github.com/jackc/pgx"
)

type Part struct {
  Code String
  Description String
}

func CountPart(db Queryer) (int64, error) {
  var n int64
  sql := `select count(*) from "part"`
  err := db.QueryRow(sql).Scan(&n)
  return n, err
}

func SelectAllPart(db Queryer) ([]Part, error) {
  sql := `select
  "code",
  "description"
from "part"`

  var rows []Part

  dbRows, err := db.Query(sql)
  if err != nil {
    return nil, err
  }

  for dbRows.Next() {
    var row Part
    dbRows.Scan(
&row.Code,
    &row.Description,
    )
    rows = append(rows, row)
  }

  if dbRows.Err() != nil {
    return nil, dbRows.Err()
  }

  return rows, nil
}

func SelectPartByPK(
  db Queryer,
  code string,
) (*Part, error) {
  sql := `select
  "code",
  "description"
from "part"
where "code"=$1`

  var row Part
  err := db.QueryRow(sql , code).Scan(
&row.Code,
    &row.Description,
    )
  if err == pgx.ErrNoRows {
    return nil, ErrNotFound
  } else if err != nil {
    return nil, err
  }

  return &row, nil
}

func InsertPart(db Queryer, row *Part) error {
  args := pgx.QueryArgs(make([]interface{}, 0, 2))

  var columns, values []string

  row.Code.addInsert(`code`, &columns, &values, &args)
  row.Description.addInsert(`description`, &columns, &values, &args)


  sql := `insert into "part"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "code"
  `

  return db.QueryRow(sql, args...).Scan(&row.Code)
}

func UpdatePart(db Queryer,
  code string,
  row *Part,
) error {
  sets := make([]string, 0, 2)
  args := pgx.QueryArgs(make([]interface{}, 0, 2))

  row.Code.addUpdate(`code`, &sets, &args)
  row.Description.addUpdate(`description`, &sets, &args)


  if len(sets) == 0 {
    return nil
  }

  sql := `update "part" set ` + strings.Join(sets, ", ") + ` where `  + `"code"=` + args.Append(code)

  commandTag, err := db.Exec(sql, args...)
  if err != nil {
    return err
  }
  if commandTag.RowsAffected() != 1 {
    return ErrNotFound
  }
  return nil
}

func DeletePart(db Queryer,
  code string,
) error {
  args := pgx.QueryArgs(make([]interface{}, 0, 1))

  sql := `delete from "part" where `  + `"code"=` + args.Append(code)

  commandTag, err := db.Exec(sql, args...)
  if err != nil {
    return err
  }
  if commandTag.RowsAffected() != 1 {
    return ErrNotFound
  }
  return nil
}

