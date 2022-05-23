import flask
from flask import request, jsonify, render_template, Flask
import os
import json
import sqlite3
from video_stuff import *
# i dont even know if this is needed because no sleep covid brain was struggling with cross origin shit
from flask_cors import CORS, cross_origin

CLIP_KEYS = ['video_id', 'start', 'end', 'upload_channel', 'title', 'description', 'vtuber']
CLIP_DOWNLOAD_KEYS = ['video_id', 'start', 'end']
current_dir = os.path.dirname(os.path.abspath(__file__))
app = flask.Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'
# i know this isn't the best way especially using sqlite but dont need anything fancier
db = sqlite3.connect(os.path.join(os.path.dirname(os.path.abspath(__file__)), 'clips.db'), check_same_thread=False)

app.config["DEBUG"] = True
# Flask auto sorts jsons for some reason, weird ass shit
app.config['JSON_SORT_KEYS'] = False

db.cursor().execute('CREATE TABLE IF NOT EXISTS clips (id INTEGER PRIMARY KEY, vtuber TEXT, clip_title TEXT, upload_channel TEXT, description TEXT, video_id TEXT, start_s INTEGER, end_s INTEGER, completed INTEGER)')
db.cursor().execute('CREATE TABLE IF NOT EXISTS timestamps (id INTEGER PRIMARY KEY, video_id TEXT, start_s INTEGER, end_s INTEGER, status TEXT)')
db.cursor().execute('CREATE TABLE IF NOT EXISTS videos (id INTEGER PRIMARY KEY, video_title TEXT, video_id TEXT, channel_id TEXT, channel_name TEXT, affiliation TEXT, parsed INTEGER)')

@app.route('/api/v1/manifest/<videoid>', methods=['GET'])
def _get_manifest_url(videoid):
    return jsonify(get_manifest_url(videoid))

@app.route('/api/v1/upload', methods=['POST'])
def _upload_clip():
    # TODO: take request json data with title, description, target channel, vtuber info etc and upload to youtube
    # surely a like 2 minute job copium
    return 'yea', 200

@app.route('/api/v1/download', methods=['POST'])
def _download_video():
    request_json = request.get_json(force=True)
    print(request_json)
    if not all(key in request_json.keys() for key in CLIP_DOWNLOAD_KEYS):
        return jsonify({'status': 'key error'})
    manifest_urls = get_manifest_url(request_json.get('video_id'))
    x = download_clip(manifest_urls[0], request_json)
    if x:
        return jsonify({'status':'success'})
    return jsonify({'status':'error'})

@app.route('/api/v1/search/<videoid>', methods=['GET'])
def _search_for_clips(videoid):
    # TODO: Pass search params through request & check for existing clips
    curr = db.cursor()
    # video_id, search_length=20, buffer=300, threshold=1.65, clip_count=-10
    # default search in 20s segments and skips the first and last 300s
    # any segment over average (message/search_lenghth)*threshold is detected as a clip
    # clip count negative because lazy and it gets the top x clips of the list sorted by 
    # average message/s speed to only get the "best" clips but mfs still randomly spam for no reason
    # this takes a LONG time to return depending on size of the chat, so maybe i should add a queue system?
    # idk my brain and throat hurt i am 死ぬing
    timestamps = get_stamps_by_vid(videoid)
    if isinstance(timestamps, dict):
        return jsonify(timestamps)
    for timestamp in timestamps:
        curr.execute('INSERT INTO timestamps (video_id, start_s, end_s, status) VALUES (?, ?, ?, ?)', (videoid, timestamp[1][0], timestamp[1][1], 'new'))
    db.commit()
    return jsonify(timestamps)

@app.route('/api/v1/clips', methods=['POST'])
def _add_clip_to_db():
    # TODO: everything lol
    # update timestamp to "used" or "deleted" to avoid showing up in GET requests
    # input clip data into clips table with clip title/description etc 
    return ({'status': 'error'})

@app.route('/api/v1/clips', methods=['GET'])
@cross_origin()
def _get_clip_from_db():
    
    # TODO: clean this shit up bro copy paste moment
    
    if not request.args:
        curr = db.cursor()
        curr.execute('SELECT * FROM timestamps WHERE status = "new"',)
        clips = curr.fetchall()
        # Reverse list so newest clips first
        clips = sorted(clips, reverse=True)
        clips_data = []
        for clip in clips:
            clips_data.append({
                'key': clip[0],
                'video_id': clip[1],
                'start_s': clip[2],
                'end_s': clip[3],
                'status': clip[4]
            })
        return jsonify(clips_data)
    
    if request.args.get('video_id'):
        print(request.args.get('video_id'))
        curr = db.cursor()
        curr.execute('SELECT * FROM timestamps WHERE ? = video_id AND status = "new"', (request.args.get('video_id'),))
        clips = sorted(curr.fetchall(), reverse=True)
        clips_data = []
        for clip in clips:
            clips_data.append({
                'key': clip[0],
                'video_id': clip[1],
                'start_s': clip[2],
                'end_s': clip[3],
                'status': clip[4]
            })
        return jsonify(clips_data)
    return {'status': 'error'}

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5050)