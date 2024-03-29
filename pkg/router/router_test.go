//--------------------------------------------------------------------------
// Copyright 2018 infinimesh
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package router

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoute(t *testing.T) {
	const dlq = "default_or_dlq"
	const shadow = "shadow.reported-state.delta"
	router := New(dlq)

	require.Equal(t, shadow, router.Route("devices/0x1/state/reported/delta", "0x1"))
	require.Equal(t, shadow, router.Route("devices/0x1/state/reported/delta", "0x2"))
	require.Equal(t, dlq, router.Route("/devices/0x1/state/reported/delta", "0x1"))
	require.Equal(t, dlq, router.Route("", "0x1"))
	require.Equal(t, dlq, router.Route("/", "0x1"))
	require.Equal(t, dlq, router.Route("/abc", "0x1"))
	require.Equal(t, dlq, router.Route("/devices/0x1", "0x1"))
	require.Equal(t, dlq, router.Route("/devices/0x1/", "0x1"))
}
