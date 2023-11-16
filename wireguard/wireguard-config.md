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
