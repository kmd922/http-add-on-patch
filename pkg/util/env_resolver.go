/*
Copyright 2023 The KEDA Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"os"
	"strconv"
	"time"
)

func ResolveOsEnvBool(envName string, defaultValue bool) (bool, error) {
	valueStr, found := os.LookupEnv(envName)

	if found && valueStr != "" {
		return strconv.ParseBool(valueStr)
	}

	return defaultValue, nil
}

func ResolveOsEnvInt(envName string, defaultValue int) (int, error) {
	valueStr, found := os.LookupEnv(envName)

	if found && valueStr != "" {
		return strconv.Atoi(valueStr)
	}

	return defaultValue, nil
}

func ResolveOsEnvDuration(envName string) (*time.Duration, error) {
	valueStr, found := os.LookupEnv(envName)

	if found && valueStr != "" {
		value, err := time.ParseDuration(valueStr)
		return &value, err
	}

	return nil, nil
}
