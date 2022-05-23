import json
from pprint import pprint
from chat_downloader import ChatDownloader
from chat_downloader.sites.youtube import YouTubeChatDownloader
import youtube_dl
from datetime import timedelta
import os

def download_clip(manifest_url, clip):
    start_time = timedelta(seconds=int(clip['start']))
    end_time = timedelta(seconds=int(clip['end']))
    clip_name = f'{clip["video_id"]}_{clip["start"]}_{clip["end"]}'
    # REALLY gotta change this from using os.system but man i honestly cbf
    # for some reason it dies if i put -ss after -i which shouldnt happen??
    # so gotta do end_time - start_time or else it'll download a fuckin huge clip
    os.system(f'cmd /c ffmpeg -ss {start_time} -i "{manifest_url}" -t {end_time-start_time} -y -c copy "{clip_name}.mp4"')
    return True

# TODO: everything...
def upload_clip(video_data):
    return

def get_msg_count(elem):
    return elem[0]

def get_video_data(video_id):
    return ChatDownloader().create_session(YouTubeChatDownloader).get_video_data(video_id)

def parse_data(video_data):
    # i know this is kinda useless but it just makes it easier for my peanut brain
    parsed_data = {
    'video_title': video_data.get('title'),
    'channel_name' : video_data.get('author'),
    'channel_id' : video_data.get('author_id'),
    'video_id' : video_data.get('original_video_id'),
    'video_type' : video_data.get('video_type'),
    'start_time' : video_data.get('start_time'),
    'end_time' : video_data.get('end_time'),
    'duration' : video_data.get('duration')
    }
    return parsed_data

def get_stamps_by_vid(video_id, search_length=20, buffer=300, threshold=1.65, clip_count=-10):
    video_data = get_video_data(video_id)
    parsed_data = parse_data(video_data)
    
    if video_id + '.json' not in os.listdir('chat-files'):
        print('chat file not found')
        # just makes sure it's not a waiting room and is actually an archive
        if parsed_data.get('video_type') == 'video' and parsed_data.get('duration'):
            downloader = ChatDownloader()
            chat = downloader.get_chat(video_id, message_groups=['messages', 'superchat'], output=os.path.join('chat-files', video_id + '.json'))
            # need to do this so it actually saves the chat. generator moment 
            for msg in chat:
                pass
        else:
            return {'error': 'invalid video'}
    with open(os.path.join('chat-files', video_id + '.json'), 'r') as f:
        chat_data = json.load(f)
    
    start_time = buffer
    search_segment = start_time + search_length
    end_time = parsed_data.get('duration') - buffer
    # i know this should be done in a complete different way and use the buffer but man i really cbf
    msg_per_second = round(len(chat_data)/parsed_data.get('duration'), 2)
    msg_per_search =  msg_per_second*search_length

    msg_in_search = 0
    segment_streak = 0
    msg_in_streak = 0
    possible_clips = []
    
    for message in chat_data:
        chat_time = message.get('time_in_seconds')
        # check to skip chat message if its not between start/end 
        if chat_time < start_time or chat_time > end_time:
            continue
        # add message to counter if its in search segment
        if chat_time <= search_segment:
            msg_in_search += 1
            continue
        # this is where the scuffed shit starts
        # checks if timestamp of message is outside of segment
        if chat_time > search_segment:
            # if message count is above the average*threshold
            if msg_in_search > (msg_per_search*threshold):
                msg_in_streak += msg_in_search
                segment_streak += 1
            
            # adds timestamp to list when chat activity drops, leaves a bit of a buffer on both sides for context
            # a lot of ways to do this better but i am covid right now gomen ne
            if segment_streak and msg_in_search < (msg_per_search*threshold):
                clip_start = search_segment-((segment_streak+2)*search_length)
                clip_end = search_segment-(search_length/2)
                # append message rate to sort for top clips
                avg_msg = round(msg_in_streak/search_length, 2)
                possible_clips.append((avg_msg, (clip_start, clip_end)))
                segment_streak = 0
                msg_in_streak = 0
            
            # reset message count and change segement time
            msg_in_search = 1
            search_segment += search_length
    
    # return top timestamps
    return sorted(possible_clips, key=get_msg_count)[clip_count:]

def get_manifest_url(video_id):
    ydl_opts = {
        'forceurl': True,
        'skip_download': True
    }
    url = f'https://www.youtube.com/watch?v={video_id}'
    with youtube_dl.YoutubeDL(ydl_opts) as ydl:
        r = ydl.extract_info(url, download=False)
    urls = [f['url'] for f in r['formats'] if f['acodec'] != 'none' and f['vcodec'] != 'none' and f['width'] >= 1280]
    return urls