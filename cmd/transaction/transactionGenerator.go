package transaction

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/zoobc/zoobc-core/common/transaction"

	"github.com/zoobc/zoobc-core/common/crypto"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/util"
)

var txTypeMap = map[string][]byte{
	"sendMoney":    {1, 0, 0, 0},
	"registerNode": {2, 0, 0, 0},
}

func GenerateTransactionBytes(logger *logrus.Logger,
	signature crypto.SignatureInterface) *cobra.Command {
	var (
		txType string
	)
	var txCmd = &cobra.Command{
		Use:   "tx",
		Short: "tx command used to generate transaction.",
		Long: `tx command generate signed transaction bytes in form of hex or []bytes
		`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "generate" {
				seed := "prune filth cleaver removable earthworm tricky sulfur citation hesitate stout snort guy"

				tx := getTransaction(txTypeMap[txType])
				unsignedTxBytes, _ := util.GetTransactionBytes(tx, false)
				tx.Signature = signature.Sign(
					unsignedTxBytes,
					tx.SenderAccountType,
					tx.SenderAccountAddress,
					seed,
				)
				signedTxBytes, _ := util.GetTransactionBytes(tx, true)
				var signedTxByteString string
				for _, b := range signedTxBytes {
					signedTxByteString += fmt.Sprintf("%v, ", b)
				}
				logger.Printf("tx-bytes:byte = %v", signedTxByteString)
				logger.Printf("tx-bytes:hex = %s", hex.EncodeToString(signedTxBytes))
			} else {
				logger.Error("unknown command")
			}
		},
	}
	txCmd.Flags().StringVarP(&txType, "type", "t", "sendMoney", "number of account to generate")
	return txCmd
}

func getTransaction(txType []byte) *model.Transaction {
	switch util.ConvertBytesToUint32(txType) {
	case util.ConvertBytesToUint32(txTypeMap["sendMoney"]):
		amount := int64(10000)
		return &model.Transaction{
			Version:                 1,
			TransactionType:         util.ConvertBytesToUint32(txTypeMap["sendMoney"]),
			Timestamp:               time.Now().Unix(),
			SenderAccountType:       0,
			SenderAccountAddress:    "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgtN",
			RecipientAccountType:    0,
			RecipientAccountAddress: "BCZKLvgUYZ1KKx-jtF9KoJskjVPvB9jpIjfzzI6zDW0J",
			Fee:                     1,
			TransactionBodyLength:   8,
			TransactionBody: &model.Transaction_SendMoneyTransactionBody{
				SendMoneyTransactionBody: &model.SendMoneyTransactionBody{
					Amount: amount,
				},
			},
			TransactionBodyBytes: util.ConvertUint64ToBytes(uint64(amount)),
		}
	case util.ConvertBytesToUint32(txTypeMap["registerNode"]):
		txBody := &model.NodeRegistrationTransactionBody{
			AccountType:    0,
			AccountAddress: "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgtN",
			NodePublicKey: []byte{
				0, 14, 6, 218, 170, 54, 60, 50, 2, 66, 130, 119, 226, 235, 126, 203, 5, 12, 152, 194, 170, 146, 43, 63, 224,
				101, 127, 241, 62, 152, 187, 255,
			},
			NodeAddressLength: uint32(len([]byte("127.0.0.1"))),
			NodeAddress:       "127.0.0.1",
			LockedBalance:     100000,
		}
		txBodyBytes := (&transaction.NodeRegistration{
			Body: txBody,
		}).GetBodyBytes()
		return &model.Transaction{
			Version:                 1,
			TransactionType:         util.ConvertBytesToUint32(txTypeMap["registerNode"]),
			Timestamp:               time.Now().Unix(),
			SenderAccountType:       0,
			SenderAccountAddress:    "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgtN",
			RecipientAccountType:    0,
			RecipientAccountAddress: "BCZKLvgUYZ1KKx-jtF9KoJskjVPvB9jpIjfzzI6zDW0J",
			Fee:                     1,
			TransactionBodyLength:   uint32(len(txBodyBytes)),
			TransactionBody: &model.Transaction_NodeRegistrationTransactionBody{
				NodeRegistrationTransactionBody: txBody,
			},
			TransactionBodyBytes: txBodyBytes,
		}
	}
	return nil
}