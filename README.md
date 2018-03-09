# Aaron's first code of Golang

For checking health of management server.
For creating instance by using WEB API with Golang.

**Linux distribution : Ubuntu 16.04**



## How to set up the environment.

        $ sudo apt-get install golang

        $ sudo mkdir /home/gohome/
        $ sudo export GOPATH=/home/gohome/
        $ sudo echo "export GOPATH=/home/gohome/" >> /etc/rc.local
        $ sudo go get github.com/revel/revel
        $ sudo go get github.com/revel/cmd/revel

        $ sudo $GOPATH/bin/revel new MyWeb
        $ sudo echo "$GOPATH/bin/revel run MyWeb" >> /etc/rc.local

        $ sudo reboot

        $ sudo cd /usr/local/src/     
        $ sudo wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-192.0.0-linux-x86_64.tar.gz
        $ sudo gunzip google-cloud-sdk-192.0.0-linux-x86_64.tar.gz 
        $ sudo tar xvf google-cloud-sdk-192.0.0-linux-x86_64.tar 
        $ sudo google-cloud-sdk/install.sh



## Help

* The [GCE API documentation for creating instance.](https://cloud.google.com/compute/docs/reference/rest/beta/instances/insert).
* The [How to code "hello world" with Golang.](https://revel.github.io/tutorial/firstapp.html).

