# Osci

Go package with command line interface to oscilloscopes through telnet over ethernet.

## Usage

```shell
go run ./cmd --host 10.163.103.196 --port 4000 identity
TEKTRONIX,MSO54B,C052474,CF:91.1CT FV:1.44.3.433
```

```shell
 go run ./cmd --host 10.163.103.196 --port 4000 measurements
MEAS1, MEAS2
-8.5897227946348E-3, -1.7879619947382E-3
```