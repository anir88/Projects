#Python code to print hello world

import datetime

print("Hello World")
print("I am a new Python program!! Simple but effective 😊")
print(f'The current time is: {datetime.datetime.now()}')

if datetime.datetime.now().hour < 12:
    print("Good Morning! Time to get a good exercise, healthy breakfast and kick start the day to win!!")
elif datetime.datetime.now().hour <= 17 & datetime.datetime.now().hour >= 13:
    print("Good Afternoon! Hope you are having a productive day. Don't forget to take short breaks and stay hydrated.")
else:
    print("Good Evening! Time to relax and unwind. Reflect on the day's achievements and plan for tomorrow.")
