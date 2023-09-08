package chainsync

import (
	"github.com/stretchr/testify/assert"
	"github.com/thuannguyen2010/ogmigo/ouroboros/chainsync/num"
	"testing"
)

func TestConvertToV5Response(t *testing.T) {
	t.Run("Convert Intersection not Found", func(t *testing.T) {
		resV6 := ResponseV6{
			Jsonrpc: "",
			Method:  "",
			Result: &ResultV6{
				Error: &Error{},
			},
		}
		resV5 := resV6.ConvertToV5()
		assert.NotNil(t, resV5.Result.IntersectionNotFound)
	})

	t.Run("Convert Intersection Found", func(t *testing.T) {
		blockHash := "blockHash"
		slot := uint64(123123)
		resV6 := ResponseV6{
			Jsonrpc: "",
			Method:  "",
			Result: &ResultV6{
				Intersection: &PointV6{
					pointType: 1,
					pointStruct: &PointStructV6{
						ID:   blockHash,
						Slot: slot,
					},
				},
			},
		}
		resV5 := resV6.ConvertToV5()
		assert.NotNil(t, resV5.Result.IntersectionFound)
		ps, _ := resV5.Result.IntersectionFound.Point.PointStruct()
		assert.Equal(t, blockHash, ps.Hash)
	})

	t.Run("Convert backward block", func(t *testing.T) {
		resV6 := ResponseV6{
			Jsonrpc: "",
			Method:  "",
			Result: &ResultV6{
				Direction: "backward",
			},
		}
		resV5 := resV6.ConvertToV5()
		assert.NotNil(t, resV5.Result.RollBackward)
	})

	t.Run("Convert forward block", func(t *testing.T) {
		txID := "txID"
		fee := num.Int64(10)
		addr := "addr"
		datum := "datum"
		lovelace := num.Int64(11)
		inputTxID := "inputTxID"

		resV6 := ResponseV6{
			Jsonrpc: "",
			Method:  "",
			Result: &ResultV6{
				Direction: "forward",
				BlockV6: &BlockV6{
					Era: "babbage",
					Transactions: []Transaction{
						{
							ID: txID,
							TxInputs: []TxInV6{{
								Transaction: struct {
									ID string `json:"id,omitempty"`
								}(struct{ ID string }{ID: inputTxID}),
								Index: 0,
							}},
							TxOutputs: []TxOutV6{{
								Address: addr,
								Value: map[string]map[string]num.Int{
									"ada": {
										"lovelace": lovelace,
									},
								},
								Datum: datum,
							}},
							Fee: Fee{
								Lovelace: fee,
							},
						},
					},
				},
			},
		}
		resV5 := resV6.ConvertToV5()
		assert.NotNil(t, resV5.Result.RollForward)
		tx := resV5.Result.RollForward.Block.Babbage.Body[0]
		txIn := tx.Body.Inputs[0]
		txOut := tx.Body.Outputs[0]
		assert.Equal(t, inputTxID, txIn.TxHash)
		assert.Equal(t, lovelace, txOut.Value.Coins)
	})
}
