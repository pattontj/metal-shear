# Name

## Requirements

- Python 3.10+ (only tested on 3.10 at least)
- requirements.txt
- ffmpeg in your PATH (will change later on)

### Install

Run `python -m pip install -r requirements.txt`, preferably in a venv. 

Then run `python main.py` to start the API server.

## Using

### API

The server runs on port `:5050` with multiple routes.

| path | method | description |
| --- | --- | --- |
| /api/v1/manifest/\<videoid\> |  GET |Returns the manifest URL for downloading directly from YouTube servers. (this includes the host's IP address in the URL and I think it's required so big NO NO for public use) |
| /api/v1/search/\<videoid\>| GET |Downloads the chat if it hasn't already been downloaded and then finds peaks of chat activity. Adds the top 10 timestamps to the db. |
|/api/v1/clips |  GET | Returns list of all "new" timestamps by default. Optional `video_id` param for target video timestamps.  |
|/api/v1/clips |  POST | Updates / inserts a clips data (wip)  |
| /api/v1/download |  POST |Downloads the clip at atleast 720p using `ffmpeg` and saves it as `{video_id}_{start}_{end}.mp4`  |
| /api/v1/upload | POST | Upload clip to YouTube (wip)|

---

Yeah so look `video_stuff.py` and especially `get_stamps_by_vid()` are scuffed but if it works it works

If the video's chat hasn't been downloaded it checks to make sure the video has a duration and is type video. By default it searches in 20 second segments while skipping the first and last 5 minutes of the video to avoid こん／おつ message spam.