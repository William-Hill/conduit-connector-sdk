// Copyright © 2022 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import "errors"

var (
	// ErrBackoffRetry can be returned by Source.Read to signal the SDK there is
	// no record to fetch right now and it should try again later.
	ErrBackoffRetry = errors.New("backoff retry")

	// ErrUnimplemented is returned in functions of plugins that don't implement
	// a certain method.
	ErrUnimplemented = errors.New("action not implemented")

	// ErrMetadataFieldNotFound is returned in metadata utility functions when a
	// metadata field is not found.
	ErrMetadataFieldNotFound = errors.New("metadata field not found")
)
