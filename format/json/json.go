package json

import (
	"context"
	"fmt"
	"github.com/netsampler/goflow2/format"
	"github.com/netsampler/goflow2/format/common"
)

type JsonDriver struct {
}

func (d *JsonDriver) Prepare() error {
	common.HashFlag()
	common.SelectorFlag()
	return nil
}

func (d *JsonDriver) Init(context.Context) error {
	err := common.ManualHashInit()
	if err != nil {
		return err
	}
	return common.ManualSelectorInit()
}

func (d *JsonDriver) Format(data interface{}) ([]byte, []byte, error) {
	if dataIf, ok := data.(interface{ MarshalJSON() ([]byte, error) }); ok {
		d, err := dataIf.MarshalJSON()
		return []byte("sth"), d, err
	}
	return nil, nil, fmt.Errorf("message is not serializable in json")

	/*
	   msg, ok := data.(proto.Message)

	   	if !ok {
	   		return nil, nil, fmt.Errorf("message is not protobuf")
	   	}

	   key := common.HashProtoLocal(msg)
	   return []byte(key), []byte(common.FormatMessageReflectJSON(msg, "")), nil
	*/
}

func init() {
	d := &JsonDriver{}
	format.RegisterFormatDriver("json", d)
}
