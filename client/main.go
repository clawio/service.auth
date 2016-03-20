package main

import (
	"github.com/clawio/service-auth/client/commands"
	"github.com/codegangsta/cli"
	"os"
)

var VERSION string

//TODO(labkode) Add flag to produce json output
func main() {

	app := cli.NewApp()
	app.Version = VERSION
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Hugo González Labrador",
			Email: "contact@hugo.labkode.com",
		},
	}
	app.Copyright = `
	ClawIO Service Auth Client CLI
	Copyright (C) 2016  Hugo González Labrador

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
	Universidade de Vigo owns the copyright of this tool which is supplied
	in confidence and which shall not be used for any purpose other than that
	for which it is supplied and shall not in whole or in part be reproduced,
	copied or communicated to any person without permission from the owner.
	`

	app.Name = "ClawIO Service Auth CLI"
	app.Usage = `
		
	ClawIO Service Auth Client CLI is an end-user tool to make
	calls to a ClawIO Service Auth 
	`

	app.Commands = []cli.Command{
		commands.AuthenticateCommand,
		commands.VerifyCommand,
	}

	app.Run(os.Args)
}