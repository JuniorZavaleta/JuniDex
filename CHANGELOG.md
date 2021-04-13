# CHANGELOG

## UNRELEASED

## [0.4.0] XXXX-XX-XX
- Using project junidex-ui to handle all the UI

## [0.3.0] XXXX-XX-XX
> Added
- Load all pokemon available in json file in admin view
- Using project junidex-ui to handle UI for APP

## [0.2.2] 2021-04-04
> Fixed
- Use blank.gif instead of blank.jpg in `client/team.html`

## [0.2.1] 2021-04-03
> Fixed
- Saving and loading team using build for Windows - raise errors
- Index in template
- Admin was reseting team (array of zeros) - noticed when the client enters first (Now both load current team)

## [0.2.0] 2021-04-03
> Added
- Save Team in a json file locally
- Load team from json previously saved

> Changed
- Team Manager handle up to 6 members
- Change styles for Pokemon Team view
- Add more pokemon to json file

## [0.1.1] 2021-04-02
> Added
- Team Manager for Kanto's starters (3 members) using local JSON
- Connection between Admin and Client with WebSockets

> Fixed
- Show Starter was failing after changes for Team Manager
