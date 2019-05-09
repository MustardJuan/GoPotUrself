# GoPotUrself
Dockered Golang Honeypot

Inspiration derived from the popular Cowrie honeypot

## Basic Environment Needed
1. An OS that has Docker present on the system, the dockerfile itelf will run everything needed in order to not only configure the site and golang code but also to get the server up in the container. 

### Running the container

Basic commands follow with brief explanantions to getting the container running, assumption is that the user is at the directory where the dockerfile and relevant code is present

To compile the docker container:   
```` docker build <name of container> . ````

To run the container:   
```` docker run -p 80:80 -t <name of container> ````

From there you can open a browser and connect to localhost to see the site on the machine your hosting it on!
Important to note that presentlt to get the full exploit running when the php reverse shell is uploaded you need to change the IP to the one you want the reverse shell to connect to and the port to 8080. We have all ports closed except for 80 and 8080 in this current uild for security reasons but will be removing this safeguard. 

