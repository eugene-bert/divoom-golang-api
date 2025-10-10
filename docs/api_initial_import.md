command format
All command formats only support JSON format
the body is below:
{
“Command”:”Channel/SetCloudIndex”,
“Index”:2
}

DIY Net Data Clock
We will introduce DataParsingRules and InputUrlAddress in DIY Net Data Clock。
1： InputUrlAddress will be the URL address of the data request.The data format is JSON format.
eg: InputUrlAddress is “http://app.divoom-gz.com/User/GetUserData?DeviceId=300043458“.
The data returned by the URL request is as follows：
{
“ReturnCode”: 0,
“ReturnMessage”: “”,
“Nickname”: “Divoom-Developer”,
“HeadId”: “group1/M00/12/9E/eEwpPWLpNhWEI97CAAAAAPpA5X01371922”,
“Level”: 20,
“Score”: 1862,
“LikeCnt”: 0,
“FansCnt”: 44,
“DeviceId”: 300043458
}

2：DataParsingRules are rules for extracting display data from URL returned data。
DataParsingRules will be a string：
eg:object1,object1-2,object1-3,n:dispNumber;
object1,object1-2,object1-3,s:dispString;
object1,object1-2,object1-3,object4,n:dispNumber;
n:dispNumber;
s:dispString;

Different objects are separated by ‘,’, and the last level of data is represented by ‘n’ for numbers and’s’ for strings. The rule is case sensitive.

1):the Rule of the divoom respone data:
{
“ReturnCode”: 0,
“ReturnMessage”: “”,
“Nickname”: “DivoomTest”,
“HeadId”: “group1/M00/05/2A/L1ghbmGdopOEP45jAAAAAIB05eo4227145”,
“DispData”: {
“LikeCnt”: 519947,
“FansCnt”: 14654,
“UserInfo”: {
“Level”: 21,
“Score”: 603141
}
}
}
The DataParsingRules of “Nickname” is “s:Nickname”
The DataParsingRules of “LikeCnt” is “dispData,n:LikeCnt”
The DataParsingRules of “FansCnt” is “dispData,n:FansCnt”
The DataParsingRules of “Level” is “dispData,UserInfo,n:Level”
The DataParsingRules of “Score” is “dispData,UserInfo,n:Score”

2):the Rule of the divoom respone data:
{
“ReturnCode”: 0,
“ReturnMessage”: “”,
“Nickname”: “Divoom-Developer”,
“HeadId”: “group1/M00/12/9E/eEwpPWLpNhWEI97CAAAAAPpA5X01371922”,
“Level”: 20,
“Score”: 1862,
“LikeCnt”: 0,
“FansCnt”: 44,
“DeviceId”: 300043458
}
The DataParsingRules of “Nickname” is “s:Nickname”
The DataParsingRules of “LikeCnt” is “n:LikeCnt”
The DataParsingRules of “FansCnt” is “n:FansCnt”
The DataParsingRules of “Level” is “n:Level”
The DataParsingRules of “Score” is “n:Score”

1):the Rule of the divoom respone data:
{
“ReturnCode”: 0,
“ReturnMessage”: “”,
“Nickname”: “DivoomTest”,
“HeadId”: “group1/M00/05/2A/L1ghbmGdopOEP45jAAAAAIB05eo4227145”,
“RetData”:
{
“UserInfo”:
{
“DispData”: {
“LikeCnt”: 519947,
“FansCnt”: 14654,
“UserInfo”: {
“Level”: 21,
“Score”: 603141
}
}
}
}
}
The DataParsingRules of “Nickname” is “s:Nickname”
The DataParsingRules of “LikeCnt” is “RetData,UserInfo,dispData,n:LikeCnt”
The DataParsingRules of “FansCnt” is “RetData,UserInfo,dispData,n:FansCnt”
The DataParsingRules of “Level” is “RetData,UserInfo,dispData,UserInfo,n:Level”
The DataParsingRules of “Score” is “RetData,UserInfo,dispData,UserInfo,n:Score”

Find device
command description
get the device list in local network.
Request URL
https://app.divoom-gz.com/Device/ReturnSameLANDevice
Request method
POST
para
Name DATA TYPE DESCRIPTION
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"DeviceList": [
{
"DeviceName": "Pixoo64",
"DeviceId": 300000020,
"DevicePrivateIP": "10.0.0.100",
"DeviceMac": "a8032aff46b1"
},
...

    ]

}
Return value
Name DATA TYPE DESCRIPTION
DeviceList ARRAY 0
DeviceName STRING the device name
DeviceId NUMBER the device ID
DevicePrivateIP STRING the device IP
DeviceMac STRING the device mac

