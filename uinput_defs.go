package uinput

import "syscall"

const uinputDevPath = "/dev/uinput"

const uinputMaxNameSize = 80

// ioctls from uinput.h
const (
	uiDevCreate  = 0x5501
	uiDevDestroy = 0x5502
	uiDevSetup   = 0x405c5503
	uiAbsSetup   = 0x401c5504
	uiSetEvBit   = 0x40045564
	uiSetKeyBit  = 0x40045565
	uiSetRelBit  = 0x40045566
	uiSetAbsBit  = 0x40045567
	uiSetMscBit  = 0x40045568
	uiSetLedBit  = 0x40045569
	uiSetSndBit  = 0x4004556a
	uiSetFfBit   = 0x4004556b
	uiSetPhys    = 0x4004556c
	uiSetSwBit   = 0x4004556d
	uiSetPropBit = 0x4004556e
)

// IDs from input.h
const (
	IDBus     = 0
	IDVendor  = 1
	IDProduct = 2
	IDVersion = 3

	BusPCI       = 0x01
	BusISANPN    = 0x02
	BusUSB       = 0x03
	BusHIL       = 0x04
	BusBluetooth = 0x05
	BusVirtual   = 0x06
)

// Event types from input-event-codes.h
const (
	EvSyn = 0x00
	EvKey = 0x01
	EvRel = 0x02
)

// Synchronization events from input-event-codes.h
const (
	SynReport   = 0
	SynConfig   = 1
	SynMtReport = 2
	SynDropped  = 3
	SynMax      = 0xf
	SynCnt      = SynMax + 1
)

