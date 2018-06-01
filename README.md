# go-clock-skew
## Description

This project is collecting tcp timestamps using a passive method and preparing for estimation of [clock-skew](https://github.com/aslanvaroqua/go-clock-skew).

clock-skew can tell us a lot about a network topology and is useful for understanding what type of systems are utilizing your network. 

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
