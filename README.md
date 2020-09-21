# asnap

## Description
asnap aims to render recon phase easier by providing updated data about which companies holds which ipv4 or ipv6 addresses and allows the user to automate initial port and service scanning.

video/screenshot/banner here.



## Installation 
### Precompiled Binary

If you have Go installed and configured in your $PATH enviroment variable, simply run:
```
go get -u github.com/paradoxxer/asngo
```
If you want to use precompiled binary, you also need to download "move.sh" and "nmap.sh" files and place them in the same directory with asnap.

To download the database that asnap searches from, you need to provide a key. To get your free key, sign up here -> https://www.maxmind.com/en/geolite2/signup
after signing up navigate -> services -> My license key and create new key. Then create "asnap_conf.txt" inside the same directory with asnap, and paste your key to first line of "asnap_conf.txt":

```
echo 'insert key' > asnap_conf.txt 
```

To be able to use port scanning functionality, you need to install nmap to your machine:
```
* Debian Based Distros:
    sudo apt install nmap
* MacOS
    brew install nmap 
* Arch Based Distros
    sudo pacman -S nmap
```

### Build Yourself (Recommended)

Install golang here -> https://golang.org/doc/install
or you can install go if it is provided with your package manager:
```
sudo apt install golang
```
After downloading the source code, navigate through the project directory and run:
```
go build
```

this will produce asnap binary. After you build it, create asnap_conf.txt inside the same directory with the asnap, insert your key to first line and you are good to go.

```
echo 'insert key' > asnap_conf.txt 
```



usage and examples

follow me on x.





