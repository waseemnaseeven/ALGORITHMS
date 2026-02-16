import requests 
import os
from instabot import Bot
import webbrowser
import subprocess

ID = os.getenv("ID")
S_KEY = os.getenv("SECRET_KEY")
MY_WEBSITE = "https://waseemnaseeven.github.io/"
authorization_code = "AQCayEdPxiqMVL7yIQIVFFN3pQRJ3wuU_Wkwxh4w5211FHiayAuX5A14qgRkQ2sLs8dEd_7teW-ZR12_bHyqQzTXxc3ptftjwSSHfN2Th-faCy2SgUEbraQu5MH6CPjFEnK17DIRjyyU34kt_jPKRUF1KDq2JQECMVxjr5M0bn3eukvINMuSO-akIoqZFY1ovreOruF8OXAinN3ekYgg-8Tx8ndzgVuEWKvODMGcECb0lg"

# Ici, on va récolter access_token et user_id
def get_short_access_token():
    # CODE = os.getenv("code")
    CODE = authorization_code
    curl_command = (
        f"curl -X POST "
        f"https://api.instagram.com/oauth/access_token "
        f"-F client_id={ID} "
        f"-F client_secret={S_KEY} "
        f"-F grant_type=authorization_code "
        f"-F redirect_uri={MY_WEBSITE} "
        f"-F code={CODE}"
    )
    cmd_exe = subprocess.run(curl_command, shell=True, capture_output=True, text=True)
    print(cmd_exe.stdout)

def try_api():
    token = os.getenv("access_token")
    user_id = os.getenv("user_id")
    other_cmd = (
        f"curl -X GET "
        f"https://graph.instagram.com/{user_id}?fields=id,username&access_token={token} "
    )
    ask = subprocess.run(other_cmd, shell=True, capture_output=True, text=True)
    print(ask.stdout)

def get_long_access_token():
    cmd = (
        "curl -i -X GET \"https://graph.instagram.com/access_token"
        "?grant_type=ig_exchange_token"
        "&client_secret=e2f2fca6ae0c297499a0474edc671bd9"
        "&access_token=IGQWROd0piOWhmUTZAseDZAvOUxUYXJtSTV1cjhYTDZA3S3ZANZAGxhQkhSNkd1dGtqVlNLVUp4eXNpM0ZAGWVNZATS1CZA2o5SmFVX3J2ZA1ZAOZAnVEYnYzVGt1bkpOa0xtb0Y2WXNQWDZAqbzlDTkRRdEpfaUVRVlVQSy1fNmRQQU9zVlRTa1gwUQZDZD\""
    )
    long = subprocess.run(cmd, shell=True, capture_output=True, text=True)
    print(long.stdout)


def process_msgs():
  print("Start the bot")
  authorize_url = f"https://api.instagram.com/oauth/authorize?client_id={ID}&redirect_uri={MY_WEBSITE}&scope=user_profile,user_media&response_type=code"
  # Effectuez une requête GET à l'URL d'autorisation
  response = requests.get(authorize_url)
  # webbrowser.open(authorize_url)
  if response.status_code == 200:
    print("Success")
    # get_short_access_token()
    get_long_access_token()
    # try_api()
  else:
    print("Game Over")

if __name__ == '__main__':
    process_msgs()