package main

const usage = `
Usage:
        clientpost <command>
The commands are:
		seek		Sends a seek request to the SDNC, specidied by the work order json file.
		teardown	Sends a teardown request to the SDNC, which terminates the collider environments.
		genjson		Outputs the json formatting for a work order as implemented inside the application.
`