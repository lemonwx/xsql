// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"bytes"
	"fmt"
)

// TrackedBuffer is used to rebuild a query from the ast.
// bindLocations keeps track of locations in the buffer that
// use bind variables for efficient future substitutions.
// nodeFormatter is the formatting function the buffer will
// use to format a node. By default(nil), it's FormatNode.
// But you can supply a different formatting function if you
// want to generate a query that's different from the default.
type TrackedBuffer struct {
	*bytes.Buffer
	nodeFormatter func(buf *TrackedBuffer, node SQLNode)
}

func NewTrackedBuffer(nodeFormatter func(buf *TrackedBuffer, node SQLNode)) *TrackedBuffer {
	buf := &TrackedBuffer{
		Buffer:        bytes.NewBuffer(make([]byte, 0, 128)),
		nodeFormatter: nodeFormatter,
	}
	return buf
}

// Fprintf mimics fmt.Fprintf, but limited to Node(%v), Node.Value(%s) and string(%s).
// It also allows a %a for a value argument, in which case it adds tracking info for
// future substitutions.
func (buf *TrackedBuffer) Fprintf(format string, values ...interface{}) {
	end := len(format)
	fieldnum := 0
	for i := 0; i < end; {
		lasti := i
		for i < end && format[i] != '%' {
			i++
		}
		if i > lasti {
			buf.WriteString(format[lasti:i])
		}
		if i >= end {
			break
		}
		i++ // '%'
		switch format[i] {
		case 'c':
			switch v := values[fieldnum].(type) {
			case byte:
				buf.WriteByte(v)
			case rune:
				buf.WriteRune(v)
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}
		case 's':
			switch v := values[fieldnum].(type) {
			case []byte:
				buf.Write(v)
			case string:
				buf.WriteString(v)
			default:
				panic(fmt.Sprintf("unexpected type %T", v))
			}
		case 'v':
			node := values[fieldnum].(SQLNode)
			if buf.nodeFormatter == nil {
				node.Format(buf)
			} else {
				buf.nodeFormatter(buf, node)
			}
		case 'a':
			buf.WriteArg(values[fieldnum].(string))
		default:
			panic("unexpected")
		}
		fieldnum++
		i++
	}
}

func (buf *TrackedBuffer) WriteArg(arg string) {
	buf.WriteString(arg)
}
