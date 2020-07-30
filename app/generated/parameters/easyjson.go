// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package parameters

import (
	models "battleship/app/generated/models"
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters(in *jlexer.Lexer, out *UserGetGoOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters(out *jwriter.Writer, in UserGetGoOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserGetGoOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserGetGoOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserGetGoOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserGetGoOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters1(in *jlexer.Lexer, out *UserGetGoBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			(out.Params).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters1(out *jwriter.Writer, in UserGetGoBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Params).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserGetGoBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserGetGoBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserGetGoBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserGetGoBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters1(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters2(in *jlexer.Lexer, out *UserCreateOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters2(out *jwriter.Writer, in UserCreateOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCreateOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters2(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters3(in *jlexer.Lexer, out *UserCreateBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			(out.Params).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters3(out *jwriter.Writer, in UserCreateBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Params).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserCreateBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserCreateBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserCreateBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserCreateBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters3(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters4(in *jlexer.Lexer, out *ShipListGoOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters4(out *jwriter.Writer, in ShipListGoOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ShipListGoOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ShipListGoOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShipListGoOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ShipListGoOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters4(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters5(in *jlexer.Lexer, out *ShipListGoBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				if out.Params == nil {
					out.Params = new(models.ShipListParams)
				}
				(*out.Params).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters5(out *jwriter.Writer, in ShipListGoBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Params == nil {
			out.RawString("null")
		} else {
			(*in.Params).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ShipListGoBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ShipListGoBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShipListGoBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ShipListGoBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters5(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters6(in *jlexer.Lexer, out *GameResetOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters6(out *jwriter.Writer, in GameResetOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameResetOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameResetOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameResetOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameResetOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters6(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters7(in *jlexer.Lexer, out *GameResetBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters7(out *jwriter.Writer, in GameResetBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GameResetBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GameResetBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GameResetBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GameResetBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters7(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters8(in *jlexer.Lexer, out *FightAddOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters8(out *jwriter.Writer, in FightAddOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FightAddOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FightAddOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FightAddOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FightAddOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters8(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters9(in *jlexer.Lexer, out *FightAddBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			(out.Params).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters9(out *jwriter.Writer, in FightAddBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Params).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FightAddBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FightAddBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FightAddBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FightAddBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters9(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters10(in *jlexer.Lexer, out *CoordinatesAddOKBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "result":
			(out.Result).UnmarshalEasyJSON(in)
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(models.Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters10(out *jwriter.Writer, in CoordinatesAddOKBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"result\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Result).MarshalEasyJSON(out)
	}
	if in.Error != nil {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Error).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CoordinatesAddOKBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CoordinatesAddOKBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CoordinatesAddOKBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CoordinatesAddOKBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters10(l, v)
}
func easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters11(in *jlexer.Lexer, out *CoordinatesAddBody) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "jsonrpc":
			out.Jsonrpc = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			(out.Params).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters11(out *jwriter.Writer, in CoordinatesAddBody) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"jsonrpc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Jsonrpc))
	}
	{
		const prefix string = ",\"method\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Method))
	}
	{
		const prefix string = ",\"params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Params).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CoordinatesAddBody) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CoordinatesAddBody) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeBattleshipAppGeneratedParameters11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CoordinatesAddBody) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CoordinatesAddBody) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeBattleshipAppGeneratedParameters11(l, v)
}
