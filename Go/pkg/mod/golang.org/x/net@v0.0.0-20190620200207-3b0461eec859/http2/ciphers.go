// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

// A list of the possible cipher suite ids. Taken from
// https://www.iana.org/assignments/tls-parameters/tls-parameters.txt

const (
	cipher_TLS_NULL_WITH_NULL_NULL               uint16 = 0x0000
	cipher_TLS_RSA_WITH_NULL_MD5                 uint16 = 0x0001
	cipher_TLS_RSA_WITH_NULL_SHA                 uint16 = 0x0002
	cipher_TLS_RSA_EXPORT_WITH_RC4_40_MD5        uint16 = 0x0003
	cipher_TLS_RSA_WITH_RC4_128_MD5              uint16 = 0x0004
	cipher_TLS_RSA_WITH_RC4_128_SHA              uint16 = 0x0005
	cipher_TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5    uint16 = 0x0006
	cipher_TLS_RSA_WITH_IDEA_CBC_SHA             uint16 = 0x0007
	cipher_TLS_RSA_EXPORT_WITH_DES40_CBC_SHA     uint16 = 0x0008
	cipher_TLS_RSA_WITH_DES_CBC_SHA              uint16 = 0x0009
	cipher_TLS_RSA_WITH_3DES_EDE_CBC_SHA         uint16 = 0x000A
	cipher_TLS_DH_DSS_EXPORT_WITH_DES40_CBC_SHA  uint16 = 0x000B
	cipher_TLS_DH_DSS_WITH_DES_CBC_SHA           uint16 = 0x000C
	cipher_TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA      uint16 = 0x000D
	cipher_TLS_DH_RSA_EXPORT_WITH_DES40_CBC_SHA  uint16 = 0x000E
	cipher_TLS_DH_RSA_WITH_DES_CBC_SHA           uint16 = 0x000F
	cipher_TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA      uint16 = 0x0010
	cipher_TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0011
	cipher_TLS_DHE_DSS_WITH_DES_CBC_SHA          uint16 = 0x0012
	cipher_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA     uint16 = 0x0013
	cipher_TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0014
	cipher_TLS_DHE_RSA_WITH_DES_CBC_SHA          uint16 = 0x0015
	cipher_TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA     uint16 = 0x0016
	cipher_TLS_DH_anon_EXPORT_WITH_RC4_40_MD5    uint16 = 0x0017
	cipher_TLS_DH_anon_WITH_RC4_128_MD5          uint16 = 0x0018
	cipher_TLS_DH_anon_EXPORT_WITH_DES40_CBC_SHA uint16 = 0x0019
	cipher_TLS_DH_anon_WITH_DES_CBC_SHA          uint16 = 0x001A
	cipher_TLS_DH_anon_WITH_3DES_EDE_CBC_SHA     uint16 = 0x001B
	// Reserved uint16 =  0x001C-1D
	cipher_TLS_KRB5_WITH_DES_CBC_SHA             uint16 = 0x001E
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_SHA        uint16 = 0x001F
	cipher_TLS_KRB5_WITH_RC4_128_SHA             uint16 = 0x0020
	cipher_TLS_KRB5_WITH_IDEA_CBC_SHA            uint16 = 0x0021
	cipher_TLS_KRB5_WITH_DES_CBC_MD5             uint16 = 0x0022
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_MD5        uint16 = 0x0023
	cipher_TLS_KRB5_WITH_RC4_128_MD5             uint16 = 0x0024
	cipher_TLS_KRB5_WITH_IDEA_CBC_MD5            uint16 = 0x0025
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_SHA   uint16 = 0x0026
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_SHA   uint16 = 0x0027
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_SHA       uint16 = 0x0028
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_MD5   uint16 = 0x0029
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_MD5   uint16 = 0x002A
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_MD5       uint16 = 0x002B
	cipher_TLS_PSK_WITH_NULL_SHA                 uint16 = 0x002C
	cipher_TLS_DHE_PSK_WITH_NULL_SHA             uint16 = 0x002D
	cipher_TLS_RSA_PSK_WITH_NULL_SHA             uint16 = 0x002E
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA          uint16 = 0x002F
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA       uint16 = 0x0030
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA       uint16 = 0x0031
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA      uint16 = 0x0032
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA      uint16 = 0x0033
	cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA      uint16 = 0x0034
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA          uint16 = 0x0035
	cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA       uint16 = 0x0036
	cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA       uint16 = 0x0037
	cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA      uint16 = 0x0038
	cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA      uint16 = 0x0039
	cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA      uint16 = 0x003A
	cipher_TLS_RSA_WITH_NULL_SHA256              uint16 = 0x003B
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA256       uint16 = 0x003C
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA256       uint16 = 0x003D
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA256    uint16 = 0x003E
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA256    uint16 = 0x003F
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256   uint16 = 0x0040
	cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA     uint16 = 0x0041
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA  uint16 = 0x0042
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA  uint16 = 0x0043
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0044
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0045
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA uint16 = 0x0046
	// Reserved uint16 =  0x0047-4F
	// Reserved uint16 =  0x0050-58
	// Reserved uint16 =  0x0059-5C
	// Unassigned uint16 =  0x005D-5F
	// Reserved uint16 =  0x0060-66
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 uint16 = 0x0067
	cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA256  uint16 = 0x0068
	cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA256  uint16 = 0x0069
	cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA256 uint16 = 0x006A
	cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 uint16 = 0x006B
	cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA256 uint16 = 0x006C
	cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA256 uint16 = 0x006D
	// Unassigned uint16 =  0x006E-83
	cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA        uint16 = 0x0084
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA     uint16 = 0x0085
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA     uint16 = 0x0086
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA    uint16 = 0x0087
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA    uint16 = 0x0088
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA    uint16 = 0x0089
	cipher_TLS_PSK_WITH_RC4_128_SHA                 uint16 = 0x008A
	cipher_TLS_PSK_WITH_3DES_EDE_CBC_SHA            uint16 = 0x008B
	cipher_TLS_PSK_WITH_AES_128_CBC_SHA             uint16 = 0x008C
	cipher_TLS_PSK_WITH_AES_256_CBC_SHA             uint16 = 0x008D
	cipher_TLS_DHE_PSK_WITH_RC4_128_SHA             uint16 = 0x008E
	cipher_TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA        uint16 = 0x008F
	cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA         uint16 = 0x0090
	cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA         uint16 = 0x0091
	cipher_TLS_RSA_PSK_WITH_RC4_128_SHA             uint16 = 0x0092
	cipher_TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA        uint16 = 0x0093
	cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA         uint16 = 0x0094
	cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA         uint16 = 0x0095
	cipher_TLS_RSA_WITH_SEED_CBC_SHA                uint16 = 0x0096
	cipher_TLS_DH_DSS_WITH_SEED_CBC_SHA             uint16 = 0x0097
	cipher_TLS_DH_RSA_WITH_SEED_CBC_SHA             uint16 = 0x0098
	cipher_TLS_DHE_DSS_WITH_SEED_CBC_SHA            uint16 = 0x0099
	cipher_TLS_DHE_RSA_WITH_SEED_CBC_SHA            uint16 = 0x009A
	cipher_TLS_DH_anon_WITH_SEED_CBC_SHA            uint16 = 0x009B
	cipher_TLS_RSA_WITH_AES_128_GCM_SHA256          uint16 = 0x009C
	cipher_TLS_RSA_WITH_AES_256_GCM_SHA384          uint16 = 0x009D
	cipher_TLS_DHE_RSA_WITH_AES_128_GCM_SHA256      uint16 = 0x009E
	cipher_TLS_DHE_RSA_WITH_AES_256_GCM_SHA384      uint16 = 0x009F
	cipher_TLS_DH_RSA_WITH_AES_128_GCM_SHA256       uint16 = 0x00A0
	cipher_TLS_DH_RSA_WITH_AES_256_GCM_SHA384       uint16 = 0x00A1
	cipher_TLS_DHE_DSS_WITH_AES_128_GCM_SHA256      uint16 = 0x00A2
	cipher_TLS_DHE_DSS_WITH_AES_256_GCM_SHA384      uint16 = 0x00A3
	cipher_TLS_DH_DSS_WITH_AES_128_GCM_SHA256       uint16 = 0x00A4
	cipher_TLS_DH_DSS_WITH_AES_256_GCM_SHA384       uint16 = 0x00A5
	cipher_TLS_DH_anon_WITH_AES_128_GCM_SHA256      uint16 = 0x00A6
	cipher_TLS_DH_anon_WITH_AES_256_GCM_SHA384      uint16 = 0x00A7
	cipher_TLS_PSK_WITH_AES_128_GCM_SHA256          uint16 = 0x00A8
	cipher_TLS_PSK_WITH_AES_256_GCM_SHA384          uint16 = 0x00A9
	cipher_TLS_DHE_PSK_WITH_AES_128_GCM_SHA256      uint16 = 0x00AA
	cipher_TLS_DHE_PSK_WITH_AES_256_GCM_SHA384      uint16 = 0x00AB
	cipher_TLS_RSA_PSK_WITH_AES_128_GCM_SHA256      uint16 = 0x00AC
	cipher_TLS_RSA_PSK_WITH_AES_256_GCM_SHA384      uint16 = 0x00AD
	cipher_TLS_PSK_WITH_AES_128_CBC_SHA256          uint16 = 0x00AE
	cipher_TLS_PSK_WITH_AES_256_CBC_SHA384          uint16 = 0x00AF
	cipher_TLS_PSK_WITH_NULL_SHA256                 uint16 = 0x00B0
	cipher_TLS_PSK_WITH_NULL_SHA384                 uint16 = 0x00B1
	cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA256      uint16 = 0x00B2
	cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA384      uint16 = 0x00B3
	cipher_TLS_DHE_PSK_WITH_NULL_SHA256             uint16 = 0x00B4
	cipher_TLS_DHE_PSK_WITH_NULL_SHA384             uint16 = 0x00B5
	cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA256      uint16 = 0x00B6
	cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA384      uint16 = 0x00B7
	cipher_TLS_RSA_PSK_WITH_NULL_SHA256             uint16 = 0x00B8
	cipher_TLS_RSA_PSK_WITH_NULL_SHA384             uint16 = 0x00B9
	cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256     uint16 = 0x00BA
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA256  uint16 = 0x00BB
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA256  uint16 = 0x00BC
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256 uint16 = 0x00BD
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256 uint16 = 0x00BE
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA256 uint16 = 0x00BF
	cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256     uint16 = 0x00C0
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA256  uint16 = 0x00C1
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA256  uint16 = 0x00C2
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256 uint16 = 0x00C3
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256 uint16 = 0x00C4
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA256 uint16 = 0x00C5
	// Unassigned uint16 =  0x00C6-FE
	cipher_TLS_EMPTY_RENEGOTIATION_INFO_SCSV uint16 = 0x00FF
	// Unassigned uint16 =  0x01-55,*
	cipher_TLS_FALLBACK_SCSV uint16 = 0x5600
	// Unassigned                                   uint16 = 0x5601 - 0xC000
	cipher_TLS_ECDH_ECDSA_WITH_NULL_SHA                 uint16 = 0xC001
	cipher_TLS_ECDH_ECDSA_WITH_RC4_128_SHA              uint16 = 0xC002
	cipher_TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA         uint16 = 0xC003
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA          uint16 = 0xC004
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA          uint16 = 0xC005
	cipher_TLS_ECDHE_ECDSA_WITH_NULL_SHA                uint16 = 0xC006
	cipher_TLS_ECDHE_ECDSA_WITH_RC4_128_SHA             uint16 = 0xC007
	cipher_TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA        uint16 = 0xC008
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA         uint16 = 0xC009
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA         uint16 = 0xC00A
	cipher_TLS_ECDH_RSA_WITH_NULL_SHA                   uint16 = 0xC00B
	cipher_TLS_ECDH_RSA_WITH_RC4_128_SHA                uint16 = 0xC00C
	cipher_TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA           uint16 = 0xC00D
	cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA            uint16 = 0xC00E
	cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA            uint16 = 0xC00F
	cipher_TLS_ECDHE_RSA_WITH_NULL_SHA                  uint16 = 0xC010
	cipher_TLS_ECDHE_RSA_WITH_RC4_128_SHA               uint16 = 0xC011
	cipher_TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA          uint16 = 0xC012
	cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA           uint16 = 0xC013
	cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA           uint16 = 0xC014
	cipher_TLS_ECDH_anon_WITH_NULL_SHA                  uint16 = 0xC015
	cipher_TLS_ECDH_anon_WITH_RC4_128_SHA               uint16 = 0xC016
	cipher_TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA          uint16 = 0xC017
	cipher_TLS_ECDH_anon_WITH_AES_128_CBC_SHA           uint16 = 0xC018
	cipher_TLS_ECDH_anon_WITH_AES_256_CBC_SHA           uint16 = 0xC019
	cipher_TLS_SRP_SHA_WITH_3DES_EDE_CBC_SHA            uint16 = 0xC01A
	cipher_TLS_SRP_SHA_RSA_WITH_3DES_EDE_CBC_SHA        uint16 = 0xC01B
	cipher_TLS_SRP_SHA_DSS_WITH_3DES_EDE_CBC_SHA        uint16 = 0xC01C
	cipher_TLS_SRP_SHA_WITH_AES_128_CBC_SHA             uint16 = 0xC01D
	cipher_TLS_SRP_SHA_RSA_WITH_AES_128_CBC_SHA         uint16 = 0xC01E
	cipher_TLS_SRP_SHA_DSS_WITH_AES_128_CBC_SHA         uint16 = 0xC01F
	cipher_TLS_SRP_SHA_WITH_AES_256_CBC_SHA             uint16 = 0xC020
	cipher_TLS_SRP_SHA_RSA_WITH_AES_256_CBC_SHA         uint16 = 0xC021
	cipher_TLS_SRP_SHA_DSS_WITH_AES_256_CBC_SHA         uint16 = 0xC022
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256      uint16 = 0xC023
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384      uint16 = 0xC024
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256       uint16 = 0xC025
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384       uint16 = 0xC026
	cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256        uint16 = 0xC027
	cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384        uint16 = 0xC028
	cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256         uint16 = 0xC029
	cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384         uint16 = 0xC02A
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256      uint16 = 0xC02B
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384      uint16 = 0xC02C
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256       uint16 = 0xC02D
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384       uint16 = 0xC02E
	cipher_TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256        uint16 = 0xC02F
	cipher_TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384        uint16 = 0xC030
	cipher_TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256         uint16 = 0xC031
	cipher_TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384         uint16 = 0xC032
	cipher_TLS_ECDHE_PSK_WITH_RC4_128_SHA               uint16 = 0xC033
	cipher_TLS_ECDHE_PSK_WITH_3DES_EDE_CBC_SHA          uint16 = 0xC034
	cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA           uint16 = 0xC035
	cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA           uint16 = 0xC036
	cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256        uint16 = 0xC037
	cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA384        uint16 = 0xC038
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA                  uint16 = 0xC039
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA256               uint16 = 0xC03A
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA384               uint16 = 0xC03B
	cipher_TLS_RSA_WITH_ARIA_128_CBC_SHA256             uint16 = 0xC03C
	cipher_TLS_RSA_WITH_ARIA_256_CBC_SHA384             uint16 = 0xC03D
	cipher_TLS_DH_DSS_WITH_ARIA_128_CBC_SHA256          uint16 = 0xC03E
	cipher_TLS_DH_DSS_WITH_ARIA_256_CBC_SHA384          uint16 = 0xC03F
	cipher_TLS_DH_RSA_WITH_ARIA_128_CBC_SHA256          uint16 = 0xC040
	cipher_TLS_DH_RSA_WITH_ARIA_256_CBC_SHA384          uint16 = 0xC041
	cipher_TLS_DHE_DSS_WITH_ARIA_128_CBC_SHA256         uint16 = 0xC042
	cipher_TLS_DHE_DSS_WITH_ARIA_256_CBC_SHA384         uint16 = 0xC043
	cipher_TLS_DHE_RSA_WITH_ARIA_128_CBC_SHA256         uint16 = 0xC044
	cipher_TLS_DHE_RSA_WITH_ARIA_256_CBC_SHA384         uint16 = 0xC045
	cipher_TLS_DH_anon_WITH_ARIA_128_CBC_SHA256         uint16 = 0xC046
	cipher_TLS_DH_anon_WITH_ARIA_256_CBC_SHA384         uint16 = 0xC047
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_128_CBC_SHA256     uint16 = 0xC048
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_256_CBC_SHA384     uint16 = 0xC049
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_CBC_SHA256      uint16 = 0xC04A
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_CBC_SHA384      uint16 = 0xC04B
	cipher_TLS_ECDHE_RSA_WITH_ARIA_128_CBC_SHA256       uint16 = 0xC04C
	cipher_TLS_ECDHE_RSA_WITH_ARIA_256_CBC_SHA384       uint16 = 0xC04D
	cipher_TLS_ECDH_RSA_WITH_ARIA_128_CBC_SHA256        uint16 = 0xC04E
	cipher_TLS_ECDH_RSA_WITH_ARIA_256_CBC_SHA384        uint16 = 0xC04F
	cipher_TLS_RSA_WITH_ARIA_128_GCM_SHA256             uint16 = 0xC050
	cipher_TLS_RSA_WITH_ARIA_256_GCM_SHA384             uint16 = 0xC051
	cipher_TLS_DHE_RSA_WITH_ARIA_128_GCM_SHA256         uint16 = 0xC052
	cipher_TLS_DHE_RSA_WITH_ARIA_256_GCM_SHA384         uint16 = 0xC053
	cipher_TLS_DH_RSA_WITH_ARIA_128_GCM_SHA256          uint16 = 0xC054
	cipher_TLS_DH_RSA_WITH_ARIA_256_GCM_SHA384          uint16 = 0xC055
	cipher_TLS_DHE_DSS_WITH_ARIA_128_GCM_SHA256         uint16 = 0xC056
	cipher_TLS_DHE_DSS_WITH_ARIA_256_GCM_SHA384         uint16 = 0xC057
	cipher_TLS_DH_DSS_WITH_ARIA_128_GCM_SHA256          uint16 = 0xC058
	cipher_TLS_DH_DSS_WITH_ARIA_256_GCM_SHA384          uint16 = 0xC059
	cipher_TLS_DH_anon_WITH_ARIA_128_GCM_SHA256         uint16 = 0xC05A
	cipher_TLS_DH_anon_WITH_ARIA_256_GCM_SHA384         uint16 = 0xC05B
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_128_GCM_SHA256     uint16 = 0xC05C
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_256_GCM_SHA384     uint16 = 0xC05D
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_GCM_SHA256      uint16 = 0xC05E
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_GCM_SHA384      uint16 = 0xC05F
	cipher_TLS_ECDHE_RSA_WITH_ARIA_128_GCM_SHA256       uint16 = 0xC060
	cipher_TLS_ECDHE_RSA_WITH_ARIA_256_GCM_SHA384       uint16 = 0xC061
	cipher_TLS_ECDH_RSA_WITH_ARIA_128_GCM_SHA256        uint16 = 0xC062
	cipher_TLS_ECDH_RSA_WITH_ARIA_256_GCM_SHA384        uint16 = 0xC063
	cipher_TLS_PSK_WITH_ARIA_128_CBC_SHA256             uint16 = 0xC064
	cipher_TLS_PSK_WITH_ARIA_256_CBC_SHA384             uint16 = 0xC065
	cipher_TLS_DHE_PSK_WITH_ARIA_128_CBC_SHA256         uint16 = 0xC066
	cipher_TLS_DHE_PSK_WITH_ARIA_256_CBC_SHA384         uint16 = 0xC067
	cipher_TLS_RSA_PSK_WITH_ARIA_128_CBC_SHA256         uint16 = 0xC068
	cipher_TLS_RSA_PSK_WITH_ARIA_256_CBC_SHA384         uint16 = 0xC069
	cipher_TLS_PSK_WITH_ARIA_128_GCM_SHA256             uint16 = 0xC06A
	cipher_TLS_PSK_WITH_ARIA_256_GCM_SHA384             uint16 = 0xC06B
	cipher_TLS_DHE_PSK_WITH_ARIA_128_GCM_SHA256         uint16 = 0xC06C
	cipher_TLS_DHE_PSK_WITH_ARIA_256_GCM_SHA384         uint16 = 0xC06D
	cipher_TLS_RSA_PSK_WITH_ARIA_128_GCM_SHA256         uint16 = 0xC06E
	cipher_TLS_RSA_PSK_WITH_ARIA_256_GCM_SHA384         uint16 = 0xC06F
	cipher_TLS_ECDHE_PSK_WITH_ARIA_128_CBC_SHA256       uint16 = 0xC070
	cipher_TLS_ECDHE_PSK_WITH_ARIA_256_CBC_SHA384       uint16 = 0xC071
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256 uint16 = 0xC072
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384 uint16 = 0xC073
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_CBC_SHA256  uint16 = 0xC074
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_CBC_SHA384  uint16 = 0xC075
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256   uint16 = 0xC076
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384   uint16 = 0xC077
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_CBC_SHA256    uint16 = 0xC078
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_CBC_SHA384    uint16 = 0xC079
	cipher_TLS_RSA_WITH_CAMELLIA_128_GCM_SHA256         uint16 = 0xC07A
	cipher_TLS_RSA_WITH_CAMELLIA_256_GCM_SHA384         uint16 = 0xC07B
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_GCM_SHA256     uint16 = 0xC07C
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_GCM_SHA384     uint16 = 0xC07D
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_GCM_SHA256      uint16 = 0xC07E
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_GCM_SHA384      uint16 = 0xC07F
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_GCM_SHA256     uint16 = 0xC080
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_GCM_SHA384     uint16 = 0xC081
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_GCM_SHA256      uint16 = 0xC082
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_GCM_SHA384      uint16 = 0xC083
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_GCM_SHA256     uint16 = 0xC084
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_GCM_SHA384     uint16 = 0xC085
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_GCM_SHA256 uint16 = 0xC086
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_GCM_SHA384 uint16 = 0xC087
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_GCM_SHA256  uint16 = 0xC088
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_GCM_SHA384  uint16 = 0xC089
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_128_GCM_SHA256   uint16 = 0xC08A
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_256_GCM_SHA384   uint16 = 0xC08B
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_GCM_SHA256    uint16 = 0xC08C
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_GCM_SHA384    uint16 = 0xC08D
	cipher_TLS_PSK_WITH_CAMELLIA_128_GCM_SHA256         uint16 = 0xC08E
	cipher_TLS_PSK_WITH_CAMELLIA_256_GCM_SHA384         uint16 = 0xC08F
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_128_GCM_SHA256     uint16 = 0xC090
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_256_GCM_SHA384     uint16 = 0xC091
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_GCM_SHA256     uint16 = 0xC092
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_GCM_SHA384     uint16 = 0xC093
	cipher_TLS_PSK_WITH_CAMELLIA_128_CBC_SHA256         uint16 = 0xC094
	cipher_TLS_PSK_WITH_CAMELLIA_256_CBC_SHA384         uint16 = 0xC095
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_128_CBC_SHA256     uint16 = 0xC096
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_256_CBC_SHA384     uint16 = 0xC097
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_CBC_SHA256     uint16 = 0xC098
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_CBC_SHA384     uint16 = 0xC099
	cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_128_CBC_SHA256   uint16 = 0xC09A
	cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_256_CBC_SHA384   uint16 = 0xC09B
	cipher_TLS_RSA_WITH_AES_128_CCM                     uint16 = 0xC09C
	cipher_TLS_RSA_WITH_AES_256_CCM                     uint16 = 0xC09D
	cipher_TLS_DHE_RSA_WITH_AES_128_CCM                 uint16 = 0xC09E
	cipher_TLS_DHE_RSA_WITH_AES_256_CCM                 uint16 = 0xC09F
	cipher_TLS_RSA_WITH_AES_128_CCM_8                   uint16 = 0xC0A0
	cipher_TLS_RSA_WITH_AES_256_CCM_8                   uint16 = 0xC0A1
	cipher_TLS_DHE_RSA_WITH_AES_128_CCM_8               uint16 = 0xC0A2
	cipher_TLS_DHE_RSA_WITH_AES_256_CCM_8               uint16 = 0xC0A3
	cipher_TLS_PSK_WITH_AES_128_CCM                     uint16 = 0xC0A4
	cipher_TLS_PSK_WITH_AES_256_CCM                     uint16 = 0xC0A5
	cipher_TLS_DHE_PSK_WITH_AES_128_CCM                 uint16 = 0xC0A6
	cipher_TLS_DHE_PSK_WITH_AES_256_CCM                 uint16 = 0xC0A7
	cipher_TLS_PSK_WITH_AES_128_CCM_8                   uint16 = 0xC0A8
	cipher_TLS_PSK_WITH_AES_256_CCM_8                   uint16 = 0xC0A9
	cipher_TLS_PSK_DHE_WITH_AES_128_CCM_8               uint16 = 0xC0AA
	cipher_TLS_PSK_DHE_WITH_AES_256_CCM_8               uint16 = 0xC0AB
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CCM             uint16 = 0xC0AC
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CCM             uint16 = 0xC0AD
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8           uint16 = 0xC0AE
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8           uint16 = 0xC0AF
	// Unassigned uint16 =  0xC0B0-FF
	// Unassigned uint16 =  0xC1-CB,*
	// Unassigned uint16 =  0xCC00-A7
	cipher_TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   uint16 = 0xCCA8
	cipher_TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 uint16 = 0xCCA9
	cipher_TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256     uint16 = 0xCCAA
	cipher_TLS_PSK_WITH_CHACHA20_POLY1305_SHA256         uint16 = 0xCCAB
	cipher_TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256   uint16 = 0xCCAC
	cipher_TLS_DHE_PSK_WITH_CHACHA20_POLY1305_SHA256     uint16 = 0xCCAD
	cipher_TLS_RSA_PSK_WITH_CHACHA20_POLY1305_SHA256     uint16 = 0xCCAE
)

