# Self-Host Jitsi Infrastructure Setup and Configuration

# Introduction

In this guide, I will walk you through my entire setup and configuration of my Jitsi infrastructure. Jitsi is an open-source video conferencing platform that allows you to host your own meetings and webinars.

## System Requirements

Before starting, makke sure your server meets the following minimum requirements:
-  Operating System: Ubuntu 20.04 LTS
-  Memory: 4GB RAM or higher
-  CPU: Dual-core processor or better
-  Public Domain or IP address for your server

# Tech Stack

The following technologies were used in this implementation. It will be broken down based on the following criteria; On premises resources, cloud resources, 3rd party vendors
The goal was to make it as inexpensive as possible using free tier offerings where possible.

![Jitsi_infrastructure_diagram](https://github.com/mark76ped/Infrastructure/assets/52715459/46eece5a-d8af-4a29-bb68-792cb120d32c)

  ### On Premises Resources:
  -  Ubuntu 22.04.3 LTS
  -  Docker 20.10.24
      -  jitsi container jicofo:stable-8960-1
      -  jitsi container jvb:stable-8960-1
      -  jitsi container web:stable-8960-1
      -  jitsi container prosody:stable-8960-1
  -  Wireguard
  ### OS Cloud Resources:
  -  Ubuntu 22.04.3 LTS
  -  Docker 20.10.24
      -  nginx-proxy-manager
  -  Wireguard
  ### DNS/WAF Cloud Resources:
  -  Cloudflare Free Tier
  -  Registration of a domain (cost ~€15/yr)


# Configuration Guide

## Establishing a link with Wireguard
Assuming two servers are step up and ready to configure, we need to link them together. One way to accomplish this is to utilize Wireguard. 
I chose the Oracle Cloud Ubuntu Server as the Wireguard Server.
###Oracle Cloud Ubuntu Server:
Switch to root and download this script to make the install super simple. It's great and all the credit goes to its author.
```python
sudo su -
curl -O https://raw.githubusercontent.com/angristan/wireguard-install/master/wireguard-install.sh
```
Change the permissions on the script file to allow it to become executable:
```python
#chmod is the command to change permission on a file within Linux
chmod +x wireguard-install.sh
```
Execute the script:
```python
# using the './' executes a shell script file within Linux.
./wireguard-install.sh
```
After the script has been executed it will ask you for input variables
> IPV4 or IPv6 public addess: 0.0.0.0

Search within the dashboard of the Oracle Cloud instance created and fill in the prompt with the same IP address info
![Screenshot 2023-11-10 at 15 50 03](https://github.com/mark76ped/Infrastructure/assets/52715459/3d1ed478-4a3b-4415-99e2-b298e6d79e8a)

Optional: Leave the rest of the script to the default input.
Make note of the "Server WireGuard port [1-65535]:" as you'll need it for your ACL's and configuration.
Open the listed port on the Oracle Cloud server.

Then follow the instructions in the script and save the information for the client. Copy the .conf and place it on the WireGuard client at /etc/wireguard/wg0.conf
```
sudo cat /root/wg0-client-<name of server>.conf
```
Confirm the WireGuard connection is successful
```
$sudo wg show
interface: wg0
  public key: <public key info appears here>
  private key: (hidden)
  listening port: <port must match ACL>

peer: <peer info appears here>
  preshared key: (hidden)
  endpoint: <on-prem IP appears here>:<random port>
  allowed ips: 10.66.66.2/32
  latest handshake: 1 minute, 3 seconds ago
  transfer: 475.29 MiB received, 1.05 GiB sent
```
If successful you should see a handshake value, as well as data received and sent.
You can also confirm with pinging between hosts.

Tip: 
-  If WireGuard goes down, first try a `sudo wg-quick down`, `sudo wg-quick up` on the client side. Restarting the service resets any iptable/ufw change that may be preventing the connection from establishing.

## Jitsi Docker Containers 
On the on premises Ubuntu Server 
  -  Download the latest zip files:
```
sudo wget https://github.com/jitsi/docker-jitsi-meet/archive/refs/tags/stable-8960-1.tar.gz
```
-  Extract the tar file:
```
tar -zxvf stable-8960-1.tar.gz
```
-  cd into the newly created folder
```
cd docker-jitsi-meet-stable-8960-1/
```
-  Follow steps 2-5 from the intructions here: https://jitsi.github.io/handbook/docs/devops-guide/devops-guide-docker/

Confirm 4 containers are up and running with the following command
```
ubuntu_user@my_ubuntu_srv:~$ docker ps
CONTAINER ID   IMAGE                         COMMAND      CREATED       STATUS                 PORTS                                                                                                                 NAMES
6c3ff054a59e   jitsi/jicofo:stable-8960-1    "/init"      13 days ago   Up 12 days             127.0.0.1:8888->8888/tcp                                                                                              docker-jitsi-meet-stable-8960-1-jicofo-1
41596af4873d   jitsi/jvb:stable-8960-1       "/init"      13 days ago   Up 12 days             127.0.0.1:8080->8080/tcp, 0.0.0.0:10000->10000/udp, :::10000->10000/udp                                               docker-jitsi-meet-stable-8960-1-jvb-1
20dedcec8669   jitsi/web:stable-8960-1       "/init"      13 days ago   Up 12 days             0.0.0.0:8443->443/tcp, :::8443->443/tcp                                                                               docker-jitsi-meet-stable-8960-1-web-1
3f934254f826   jitsi/prosody:stable-8960-1   "/init"      13 days ago   Up 12 days             5222/tcp, 5280/tcp, 5347/tcp                                                                                          docker-jitsi-meet-stable-8960-1-prosody-1
```

Tips:
-  Add your user to the docker group to run docker without sudo
-  Take note of the port 8443 for container web:stable-8960-1 for later on when we need the nginx-proxy-manager to direct web traffic to this location

## Oracle Cloud Ubuntu Configuration
For the in the Oracle Cloud, the following ports must be forwarded to the on prem Ubuntu server:
```
sudo iptables -t nat -A PREROUTING -p udp --dport 4443 -j DNAT --to 10.66.66.2:4443
sudo iptables -t nat -A PREROUTING -p udp --dport 8080 -j DNAT --to 10.66.66.2:8080
sudo iptables -t nat -A PREROUTING -p udp --dport 8443 -j DNAT --to 10.66.66.2:8443
sudo iptables -t nat -A PREROUTING -p udp --dport 10000 -j DNAT --to 10.66.66.2:10000
```

Now within the Oracle Cloud interface, allow for port tcp/443, optional: tcp/80, and udp/10000 through. 80 and 443 are open for web traffic and 10000 is for the jvb on jitsi. The jvb is used when 3 or more users join a Jitsi meeting.

![Screenshot 2023-11-12 at 23 49 41](https://github.com/mark76ped/Infrastructure/assets/52715459/28534837-cff4-4723-a9d0-8b6e6cc8ec62)

### Install nginx-proxy-manager on Oracle Cloud Ubuntu via Docker
Following the instructions found on https://nginxproxymanager.com/guide/#quick-setup
- Install Docker and Docker-Compose
- Create a docker-compose.yaml file such as this:
```
version: '3.8'
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    ports:
      - '80:80'
      - '81:81'
      - '443:443'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
```
- Spin up the docker container
```
docker-compose up -d
```
- Check Oracle Cloud ACL for port tcp/81 connections from your public IP. This restricts access to the admin portal of the nginx-proxy-manager page.
- Log into the Admin Portal
```
http://$Your_Oracle_Cloud_Public_IP:81
```
- Enter the default admin creds:
```
Email:    admin@example.com
Password: changeme
```
- You will be prompted to change the default credintials upon your first successful log in.
- Proxy Host: After logging in, add a Proxy Host with the information regarding your network.
  - Domain Name: meet.mark-pedersen.com
  - Scheme: https
  - Forward Hostname / IP: 10.66.66.2
  - Forward Port: 8443

 *The Forward Hostname / IP must match the Wireguard IP of the on-prem Ubuntu Server
 *The Forward Port must match the configured above in the `docker ps` output from the on-prem jitsi docker setup.

 <img width="494" alt="Screenshot 2023-11-14 at 18 43 59" src="https://github.com/mark76ped/Infrastructure/assets/52715459/073655fb-ec22-430f-841d-792347030956">

- Add the following text to the Advanced tab according to Jitsi's offical documentation
```
location /xmpp-websocket {
    proxy_pass https://localhost:8443;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
location /colibri-ws {
    proxy_pass https://localhost:8443;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```
<img width="494" alt="Screenshot 2023-11-14 at 18 44 27" src="https://github.com/mark76ped/Infrastructure/assets/52715459/1948a98a-40df-4bb7-bd8b-d01a5ea63f5b">

- Access List Security
  - Add an access list for basic control over who has access to your meet services and reources with a `username` and `password`
     <img width="1174" alt="Screenshot 2023-11-14 at 18 44 56" src="https://github.com/mark76ped/Infrastructure/assets/52715459/cf46d845-578d-4a69-98fc-b27890fcdf29">
  - Make sure to add it to the Proxy Host configuation

### CloudFlare DNS / WAF Config

The final piece is to add CloudFlare free tier WAF protection, and DNS propagation.
- Make sure CloudFlare is hosting your root DNS name you purchased (usually for €7-€15/yr)
- Add an A record with your Oracle Cloud Ubuntu Public IP Address and put the Name you purchased in the Name section as shown below
- Create a CNAME record to allow to seperate traffic from your main site and meet traffic as also reflected in the nginx-proxy-manager proxy host config
<img width="1033" alt="Screenshot 2023-11-14 at 19 07 41" src="https://github.com/mark76ped/Infrastructure/assets/52715459/ab23cf16-711a-480e-af47-d251071c903d">
- (Optional) CloudFlare free tier allows for 5 WAF rules; add rules for blocking bots and threatscore ratings.
<img width="1532" alt="Screenshot 2023-11-14 at 19 12 29" src="https://github.com/mark76ped/Infrastructure/assets/52715459/6b5d3c8a-b60e-411a-9e3a-27d5cacb0219">

## Test out your Jitsi Meet setup
-  Create a meeting by going to your site (https://meet.thedomainyoupurchased.com)
-  The page should prompt you for the `username` and `password` you configured in the nginx-proxy-manager access list
-  Create meeting
![Screenshot 2023-11-14 at 19 17 33](https://github.com/mark76ped/Infrastructure/assets/52715459/7daff9fa-f04d-44b1-9f3b-32b33b43f8f7)
-  Then have two other devices join the meeting, this can be on the same network or different networks.
  - Why a total of three devices? Because thats when the jvb connection kicks in. Two connections can work fine with jvb, but if the configuration with iptables,ACL,nat, or pre/post routing is wrong 3 or more users cannot use the jitsi service
  - If the 3rd device/person joins the meeting and cameras AND audio cuts from the meeting, you have a jvb access issue mentioned above.
    - If just the video cuts off, but audio still works, that could be a bug specific to firefox/same public IP and can be ignored as it will function normally when users from other public IP addresses access the meeting.
-  If three devices connect to the meeting and can communication, then you have successfully configured your own jitsi meet setup with docker, ubuntu, wireguard, oracle cloud, cloudflare, and nginx-proxy-manager.
-  Now time limitations from Google Meet can never hassle you again.
