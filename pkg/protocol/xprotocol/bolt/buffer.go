/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bolt

import (
	"context"

	mbuffer "mosn.io/mosn/pkg/buffer"
	"mosn.io/pkg/buffer"
	"mosn.io/pkg/log"
)

var ins boltBufferCtx

func init() {
	mbuffer.RegisterBuffer(&ins)
}

type boltBufferCtx struct {
	mbuffer.TempBufferCtx
}

func (ctx boltBufferCtx) New() interface{} {
	return new(boltBuffer)
}

func (ctx boltBufferCtx) Reset(i interface{}) {
	buf, _ := i.(*boltBuffer)

	// recycle ioBuffer
	if buf.request.Data != nil {
		if e := buffer.PutIoBuffer(buf.request.Data); e != nil {
			log.DefaultLogger.Errorf("[protocol] [bolt] [buffer] [reset] PutIoBuffer error: %v", e)
		}
	}

	if buf.response.Data != nil {
		if e := buffer.PutIoBuffer(buf.response.Data); e != nil {
			log.DefaultLogger.Errorf("[protocol] [bolt] [buffer] [reset] PutIoBuffer error: %v", e)
		}
	}

	*buf = boltBuffer{}
}

type boltBuffer struct {
	request  Request
	response Response
}

func bufferByContext(ctx context.Context) *boltBuffer {
	poolCtx := mbuffer.PoolContext(ctx)
	return poolCtx.Find(&ins, nil).(*boltBuffer)
}
