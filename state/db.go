package state

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type State struct {
	database *bolt.DB
}

func New() *State {

	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	state := &State{
		database: db,
	}

	state.database.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Service"))
		return err
	})

	return state
}

func (state *State) List() {
	state.database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Service"))

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			fmt.Println(string(k[:]))
		}
		return nil

	})
}

func (state *State) Remove(service string) {
	state.database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Service"))
		err := b.Delete([]byte(service))
		return err
	})
}

func (state *State) RemoveAll() {
	state.database.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("Service"))
		if err == nil {
			_, er := tx.CreateBucket([]byte("Service"))
			return er
		}
		return err
	})
}

func (state *State) GetService(service string) (command string) {

	var commander string
	state.database.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Service"))
		command := b.Get([]byte(service))
		commander = string(command[:])
		return nil
	})

	return commander

}

func (state *State) AddService(service string, command string) {
	state.database.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Service"))
		error_on_put := b.Put([]byte(service), []byte(command))
		return error_on_put
	})
}

func (state *State) Close() {
	state.database.Close()
}