PIXOO64

System reboots
command description
System reboots!
command
Device/SysReboot
request
post
para
Name DATA TYPE DESCRIPTION
command string “Device/SysReboot”

Welcome to the Divoom API

command description
select dial type.
Request URL
https://app.divoom-gz.com/Channel/GetDialType
Request method
POST
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"DialTypeList": [
"Social",
"normal",
"financial",
"Game",
"HOLIDAYS",
"TOOLS",
"DESIGN-64"
]
}
Return value
Name DATA TYPE DESCRIPTION
ReturnCode number 0 : Success

Welcome to the Divoom API

command description
select dial List.
Request URL
https://app.divoom-gz.com/Channel/GetDialList
Request method
POST
para
Name DATA TYPE DESCRIPTION
DialType string dial type
Page number the number of pages,for example: 1, Notes: 30 per page
request example
{
"DialType":"Social",
"Page":1
}
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"TotalNum": 100,
"DialList": [
{
"ClockId": 10,
"Name": "Classic Digital Clock"
},
{
"ClockId": 12,
"Name": "US Stock - 2"
},
{
"ClockId": 24,
"Name": "Twitter Post"
},
{
"ClockId": 38,
"Name": "YouTube Account"
},
{
"ClockId": 40,
"Name": "YouTube Video"
},
{
"ClockId": 42,
"Name": "Twitch Account"
},
{
"ClockId": 44,
"Name": "Twitch Stream"
},
{
"ClockId": 46,
"Name": "Bilibili Account"
},
{
"ClockId": 48,
"Name": "Bilibili-works"
},
...
]
}
Return value
Name DATA TYPE DESCRIPTION
ReturnCode number 0 : Success
TotalNum number Total number of dials
ClockId number Dial ID
Name string Dial name

Welcome to the Divoom API

command description
Get working Faces id .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/GetClockInfo
respone example
{
"ClockId": 12,
"Brightness":100
}
Return value
Name DATA TYPE DESCRIPTION
ClockId number face id
Brightness number the device’s brightness

channel control

elect channel
Welcome to the Divoom API

command description
select chnanel, device will move to the selected channel.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetIndex
SelectIndex number channel id 0~3 ；0:Faces;1:Cloud Channdle;2:Visualizer;3:Custom;4:black screen
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

control custom channel
Welcome to the Divoom API

command description
select Custom, device will work on the custom channel.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetCustomPageIndex
CustomPageIndex number custom index 0~２
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Visualizer Channel
Welcome to the Divoom API

command description
select Visualizer, device will work on visualizer channel.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetEqPosition
EqPosition number index :start from 0
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Cloud Channel
Welcome to the Divoom API

command description
select chnanel, device will move to the selected channel.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/CloudIndex
Index number cloud index 0~2; 0:Recommend gallery;1:Favourite;2:Subscribe artist;3:album
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

get cureent channel
Welcome to the Divoom API

command description
get current chnanel.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/GetIndex
respone example
{
"SelectIndex": 1
}
Return value
Name DATA TYPE DESCRIPTION
SelectIndex number channel id 0~3 ；0:Faces;1:Cloud Channdle;2:Visualizer;3:Custom

et brightness
Welcome to the Divoom API

