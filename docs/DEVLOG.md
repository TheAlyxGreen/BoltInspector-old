# Developer Log

### March 3rd, 2017
I'm beginning to recognize the need for a central error message system. Currently all error messages are hard-coded into the functions that call them. Unfortunately, this leads to some messages with the same meaning having different wording. I'll look for existing solutions. If none that I like exist, I may make my own.

### February 19th, 2017
As of right now, basically all of the main functions exist for handling data in the database. The main thing that is missing is to make the copy/move function support buckets. But that seems like a somewhat rare use case, so I'm going to go ahead and push this update without that. I'll probably come back at some point and add them.
 
 The next thing I want to do is add a web interface to manage all of this, sort of like PHPMyAdmin. I think it'd make this far more useful. Because of the way it's written, whenever a command isn't actively running, BoltInspector doesn't lock the file. This means that you could theoretically leave it running while another app is using the database file and the two not interfere with one another, offering a live view of the changes. A web app could leverage this by remembering the information inside the web app and highlighting changes. That's a bit down the road, however.
 
 Another thing I'm considering is adding a mode that monitors for changes and notifies (probably via webhook) when there is one. What I see as one of the strongest areas for something like BoltDB is in an IoT/embedded device or in other devices where you won't have a monitor or I/O attached at all times. As such, these sorts of features would be an excellent way to monitor their activities and performance automatically. This, however, might also be best done as separate app, rather than bundling it into a larger administration toolkit.