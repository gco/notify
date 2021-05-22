// Copyright (c) 2014-2015 The Notify Authors. All rights reserved.
// Use of this source code is governed by the MIT license that can be
// found in the LICENSE file.

package notify

// #include <port.h>
// #include <stdio.h>
// #include <stdlib.h>
import "C"

const (
	fileAccess     = Event(C.FILE_ACCESS)
	fileModified   = Event(C.FILE_MODIFIED)
	fileAttrib     = Event(C.FILE_ATTRIB)
	fileDelete     = Event(C.FILE_DELETE)
	fileRenameTo   = Event(C.FILE_RENAME_TO)
	fileRenameFrom = Event(C.FILE_RENAME_FROM)
	fileNoFollow   = Event(C.FILE_NOFOLLOW)
	// Use the illumos constant to minimize code changes, but this
	// event will not be generated on Solaris
	fileTrunc      = Event(0x00100000)
	unmounted      = Event(C.UNMOUNTED)
	mountedOver    = Event(C.MOUNTEDOVER)
)
