# Go calendar
This is a simple go web application that display the current month, like a calendar would. This application has been written in the context of a technical challenge.
It display the current month for any GET requests and return a 501 otherwise.

## Build
```
docker build -t go-calendar .
docker tag go-calendar:latest maximerenou50/go-calendar:0.1.0
docker push maximerenou50/go-calendar:0.1.0
```
Image can be found [here](https://hub.docker.com/r/maximerenou50/go-calendar/dockerfile)

## Run
```
docker run -p 1080:1080 maximerenou50/go-calendar:0.1.0
```

## Test
GET request:
```
curl -i -XGET localhost:1080
HTTP/1.1 200 OK
Date: Sun, 03 Nov 2019 19:30:59 GMT
Content-Length: 358
Content-Type: text/plain; charset=utf-8

November 2019
----------
1 Friday
2 Saturday
3 Sunday (today)
4 Monday
5 Tuesday
6 Wednesday
7 Thursday
8 Friday
9 Saturday
10 Sunday
11 Monday
12 Tuesday
13 Wednesday
14 Thursday
15 Friday
16 Saturday
17 Sunday
18 Monday
19 Tuesday
20 Wednesday
21 Thursday
22 Friday
23 Saturday
24 Sunday
25 Monday
26 Tuesday
27 Wednesday
28 Thursday
29 Friday
30 Saturday
```

POST request:
```
curl -i -XPOST localhost:1080
HTTP/1.1 501 Not Implemented
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sun, 03 Nov 2019 19:34:18 GMT
Content-Length: 20

501 Not Implemented
```