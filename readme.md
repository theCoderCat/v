# In this repo

This repo contains:

- Source code for 2 coding problems
- Compiled program to demontrate problems and solutions (for MacOS and Windows)
- Architecture design for to do app
- Explainations and answers for other questions
- Benchmark results to compare 2 solutions for merging problem

## How to use compiled program:
- Please download compiled program from here: https://github.com/theCoderCat/v/releases/tag/1.0
- Run program from terminal/cmd/shell... according to your platform
  - `v_linux` on Linux
  - `v_macos` on MacOS
  - `v.exe` on Windows
- Pick an option when prompted

## 1. Coding problems

### Palindrome validator

Please take a look at `palindrome.go` for the code.  
Assumed the question should be "ignoring case, spaces and special characters".  
I also assumed that we only need to validate strings written in alphanumeric writting system.

The solution is quite simple:
- take a string
- keep only alphanumeric characters, strip all others
- make all the characters lower case (UPPER CASE also okay)
- loop. compare first and last characters, then second and second last, and so on until we reach the middle of the string...
- return as soon as the compare operator yield `false`

Let's say the given string has n characters.  

- Time complexity: O(n/2)
- Space complexity: O(n)

### Merge 2 arrays in sorted order

My chosen solution is ***sort first, then merge***
Please take a look at `sortedmerge.go` for the code.

This solution will sort 2 smaller arrays before merging.

Let's say sA is the size of the first array and sB is the size of second array

- Time complexity: O(sA log sA + sB log sB + sA + sB)
- Space complexity: O(sA + sB)

> Another solution is ***merge 2 arrays first then sort*** the big array
> - Time complexity: 0((sA+sB) log (sA+sB))
> - Space complexity: O(sA + sB)

Take a look at the time complexity (I had to use a calculator actually), we can see that the first solution (sort then merge) is faster than the second solution (merge then sort).

But how much faster? I was curious, so I decided to implement the second solution (`mergedsort.go`) and run a benchmark with 2 very big arrays (randomly generated, size from 10 millions to 100 millions integers), do it 10 times each run, each time with a different set of data. Please take a look at `benchmark_....txt` files for the results. I ran it on 2 different machines, Mac OS and Windows.

Sorted then merge solution run faster most of the test. But only slightly faster.

## 2. Websocket

WebSocket is a computer communications protocol, providing full-duplex communication channels over a single TCP connection.

The WebSocket protocol enables interaction between a web browser (or other client application) and a web server with lower overhead than half-duplex alternatives such as HTTP polling, facilitating real-time data transfer from and to the server. This is made possible by providing a standardized way for the server to send content to the client without being first requested by the client, and allowing messages to be passed back and forth while keeping the connection open. In this way, a two-way ongoing conversation can take place between the client and the server. The communications are done over TCP port number 80 (or 443 in the case of TLS-encrypted connections)

Before WebSocket, port 80 full-duplex communication was attainable using Comet channels; however, Comet implementation is nontrivial, and due to the TCP handshake and HTTP header overhead, it is inefficient for small messages.

## 3. Hey! Your website is slow!

### Where did you access from?

  -- *Europe. Why?*  
  -- Our web servers are located in Japan and internet connection between Europe and Japan is not very good. May be that's why. We will try to serve content from servers in Europe.

### Which part of the website is slow or it's slow across all parts of the website?

  -- *The homepage where you introduce the company with some text and image was fine, I could load it in matter of seconds. But tours listing with 50 tours, 4 or 5 customization options each, all showed at the same time and tour detail pages which display every possibly related tours presented in the system, they took forever to load.*  
  -- The homepage only contain some static information and very little dynamic content, while tours listing and and tour detail pages have complex data and required complicated operations to fetch all variations, related tours, status, price calculations... and they need to be fetched from database, recalculated again every time you access the page because we didnt have any methods to cache these information We'll try to cache the content when there're changes happen to the database so they're ready to be served next time you access.

### Is it always slow or only slow at certain time?

  -- *It's usually slow at around 10am~16pm, everything are quite fast at 1am.*  
  -- Might be because there're too many traffics happen during a certain time (day time, when people try to find a tour all at once), this happen because load balancing is not very effective or doesnt work at all and overload the server. We'll try to improve/implement load balancing solution.

### What kind of internet connection you're using?

  -- *Dial-up connection with 64kbps bandwidth. I've been using it for 20 years now.*  
  -- Tch! Tough luck.

## 4. Todo web app

### Architect diagram

*Diagram goes here*
![To do web app diagram](https://github.com/theCoderCat/v/raw/master/todo.png "To do web app diagram")

### How the system works

- A web application
- Front-end with VueJS
- Back-end with Go
- Data storage is MySQL

- User go to the website
- User adds tasks
- User marks tasks as finished
- User views the content of the task
- User deletes task when he doesn't feel like finishing it
- User's amazed because the page doesn't try to reload every time he click on something

### Things to consider

- Can user share their tasks with others?
- Does task have deadline?
- How many people will use this system?
- Do I need a mobile app later?

### Limitation

- Cannot remind user when a task due in tomorrow
- 

## Source of information

https://en.wikipedia.org/wiki/Palindrome  
https://golang.org/pkg/sort/  
https://en.wikipedia.org/wiki/WebSocket  
