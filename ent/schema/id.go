package schema

import (
	"crypto/rand"
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

type IDMixin struct {
	mixin.Schema
}

type ID string

func (i ID) Value() (driver.Value, error) {

	return i.ToInt64()
}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (i *ID) ToInt64() (int64, error) {
	str := string(*i)
	result := int64(0)
	for i := 0; i < len(str); i++ {
		result *= int64(len(alphabet))
		value := int64(strings.IndexByte(alphabet, str[i]))
		if value == -1 {
			return 0, errors.New("not a valid base62")
		}
		result += value
	}
	return result, nil
}

func NewID() ID {
	return ID(RandomString(8))
}

func (i *ID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case int64:
		*i = ID(Int64ToString(v))
		return nil
	}
	return errors.New("not a valid base62")
}

func Int64ToString(value int64) string {
	digit := int(math.Log(float64(value))/math.Log(float64(len(alphabet)))) + 1
	data := make([]uint8, digit)
	r := int64(len(alphabet))
	for i := 0; i < len(data); i++ {
		data[len(data)-i-1] = alphabet[value%r]
		value = value / r
	}
	if value != 0 {
		fmt.Println("something wrong at id")
	}
	return string(data)
}

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(n int) string {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret)
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", ID("")).
			SchemaType(map[string]string{
				dialect.MySQL:    "bigint",
				dialect.Postgres: "bigint",
				dialect.SQLite:   "integer",
			}).
			Default(NewID).
			Immutable(),
	}
}
