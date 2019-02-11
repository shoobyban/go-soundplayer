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
            if (!p) {
                break;
            }
            //NSTimeInterval t = CMTimeGetSeconds([p currentTime]);
            // if (t > 0. && t == played) {
            //     return 2;
            // }
            //played = t;

            [[NSRunLoop currentRunLoop] runMode:NSDefaultRunLoopMode
                beforeDate:[[NSDate date] dateByAddingTimeInterval:.1]];

            [[NSNotificationCenter defaultCenter]
                addObserverForName:AVPlayerItemDidPlayToEndTimeNotification
                object:nil
                queue:nil
                usingBlock: ^ (NSNotification * note) {
                    [p pause];
                    p = nil;
                }];
        }

        return 0;
    }
}

static void _resume() {
    [p play];
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

static void _seek(float secs) {
    int32_t timeScale = p.currentItem.asset.duration.timescale;
    [p seekToTime:CMTimeMakeWithSeconds(secs,timeScale)];
}

static float _getsecs() {
    return CMTimeGetSeconds([p currentTime]);
}

static float _getduration() {
    return CMTimeGetSeconds([p currentItem].duration);
}

static int32_t _getTimeScale() {
    return p.currentItem.asset.duration.timescale;
}

*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

func Play(filename string) error {
	c := C.CString(filename)
	defer C.free(unsafe.Pointer(c))

	if r := C._play(c); r != 0 {
		return fmt.Errorf("play error %d", r)
	}
	return nil
}

func Pause() {
	C._pause()
}

func Stop() {
	C._stop()
}

func Seek(seconds float64) {
	C._seek(C.float(seconds))
}

func GetSecs() float64 {
	return float64(C._getsecs())
}

func GetDuration() float64 {
	return float64(C._getduration())
}

func TimeScale() int32 {
	return int32(C._getTimeScale())
}

func Resume() {
	C._resume()
}
