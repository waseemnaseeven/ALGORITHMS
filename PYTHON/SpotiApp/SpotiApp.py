import json
import os
import requests
from flask import Flask, request, redirect, session, url_for
from dotenv import load_dotenv
import base64
import urllib.parse

load_dotenv()

app = Flask(__name__)
app.secret_key = os.urandom(16)

class SpotiApp():
    # attributs
    def __init__(self) -> None:
        self.token_url = "https://accounts.spotify.com/api/token"
        self.profile_url = "https://api.spotify.com/v1/me"
        self.search_endpoint = "https://api.spotify.com/v1/search"
        self.client_id = os.getenv("CLIENT_ID")
        self.client_secret = os.getenv("CLIENT_SECRET")
        self.auth_url = "https://accounts.spotify.com/authorize"
        self.profile_url = "https://api.spotify.com/v1/me"
        self.scope = "user-read-private user-read-email"
        self.access_token = None
        self.token_type = None
        self.token_expires_in = None
        pass

    # methods
    def authenticate(self):
        auth_headers = {
            "Content-Type" : "application/x-www-form-urlencoded"
        }
        auth_data = {
            "grant_type": "client_credentials",
            "client_id": self.client_id,
            "client_secret": self.client_secret
        }

        response = requests.post(self.token_url, headers=auth_headers, data=auth_data)
        if response.status_code == 200:
            response_data = response.json()
            self.access_token = response_data['access_token']
            self.token_type = response_data['token_type']
            self.token_expires_in = response_data['expires_in']
            print(self)
        else:
            print(f"Something went wrong during access_token, \nstatus_code: {req.status_code}, {req.text}")
            exit()

    #TODO: doesnt work...
    def get_profile(self):
        if self.access_token is None:
            print("You need to authenticate first!")
            return

        profile_headers = {
            "Authorization": f"Bearer {self.access_token}"
        }

        response = requests.get(self.profile_url, headers=profile_headers)
        if response.status_code == 200:
            profile_data = response.json()
            print(json.dumps(profile_data, indent=4))
        else:
            print(f"Failed to get profile data, status_code: {response.status_code}, {response.text}")

    def __str__(self):
        return (
            f"SpotiApp(client_id={self.client_id},\n "
            f"client_secret={self.client_secret},\n "
            f"access_token={self.access_token},\n "
            f"token_type={self.token_type},\n "
            f"token_expires_in={self.token_expires_in})"
        )
