# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

i = 0

def adder():
# In Python you "import" a global variable, instead of "export"ing it when you declare it
    	global i
    
# In Python 2 this generates a list of integers (which takes time),
# while in Python 3 this is an iterable (which is much faster to generate).
# In python 2, an iterable is created with xrange()
    	for x in range(0, 1000000):
		Lock.aqcuire(False)
        	i += 1
		Lock.aqcuire(True)
def sub():#
	global i
    	for x in range(0, 1000000):
		Lock.acquire(False)
		i-=1
		Lock.aqcuire(True)
def main():
    	adder_thr = Thread(target = adder)

	sub_thr = Thread(target = sub)#

    	adder_thr.start()

	sub_thr.start()#

    	for x in range(0, 50):
        	print(i)

   	adder_thr.join()

	sub_thr.join() #

    	print("Done: " + str(i))

main()
