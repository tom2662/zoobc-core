package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	rpc_model "github.com/zoobc/zoobc-core/common/model"
	rpc_service "github.com/zoobc/zoobc-core/common/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := rpc_service.NewTransactionServiceClient(conn)

	response, err := c.PostTransaction(context.Background(), &rpc_model.PostTransactionRequest{
		// TransactionBytes: []byte{
		//	1, 0, 1, 82, 108, 58, 93, 0, 0, 0, 0, 0, 0, 66, 67, 90, 110, 83, 102, 113, 112, 80, 53, 116, 113, 70, 81, 108, 77, 84, 89,
		//	107, 68, 101, 66, 86, 70, 87, 110, 98, 121, 86, 75, 55, 118, 76, 114, 53, 79, 82, 70, 112, 84, 106, 103, 116, 78, 0, 0, 66,
		//	67, 90, 75, 76, 118, 103, 85, 89, 90, 49, 75, 75, 120, 45, 106, 116, 70, 57, 75, 111, 74, 115, 107, 106, 86, 80, 118, 66, 57,
		//	106, 112, 73, 106, 102, 122, 122, 73, 54, 122, 68, 87, 48, 74, 1, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 16, 39, 0, 0, 0, 0, 0, 0, 140,
		//	134, 217, 86, 232, 251, 81, 174, 86, 44, 221, 44, 226, 73, 245, 19, 170, 94, 47, 160, 53, 20, 225, 192, 19, 200, 196, 217, 96, 64,
		//	66, 6, 146, 16, 61, 104, 106, 112, 122, 96, 233, 224, 208, 119, 245, 148, 60, 9, 131, 211, 110, 68, 167, 115, 243, 251, 90, 64, 234,
		//	66, 108, 30, 116, 9,
		// },
		TransactionBytes: []byte{
			1, 0, 1, 53, 119, 58, 93, 0, 0, 0, 0, 0, 0, 66, 67, 90, 110, 83, 102, 113, 112, 80, 53, 116, 113, 70, 81, 108, 77, 84, 89, 107, 68,
			101, 66, 86, 70, 87, 110, 98, 121, 86, 75, 55, 118, 76, 114, 53, 79, 82, 70, 112, 84, 106, 103, 116, 78, 0, 0, 66, 67, 90, 75, 76,
			118, 103, 85, 89, 90, 49, 75, 75, 120, 45, 106, 116, 70, 57, 75, 111, 74, 115, 107, 106, 86, 80, 118, 66, 57, 106, 112, 73, 106,
			102, 122, 122, 73, 54, 122, 68, 87, 48, 74, 1, 0, 0, 0, 0, 0, 0, 0, 8, 0, 0, 0, 16, 39, 0, 0, 0, 0, 0, 0, 32, 85, 34, 198, 89, 78,
			166, 142, 59, 148, 243, 133, 69, 66, 123, 219, 2, 3, 229, 172, 221, 35, 185, 208, 43, 44, 172, 96, 166, 116, 205, 93, 78, 194, 153,
			95, 243, 145, 108, 96, 42, 6, 186, 128, 59, 117, 83, 196, 26, 9, 15, 157, 215, 108, 180, 35, 195, 100, 7, 142, 47, 96, 108, 10,
		},
	})

	if err != nil {
		log.Fatalf("error calling rpc_service.PostTransaction: %s", err)
	}

	log.Printf("response from remote rpc_service.PostTransaction(): %s", response)

}
