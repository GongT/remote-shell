package constants

import "time"

// const MulticastAddress = "[ff12::feed:a:dead:beef]:2333"
const MulticastAddress = "10.0.255.255:23333"
const BoardcastListenAddress = "224.1.2.3:23333"
const Timeout = 3 * time.Second

const FileOpenLocalPrefix = "/data/Volumes/"
const FileOpenBase = "//shabao-share/"
