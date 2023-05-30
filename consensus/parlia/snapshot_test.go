package parlia

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
)

func TestValidatorSetSort(t *testing.T) {
	size := 100
	validators := make([]common.Address, size)
	for i := 0; i < size; i++ {
		validators[i] = randomAddress()
	}
	sort.Sort(validatorsAscending(validators))

	for i := 0; i < size-1; i++ {
		assert.True(t, bytes.Compare(validators[i][:], validators[i+1][:]) < 0)
	}
}

func TestValidatorSetDecode(t *testing.T) {
	//code := "0x00000000000000000000000000000000000000000000000000000000000000002a7cdd959bfe8d9487b2a43b33565295a698f7e26488aa4d1955ee33403f8ccb1d4de5fb97c7ade29ef9f4360c606c7ab4db26b016007d3ad0ab86a0ee01c3b1283aa067c58eab4709f85e99d46de5fe685b1ded8013785d6623cc18d214320b6bb6475978f3adfc719c99674c072166708589033e2d9afec2be4ec20253b8642161bc3f444f53679c1f3d472f7be8361c80a4c1e7e9aaf001d0877f1cfde218ce2fd7544e0b2cc94692d4a704debef7bcb61328b8f7166496996a7da21cf1f1b04d9b3e26a3d0772d4c407bbe49438ed859fe965b140dcf1aab71a96bbad7cf34b5fa511d8e963dbba288b1960e75d64430b3230294d12c6ab2aac5c2cd68e80b16b581ea0a6e3c511bbd10f4519ece37dc24887e11b55d7ae2f5b9e386cd1b50a4550696d957cb4900f03a82012708dafc9e1b880fd083b32182b869be8e0922b81f8e175ffde54d797fe11eb03f9e3bf75f1d68bf0b8b6fb4e317a0f9d6f03eaf8ce6675bc60d8c4d90829ce8f72d0163c1d5cf348a862d55063035e7a025f4da968de7e4d7e4004197917f4070f1d6caa02bbebaebb5d7e581e4b66559e635f805ff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	code := "0x00000000000000000000000000000000000000000000000000000000000000002a7cdd959bfe8d9487b2a43b33565295a698f7e20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	//validatorBytes := common.Hex2Bytes(code)
	validatorBytes, error := hexutil.Decode(code)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(len(validatorBytes))
	validatorBytes1 := validatorBytes[extraVanity : len(validatorBytes)-extraSeal]
	validators, err := ParseValidators(validatorBytes1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validators)
	fmt.Println(math.MaxUint64 / 2)

}

func TestValidatorSetEncode(t *testing.T) {
	//address := "0x431E1C7c0a8D0A41D413dc7b6171d2f6858f5138"
	//address := "0xc790c2d1514288f1367ef3b930b42e2d5e4caac1"
	validators := make([]common.Address, 5)
	//validators[0] = common.HexToAddress(address)
	//validators[1] = common.HexToAddress("0x431E1C7c0a8D0A41D413dc7b6171d2f6858f5138")
	//validators[0] = common.HexToAddress("01ce6113c76ba2b8a222052a4288eb72891d3b3a")
	//validators[1] = common.HexToAddress("155f849a5ba6b29f69e49521345cf642c591ca1e")
	//validators[2] = common.HexToAddress("18723dbecbf14ba6c249a5d28cc108c79a9df9ea")
	//validators[3] = common.HexToAddress("fe28f56e35f97d2d7209a8e8926a9c2e7f58458c")
	//validators[4] = common.HexToAddress("b2997c28f48fd9fb9fc841b97d299710a54fb16f")
	validators[0] = common.HexToAddress("fc899cc9c97334d2cbbd0a8149f7e58169558842")
	validators[1] = common.HexToAddress("18e474ef0f3287ced6ca84308b03e784818c1aa5")
	validators[2] = common.HexToAddress("47f5a7b1c987cb509776704f00a7e0d65f85ecd6")
	validators[3] = common.HexToAddress("05f49769eca13a295dfde7843a59310e577a834d")
	validators[4] = common.HexToAddress("28afaff411702e4ff616e85e9625d91e199bbaa9")
	fmt.Println(validators)
	//validatorBytes := common.Hex2Bytes(code)
	extra := make([]byte, extraVanity)
	for _, validator := range validators {
		extra = append(extra, validator.Bytes()...)
	}
	extra = append(extra, make([]byte, extraSeal)...)
	extraStr := hexutil.Encode(extra)
	fmt.Println(extraStr)
	fmt.Println(len(extra))
}