// isBadCipher reports whether the cipher is blacklisted by the HTTP/control spec.
// References:
// https://tools.ietf.org/html/rfc7540#appendix-A
// Reject cipher suites from Appendix A.
// "This list includes those cipher suites that do not
// offer an ephemeral key exchange and those that are
// based on the TLS null, stream or block cipher type"
func isBadCipher(cipher uint16) bool {
	switch cipher {
	case cipher_TLS_NULL_WITH_NULL_NULL,
		cipher_TLS_RSA_WITH_NULL_MD5,
		cipher_TLS_RSA_WITH_NULL_SHA,
		cipher_TLS_RSA_EXPORT_WITH_RC4_40_MD5,
		cipher_TLS_RSA_WITH_RC4_128_MD5,
		cipher_TLS_RSA_WITH_RC4_128_SHA,
		cipher_TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5,
		cipher_TLS_RSA_WITH_IDEA_CBC_SHA,
		cipher_TLS_RSA_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_RSA_WITH_DES_CBC_SHA,
		cipher_TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DH_DSS_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_DES_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DH_RSA_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_DES_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_DES_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_DES_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DH_anon_EXPORT_WITH_RC4_40_MD5,
		cipher_TLS_DH_anon_WITH_RC4_128_MD5,
		cipher_TLS_DH_anon_EXPORT_WITH_DES40_CBC_SHA,
		cipher_TLS_DH_anon_WITH_DES_CBC_SHA,
		cipher_TLS_DH_anon_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_KRB5_WITH_DES_CBC_SHA,
		cipher_TLS_KRB5_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_KRB5_WITH_RC4_128_SHA,
		cipher_TLS_KRB5_WITH_IDEA_CBC_SHA,
		cipher_TLS_KRB5_WITH_DES_CBC_MD5,
		cipher_TLS_KRB5_WITH_3DES_EDE_CBC_MD5,
		cipher_TLS_KRB5_WITH_RC4_128_MD5,
		cipher_TLS_KRB5_WITH_IDEA_CBC_MD5,
		cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_SHA,
		cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_SHA,
		cipher_TLS_KRB5_EXPORT_WITH_RC4_40_SHA,
		cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_MD5,
		cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_MD5,
		cipher_TLS_KRB5_EXPORT_WITH_RC4_40_MD5,
		cipher_TLS_PSK_WITH_NULL_SHA,
		cipher_TLS_DHE_PSK_WITH_NULL_SHA,
		cipher_TLS_RSA_PSK_WITH_NULL_SHA,
		cipher_TLS_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA,
		cipher_TLS_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA,
		cipher_TLS_RSA_WITH_NULL_SHA256,
		cipher_TLS_RSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_RSA_WITH_AES_256_CBC_SHA256,
		cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA256,
		cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256,
		cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA256,
		cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA256,
		cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA256,
		cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA256,
		cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA256,
		cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA256,
		cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA,
		cipher_TLS_PSK_WITH_RC4_128_SHA,
		cipher_TLS_PSK_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_PSK_WITH_AES_128_CBC_SHA,
		cipher_TLS_PSK_WITH_AES_256_CBC_SHA,
		cipher_TLS_DHE_PSK_WITH_RC4_128_SHA,
		cipher_TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA,
		cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA,
		cipher_TLS_RSA_PSK_WITH_RC4_128_SHA,
		cipher_TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA,
		cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA,
		cipher_TLS_RSA_WITH_SEED_CBC_SHA,
		cipher_TLS_DH_DSS_WITH_SEED_CBC_SHA,
		cipher_TLS_DH_RSA_WITH_SEED_CBC_SHA,
		cipher_TLS_DHE_DSS_WITH_SEED_CBC_SHA,
		cipher_TLS_DHE_RSA_WITH_SEED_CBC_SHA,
		cipher_TLS_DH_anon_WITH_SEED_CBC_SHA,
		cipher_TLS_RSA_WITH_AES_128_GCM_SHA256,
		cipher_TLS_RSA_WITH_AES_256_GCM_SHA384,
		cipher_TLS_DH_RSA_WITH_AES_128_GCM_SHA256,
		cipher_TLS_DH_RSA_WITH_AES_256_GCM_SHA384,
		cipher_TLS_DH_DSS_WITH_AES_128_GCM_SHA256,
		cipher_TLS_DH_DSS_WITH_AES_256_GCM_SHA384,
		cipher_TLS_DH_anon_WITH_AES_128_GCM_SHA256,
		cipher_TLS_DH_anon_WITH_AES_256_GCM_SHA384,
		cipher_TLS_PSK_WITH_AES_128_GCM_SHA256,
		cipher_TLS_PSK_WITH_AES_256_GCM_SHA384,
		cipher_TLS_RSA_PSK_WITH_AES_128_GCM_SHA256,
		cipher_TLS_RSA_PSK_WITH_AES_256_GCM_SHA384,
		cipher_TLS_PSK_WITH_AES_128_CBC_SHA256,
		cipher_TLS_PSK_WITH_AES_256_CBC_SHA384,
		cipher_TLS_PSK_WITH_NULL_SHA256,
		cipher_TLS_PSK_WITH_NULL_SHA384,
		cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA256,
		cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA384,
		cipher_TLS_DHE_PSK_WITH_NULL_SHA256,
		cipher_TLS_DHE_PSK_WITH_NULL_SHA384,
		cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA256,
		cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA384,
		cipher_TLS_RSA_PSK_WITH_NULL_SHA256,
		cipher_TLS_RSA_PSK_WITH_NULL_SHA384,
		cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA256,
		cipher_TLS_EMPTY_RENEGOTIATION_INFO_SCSV,
		cipher_TLS_ECDH_ECDSA_WITH_NULL_SHA,
		cipher_TLS_ECDH_ECDSA_WITH_RC4_128_SHA,
		cipher_TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_NULL_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDH_RSA_WITH_NULL_SHA,
		cipher_TLS_ECDH_RSA_WITH_RC4_128_SHA,
		cipher_TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDHE_RSA_WITH_NULL_SHA,
		cipher_TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		cipher_TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDH_anon_WITH_NULL_SHA,
		cipher_TLS_ECDH_anon_WITH_RC4_128_SHA,
		cipher_TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDH_anon_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDH_anon_WITH_AES_256_CBC_SHA,
		cipher_TLS_SRP_SHA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_SRP_SHA_RSA_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_SRP_SHA_DSS_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_SRP_SHA_WITH_AES_128_CBC_SHA,
		cipher_TLS_SRP_SHA_RSA_WITH_AES_128_CBC_SHA,
		cipher_TLS_SRP_SHA_DSS_WITH_AES_128_CBC_SHA,
		cipher_TLS_SRP_SHA_WITH_AES_256_CBC_SHA,
		cipher_TLS_SRP_SHA_RSA_WITH_AES_256_CBC_SHA,
		cipher_TLS_SRP_SHA_DSS_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384,
		cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,
		cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256,
		cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384,
		cipher_TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256,
		cipher_TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384,
		cipher_TLS_ECDHE_PSK_WITH_RC4_128_SHA,
		cipher_TLS_ECDHE_PSK_WITH_3DES_EDE_CBC_SHA,
		cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA,
		cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA,
		cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256,
		cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA384,
		cipher_TLS_ECDHE_PSK_WITH_NULL_SHA,
		cipher_TLS_ECDHE_PSK_WITH_NULL_SHA256,
		cipher_TLS_ECDHE_PSK_WITH_NULL_SHA384,
		cipher_TLS_RSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_RSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DH_DSS_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DH_DSS_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DH_RSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DH_RSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DHE_DSS_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DHE_DSS_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DHE_RSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DHE_RSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DH_anon_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DH_anon_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_ECDHE_ECDSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_ECDSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_ECDHE_RSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_RSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_ECDH_RSA_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_ECDH_RSA_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_RSA_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_RSA_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_DH_RSA_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_DH_RSA_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_DH_DSS_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_DH_DSS_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_DH_anon_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_DH_anon_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_ECDH_RSA_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_ECDH_RSA_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_PSK_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_PSK_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_DHE_PSK_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_DHE_PSK_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_RSA_PSK_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_RSA_PSK_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_PSK_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_PSK_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_RSA_PSK_WITH_ARIA_128_GCM_SHA256,
		cipher_TLS_RSA_PSK_WITH_ARIA_256_GCM_SHA384,
		cipher_TLS_ECDHE_PSK_WITH_ARIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_PSK_WITH_ARIA_256_CBC_SHA384,
		cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_RSA_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_RSA_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_DH_RSA_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_DH_DSS_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_DH_anon_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_DH_anon_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_PSK_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_PSK_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_GCM_SHA256,
		cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_GCM_SHA384,
		cipher_TLS_PSK_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_PSK_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_DHE_PSK_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_DHE_PSK_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_128_CBC_SHA256,
		cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_256_CBC_SHA384,
		cipher_TLS_RSA_WITH_AES_128_CCM,
		cipher_TLS_RSA_WITH_AES_256_CCM,
		cipher_TLS_RSA_WITH_AES_128_CCM_8,
		cipher_TLS_RSA_WITH_AES_256_CCM_8,
		cipher_TLS_PSK_WITH_AES_128_CCM,
		cipher_TLS_PSK_WITH_AES_256_CCM,
		cipher_TLS_PSK_WITH_AES_128_CCM_8,
		cipher_TLS_PSK_WITH_AES_256_CCM_8:
		return true
	default:
		return false
	}
}
