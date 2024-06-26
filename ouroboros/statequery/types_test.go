package statequery

import (
	"encoding/json"
	"github.com/thuannguyen2010/ogmigo/ouroboros/chainsync"
	"github.com/thuannguyen2010/ogmigo/ouroboros/chainsync/num"
	"reflect"
	"testing"
)

func TestUtxo_MarshalJSON(t *testing.T) {
	want := Utxo{
		TxIn: chainsync.TxIn{
			TxHash: "hash",
			Index:  0,
		},
		TxOut: chainsync.TxOut{
			Address: "address",
			Datum:   "datum",
			Value: chainsync.Value{
				Coins: num.Int64(123),
			},
		},
	}

	data, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	var got Utxo
	err = json.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("got %v; want nil", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
