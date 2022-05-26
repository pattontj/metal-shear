import sys
import json
import requests


import youtube_dl
import chat_downloader
from chat_downloader import ChatDownloader
from chat_downloader.sites.youtube import YouTubeChatDownloader





def main():
    downloader = ChatDownloader()
    session = downloader.create_session(YouTubeChatDownloader)

    channel_id = sys.argv[1]
    DB_vtuber_id = sys.argv[2]

    # TODO: pass channel ID from cmd
    videos = session.get_user_videos( channel_id, video_status='past' )


    try:
      
        for i, video in enumerate(videos):
            # skips any live streams if video_status=all
            if 'watching' in video.get('view_count'):
                continue
            
            
            video_data = session.get_video_data(video['video_id'])

            parsed_data = {
                'link' : video_data.get('original_video_id'),
                'tsBegin' : '0',
                'tsEnd' : '10',
                'vtuberID': str(DB_vtuber_id)
            }

            r = requests.post( 
                url =   "http://localhost:8080/api/clips/post", 
                headers = {'Content-type': 'application/json', 'Accept': 'text/plain'},
                json = parsed_data
            )

            print(r)

            if i == 1:
                break

            # gets video data here
            # video_data = session.get_video_data(video['video_id'])
            # print( video.get('title') )
            
    except chat_downloader.errors.VideoNotFound:
        print(' can\'t find channel ')
    except chat_downloader.errors.NoVideos:
        print('has no videos.')
    # except Exception as e:
    #     print('a')


    # except chat_downloader.errors.VideoNotFound:
    #     print(f'{slug}: can\'t find channel {channel_id}')
    # except chat_downloader.errors.NoVideos:
    #     print(f'{slug} has no videos.')
    # except Exception as e:
    #     print(f'{slug}: {e}')


    return 0



def parse_data(video_data):
    parsed_data = {
    'video_id' : video_data.get('original_video_id'),
    'start_time' : video_data.get('start_time'),
    'end_time' : video_data.get('end_time'),
    }
    return parsed_data



if __name__ == "__main__":
    main()