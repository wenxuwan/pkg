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

package types

import (
	"errors"

	"mosn.io/api"
)

// MOSN Header keys
const (
	HeaderGlobalTimeout = "x-mosn-global-timeout"
	HeaderTryTimeout    = "x-mosn-try-timeout"
	HeaderOriginalPath  = "x-mosn-original-path"
)

// Error messages
const (
	ChannelFullException = "Channel is full"
	CodecException       = "codec exception occurs"
	SerializeException   = "serialize exception occurs"
	DeserializeException = "deserialize exception occurs"

	NoStatusCodeForHijackException = "no status code found for hijack reply"
)

// Errors
var (
	ErrChanFull             = errors.New(ChannelFullException)
	ErrCodecException       = errors.New(CodecException)
	ErrSerializeException   = errors.New(SerializeException)
	ErrDeserializeException = errors.New(DeserializeException)

	ErrNoStatusCodeForHijack = errors.New(NoStatusCodeForHijackException)
)

// Error codes, used by top level logic code(like proxy logic).
const (
	CodecExceptionCode    = 0
	UnknownCode           = 2
	DeserialExceptionCode = 3
	SuccessCode           = 200
	PermissionDeniedCode  = 403
	RouterUnavailableCode = 404
	InternalErrorCode     = 500
	NoHealthUpstreamCode  = 502
	UpstreamOverFlowCode  = 503
	TimeoutExceptionCode  = 504
	LimitExceededCode     = 509
)

// ResponseFlags sets
const (
	MosnProcessFailedFlags = api.NoHealthyUpstream | api.NoRouteFound | api.UpstreamLocalReset |
		api.FaultInjected | api.RateLimited | api.DownStreamTerminate | api.ReqEntityTooLarge
)
