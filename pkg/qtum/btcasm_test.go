package qtum

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestParseCallSenderASM(t *testing.T) {
	samStr := "1 7926223070547d2d15b2ef5e7383e541c338ffe9 69463043021f3ba540f52e0bae0c608c3d7135424fb683c77ee03217fcfe0af175c586aadc02200222e460a42268f02f130bc46f3ef62f228dd8051756dc13693332423515fcd401210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112 OP_SENDER 4 40000000 40 60fe47b10000000000000000000000000000000000000000000000000000000000000319 9e11fba86ee5d0ba4996b0d1973de6b694f4fc95 OP_CALL"
	got, err := ParseCallSenderASM(strings.Split(samStr, " "))
	if err != nil {
		t.Error(err)
	}
	want := &ContractInvokeInfo{
		From:     "7926223070547d2d15b2ef5e7383e541c338ffe9",
		GasLimit: "2625a00",
		GasPrice: "28",
		CallData: "60fe47b10000000000000000000000000000000000000000000000000000000000000319",
		To:       "9e11fba86ee5d0ba4996b0d1973de6b694f4fc95",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(
			"parse transaction call sam error\ninput: %s\nwant: %s\ngot: %s",
			samStr,
			string(mustMarshalIndent(want, "", "  ")),
			string(mustMarshalIndent(got, "", "  ")),
		)
	}
}

func TestParseParseCreateSenderASM(t *testing.T) {
	samStr := "1 7926223070547d2d15b2ef5e7383e541c338ffe9 6a473044022067ca66b0308ae16aeca7a205ce0490b44a61feebe5632710b52aabde197f9e4802200e8beec61a58dbe1279a9cdb68983080052ae7b9997bc863b7c5623e4cb55fdb01210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112 OP_SENDER 4 6721975 100 6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101de8061003b6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630900f010811461005d578063445df0ac1461007e5780638da5cb5b146100a3578063fdacd576146100d257600080fd5b341561006857600080fd5b61007c600160a060020a03600435166100e8565b005b341561008957600080fd5b61009161017d565b60405190815260200160405180910390f35b34156100ae57600080fd5b6100b6610183565b604051600160a060020a03909116815260200160405180910390f35b34156100dd57600080fd5b61007c600435610192565b6000805433600160a060020a03908116911614156101795781905080600160a060020a031663fdacd5766001546040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401600060405180830381600087803b151561016457600080fd5b6102c65a03f1151561017557600080fd5b5050505b5050565b60015481565b600054600160a060020a031681565b60005433600160a060020a03908116911614156101af5760018190555b505600a165627a7a72305820b6a912c5b5115d1a5412235282372dc4314f325bac71ee6c8bd18f658d7ed1ad0029 OP_CREATE"
	got, err := ParseCreateSenderASM(strings.Split(samStr, " "))
	if err != nil {
		t.Error(err)
	}
	want := &ContractInvokeInfo{
		From:     "7926223070547d2d15b2ef5e7383e541c338ffe9",
		GasLimit: "6691b7",
		GasPrice: "64",
		CallData: "6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a03199091161790556101de8061003b6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630900f010811461005d578063445df0ac1461007e5780638da5cb5b146100a3578063fdacd576146100d257600080fd5b341561006857600080fd5b61007c600160a060020a03600435166100e8565b005b341561008957600080fd5b61009161017d565b60405190815260200160405180910390f35b34156100ae57600080fd5b6100b6610183565b604051600160a060020a03909116815260200160405180910390f35b34156100dd57600080fd5b61007c600435610192565b6000805433600160a060020a03908116911614156101795781905080600160a060020a031663fdacd5766001546040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401600060405180830381600087803b151561016457600080fd5b6102c65a03f1151561017557600080fd5b5050505b5050565b60015481565b600054600160a060020a031681565b60005433600160a060020a03908116911614156101af5760018190555b505600a165627a7a72305820b6a912c5b5115d1a5412235282372dc4314f325bac71ee6c8bd18f658d7ed1ad0029",
		To:       "",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(
			"parse transaction call sam error\ninput: %s\nwant: %s\ngot: %s",
			samStr,
			string(mustMarshalIndent(want, "", "  ")),
			string(mustMarshalIndent(got, "", "  ")),
		)
	}
}

func mustMarshalIndent(v interface{}, prefix, indent string) []byte {
	res, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return res
}
