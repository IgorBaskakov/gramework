// Copyright 2017 Kirill Danshin and Gramework contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//

package gramework

import (
	"net"
)

// Serve app on given listener
func (app *App) Serve(ln net.Listener) error {
	if len(app.name) == 0 {
		app.name = "gramework/" + Version
	}

	var err error
	srv := app.copyServer()
	if err = srv.Serve(ln); err != nil {
		app.Logger.Errorf("ListenAndServe failed: %s", err)
	}

	return err
}
