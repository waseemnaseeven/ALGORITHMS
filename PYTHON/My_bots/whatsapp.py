import pywhatkit
import keyboard
from datetime import datetime
import time

#One Contact
def one_Contact():
	phone_number = input("Enter the phone number, start with '+ something': ")
	if phone_number != "":
		pywhatkit.sendwhatmsg(phone_number, "Test2", datetime.now().hour, datetime.now().minute + 1, 15, True, 2)
		time.sleep(1)
	else:
		print("Please try again!")
		one_Contact()

#Group
def group_Contact():
	group_id = input("Enter the group id: ")
	if group_id != "":
		pywhatkit.sendwhatmsg_to_group(group_id, "test group", datetime.now().hour, datetime.now().minute + 1)


if __name__ == '__main__':
    one_Contact()
