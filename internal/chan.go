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

package internal

import (
	"context"
	"time"
)

type Chan[T any] chan T

func (c Chan[T]) Recv(ctx context.Context) (T, bool, error) {
	select {
	case val, ok := <-c:
		return val, ok, nil
	case <-ctx.Done():
		var empty T
		return empty, false, ctx.Err()
	}
}

func (c Chan[T]) RecvTimeout(ctx context.Context, timeout time.Duration) (T, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return c.Recv(ctx)
}
