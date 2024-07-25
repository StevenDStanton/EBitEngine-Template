# EbitEngineTemplate

This is a quick start EBitEngine Template for my projects. It is designed to get your game built and deployed for HTML5 to itch.io

I am running this on Pop!\_OS now Windows. If you are using Windows this may not work for you.

## Get Started

- [ ] `git clone git@github.com:StevenDStanton/EBitEngine-Template.git gameName`
- [ ] `cd gameName`
- [ ] `rm -rf .git`
- [ ] `git init`
- [ ] Finish your github repo setup and configure it here
- [ ] open go.mod and set your module name
  - Recomended: github.com/{profileName}/{gameName}
- [ ] `go mod tidy`
- [ ] [Install Butler](https://itch.io/docs/butler/installing.html)
- [ ] `butler login`
- [ ] update Makefile line 25 to match your module name above
- [ ] update Makefile line 35 with itch.io profile and game name and version
- [ ] update Makefile line 45 with itch.io profile and gane name

Reminder: Update version for each deploy

## Commands

- `make`: build the whole project and deploy
- `make build-test` build but do not deploy
- `make status` check deploy status
