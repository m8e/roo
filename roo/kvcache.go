// Copyright 2017-2019 Lei Ni (nilei81@gmail.com)
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
package main

import (
	"context"
	"errors"
)

// ErrCacheMiss is returned when a certificate is not found in cache.
var ErrCacheMiss = errors.New("acme/autocert: certificate cache miss")

// KvCache implements Cache using distributed rocksdb key-value store.
type KvCache struct {
}

// Get reads a certificate data from the specified file name.
func (d KvCache) Get(ctx context.Context, name string) ([]byte, error) {
	return nil, nil
}

// Put writes the certificate data to the specified file name.
// The file will be created with 0600 permissions.
func (d KvCache) Put(ctx context.Context, name string, data []byte) error {
	return nil
}

// Delete removes the specified file name.
func (d KvCache) Delete(ctx context.Context, name string) error {
	return nil
}
