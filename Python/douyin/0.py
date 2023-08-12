import re

import requests

url = ("https://www.douyin.com/user/MS4wLjABAAAA9fF5EnpyORa0KYx5CkZG5gIPVxeYgc1YcTyk4CSrL5o?modal_id"
       "=7237357898685336867&vid=71")
headers = {
    "cookie": "douyin.com; webcast_local_quality=null; "
              "ttwid=1%7CHMWmBGjEnZUvVHYDuFmcHVBwj0_Nq1ueFyLbrwH1J1Q%7C1691297703"
              "%7C46262630734ad4f1166731c98a96beaa4d49c330c67159d1f18e8e6af06e8fc9; "
              "strategyABtestKey=%221691297725.169%22; passport_csrf_token=26c126b8113d9bebf90d5094831cc4a3; "
              "passport_csrf_token_default=26c126b8113d9bebf90d5094831cc4a3; "
              "__bd_ticket_guard_local_probe=1691297726163; "
              "s_v_web_id=verify_lkyyzxn9_FxlRSBUk_5sDV_4Sxr_AY8p_YfoHRa1IYplv; "
              "volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Afalse%2C%22volume%22%3A0.5%7D; "
              "bd_ticket_guard_client_data"
              "=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQ"
              "tZ3VhcmQtY2xpZW50LWNzciI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tXHJcbk1JSUJEakNCdFFJQkFEQW5NUXN3"
              "Q1FZRFZRUUdFd0pEVGpFWU1CWUdBMVVFQXd3UFltUmZkR2xqYTJWMFgyZDFcclxuWVhKa01Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqM"
              "ERBUWNEUWdBRUViVC8yMG1TRWlmVTBPR1kyd1dsbE41RlxyXG5WZ2xSeitUQndKYldudmk1SmNZNklVT0ExVHhCVmhhdEUzNXR0cTlNc"
              "GFrb01IYWsyNkNRUXBoNUpVZDlFNkFzXHJcbk1Db0dDU3FHU0liM0RRRUpEakVkTUJzd0dRWURWUjBSQkJJd0VJSU9kM2QzTG1SdmRYbH"
              "BiaTVqYjIwd0NnWUlcclxuS29aSXpqMEVBd0lEU0FBd1JRSWdVQkc3aGZNbmJESjhqMnFkVU4xbHZucUM1eEVvdmZrTi9DdUZXcXpVTHp"
              "zQ1xyXG5JUUR4dHI0alRpY3dEd3d4TE5SODJaUXJiakg2SFdldjIxNWROcWR5SnZRSXd3PT1cclxuLS0tLS1FTkQgQ0VSVElGSUNBVEU"
              "gUkVRVUVTVC0tLS0tXHJcbiJ9; ttcid=6424f2f5c05f4701b987663eeb6bf60f96; FORCE_LOGIN=%7B%22videoConsumedRema"
              "inSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; passport_assist_user=Cj1BV1XKJPPY7LvwpcjuHdX9Mj62aNCN"
              "cNcu2nleYt7K-95P2zIkiRYG3-YMR1t9a3mmRQyoE068B0EQK_BUGkgKPLx2WOQ3fFpwXFZLSGlk3NgYjR6bwRav1xMbYtZS-8xXyVhU"
              "pJB-8bmEEFwl1LDZyt1fhQ6gJtx7tiu-XRCjvLgNGImv1lQiAQMq0Qqn; n_mh=LdeZtk6a5o9cP5ln0p59Y1HdRdPgy-V0M3hRZwNxy"
              "Nk; sso_uid_tt=ae114840233269f9a22e82bbd1531360; sso_uid_tt_ss=ae114840233269f9a22e82bbd1531360; toutiao"
              "_sso_user=fb347a9c6cd6bc0357362ed143da9c83; toutiao_sso_user_ss=fb347a9c6cd6bc0357362ed143da9c83; sid_u"
              "cp_sso_v1=1.0.0-KDhkODU0ZGM3YTY1ZGViMDM0Y2MyYjJlNGRhNGRhOTY2ZDU0ZTA3OTMKHQiXiZTXkwMQztG8pgYY7zEgDDDvhqfg"
              "BTgGQPQHGgJscSIgZmIzNDdhOWM2Y2Q2YmMwMzU3MzYyZWQxNDNkYTljODM; ssid_ucp_sso_v1=1.0.0-KDhkODU0ZGM3YTY1ZGViM"
              "DM0Y2MyYjJlNGRhNGRhOTY2ZDU0ZTA3OTMKHQiXiZTXkwMQztG8pgYY7zEgDDDvhqfgBTgGQPQHGgJscSIgZmIzNDdhOWM2Y2Q2YmMwM"
              "zU3MzYyZWQxNDNkYTljODM; passport_auth_status=6a9d61bb28414bf2006bab011ef58d73%2C; passport_auth_status_s"
              "s=6a9d61bb28414bf2006bab011ef58d73%2C; uid_tt=35c53472629e3e4bc8580db82751461d; uid_tt_ss=35c53472629e3e"
              "4bc8580db82751461d; sid_tt=ec7fd941003f228cbacfe7105f51291d; sessionid=ec7fd941003f228cbacfe7105f51291d;"
              " sessionid_ss=ec7fd941003f228cbacfe7105f51291d; publish_badge_show_info=%220%2C0%2C0%2C1691298009097%22; "
              "LOGIN_STATUS=1; __security_server_data_status=1; store-region=cn-sn; store-region-src=uid; d_ticket=c69f8"
              "a76dd42e7f4e02c6fa8a567254f46fbe; sid_guard=ec7fd941003f228cbacfe7105f51291d%7C1691298027%7C5183975%7CTh"
              "u%2C+05-Oct-2023+05%3A00%3A02+GMT; sid_ucp_v1=1.0.0-KGRjZjM4MWQxNDQyNTQwYjQyY2M4N2NlOTc5ZTllYzQ3ZmFhYTFl"
              "YTcKGQiXiZTXkwMQ69G8pgYY7zEgDDgGQPQHSAQaAmxxIiBlYzdmZDk0MTAwM2YyMjhjYmFjZmU3MTA1ZjUxMjkxZA; ssid_ucp_v1="
              "1.0.0-KGRjZjM4MWQxNDQyNTQwYjQyY2M4N2NlOTc5ZTllYzQ3ZmFhYTFlYTcKGQiXiZTXkwMQ69G8pgYY7zEgDDgGQPQHSAQaAmxxIiBlYzdmZDk0MTAwM2YyMjhjYmFjZmU3MTA1ZjUxMjkxZA; SEARCH_RESULT_LIST_TYPE=%22single%22; _bd_ticket_crypt_cookie=75a0f8491404c5b49bde7883120ce348; pwa2=%220%7C0%7C3%7C0%22; download_guide=%223%2F20230806%2F0%22; douyin.com; device_web_cpu_core=8; device_web_memory_size=8; webcast_local_quality=null; VIDEO_FILTER_MEMO_SELECT=%7B%22expireTime%22%3A1691908811236%2C%22type%22%3Anull%7D; csrf_session_id=5e0382dda0a4d6e33e8781f49f324a99; __ac_nonce=064cf49c600e64aa716d0; __ac_signature=_02B4Z6wo00f01QXeDzAAAIDDQb0QMyFwEF0F.guAACW7x5DauHdiSmtihyQvnZ0NFpzibLEA-HfD3Al4kfub1UTrpHvLoMTGr-74K4IFX62hZga.T4Ag33f9v2Tc5hiaAOXnHgX3honHvb046b; odin_tt=747537619d4aee8fdc7302a48d243cd9a9e97b042dd9b377e9cef66c6a4a318aa6d00e71637ca9ff37a2afd82d982b41; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A2048%2C%5C%22screen_height%5C%22%3A1152%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A8%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A0%7D%22; home_can_add_dy_2_desktop=%221%22; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAunQb9TiNmNFMg0DAtkQW01TVs5Qh20LFP9ZXk0GBl9E%2F1691337600000%2F0%2F0%2F1691307215498%22; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAunQb9TiNmNFMg0DAtkQW01TVs5Qh20LFP9ZXk0GBl9E%2F1691337600000%2F0%2F0%2F1691307815498%22; msToken=p9UM3q1onhjUfEsUMyO_NcOhgjQ_LqXQuzDMJrJP957sixBToFUyQ-jlakiYabs1fghpsgjS12EFvVy1PF6FXlpyCyCjhXjxEUwO33zjdsPE-FoOnw==; msToken=_YDAWNt0_xZlEaRx2Cs2AttAKV-xuxaF3G00tJRRr8VS2uYNzJLl5RJkmrd3x8-Ck32qpxvR30J9gUe4XHEGSCXJRL1NYq66-I48NtZ610Np0PtVrn3z_L2hljA=; tt_scid=qmcyHogaSCAeh9yAdoWFntE4h.w63SyOdYsCsZTKglBjHbXkFitp2PIrKPtAlEHQa645; passport_fe_beating_status=false",
    "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 "
                  "Safari/537.36"
}

response = requests.get(url=url, headers=headers)

title = re.findall('<meta data-rh="true" name="lark:url:video_title" content="(.*?)/><meta data-rh="true"',
                   response.text)[0]
html_data = re.findall('<meta data-rh="true" name="lark:url:video_iframe_url" content="(.*?)"/><meta data-rh="true"',
                       response.text)[0]

print(title)
print(html_data)
# video = requests.get(html_data, headers=headers).content
# with open('./' + '1' + '.mp4', mode='wb') as f:
#     f.write(video)
# print(video)
