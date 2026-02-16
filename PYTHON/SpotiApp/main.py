from SpotiApp import SpotiApp
import webbrowser
import sys
from urllib.parse import urlparse, parse_qs

#TODO: rebuild SpotiApp.py to get the user data

def get_code_from_url(url):
    """ Extract authorization code from redirect URL """
    parsed_url = urlparse(url)
    return parse_qs(parsed_url.query).get("code", [None])[0]

if __name__ == '__main__':
    app = SpotiApp()
    app.authenticate()
    app.get_profile()