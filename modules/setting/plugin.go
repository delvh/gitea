// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package setting

import "code.gitea.io/gitea/modules/log"

// Project settings
var (
	Plugin = struct {
		BaseDirectory string // The base directory where all plugins must be present
		PluginGlobs   []string // The globs that plugins can match inside the base directory to be loaded
	}{}
)

func newPlugin() {
	if err := Cfg.Section("plugin").MapTo(&Plugin); err != nil {
		log.Fatal("Failed to map Plugin settings: %v", err)
	}
}
