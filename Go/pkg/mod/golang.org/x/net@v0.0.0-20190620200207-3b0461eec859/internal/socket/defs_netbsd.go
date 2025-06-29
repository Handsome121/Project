// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// +godefs map struct_in_addr [memory]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

package socket

/*
#include <sys/socket.h>

#include <netinet/in.h>
*/
import "C"

type iovec C.struct_iovec

type msghdr C.struct_msghdr

type mmsghdr C.struct_mmsghdr

type cmsghdr C.struct_cmsghdr

type sockaddrInet C.struct_sockaddr_in

type sockaddrInet6 C.struct_sockaddr_in6

const (
	sizeofIovec   = C.sizeof_struct_iovec
	sizeofMsghdr  = C.sizeof_struct_msghdr
	sizeofMmsghdr = C.sizeof_struct_mmsghdr
	sizeofCmsghdr = C.sizeof_struct_cmsghdr

	sizeofSockaddrInet  = C.sizeof_struct_sockaddr_in
	sizeofSockaddrInet6 = C.sizeof_struct_sockaddr_in6
)
