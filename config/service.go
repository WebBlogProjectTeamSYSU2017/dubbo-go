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

package config

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

var (
	conServices = map[string]common.RPCService{} // service name -> service
	proServices = map[string]common.RPCService{} // service name -> service
)

// SetConsumerService is called by init() of implement of RPCService
func SetConsumerService(service common.RPCService) {
	ref := common.GetReference(service)
	conServices[ref] = service
}

// SetProviderService is called by init() of implement of RPCService
func SetProviderService(service common.RPCService) {
	ref := common.GetReference(service)
	proServices[ref] = service
}

// GetConsumerService gets ConsumerService by @name
func GetConsumerService(name string) common.RPCService {
	return conServices[name]
}

// GetProviderService gets ProviderService by @name
func GetProviderService(name string) common.RPCService {
	return proServices[name]
}

// GetAllProviderService gets all ProviderService
func GetAllProviderService() map[string]common.RPCService {
	return proServices
}

// GetCallback gets CallbackResponse by @name
func GetCallback(name string) func(response common.CallbackResponse) {
	service := GetConsumerService(name)
	if sv, ok := service.(common.AsyncCallbackService); ok {
		return sv.CallBack
	}
	return nil
}
