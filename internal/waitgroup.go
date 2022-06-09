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
	"sync"
	"time"
)

type WaitGroup sync.WaitGroup

func (wg *WaitGroup) Add(delta int) {
	(*sync.WaitGroup)(wg).Add(delta)
}

func (wg *WaitGroup) Done() {
	(*sync.WaitGroup)(wg).Done()
}

func (wg *WaitGroup) Wait(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		(*sync.WaitGroup)(wg).Wait()
		close(done)
	}()
	_, _, err := Chan[struct{}](done).Recv(ctx)
	return err
}

func (wg *WaitGroup) WaitTimeout(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return wg.Wait(ctx)
}
