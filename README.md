# JuniDex

This application wants to give the ability to Pokemon streamers to have a local server that could be used by OBS Studio

## Features planned for v1.0

- Update Pokemon Team from a admin side and update it in OBS with one click or two
 - Evolves a Pokemon
 - Drop a Pokemon in box
 - Change positions
- Save your team locally (Done)
 - Here will be two implementations, one for me (for fun using graphs - neo4j) and the other with json for users
- Show possible counters for the current team

## Installation for users
- Download this repository (as zip or clone it) - Unzip if needed (for example in D:\pokemon\junidex)
- Download executable from `Releases`
- Set environment variable (following that you copy the repository in D:\pokemon\junidex)
    - for Windows users run the following commands in CMD or use Control Panel
        ```
        setx DIR_TEMPLATE "D:\pokemon\junidex\templates"
        setx STATIC_FOLDER "D:\pokemon\junidex\static"
        setx DATA_FOLDER "D:\pokemon\junidex\data"
        ```
- Read about sprites to fulfill your static folder, after that you can start the server executing `junidex.exe`

## Installation for developer (More details in next update)
- Install Go
- Download the project
- Install packages (gorilla/mux and gorilla/websocket)
- Set environment variable - that way that you want, idk which IDE, OS you uses xd
- Run localhost

## Integration with OBS
- After run localhost server, open a tab in your browser with `localhost:8080/admin/team`
- Add a browser source in OBS with the url `localhost:8080/client/team`
- Update your team and you'll see your team changing in your streaming with the sprites that you've added

## About sprites
Sprites could be downloaded from your prefered source or to use a script i've made [here](https://github.com/JuniorZavaleta/junidex-scrapper). The images that you want to use MUST be in the path `/static/img/`.

## Ways to contribute
- Create a PR with code changes, grammatical expressions (i'm a non-native English speaker) or adding/updating data in json files
- Send me a message with your feedback or any comments on my [Telegram](https://t.me/JuniorZavaleta)
- [![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y338J3F)
