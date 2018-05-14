// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
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

package commands

import (
	"go.thethings.network/lorawan-stack/cmd/internal/shared"
	shared_identityserver "go.thethings.network/lorawan-stack/cmd/internal/shared/identityserver"
	conf "go.thethings.network/lorawan-stack/pkg/config"
	"go.thethings.network/lorawan-stack/pkg/identityserver"
)

// Config for the ttn-lw-identity-server binary.
type Config struct {
	conf.ServiceBase `name:",squash"`
	IS               identityserver.Config `name:"is"`
}

// DefaultConfig contains the default config for the ttn-lw-identity-server binary.
var DefaultConfig = Config{
	ServiceBase: shared.DefaultServiceBase,
	IS:          shared_identityserver.DefaultIdentityServerConfig,
}

// InitConfig for the `init` command.
type InitConfig struct {
	conf.ServiceBase `name:",squash"`
	IS               identityserver.Config      `name:"is"`
	InitialData      identityserver.InitialData `name:"initial-data"`
}

// DefaultInitConfig contains the default config for the ttn-lw-identity-server binary.
var DefaultInitConfig = InitConfig{
	ServiceBase: shared.DefaultServiceBase,
	IS:          shared_identityserver.DefaultIdentityServerConfig,
	InitialData: shared_identityserver.DefaultIdentityServerInitialData,
}