// Package TestApp comment
// This file was generated by tars2go 1.1
// Generated from SayHello.tars
package TestApp

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)

//SayHello struct
type SayHello struct {
	s m.Servant
}

//EchoHello is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) EchoHello(Name string, Greeting *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Name, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*Greeting), 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()

	err = _obj.s.Tars_invoke(ctx, 0, "echoHello", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Greeting), 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//EchoHelloWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) EchoHelloWithContext(ctx context.Context, Name string, Greeting *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Name, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*Greeting), 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(ctx, 0, "echoHello", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Greeting), 2, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//EchoHelloOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) EchoHelloOneWayWithContext(ctx context.Context, Name string, Greeting *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Name, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string((*Greeting), 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(ctx, 1, "echoHello", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *SayHello) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *SayHello) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *SayHello) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *SayHello) AddServant(imp _impSayHello, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *SayHello) AddServantWithContext(imp _impSayHelloWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impSayHello interface {
	EchoHello(Name string, Greeting *string) (ret int32, err error)
}
type _impSayHelloWithContext interface {
	EchoHello(ctx context.Context, Name string, Greeting *string) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *SayHello) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "echoHello":
		var Name string
		err = _is.Read_string(&Name, 1, true)
		if err != nil {
			return err
		}

		var Greeting string
		err = _is.Read_string(&Greeting, 2, false)
		if err != nil {
			return err
		}

		if !withContext {
			_imp := _val.(_impSayHello)
			ret, err := _imp.EchoHello(Name, &Greeting)
			if err != nil {
				return err
			}

			err = _os.Write_int32(ret, 0)
			if err != nil {
				return err
			}

		} else {
			_imp := _val.(_impSayHelloWithContext)
			ret, err := _imp.EchoHello(ctx, Name, &Greeting)
			if err != nil {
				return err
			}

			err = _os.Write_int32(ret, 0)
			if err != nil {
				return err
			}

		}

		err = _os.Write_string(Greeting, 2)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}
