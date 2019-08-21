package crypto

import (
	"bytes"
	"encoding/hex"
	"github.com/zoobc/zoobc-core/common/blocker"
	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/util"
)

var LastRequestTimestamp uint64

// VerifyAuthAPI verify the request body and the signature of the request message, checking include
// request type checking, and the validity of the signature to the owner address
// return nil if valid, and Blocker object otherwise
func VerifyAuthAPI(
	ownerAddress,
	auth string,
	requestType model.RequestType) error {
	// parse
	var (
		authRequestType int32
		authTimestamp   uint64
	)
	signature := NewSignature()
	authBytes, err := hex.DecodeString(auth)
	if err != nil {
		return err
	}
	authBytesBuffer := bytes.NewBuffer(authBytes)
	authRequestType = int32(util.ConvertBytesToUint32(authBytesBuffer.Next(constant.AuthRequestType)))
	authTimestamp = util.ConvertBytesToUint64(authBytesBuffer.Next(constant.AuthTimestamp))
	if authRequestType != int32(requestType) {
		return blocker.NewBlocker(
			blocker.RequestParameterErr,
			"invalid request type",
		)
	}
	if authTimestamp <= LastRequestTimestamp {
		return blocker.NewBlocker(
			blocker.ValidationErr,
			"timestamp is in the past",
		)
	}
	LastRequestTimestamp = authTimestamp
	signatureValid := signature.VerifySignature(
		authBytes[:constant.AuthRequestType+constant.AuthTimestamp],
		authBytes[constant.AuthRequestType+constant.AuthTimestamp:],
		ownerAddress,
	)
	if !signatureValid {
		return blocker.NewBlocker(blocker.ValidationErr, "request signature invalid")
	}
	return nil
}