// Keys and buttons from input-event-codes.h
const (
	KeyReserved   = 0
	KeyEsc        = 1
	Key1          = 2
	Key2          = 3
	Key3          = 4
	Key4          = 5
	Key5          = 6
	Key6          = 7
	Key7          = 8
	Key8          = 9
	Key9          = 10
	Key0          = 11
	KeyMinus      = 12
	KeyEqual      = 13
	KeyBackspace  = 14
	KeyTab        = 15
	KeyQ          = 16
	KeyW          = 17
	KeyE          = 18
	KeyR          = 19
	KeyT          = 20
	KeyY          = 21
	KeyU          = 22
	KeyI          = 23
	KeyO          = 24
	KeyP          = 25
	KeyLeftBrace  = 26
	KeyRightBrace = 27
	KeyEnter      = 28
	KeyLeftCtrl   = 29
	KeyA          = 30
	KeyS          = 31
	KeyD          = 32
	KeyF          = 33
	KeyG          = 34
	KeyH          = 35
	KeyJ          = 36
	KeyK          = 37
	KeyL          = 38
	KeySemicolon  = 39
	KeyApostrophe = 40
	KeyGrave      = 41
	KeyLeftShift  = 42
	KeyBackslash  = 43
	KeyZ          = 44
	KeyX          = 45
	KeyC          = 46
	KeyV          = 47
	KeyB          = 48
	KeyN          = 49
	KeyM          = 50
	KeyComma      = 51
	KeyDot        = 52
	KeySlash      = 53
	KeyRightShift = 54
	KeyKpAsterisk = 55
	KeyLeftAlt    = 56
	KeySpace      = 57
	KeyCapsLock   = 58
	KeyF1         = 59
	KeyF2         = 60
	KeyF3         = 61
	KeyF4         = 62
	KeyF5         = 63
	KeyF6         = 64
	KeyF7         = 65
	KeyF8         = 66
	KeyF9         = 67
	KeyF10        = 68
	KeyNumLock    = 69
	KeyScrollLock = 70
	KeyKp7        = 71
	KeyKp8        = 72
	KeyKp9        = 73
	KeyKpMinus    = 74
	KeyKp4        = 75
	KeyKp5        = 76
	KeyKp6        = 77
	KeyKpPlus     = 78
	KeyKp1        = 79
	KeyKp2        = 80
	KeyKp3        = 81
	KeyKp0        = 82
	KeyKpDot      = 83

	KeyZenkakuHankaku   = 85
	Key102Nd            = 86
	KeyF11              = 87
	KeyF12              = 88
	KeyRo               = 89
	KeyKatakana         = 90
	KeyHiragana         = 91
	KeyHenkan           = 92
	KeyKatakanaHiragana = 93
	KeyMuhenkan         = 94
	KeyKpJpComma        = 95
	KeyKpEnter          = 96
	KeyRightCtrl        = 97
	KeyKpslash          = 98
	KeySysrq            = 99
	KeyRightAlt         = 100
	KeyLineFeed         = 101
	KeyHome             = 102
	KeyUp               = 103
	KeyPageUp           = 104
	KeyLeft             = 105
	KeyRight            = 106
	KeyEnd              = 107
	KeyDown             = 108
	KeyPageDown         = 109
	KeyInsert           = 110
	KeyDelete           = 111
	KeyMacro            = 112
	KeyMute             = 113
	KeyVolumeDown       = 114
	KeyVolumeUp         = 115
	KeyPower            = 116 /* SC System Power Down */
	KeyKpEqual          = 117
	KeyKpPlusMinus      = 118
	KeyPause            = 119
	KeyScale            = 120 /* AL Compiz Scale (Expose) */

	KeyKpComma   = 121
	KeyHangeul   = 122
	KeyHanja     = KeyHangeul
	KeyYen       = 124
	KeyLeftMeta  = 125
	KeyRightMeta = 126
	KeyCompose   = 127

	KeyStop          = 128 /* AC Stop */
	KeyAgain         = 129
	KeyProps         = 130 /* AC Properties */
	KeyUndo          = 131 /* AC Undo */
	KeyFront         = 132
	KeyCopy          = 133 /* AC Copy */
	KeyOpen          = 134 /* AC Open */
	KeyPaste         = 135 /* AC Paste */
	KeyFind          = 136 /* AC Search */
	KeyCut           = 137 /* AC Cut */
	KeyHelp          = 138 /* AL Integrated Help Center */
	KeyMenu          = 139 /* Menu (show menu) */
	KeyCalc          = 140 /* AlCalculator */
	KeySetup         = 141
	KeySleep         = 142 /* ScC System Sleep */
	KeyWakeUp        = 143 /* System Wake Up */
	KeyFile          = 144 /* AL Local Machine Browser */
	KeySendFile      = 145
	KeyDeleteFile    = 146
	KeyXfer          = 147
	KeyProg1         = 148
	KeyProg2         = 149
	KeyWww           = 150 /* AL Internet Browser */
	KeyMsDos         = 151
	KeyCoffee        = 152 /* AL Terminal Lock/Screensaver */
	KeyScreenLock    = KeyCoffee
	KeyRotateDisplay = 153 /* Display orientation for e.g. tablets */
	KeyRotation      = KeyRotateDisplay
	KeyCycleWindows  = 154
	KeyMail          = 155
	KeyBookmarks     = 156 /* AC Bookmarks */
	KeyComputer      = 157
	KeyBack          = 158 /* AC Back */
	KeyForward       = 159 /* AC Forward */
	KeyCloseCD       = 160
	KeyEjectCD       = 161
	KeyEjectCloseCd  = 162
	KeyNextSong      = 163
	KeyPlayPause     = 164
	KeyPreviousSong  = 165
	KeyStopCD        = 166
	KeyRecord        = 167
	KeyRewind        = 168
	KeyPhone         = 169 /* Media Select Telephone */
	KeyISO           = 170
	KeyConfig        = 171 /* AL Consumer Control Configuration */
	KeyHomePage      = 172 /* AC Home */
	KeyRefresh       = 173 /* AC Refresh */
	KeyExit          = 174 /* AC Exit */
	KeyMove          = 175
	KeyEdit          = 176
	KeyScrollUp      = 177
	KeyScrollDown    = 178
	KeyKpLeftParen   = 179
	KeyKpRightParen  = 180
	KeyNew           = 181 /* AC New */
	KeyRedo          = 182 /* AC Redo/Repeat */

	KeyF13 = 183
	KeyF14 = 184
	KeyF15 = 185
	KeyF16 = 186
	KeyF17 = 187
	KeyF18 = 188
	KeyF19 = 189
	KeyF20 = 190
	KeyF21 = 191
	KeyF22 = 192
	KeyF23 = 193
	KeyF24 = 194

	KeyPlayCd         = 200
	KeyPauseCd        = 201
	KeyProg3          = 202
	KeyProg4          = 203
	KeyDashBoard      = 204 /* AL Dashboard */
	KeySuspend        = 205
	KeyClose          = 206 /* AC Close */
	KeyPlay           = 207
	KeyFastForward    = 208
	KeyBassBoost      = 209
	KeyPrint          = 210 /* AC Print */
	KeyHp             = 211
	KeyCamera         = 212
	KeySound          = 213
	KeyQuestion       = 214
	KeyEmail          = 215
	KeyChat           = 216
	KeySearch         = 217
	KeyConnect        = 218
	KeyFinance        = 219 /* AL Checkbook/Finance */
	KeySport          = 220
	KeyShop           = 221
	KeyAlterase       = 222
	KeyCancel         = 223 /* AC Cancel */
	KeyBrightnessDown = 224
	KeyBrightnessUp   = 225
	KeyMedia          = 226

	KeySwitchVideoMode = 227 /* Cycle between available video outputs (Monitor/LCD/TV-out/etc) */

	KeyKbDillumToggle = 228
	KeyKbDillumDown   = 229
	KeyKbDillumUp     = 230

	KeySend        = 231 /* AC Send */
	KeyReply       = 232 /* AC Reply */
	KeyForwardMail = 233 /* AC Forward Msg */
	KeySave        = 234 /* AC Save */
	KeyDocuments   = 235

	KeyBattery = 236

	KeyBluetooth = 237
	KeyWlan      = 238
	KeyUwb       = 239

	KeyUnknown = 240

	KeyVideoNext       = 241 /* drive next video source */
	KeyVideoPrev       = 242 /* drive previous video source */
	KeyBrightnessCycle = 243 /* brightness up, after max is min */
	KeyBrightnessAuto  = 244 /* Set Auto Brightness: manual brightness control is off, rely on ambient */

	KeyBrightnessZero = KeyBrightnessAuto
	KeyDisplayOff     = 245 /* display device to off state */

	KeyWwan   = 246 /* Wireless WAN (LTE, UMTS, GSM, etc.) */
	KeyWimax  = KeyWwan
	KeyRfKill = 247 /* Key that controls all radios */

	KeyMicMute = 248 /* Mute / unmute the microphone */

	KeyMax = 0x2ff
	KeyCnt = KeyMax + 1
)

// struct uinput_setup from uinput.h
type uinputSetup struct {
	id           inputID
	name         [uinputMaxNameSize]byte
	ffEffectsMax uint32
}

// IOCTLs (0x00 - 0x7f) from input.h
type inputID struct {
	busType uint16
	vendor  uint16
	product uint16
	version uint16
}

// The event structure from input.h
type inputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}
