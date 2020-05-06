package cache

import "testing"

func TestDB(t *testing.T) {
	c, err := ConnectDB("abc.db")
	if err != "" {
		print(err)
	}
	c.CreateDB()

}
