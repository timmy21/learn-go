package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type ColScanSet[T any] []ColScan[T]

func (c ColScanSet[T]) Fields() []string {
	fields := make([]string, len(c))
	for i, item := range c {
		fields[i] = item.Col
	}
	return fields
}

type ColScan[T any] struct {
	Col  string
	Scan func(item *T) any
	Bind func(item *T, field any) error
}

// 这是一个 sql.Rows 的一个封装，用于将查询数据库返回的字段值绑定到结构体对象上。
type StructRows[T any] struct {
	*sql.Rows
	ColScanSet[T]
}

func (r *StructRows[T]) ScanStruct(item *T) error {
	var binds []struct {
		idx int
		fn  func(*T, any) error
	}
	dests := make([]any, 0, len(r.ColScanSet))
	for i, cs := range r.ColScanSet {
		dests = append(dests, cs.Scan(item))
		if cs.Bind != nil {
			binds = append(binds, struct {
				idx int
				fn  func(*T, any) error
			}{
				idx: i,
				fn:  cs.Bind,
			})
		}
	}
	err := r.Rows.Scan(dests...)
	if err != nil {
		return err
	}
	for _, b := range binds {
		if err := b.fn(item, dests[b.idx]); err != nil {
			return err
		}
	}
	return nil
}

type address struct {
	city string
	addr string
}

type person struct {
	name    string
	address address
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table person(id integer not null primary key, name text, city text, addr text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}

	_, err = db.Exec("insert into person(id, name, city, addr) values(1, 'tony', '上海', 'xxx'), (2, 'tom', '北京', 'yyy')")
	if err != nil {
		log.Fatal(err)
	}

	func() {
		// 简单迭代
		rows, err := db.Query("select name, city, addr from person")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var name, city, addr string
			err = rows.Scan(&name, &city, &addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(name, city, addr)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}()
	func() {
		// 结构体迭代
		colScans := ColScanSet[person]{
			{
				Col: "name",
				Scan: func(item *person) any {
					return &item.name
				},
				Bind: func(item *person, field any) error {
					item.name = strings.ToUpper(*field.(*string))
					return nil
				},
			},
			{
				Col: "city",
				Scan: func(item *person) any {
					return &item.address.city
				},
			},
			{
				Col: "addr",
				Scan: func(item *person) any {
					return &item.address.addr
				},
			},
		}
		rows, err := db.Query("select " + strings.Join(colScans.Fields(), ",") + " from person")
		if err != nil {
			log.Fatal(err)
		}
		// 结构体泛型当前不支持类型推断
		sRows := StructRows[person]{
			Rows:       rows,
			ColScanSet: colScans,
		}
		defer sRows.Close()
		for sRows.Next() {
			var p person
			err := sRows.ScanStruct(&p)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%+v\n", p)
		}
		err = sRows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
