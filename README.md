# go sample application for cloud foundry

## Usage

Just push the code:

	cf push

The manifest file must have GOPACKAGENAME pointed to the source.go file.  The Procfile must just have the binary output (same as the source for the standard buildpack)

## Diagnostics

You may run the app without the health check with:

	cf push -u none

and then enter in the Diego container with:

	cf ssh <app-name>
