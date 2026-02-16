from datetime import datetime
from bs4 import BeautifulSoup
import sys
import requests
import csv

def get_time():
	now = datetime.now()
	current_time = now.strftime("%H:%M:%S")
	print("Current Time =", current_time)

def extract_data(data):
	ret = []
	for elm in data:
		ret.append(elm.string)
	return ret

def fullfil_data(file_name):
	with open(file_name, 'w') as csv_file:
		writer = csv.writer(csv_file, delimiter=',')
		#writer.writerow(["", ""])

if __name__ == '__main__':
    fullfil_data("data.csv")
