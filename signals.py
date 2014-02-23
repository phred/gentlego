
import signal, sys

def handler(signum, frame):
    print 'Got signal:', signum
    sys.exit()

signal.signal(signal.SIGQUIT, handler)
signal.signal(signal.SIGINT, handler)

while True:
    pass
