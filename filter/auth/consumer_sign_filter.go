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

package auth

import (
	"context"
	"fmt"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetFilter(constant.AuthConsumerFilterKey, func() filter.Filter {
		return &ConsumerSignFilter{}
	})
	extension.SetFilter(constant.AuthProviderFilterKey, func() filter.Filter {
		return &ProviderAuthFilter{}
	})
}

// ConsumerSignFilter signs the request on consumer side
type ConsumerSignFilter struct{}

// Invoke retrieves the configured Authenticator to add signature to invocation
func (csf *ConsumerSignFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	logger.Infof("invoking ConsumerSign filter.")
	url := invoker.GetURL()

	err := doAuthWork(url, func(authenticator filter.Authenticator) error {
		return authenticator.Sign(invocation, url)
	})
	if err != nil {
		panic(fmt.Sprintf("Sign for invocation %s # %s failed", url.ServiceKey(), invocation.MethodName()))
	}
	return invoker.Invoke(ctx, invocation)
}

// OnResponse dummy process, returns the result directly
func (csf *ConsumerSignFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	return result
}
