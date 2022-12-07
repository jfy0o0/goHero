package internal

import (
	"crypto/tls"
	"crypto/x509"
)

var svrCertificate = []byte(`-----BEGIN CERTIFICATE-----
MIIDMTCCAhkCFEZwoEdB+gkEg3yW46LGePQzZ9zWMA0GCSqGSIb3DQEBCwUAMFUx
CzAJBgNVBAYTAkNOMQswCQYDVQQIDAJaSjELMAkGA1UEBwwCSFoxDDAKBgNVBAoM
A29yZzEMMAoGA1UECwwDb3JnMRAwDgYDVQQDDAdvcmcuY29tMB4XDTIyMDMxNDA4
MDkwNFoXDTM1MTEyMTA4MDkwNFowVTELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAlpK
MQswCQYDVQQHDAJIWjEMMAoGA1UECgwDb3JnMQwwCgYDVQQLDANvcmcxEDAOBgNV
BAMMB29yZy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDgkSGH
CzwLbDl5K8516CCzkdnOm9DFoNtq50x0sT1oI5lMCKQ2ED76GNY442A5OY8UjsWU
hibhnHZzFYcgFIi2rA1VLtgGuX7GnCXLY1imGQuyCsmga5jqWGj1K6zEdUX+P2H1
s5EujQMRF3A9Gma2/0N6cU7yLyedG8AD7KrZYSP9Ahx+EZfVcdzmRpw7cePv9v75
uBTrojYf4o0mTxulKEsSX20Q0bpRxS4ttxOCyEo51r/dS4KAPQHyIKdGXcL8Tpv+
zuTixpznnCIexgkRwvKzLPQ1yCuWwJkzXq4zziICN7eABbOVArJBVqgegc8TyWmP
xqQ1sswCnyx9DhQNAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAISGB83vqPh0OHrH
Ni3z6udxwBRFnzypqE+s+8GZUx2Do4u2y5MZwv7Em0GVplmqEYXdcWmdRm2FcCUM
5XsEJ5cZoolyS0CttM7TS05BERrAyXfd0qt1WNrGU7mOVQZGo8V4HRTxR330IsXu
tua+3+xT+RFBE1cC36C8V7Ndh9TSlibP045PECkVGfNmRVftu9/p4h8MT61mMGx8
E3oaDHYnvtYYfuPK7LTRquOOwZSeg1MkPk2pBxPTOLU6hMkodDhjI+OqwYdVa92w
P4jPl0JbOtDlziPMQkzphylLJ2LuZLWNMxXygy1zC5NfyT46E0T9uR+dcd0lKUko
QrN890w=
-----END CERTIFICATE-----`)

var svrPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4JEhhws8C2w5eSvOdeggs5HZzpvQxaDbaudMdLE9aCOZTAik
NhA++hjWOONgOTmPFI7FlIYm4Zx2cxWHIBSItqwNVS7YBrl+xpwly2NYphkLsgrJ
oGuY6lho9SusxHVF/j9h9bORLo0DERdwPRpmtv9DenFO8i8nnRvAA+yq2WEj/QIc
fhGX1XHc5kacO3Hj7/b++bgU66I2H+KNJk8bpShLEl9tENG6UcUuLbcTgshKOda/
3UuCgD0B8iCnRl3C/E6b/s7k4sac55wiHsYJEcLysyz0NcgrlsCZM16uM84iAje3
gAWzlQKyQVaoHoHPE8lpj8akNbLMAp8sfQ4UDQIDAQABAoIBAQCHQqNBJeEquCr9
8wbOOdzIjAb2zR84iiTfBSYl+0OTYC3LgED4V1y3YxJU/Y9NqH+n0D7tsMfmiElI
FCVpizS1Yng9YJO5qSzVD+Qr11Dj6p4CxUCgoW5faJT/ZykOw9w+wYqNZXUC5uh7
+PQfubE4dBP+whTdBs4DcV18f6aAe0NJisfWFquvKqrZhVInPXfMCgjntKiKNPib
GqsRBj5Tokoz5Y4na45LZh2kAhof54itwXCDcWMST61qnTbtYLGf5fOwLOLjdyIa
0nyp15yRPUjy2Z7loj2s0Q9GqQhmY/65xE8B7v/qKj8ZdkFngGosweI4Uv5Sxeyt
D5ft0VZ5AoGBAPJBwJvSoDhdcqXlGgHakdsdc/U1Lx6N7D+g15FvzJHhQcYrznIR
UGkr6i5+0JiU74eWy9q7P5N25y3VY5J7hmClORUQnZOC6vne3dfKlqX/3bF27Gb0
oxaa7WS8uAd2PfUV3EGvMXWmkP9vzfuORBE94EWWZlYcQEdbQ1JHbj8TAoGBAO1O
eKhQqJph0KMw2liJRJaomuXamcvrtzc7Wq+udazxM+Brc4l13QifjjtzZJ440ZwX
zYRACQMT9AuV4SsVvF94R0w84Z6tPX1UOmi0Kl1ezIgeXUHchlpKoR//bG68waaj
/N+0siSu4ygjnaRQ1Er/5WK8wcL923rDh3P3YiRfAoGAD0rZRrzYAlbbyt8Yci+C
74r/a+YskEgzNp0HnpWpGzhV95WCVa8EHW95O/AjaM5WBIAfZDJXxA5Ib5s9ytxj
Uix72vmOiOZFwPoxlNKkCdyJ+Q7hw2oD9YRYIfjLwSTUqmjz222cuKmIE6fSH+QO
Jex7AMmKsjrcT+TR7hq2OeMCgYEAnnTkind5bcTc3Os1ESNvoi4paEwkoA7gg0fl
0SNsPjO9USBOwL3r/uMTxUzTJfeIt3MJb0KXAfAj6EyHUw6rfxBAg36vpRXG1lYU
vm93TmuLpWSpHt5S9B2bI1OaIwN1R4F+zB3LSkVhNx85F0xFxWLD1790nnCGSp4f
Z2cLqH8CgYAK6bVzLgSzdcOydFU2WxvqCfyO92Telx40Qc8VnM3/lyVRXYZ6FBWE
iLJvniQHfSzEFZMteAGF4v3EbdEM5GBHIoKZWA1fAz0tqjBPRuHV4zy6n4gEfrdj
CTG59PLFnVRsrrAYHZWnh7nF1oUJwimQjDvjntONUgyrIO99nq974A==
-----END RSA PRIVATE KEY-----`)

var svrCertificate2 = []byte(`-----BEGIN CERTIFICATE-----
MIIC6zCCAdOgAwIBAgIUeaiXr0uM2TYo3MombRr+NrMW/yIwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAwwHb3JnLmNvbTAeFw0yMjAzMTYwMzA5NDFaFw0zMjAzMTMw
MzA5NDFaMBIxEDAOBgNVBAMMB29yZy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQC6QEdbNdOP6Ltx+8CyuvweyvmqX75Dv03kj+/33ZYVd8+ML34Q
NOFha9BByw/JuXLUvpLdeydw1XoVxVLXSSKYQoV5f9I3mMcyLRWpodDxQOdXQICC
h1SnX8CWSaSOPltQAXX5pXD3yGwZbPkXNPIuEhbbQNGMPUTTdYxH0UKP89X8lBON
K0luZOLTWtpuBlik5KwppcjZyqUR/p54kg/vQORDq99EOg368Eaw2MQ8DXl6zV1X
w0o8GKgtVklbGkN571DZv/IveoOVAQe9Pj7Hhr/VCoMaVfGuFnkzFqONClqjwU2N
hpLBxNTgsj0ooFoz4l3fdK0NzVQX6Jx2tQorAgMBAAGjOTA3MB0GA1UdEQQWMBSC
B29yZy5jb22CCSoub3JnLmNvbTAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DANBgkq
hkiG9w0BAQsFAAOCAQEAcA2MEipRDYKmCeiDbj3aWWEpXi5aqqQdT3S15hCbonjs
3c3MM8oU5zWhUgqHqfgmTtKrYowGdLNA6FCiSG15KZDoxwtaRn/ZP8ShCk8fqN0h
O/FLv7styNbxXP4MuWke9MLdiplQ1idlupSG9c5qRgR0j2IdsojiwFrlgIVe0K1+
SlT44Nzy2LyTpj1rbyeH4gP5aMwjLKWFSFdmCehcOM4W0zBzJXRLp1wLDHGsRZOz
AuMP32iATqZ5ADxWYXYxvBJNnfcWQXTaNLe/L4t/saIAuT7kLW7qaHL5GIJ3N5OV
cDE6m7h4BKAmLNiShsv0lESxVKbP4eNK3Uu5uAcOzw==
-----END CERTIFICATE-----`)

var svrPrivateKey2 = []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC6QEdbNdOP6Ltx
+8CyuvweyvmqX75Dv03kj+/33ZYVd8+ML34QNOFha9BByw/JuXLUvpLdeydw1XoV
xVLXSSKYQoV5f9I3mMcyLRWpodDxQOdXQICCh1SnX8CWSaSOPltQAXX5pXD3yGwZ
bPkXNPIuEhbbQNGMPUTTdYxH0UKP89X8lBONK0luZOLTWtpuBlik5KwppcjZyqUR
/p54kg/vQORDq99EOg368Eaw2MQ8DXl6zV1Xw0o8GKgtVklbGkN571DZv/IveoOV
AQe9Pj7Hhr/VCoMaVfGuFnkzFqONClqjwU2NhpLBxNTgsj0ooFoz4l3fdK0NzVQX
6Jx2tQorAgMBAAECggEAODrz0Bu+FbKD9PO2x78gIwPQwJlLThTRpXG3VzP0/eJA
LOihvK7aZJkyNMrjp1lvy78FhAyMacppo20Bh7Gy2hkrRvVKODLN4N4vrhkGhe+8
aGE26zGEY0vLTxRItBErbEokngxOwOAnkr2Nrm3zt+kKMYOuQBK3VsgkZAJPgDnv
xV/38DEcDmof3O42QR8+PLP6PKJAmtusaMXQyvVaMmSWQZrl31f4tAsU4QduRhaJ
9qt1O0RjeV0NUpOWsa1ALdJwuatfhv0GWHSSoExLI47IAmW6dNMqL1k7H/zeE29w
AkmMqJftop27XprC1NsTDoVYFLRIKRegyOPFvmmzQQKBgQDa8oBBmuw4oBaKt54n
HllaaQOcLkno5aj+GqbfNOStnIgiY87n5wMedCvY2QjN6gGdXlJJ+mwKTEOyDs2G
E57Lm0/JiPCOQVFrtqE1/eTUrZPbsACNm/K6c6yUjklDQGDRF+5wvosLwgLq+z8F
S2xLaoitPeT25bDo5Iw4U0aV0QKBgQDZxUYvdpfR9QP4bYhwjrflQ26zSHCWD46u
VdecRNBha7VpKoDGEMa7ngHtLoWTPn+vHyCwYEuy9nm17Gy/YZ5gxgg4KCgA0TJ4
IytGd4DObY7bVM+fFWjyhLa3KYz+YJ6f2wndUv+wyH4dzQvmLPultKtmwIFj9DRo
gLMfTbITOwKBgQCKAzijoiNrIkZzVIFFyVSrr+yNo4QFYdcfre8oXfAUG8qYWu/O
sIj/xlzpSQ3Ktsojx5P5e/hkmWRGthwJew6q65DkygRQt426ZxnBrfRzt6KvZYdD
vFzJ+SZ9Uh+OW+NrijlMKl/9nmM2ef9kuAFgnJvhEFXm6CyW8ZIN4zhoMQKBgGat
MO5oCCUiVQDpBA+t+t7J8IPvf8YBB9EDAGnZuDZ2xk4TkCvS9SC420etS6WlLiav
koYWoRp8Q96W3p9Ns5MFTBLOPC5GdHlYe6r55W1vEpAvaMUlHxpZoJyzppq5i5Dr
gbcszwnXPj9m0llwlCYJDVehJGCwX5V2OP5lb7TJAoGAVvOC9j3ID0vyF/T/PDyO
+NngHvI32szBJtFQlnW6i7U6DIQntraJ0NudGbkImbeal3+oe64e+66g1Xrm3rJQ
rbzGKmraUrKP5uunRZHE7YVP5ckUS0xf8is7pb44KqksxJTu+QPlH1gidgCIq+/h
Xt+m09kNSOuMslWyXm3CBUc=
-----END PRIVATE KEY-----
`)
var svrTlsConfig *tls.Config

func GetServerTlsConfig() *tls.Config {
	if svrTlsConfig == nil {
		crt, _ := tls.X509KeyPair(svrCertificate2, svrPrivateKey2)
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caCertificate2)
		svrTlsConfig = &tls.Config{
			ClientCAs:    pool,
			Certificates: []tls.Certificate{crt},
			ClientAuth:   tls.RequireAndVerifyClientCert,
		}
	}
	return svrTlsConfig
}
