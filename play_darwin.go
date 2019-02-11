// +build darwin

package soundplayer

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework AVFoundation -framework CoreMedia

#import <AVFoundation/AVFoundation.h>

static AVPlayer* p;

static int _play(const char* filename) {
    @autoreleasepool {
        NSURL* u = [NSURL fileURLWithPath:[NSString stringWithUTF8String:filename]];

        p = [AVPlayer playerWithURL:u];
        [p play];

        NSTimeInterval played = 0.;
        while (1) {
            NSTimeInterval t = CMTimeGetSeconds([p currentTime]);
            if (t > 0. && t == played) {
                break;
            }
            played = t;

            [[NSRunLoop currentRunLoop] runMode:NSDefaultRunLoopMode
                                    beforeDate:[[NSDate date] dateByAddingTimeInterval:.1]];

        }

        return 0;
    }
}

static void _pause() {
    if (p) {
        [p pause];
    }
}

static void _stop() {
    _pause();
    p = nil;
}

*/
import "C"
import "unsafe"

import (
	"errors"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func Play(filename string) error {
	c := C.CString(filename)
	defer C.free(unsafe.Pointer(c))

	if r := C._play(c); r != 0 {
		return errors.New("play error")
	}
	return nil
}

func Pause() {
	C._pause()
}

func Stop() {
	C._stop()
}