command description
it will set the device brightness.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetBrightness
Brightness Number 0~100
request example
{
"Command":"Channel/SetBrightness",
"Brightness":100
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

get all setting
Welcome to the Divoom API

command description
it will get all settings and be ok at 90104.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/GetAllConf
request example
{
"Command":"Channel/GetAllConf"
}
respone example
Name DATA TYPE DESCRIPTION
Brightness Number 0~100, the system brightness
RotationFlag Number 1: it will switch to display faces and gifs
ClockTime Number the time of displaying faces and it will be active with RotationFlag = 1
GalleryTime Number the time of displaying gifs and it will be active with RotationFlag = 1
SingleGalleyTime Number the time of displaying each gif
PowerOnChannelId Number device will display the channle when it powers on
GalleryShowTimeFlag Number 1: it will display time at right-top ;
CurClockId Number the running’s face id
Time24Flag Number the display hour flag
TemperatureMode Number the display temperature flag
GyrateAngle Number the rotation angle:0: normal;1:90;2:180;3:270
MirrorFlag Number the mirror mode
LightSwitch Number the screen switch
{
"error_code": 0,
"Brightness":100,
"RotationFlag":1,
"ClockTime":60,
"GalleryTime":60,
"SingleGalleyTime":5,
"PowerOnChannelId":1,
"GalleryShowTimeFlag":1,
"CurClockId":1,
"Time24Flag":1,
"TemperatureMode":1,
"GyrateAngle":1,
"MirrorFlag":1,
"LightSwitch":1
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Weather area setting
Welcome to the Divoom API

command description
it will set the Longitude and latitude which get weather information. All data comes from https://openweathermap.org/.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Sys/LogAndLat
Longitude string value
Latitude string value
request example
{
"Command":"Sys/LogAndLat",
"Longitude": "30.29",
"Latitude":"20.58"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set Time Zone
Welcome to the Divoom API

command description
it will set the time zone .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Sys/TimeZone
TimeZoneValue string time zone value
request example
{
"Command":"Sys/TimeZone",
"TimeZoneValue": "GMT-5"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

system Time
Welcome to the Divoom API

command description
it will set the system time when the device powers on .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetUTC
Utc number utc time
request example
{
"Command":"Device/SetUTC",
"Utc": 1672416000
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Screen switch
Welcome to the Divoom API

command description
it will switch the screen .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/OnOffScreen
OnOff number 1:on; 0:off
request example
{
"Command":"Channel/OnOffScreen",
"OnOff": 1
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Get Device Time
Welcome to the Divoom API

command description
it will get the device system time.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/GetDeviceTime
request example
{
"Command":"Device/GetDeviceTime"
}
respone example
{
"error_code": 0,
"UTCTime":1647200428,
"LocalTime":"2022-03-14 03:40:28"
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0
UTCTime number 1647200428
LocalTime string “2022-03-14 03:40:28”

Set temperature mode
Welcome to the Divoom API

command description
it will set the temperature mode with Fahrenheit or Celsius.
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetDisTempMode
Mode number 0:Celsius;1:Fahrenheit
request example
{
"Command":"Device/SetDisTempMode",
"Mode":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set Rotation angle
Welcome to the Divoom API

command description
it will set the screen Rotation angle .
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetScreenRotationAngle
Mode number 0:normal,1:90;2:180;3:270
request example
{
"Command":"Device/SetScreenRotationAngle",
"Mode":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set Mirror mode
Welcome to the Divoom API

command description
it will set the screen mirror mode.
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetMirrorMode
Mode number 0:disable;1:enalbe
request example
{
"Command":"Device/SetMirrorMode",
"Mode":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set hour mode
Welcome to the Divoom API

command description
it will set the screen hour24 mode.
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetTime24Flag
Mode number 1:24－hour;0:12-hour
request example
{
"Command":"Device/SetTime24Flag",
"Mode":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set High Ligit mode
Welcome to the Divoom API

command description
it will set the screen high light mode.
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetHighLightMode
Mode number 0:close;1:open
request example
{
"Command":"Device/SetHighLightMode",
"Mode":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set White Balance
Welcome to the Divoom API

command description
it will set the screen White Balance.
it won’t be saved and reset when the device power off.
it will be active after 90107.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SetWhiteBalance
RValue number 100 ;(0~100)
GValue number 100 ;(0~100)
BValue number 100 ;(0~100)
request example
{
"Command":"Device/SetWhiteBalance",
"RValue":100,
"GValue":100,
"BValue":100
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Get the Weather of the device
Welcome to the Divoom API

command description
it will get the display weather information of the device.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/GetWeatherInfo
request example
{
"Command":"Device/GetWeatherInfo"
}
respone example
{
"error_code": 0,
"Weather": "Cloudy",
"CurTemp": 33.680000,
"MinTemp": 31.850000,
"MaxTemp": 33.680000,
"Pressure": 1006,
"Humidity": 50,
"Visibility": 10000,
"WindSpeed": 2.540000
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0
Weather string “Sunny”,”Cloudy”,”Rainy”,”Rainy”,”Frog”
CurTemp number the current temperature
MinTemp number the minimum temperature
MaxTemp number the maximum temperature
Pressure number current pressure
Humidity number current humidity
Visibility number current visibility
WindSpeed number current wind speed m/s

SetGalleryTime
Welcome to the Divoom API

command description
it will set the time zone .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetSubscribeGalleryTime
SingleGalleyTime Number the play time
request example
{
"Command":"Channel/SetSubscribeGalleryTime",
"SingleGalleyTime": 10
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set Subscribe Gallery attribute
Welcome to the Divoom API

command description
it will set the Subscribe Gallery attribute
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Channel/SetSubscribeGalleryTime
SingleGalleyTime number The playback time of each animation
GalleryShowTimeFlag number 1: display time; 0:none
request example
{
"Command":"Channel/SetSubscribeGalleryTime",
"SingleGalleyTime":0,
"GalleryShowTimeFlag":0
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set countdown tool
Welcome to the Divoom API

command description
it will contol the the countdown tool .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Tools/SetTimer
Minute number the countdown’s minute
Second number the countdown’s second
Status number 1: start; 0: stop
request example
{
"Command":"Tools/SetTimer",
"Minute": 1,
"Second": 0,
"Status": 1
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set stopwatch tool
Welcome to the Divoom API

command description
it will contol the the stopwatch tool .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Tools/SetStopWatch
Status number 2:reset;1: start; 0: stop
request example
{
"Command":"Tools/SetStopWatch",
"Status": 1
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set scoreboard tool
Welcome to the Divoom API

command description
it will contol the the scoreboard tool .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Tools/SetScoreBoard
BlueScore number the blue score 0~999
RedScore number the red score 0~999
request example
{
"Command":"Tools/SetScoreBoard",
"BlueScore": 100,
"RedScore": 79
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Set noise tool
Welcome to the Divoom API

command description
it will contol the the noise tool .
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Tools/SetNoiseStatus
NoiseStatus number 1:start; 0:stop
request example
{
"Command":"Tools/SetNoiseStatus",
"NoiseStatus": 1
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

play gif
Welcome to the Divoom API

command description
play gif file, the command can select the net gif file. the gif files only support the size (1616 ,32 32 ,64 \* 64).
the command will be implemented after the 90096 version.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/PlayTFGif
FileType Number 2:play net file; other is err
FileName Number 2:net file address;
request example
{
"Command":"Device/PlayTFGif",
"FileType":2,
"FileName":"http://f.divoom-gz.com/64_64.gif"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Get sending animation PicId
Welcome to the Divoom API

command description
get the PicId which the command “Draw/SendHttpGif”。
It will return the PicId , it’s value is the previous gif id plus 1, the command will be implemented after the 90095 version.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/GetHttpGifId
respone example
{
"error_code": 0,
"PicId":100
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0
PicId number 100

Reset sending animation PicId
Welcome to the Divoom API

command description
it will reset gif id , “Send animation” will start from PicID=1.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/ResetHttpGifId
request example
{
"Command":"Draw/ResetHttpGifId"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Send animation
Welcome to the Divoom API

command description
send animation to device, and device will loop animation。
This method only accepts one picture of animation at a time, and the picture format must be based on Base64 encoded RGB data.
If the animation is composed of N pictures, it will be sent in N times, one picture data will be sent each time with the picture offset。
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/SendHttpGif
PicNum number the include single pictures of the animation and smaller than 60
PicWidth number the pixels of the animation, and only support the 16,32,64
PicOffset number the picture offset start from 0. eg:0,1,2,3,4,PicNum-1
PicID number the animation ID, every animation must have unique ID and auto increase,It’s getting bigger and start with 1, example: the current gif id is 100, and then next gif’s id should be greater than or equal to 101
PicSpeed number the animation speed, it bases on ms
PicData string the picutre Base64 encoded RGB data, The RGB data is left to right and up to down
request example
we will send the gif with two frames, we will send two times,the below are two packets, the “PicOffset=0” packet should be sended first， and the “PicOffset=1” should be sended later.

the first packet:
{
"Command":"Draw/SendHttpGif",
"PicNum":2,
"PicWidth":64,
"PicOffset":0,
"PicID":3,
"PicSpeed":100,
"PicData":"AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAQ+hAQ+hABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAQ+hAQ+hABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwABngABngAAIpABdwABdwABngABngABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngAAIpAAIpABngABngAQ+hAQ+hABdwABdwAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwABngABngAAIpABdwABdwABngABngABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngAAIpAAIpABngABngAQ+hAQ+hABdwABdwAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GABngABngAAIpAA9GAA9GABngABngAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hABngABngAQ+hAQ+hABdwABdwAA9GAA9GAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GABngABngAAIpAA9GAA9GABngABngAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hABngABngAQ+hAQ+hABdwABdwAA9GAA9GAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABngABngABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABngABngABdwABdwAA9GAA9GAAIpNThlAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABngABngAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpNThlAAIpAAIpAAIpAA9GAA9GABngABngAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAQ+hAQ+hABdwABdwAA9GAA9GABdwABdwAQ+hAQ+hABngABngAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAQ+hAQ+hABdwABdwAA9GAA9GABdwABdwAQ+hAQ+hABngABngAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpAAIpAAIpAQ+hAQ+hAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAAIpAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAA9GAA9GABdwABdwAQ+hAQ+hAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAAIpAQ+hAQ+hAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAAIpAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAA9GAA9GABdwABdwAQ+hAQ+hAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAAIpABdwABdwAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABdwABdwABdwABdwABdwABdwAAIpABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpABdwABdwAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABdwABdwABdwABdwABdwABdwAAIpABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAA9GAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAA9GAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////AAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbdnNzdnNzREJCREJCREJCREJCREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNz////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbdnNzdnNzREJCREJCREJCREJCREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNz////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////////////////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzREJCREJC////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////////////////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzREJCREJC////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz/9vb/9vbAAIpAAIp////////dnNzdnNzREJCREJCKCgoKCgo////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz/9vb/9vbAAIpAAIp////////dnNzdnNzREJCREJCKCgoKCgo////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////REJCREJCREJCREJCREJCREJCREJCREJC////////AAIpAAIp////////////////////////////////////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////REJCREJCREJCREJCREJCREJCREJCREJC////////AAIpAAIp////////////////////////////////////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////KCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////AAIpAAIpdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz////////dnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////KCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////AAIpAAIpdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz////////dnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpREJCREJCREJCREJCREJCREJCREJCREJC////////REJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpREJCREJCREJCREJCREJCREJCREJCREJC////////REJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNz/9vb/9vb/////////////////9vb/9vbdnNzdnNzAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////KCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNz/9vb/9vb/////////////////9vb/9vbdnNzdnNzAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////KCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNzREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNzREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpKCgoKCgoREJCREJCREJCREJCREJCREJCREJCREJCKCgoKCgoAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoREJCREJCREJCREJCREJCREJCREJCREJCKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp"
}

the second packet
{
"Command":"Draw/SendHttpGif",
"PicNum":2,
"PicWidth":64,
"PicOffset":1,
"PicID":3,
"PicSpeed":100,
"PicData":"AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAQ+hAQ+hABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAQ+hAQ+hABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwABngABngAAIpABdwABdwABngABngABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngAAIpAAIpABngABngAQ+hAQ+hABdwABdwAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwABngABngAAIpABdwABdwABngABngABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngAAIpAAIpABngABngAQ+hAQ+hABdwABdwAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpABngABngAQ+hAQ+hABdwABdwABdwABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GABngABngAAIpAA9GAA9GABngABngAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hABngABngAQ+hAQ+hABdwABdwAA9GAA9GAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GABngABngAAIpAA9GAA9GABngABngAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hABngABngAQ+hAQ+hABdwABdwAA9GAA9GAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpABngABngABdwABdwAA9GAA9GAA9GAA9GAA9GABdwABdwABngABngAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABngABngABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpAAIpAAIpABngABngABngABngABngABngABngABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABngABngABdwABdwAA9GAA9GAAIpNThlAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpABngABngAA9GAA9GAAEaAAEaAAEaAAEaAAEaAA9GAA9GABngABngAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABngABngAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpNThlAAIpAAIpAAIpAA9GAA9GABngABngAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngABdwABdwABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABngABngAAIpAAIpAAIpAAIpAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpABngABngAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaAAEaABngABngAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAA9GAA9GAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpABngABngAQ+hAQ+hABdwABdwAQ+hAQ+hABngABngAAIpAAIpAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAQ+hAQ+hABngABngAAEaAAEaAAEaAAEaAAEaABngABngAQ+hAQ+hAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAQ+hAQ+hABdwABdwAA9GAA9GABdwABdwAQ+hAQ+hABngABngAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpAAIpAAIpABngABngAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABngABngABngABngABngABngAAIpABngABngAQ+hAQ+hABdwABdwAA9GAA9GABdwABdwAQ+hAQ+hABngABngAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpABdwABdwAQ+hAQ+hABngABngABngABngABngAQ+hAQ+hABdwABdwAAIpAAIpAAIpAQ+hAQ+hAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAAIpAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAA9GAA9GABdwABdwAQ+hAQ+hAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAAIpAQ+hAQ+hAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAQ+hAQ+hAQ+hAQ+hAQ+hAQ+hAAIpAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAA9GAA9GABdwABdwAQ+hAQ+hAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAA9GAA9GABdwABdwAQ+hAQ+hAQ+hAQ+hAQ+hABdwABdwAA9GAA9GAAIpAAIpAAIpABdwABdwAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABdwABdwABdwABdwABdwABdwAAIpABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpABdwABdwAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpABdwABdwABdwABdwABdwABdwAAIpABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GABdwABdwABdwABdwABdwAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAA9GAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAA9GAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAA9GAA9GAA9GAA9GAA9GAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////AAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbdnNzdnNzREJCREJCREJCREJCREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNz////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbdnNzdnNzREJCREJCREJCREJCREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNz////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////////////////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzREJCREJC////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNYmSOJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vb/////////////////////////9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzREJCREJC////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpJCZNAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz/9vb/9vbAAIpAAIp////////dnNzdnNzREJCREJCKCgoKCgo////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////dnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz/9vb/9vbAAIpAAIp////////dnNzdnNzREJCREJCKCgoKCgo////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////REJCREJCREJCREJCREJCREJCREJCREJC////////AAIpAAIp////////////////////////////////////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////REJCREJCREJCREJCREJCREJCREJCREJC////////AAIpAAIp////////////////////////////////////////////////AAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////KCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////AAIpAAIpdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz////////dnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp////////KCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////AAIpAAIpdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNz////////dnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpREJCREJCREJCREJCREJCREJCREJCREJC////////REJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp/9vb/9vbAAIpAAIpREJCREJCREJCREJCREJCREJCREJCREJC////////REJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNz/9vb/9vb/////////////////9vb/9vbdnNzdnNzAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////KCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNz/9vb/9vb/////////////////9vb/9vbdnNzdnNzAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgo////////KCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNzREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCdnNzdnNzdnNzdnNzdnNzdnNzdnNzdnNzREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpdnNzdnNzAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpYmSOAAIpAAIpAAIpAAIpAAIpKCgoKCgoREJCREJCREJCREJCREJCREJCREJCREJCKCgoKCgoAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoREJCREJCREJCREJCREJCREJCREJCREJCKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpREJCREJCAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpNThlAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpKCgoKCgoAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIpAAIp"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Send Text
Welcome to the Divoom API

command description
send text to device, and device will add one text in current animation。the command can be runned after sending animation(the “Draw/SendHttpGif” comand). the command will be active at 90102 version.
It will use font types from app animation font, and display in one line, it will scroll if the text’s length isn’t enough 。the text height is 16 point.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/SendHttpText
TextId number the text id is is unique and will be replaced with the same ID， it is samller than 20
x number the start x postion
y number the start y postion
dir number 0:scroll left, 1:scroll right
font number 0~7, app animation’s font
TextWidth number the text width is based point and bigger than 16, smaller than 64
TextString string the text string is utf8 string and lenght is smaller than 512
speed number the scroll speed if it need scroll, the time (ms) the text move one step
color string the font color, eg:#FFFF00
align number horizontal text alignment, 1:left; 2: middle; 3:right , it will support at 90102 version
request example
{
"Command":"Draw/SendHttpText",
"TextId":4,
"x":0,
"y":40,
"dir":0,
"font":4,
"TextWidth":56,
"speed":10,
"TextString":"hello, Divoom",
"color":"#FFFF00",
"align":1
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Clear all text area
Welcome to the Divoom API

command description
it will clear all text area.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/ClearHttpText
request example
{
"Command":"Draw/ClearHttpText"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Get font list
Welcome to the Divoom API

command description
it will all font information, it will be used in Send display list.
you can login “http://dial.divoom-gz.com/dial.php/index.html“ via divoom acount and check all font in “Display font view”, these font will be used the “reivew all disp items” of the “template dail”.
Request URL
https://app.divoom-gz.com/Device/GetTimeDialFontList
Request method
POST
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"FontList": [
{
"id": 2,
"name": "8\*8 English letters, Arabic figures,punctuation",
"width": "8",
"high": "8",
"charset": "",
"type": 0
},
...

    ]

}
Return value
Name DATA TYPE DESCRIPTION
id number the font id which is used as font in Send display list
name string the font name
width number the font width
high number the font height
Type number 1: the information won’t support scroll, 0:it will scroll if the width isn’t enough
charset string the font include character setting

Send display list
Welcome to the Divoom API

command description
it can send some based elements(time,temperature….). it runs after the command “Send animation”.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/SendHttpItemList
ItemList array item list
TextId number the text id is unique and will be replaced with the same ID， it is samller than 40
type number the display type， It will be introduced below
x number the start x postion
y number the start y postion
dir number 0:scroll left, 1:scroll right
font number it is font id via https://app.divoom-gz.com/Device/GetTimeDialFontList, you shold select the font with Type=0 if you hope it to scrool
TextWidth number the text area width
Textheight number the text area height
TextString string the text string is utf8 string and lenght is smaller than 512 , it will be display string or request url string， it is Optional
speed number the scroll speed if it need scroll, the time (ms) the text move one step
color string the font color, eg:#FFFF00
update_time number the url request interval time based on seconds， it is Optional
align number horizontal text alignment, 1:left; 2: middle; 3:right , it will support at 90102 version
request example
{
"Command":"Draw/SendHttpItemList",
"ItemList":[
{
"TextId":5,
"type":6,
"x":32,
"y":32,
"dir":0,
"font":18,
"TextWidth":32,
"Textheight":16,
"speed":100,
"align":1,
"color":"#FF0000"
},

            {
            "TextId":1,
            "type":14,
            "x":0,
            "y":0,
            "dir":0,
            "font":18,
            "TextWidth":32,
            "Textheight":16,
            "speed":100,
            "align":1,
            "color":"#FF0000"

},
{
"TextId":2,
"type":22,
"x":16,
"y":16,
"dir":0,
"font":2,
"TextWidth":48,
"Textheight":16,
"speed":100,
"align":1,
"TextString":"hello, divoom",
"color":"#FFFFFF"
},
{
"TextId":20,
"type":23,
"x":0,
"y":48,
"dir":0,
"font":4,
"TextWidth":64,
"Textheight":16,
"speed":100,
"update_time":60,
"align":1,
"TextString":"http://appin.divoom-gz.com/Device/ReturnCurrentDate?test=0",
"color":"#FFF000"
}
]
}

display type:
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_SECOND = 1, //sceocnd , font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_MIN = 2, //min, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_HOUR = 3, //hour, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_TIME_AM_PM = 4, //am or pm, font should include a,m,p
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_HOUR_MIN = 5, //hour：min , font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_HOUR_MIN_SEC = 6, //hour:min:sec, , font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_YEAR = 7, //year,, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_DAY = 8, //day, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_MON = 9, //month, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_MON_YEAR = 10, //mon-year, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_ENG_MONTH_DOT_DAY = 11,//month, font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_DATE_WEEK_YEAR = 12, //day:month:year, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_ENG_WEEK = 13, ///weekday-"SU","MO","TU","WE","TH","FR","SA", font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_ENG_WEEK_THREE = 14, //weekday-"SUN","MON","TUE","WED","THU","FRI","SAT", font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_ENG_WEEK_ALL = 15, //weekday-"SUNDAY","MONDAY","TUESDAY","WEDNESDAY","THURSDAY","FRIDAY","SATURDAY", font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_ENG_MON = 16, //month-"JAN","FEB","MAR","APR","MAY","JUN","JUL","AUG","SEP","OCT","NOV","DEC", font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_TEMP_DIGIT = 17, //temperature, font should include digit and c,f
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_TODAY_MAX_TEMP = 18, //the max temperature, font should include digit and c,f
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_TODAY_MIN_TEMP = 19, //the min temperature, font should include digit and c,f
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_WEATHER_WORD = 20, //the weather, font should include english letters
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_NOISE_DIGIT = 21, //the nosie value, font should include digit
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_TEXT_MESSAGE = 22, //the text string, font should include text information
DIVOOM_DISP_CUSTOM_DIAL_SUPPORT_NET_TEXT_MESSAGE = 23, //the url request string, font should include url information, respone should be json encode including the "DispData" string element, eg:http://appin.divoom-gz.com/Device/ReturnCurrentDate?test=0 repone {"DispData": "2022-01-22 13:51:56"}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Play Buzzer
Welcome to the Divoom API

command description
it will play buzzer.
it will be active after 90109.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/PlayBuzzer
ActiveTimeInCycle number Working time of buzzer in one cycle in milliseconds
OffTimeInCycle number Idle time of buzzer in one cycle in milliseconds
PlayTotalTime number Working total time of buzzer in milliseconds
request example
{
"Command":"Device/PlayBuzzer",
"ActiveTimeInCycle":500,
"OffTimeInCycle":500,
"PlayTotalTime":3000
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

play divoom gif
Welcome to the Divoom API

command description
play divoom gif file, it will get from “Get Img Upload List” and “Get My Like Img List”.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/SendRemote
FileId string it is FileId
request example
{
"Command":"Draw/SendRemote",
"FileId":"group1/M00/1C/80/eEwpPWQZFUmEQwsOAAAAAM8RSLs0290624"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Get Img Upload List
Welcome to the Divoom API

Request URL
https://app.divoom-gz.com/Device/GetImgUploadList
Request method
POST
para
Name DATA TYPE DESCRIPTION
DeviceId number Device ID
DeviceMac string Device Mac, it can get from “Find device”
Page number paged, default:1
request example
{
"DeviceId":300000001,
"DeviceMac":"a8032aff46b1",
"Page": 1
}
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"ImgList": [
{
"FileName": "avaa",
"FileId": "group1\/M00\/10\/50\/L1ghbmLVLZ6EI5kGAAAAAHM30Do8982712"
},
...
]
}
Return value
Name DATA TYPE DESCRIPTION
ReturnCode number 0
ImgList array Img List

Get My Like Img List
Welcome to the Divoom API

Request URL
https://app.divoom-gz.com/Device/GetImgLikeList
Request method
POST
para
Name DATA TYPE DESCRIPTION
DeviceId number Device ID
DeviceMac string Device Mac , it can get from “Find device”
Page number paged, default:1
request example
{
"DeviceId":300000001,
"DeviceMac":"a8032aff46b1",
"Page": 1
}
respone example
{
"ReturnCode": 0,
"ReturnMessage": "",
"ImgList": [
{
"FileName": "avaa",
"FileId": "group1\/M00\/10\/50\/L1ghbmLVLZ6EI5kGAAAAAHM30Do8982712"
},
...
]
}
Return value
Name DATA TYPE DESCRIPTION
ReturnCode number 0
ImgList array Img List

save gif
Welcome to the Divoom API

command description
save gif file to local. the gif files only support the size (1616 ,32 32 ,64 \* 64).
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/SaveTFGif
LocalName string local path
NetName string net file address;
request example
{
"Command":"Device/SaveTFGif",
"NetName":"http://f.divoom-gz.com/energy_overdose.gif",
"LocalName":"test/gif/ok/energy_overdose.gif"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Play a frame of a GIF file
Welcome to the Divoom API

command description
Play a frame of a local GIF file. the gif files only support the size (1616 ,32 32 ,64 \* 64).
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Device/PlaySomeFrameGif
FileName string local path
FrameId number the frame id;
request example
{
"Command":"Device/PlaySomeFrameGif",
"FrameId":10,
"FileName":"test/gif/ok/campfire.gif"

}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Command list
Welcome to the Divoom API

command description
CommandList will run all commmand of the the command array.
the command will be implemented after the 90102 version.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/CommandList
CommandList Array the command array information
request example
{
"Command":"Draw/CommandList",
"CommandList":[
{
"Command":"Device/PlayTFGif",
"FileType":2,
"FileName":"http://f.divoom-gz.com/64_64.gif"
},
{
"Command":"Channel/SetBrightness",
"Brightness":100
}
]
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0

Url Command file
Welcome to the Divoom API

command description
UseHTTPCommandSource will run all commmands in url file.
the command will be implemented after the 90102 version.
Request URL
http://IP:80/post
Request method
POST
para
Name DATA TYPE DESCRIPTION
Command string Draw/UseHTTPCommandSource
CommandUrl string the url address of the command array information
request example
{
"Command":"Draw/UseHTTPCommandSource",
"CommandUrl":"http://f.divoom-gz.com/all_command.txt"
}
respone example
{
"error_code": 0
}
Return value
Name DATA TYPE DESCRIPTION
error_code number 0
