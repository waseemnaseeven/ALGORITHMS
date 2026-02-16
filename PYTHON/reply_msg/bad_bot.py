import requests 
import os
from instabot import Bot
import webbrowser
import subprocess
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By 
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time
import getpass

# LOGIN = input('login: ')
# PASSWD = getpass.getpass('ig_passwd: ')
#  fisapo2943 boolbi54
#  boubou2803 bulbizar2803

class BadBot():
    def __init__(self):
        self.login()
        self.message()

    def login(self):
        # get incognito browser to go to insta
        chrome_options = Options()
        chrome_options.add_argument("--incognito")
        self.driver = webdriver.Chrome(options=chrome_options)
        self.driver.get('https://instagram.com/')
        
        # Disabled cookies
        wait = WebDriverWait(self.driver, 10)
        disabled_cookies = wait.until(EC.presence_of_element_located((By.XPATH, '/html/body/div[2]/div/div/div[3]/div/div/div[1]/div/div[2]/div/div/div/div/div[2]/div/button[2]')))
        disabled_cookies.click()
        time.sleep(2)
        
        # Send username and password
        username_input = self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[1]/div/label/input')
        username_input.send_keys('lagrandephrase')
        time.sleep(2)
        passwd_input = self.driver.find_element(By.XPATH, '//*[@id="loginForm"]/div/div[2]/div/label/input')
        passwd_input.send_keys('Rowenta91@')
        time.sleep(2)
        
        # Click log in
        self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/section/main/article/div[2]/div[1]/div[2]/form/div/div[3]').click()
        time.sleep(5)
        
        # Click Not Now
        self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/div[1]/div[2]/section/main/div/div/div/div/div').click()
        time.sleep(3)
        
        # Click Not Now
        self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div/div[3]/div/div/div[1]/div/div[2]/div/div/div/div/div[2]/div/div/div[3]/button[2]').click()
        time.sleep(3)
        
        # Click sur message
        self.driver.find_element(By.XPATH, '/html/body/div[2]/div/div/div[2]/div/div/div/div[1]/div[1]/div[1]/div/div/div/div/div[2]/div[5]/div/div/div/span/div/a/div').click()
        time.sleep(5)

    def message(self):
        print("start messaging")
        getthrough = True
        while getthrough:
            counter = 0
            messages = self.driver.find_elements(By.CLASS_NAME, 'x1n2onr6')
            time.sleep(2)
            for message in messages:
                blue_dot = message.value_of_css_property("background-color")
                if blue_dot == "rgba(0, 149, 246, 1)":
                    counter += 1
            print(f"Unread messages : {counter}")
            time.sleep(5)



def main():
    my_bot = BadBot()

if __name__ == '__main__':
    main()
