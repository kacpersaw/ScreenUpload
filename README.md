# ScreenUpload

ScreenUpload app is used as shell script in ScreenCloud. This app allows you to upload your screens to own server by sftp. ScreenUpload creates dirs (year/month/day) and then upload your image to it. ScreenCloud generates unique screen name.

Currently binary is only available for Linux.

###Download
To download ScreenUpload binary file for Linux go to this [link](https://github.com/kacpersaw/ScreenUpload/releases).

### How to run?
To run ScreeCloud you must copy executable to /usr/bin after this execute this command:
```bash
cat MyScreen.png | ScreenCloud --server=localhost --port=22 --username=root --password=secretpassword --upload_path=/root/screens --upload_url=http://localhost/screens
```
Be sure to not add a trailing slash in ``upload_path`` and ``upload_url``

``upload_path`` is directory where ScreenUpload creates dirs (year/month/day) and uploads your image to it.
``upload_url`` is url where ScreenUpload adds ``year/month/day/unique_name.png``

To use ScreenUpload with ScreenCloud you must create bash script.

Example bash script:
```bash
 #!/bin/bash
 cat $1 | ScreenCloud --server=localhost --port=22 --username=root --password=secretpassword --upload_path=/root/screens --upload_url=http://localhost/screens
 exit 0
```

In ScreenCloud set command to like below:
```bash
sh /home/user/upload.sh {s}
```

###License
The MIT License (MIT)

Copyright (c) 2015 Kacper Sawicki

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
