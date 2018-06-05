# go-clock-view
## Description

This project collects tcp timestamps passivly and in real time compares each packets system uptime to the uptime of the packets before it. It can provide useful information such as the of the number of devices behind a firewalled network without generating a packet. This can be useful in detecting tethering and to gain insight into the network topology behind a network address translator or in situations where promiscous detection is unsuitable. 

(https://github.com/aslanvaroqua/go-clock-view).

the offset of the uptimes per ip can tell us a lot about a network topology and is useful for understanding what type of systems are utilizing a network and also to prevent unauthorized tethering/connection sharing in the telecommunications industry.

## Usage
<pre><code>
Usage of ./go-clock-skew:<br>
  -e string<br>
    	device name (default "eth0")<br>
  -f string<br>
    	storage file (default "storage.csv")<br>
  -filter string<br>
    	bpFilter (default "tcp")<br>
  -h	help<br>
</pre></code>
## Example
<pre><code>
./go-clock-skew -filter "src host 10.10.89.144" -f 144.csv
</pre></code>

where to go from here?
```
go get "github.com/aslanvaroqua/go-clock-view"
go get "github.com/google/gopacket"
go get "go get gopkg.in/mgo.v2"
su -i
apt install mongodb
apt install libpcap-dev
ifconfig -> find appropriate network interface or use default (eth0)
go build main.go
mv main.go /usr/bin/clock-view
clock-view -e eth0 {other options)
```
