# A weapon to surpass metal gear

## TODO features:
 - [ ] Automated clip creation via ffmpeg
 - [ ] Auto-upload to Youtube
 - [ ] Discord Bot integration
 - [ ] Chat stats page (Statsuki)


## Far off features:
 - [ ] Automated hard-subbing via YT captions and ffmpeg
 - [ ] Combine clips into compilation videos
 - [ ] NND style chat overlay
 - [ ] View whole chat for a VOD


## How to Run (development):

### Server
To run the serverside code, run `go install` on all required library names (can be copied from the import statement w/o quotes).
Once all the required libraries are installed, run `go run main.go`. 

### Client
For frontend development, first run the server in a seperate terminal. Once the server is running, you can run `npm run start` in the 
frontend/client folder to spin up a react development server. 

The backend server runs on port `:8080`, while the React development server runs on port `:3000`. Developing the frontend using 
`localhost:3000` will allow you to use the built-in React development tools and get instant updates.

## How to build for Release:

## Server
Before running the server, set pathvars `DBUSER` and `DBPASS` accordingly. Once set, exectute `go build main.go`.

## Frontend 
Inside the `frontend/client` folder, run `npm run build`. This will compile the client into a single `index.html` file and webpack it.
The contents of build should be ready to serve over `:8080`. 
 

 ## Links
  - Uploading to YT Via Go, YT-API: https://github.com/youtube/api-samples/blob/master/go/upload_video.go