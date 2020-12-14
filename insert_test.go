package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/volatiletech/boilbench/gorms"
	"github.com/volatiletech/boilbench/gorps"
	"github.com/volatiletech/boilbench/mimic"
	"github.com/volatiletech/boilbench/models"
	"github.com/volatiletech/boilbench/xorms"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/gorp.v1"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

func BenchmarkGORMInsert(b *testing.B) {
	store := gorms.Jet{
		ID: 1,
	}

	exec := jetQueryInsert()
	exec.NumInput = -1
	mimic.NewResult(exec)

	gormdb, err := gorm.Open(gormMimicDialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	b.Run("gorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gormdb.Create(&store).Error
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkGORPInsert(b *testing.B) {
	store := gorps.Jet{
		ID: 1,
	}

	query := jetQueryInsert()
	query.NumInput = -1
	mimic.NewQuery(query)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	gorpdb := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	if err != nil {
		panic(err)
	}

	gorpdb.AddTable(gorps.Jet{}).SetKeys(true, "ID")

	b.Run("gorp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := gorpdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkXORMInsert(b *testing.B) {
	store := xorms.Jet{
		Id: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	xormdb, err := xorm.NewEngine("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("xorm", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := xormdb.Insert(&store)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkBoilInsert(b *testing.B) {
	store := models.Jet{
		ID: 1,
	}

	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	db, err := sql.Open("mimic", "")
	if err != nil {
		panic(err)
	}

	b.Run("boil", func(b *testing.B) {
		ctx := context.Background()

		for i := 0; i < b.N; i++ {
			err := store.Insert(ctx, db, boil.Infer())
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkEntInsert(b *testing.B) {
	exec := jetExec()
	exec.NumInput = -1
	mimic.NewResult(exec)

	client := openEnt()

	b.Run("ent", func(b *testing.B) {
		ctx := context.Background()

		for i := 0; i < b.N; i++ {
			_, err := client.Jet.Create().
				SetID(1).
				Save(ctx)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
