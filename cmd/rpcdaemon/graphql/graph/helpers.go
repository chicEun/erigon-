package graph

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"

	"github.com/erigontech/erigon-lib/common"
	hexutil2 "github.com/erigontech/erigon-lib/common/hexutil"

	"github.com/holiman/uint256"

	"github.com/erigontech/erigon-lib/common/hexutil"

	"github.com/erigontech/erigon-lib/types"
)

func convertDataToStringP(abstractMap map[string]interface{}, field string) *string {
	var result string

	switch v := abstractMap[field].(type) {
	case int64:
		result = strconv.FormatInt(v, 10)
	case *hexutil2.Big:
		if reflect.ValueOf(abstractMap[field]).IsZero() {
			return nil
		}
		result = v.String()
	case hexutil.Bytes:
		result = v.String()
	case hexutil2.Uint:
		result = v.String()
	case hexutil2.Uint64:
		result = v.String()
	case *common.Address:
		if reflect.ValueOf(abstractMap[field]).IsZero() {
			return nil
		}
		result = v.String()
	case common.Address:
		result = v.String()
	case common.Hash:
		result = v.String()
	case types.Bloom:
		result = hex.EncodeToString(v.Bytes())
	case types.BlockNonce:
		result = "0x" + fmt.Sprintf("%016x", int64(v.Uint64()))
	case []uint8:
		result = "0x" + hex.EncodeToString(v)
	case *uint256.Int:
		if reflect.ValueOf(abstractMap[field]).IsZero() {
			return nil
		}
		result = v.Hex()
	case uint64:
		result = "0x" + strconv.FormatInt(int64(v), 16)
	default:
		fmt.Println("unhandled/string", reflect.TypeOf(abstractMap[field]), field, abstractMap[field])
		result = "unhandled"
	}

	return &result
}

func convertDataToIntP(abstractMap map[string]interface{}, field string) *int {
	var result int

	switch v := abstractMap[field].(type) {
	case hexutil2.Uint64:
		resultUint, err := hexutil2.DecodeUint64(v.String())
		if err != nil {
			result = 0
		} else {
			result = int(resultUint)
		}
	case hexutil2.Uint:
		resultUint, err := hexutil2.DecodeUint64(v.String())
		if err != nil {
			result = 0
		} else {
			result = int(resultUint)
		}
	case int:
		result = v
	default:
		fmt.Println("unhandled/int", reflect.TypeOf(abstractMap[field]), field, abstractMap[field])
		result = 0
	}

	return &result
}

func convertDataToUint64P(abstractMap map[string]interface{}, field string) *uint64 {
	var result uint64

	switch v := abstractMap[field].(type) {
	case hexutil2.Uint64:
		resultUint, err := hexutil2.DecodeUint64(v.String())
		if err != nil {
			result = 0
		} else {
			result = resultUint
		}
	case hexutil2.Uint:
		resultUint, err := hexutil2.DecodeUint64(v.String())
		if err != nil {
			result = 0
		} else {
			result = resultUint
		}
	case *hexutil2.Big:
		result = v.ToInt().Uint64()
	case int:
		result = abstractMap[field].(uint64)
	case uint64:
		result = abstractMap[field].(uint64)
	default:
		fmt.Println("unhandled/uint64", reflect.TypeOf(abstractMap[field]), field, abstractMap[field])
		result = 0
	}

	return &result
}

func convertStrHexToDec(hexString *string) *string {
	var result string

	resUInt64, err := hexutil2.DecodeUint64(*hexString)
	if err != nil {
		fmt.Println(err)
		result = "0"
	}
	result = strconv.FormatUint(resUInt64, 10)

	return &result
}
