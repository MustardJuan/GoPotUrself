package main

import "fmt"

func ifconfig(){
	
	fmt.Println("enp0s3: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500\n" + 
        "	inet 10.0.2.15  netmask 255.255.255.0  broadcast 10.0.2.255\n" +
        "	inet6 fe80::af79:52e8:8bd7:ab09  prefixlen 64  scopeid 0x20<link>\n" +
        "	ether 08:00:27:37:29:52  txqueuelen 1000  (Ethernet)\n" +
        "	RX packets 271  bytes 196172 (196.1 KB)\n" +
        "	RX errors 0  dropped 0  overruns 0  frame 0\n" +
        "	TX packets 238  bytes 44197 (44.1 KB)\n" +
        "	TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n\n" +
	"lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536\n" +
        "	inet 127.0.0.1  netmask 255.0.0.0\n" +
        "	inet6 ::1  prefixlen 128  scopeid 0x10<host>\n" +
        "	loop  txqueuelen 1000  (Local Loopback)\n" +
        "	RX packets 189  bytes 14962 (14.9 KB)\n" +
        "	RX errors 0  dropped 0  overruns 0  frame 0\n" +
        "	TX packets 189  bytes 14962 (14.9 KB)\n" +
        "	TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0\n")

}
