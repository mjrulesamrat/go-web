import urllib2
from multiprocessing import Pool

# ruff example of how multiprocessing works in python
# find routines.go to check how this works in Golang
def f(url):
    req = urllib2.urlopen(url)
    try:
        print len(req.read())
    finally:
        req.close()

urls = (
    "http://www.peterbe.com",
    "http://peterbe.com",
    "http://htmltree.peterbe.com",
    "http://tflcameras.peterbe.com",
)

if __name__ == '__main__':
    p = Pool(3)
    p.map(f, urls)
