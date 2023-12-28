




# Wazuh SIEM Overview for active response to potential web server attackers 

### Initial Connection:
* Clients open connections normally
* They both received response 200 from the webserver
* The attacker attempts an attack which results in a response 400 
![wazuh-stage1 drawio](https://github.com/mark76ped/Infrastructure/assets/52715459/9f4f3e9b-7d6d-49ce-8d67-709c5a464fb8)




### Logging and Active Response:
* The webserver logs the responses
* The Wazuh agent sends this logs to the Wazuh server
* The Wazuh server triggers the active response module due to the response 400
* The Wazuh server pushs the IP Table change to the webserver client to block any further request from the attacking IP for 10 minutes
![wazuh-stage2 drawio](https://github.com/mark76ped/Infrastructure/assets/52715459/38e6c42e-ce57-4e05-908f-f5246a7521bb)



### Key takeaways
* The error log on the webserver is no longer reporting any attacks/errors since implimenting this solution
* This configuration frees up the webservers resoures
* Normal users are not affected by this configuration

