# TheLinuxCommandLine

## reading notes

### 2 What Is The Shell

#### Your First Keystroke

If the last character of the prompt is a pound sign ("#") rather than a dollar sign, the terminal session has superuser privileges. This means either we are logged in as the root user or we selected a terminal emulator that provides superuser (administrative) privilege.

+ **date** display the current time and date
+ **cal** display a calendar of the current month
+ **df** see the current amount of free space on your disk drives
+ **free** display the amount of free memory
+ **exit** end a terminal session
	 
### 3 Navigation

+ **pwd**
+ **cd**
	- **cd -** change to the previous working dir
	- **cd ~username** change to the username's home dir
+ **ls**

### 4 Exploring The System

+ **ls**
+ **file** determine file type
+ **less** view file content
	- **f** forward one window
	- **b** backward one window
	- **d** forward one-half window
	- **u** backward one-half window
	- **j** forward one line
	- **k** backward one line
	- **g** goto first line
	- **G** goto last line
	- **q** quit

### 5 Manipulating Files And Directories

+ **cp**
	- **cp -u** copy only when the SOURCE file is newer
  than the destination file or when the destination
	file is missing
+ **mv** move/rename files and dirs
+ **mkdir**
+ **rm** linux __donot__ have an undelete command
	- **rm -rf** recusive rm dir
+ **ln** create hard and soft links
	- **ln file link** create a hard link
		1. *A hard link cannot reference a file outside its own file system. This means a link may not reference a file that is not on the same disk partition as the link itself.*
		2. *A hard link may not reference a directory.*
		3. *When a hard link is deleted, the link is removed but the contents of the file itself continue to exist (that is, its space is not deallocated) until all links to the file are deleted.*
	- **ln -s item link** create a soft link

#### Wildcards

+ __*__ matches any characters
+ __?__ matches any single character
+ __[characters]__ Matches any character that is a member of the set characters
+ __[!characters]__ Matches any character that is not a member of the set
characters
+ __[[:class:]]__ Matches any character that is a member of the specified
class

Commonly Used Character Classes

+ __[:alnum:]__ Matches any alphanumeric character
+ __[:alpha:]__ Matches any alphabetic character
+ __[:digit:]__ Matches any numeral
+ __[:lower:]__ Matches any lowercase letter
+ __[:upper:]__ Matches any uppercase letter

Examples

+ __*__ All files
+ __g*__ Any file beginning with "g"
+ __b*.txt__ Any file beginning with "b" followed byny characters and ending with ".txt"
+ __Data???__ Any file beginning with "Data" followed by exactly three characters
+ __[abc]*__ Any file beginning with either an "a", a "b", or a "c"
+ __BACKUP.[0-9][0-9][0-9]__ Any file beginning with "BACKUP." followed by exactly three numerals
+ __[[:upper:]]*__ Any file beginning with an uppercase letter
+ __[![:digit:]]*__ Any file not beginning with a numeral
+ __*[[:lower:]123]__ Any file ending with a lowercase letter or the numerals "1", "2", or "3"

### 6 Working With Commands

+ **type** Indicate how a command name is interpreted
+ **which** Display which executable program will be executed
+ **help** Get help for shell builtins
+ **man** Display a command's manual page
+ **apropos** Display a list of appropriate commands
+ **info** Display a command's info entry
+ **whatis** Display a very brief description of a command
+ **alias** Create an alias for a command

### 7 Redirection

+ **cat** Concatenate files. reads one or more files and copies them to standard output.
	- **cat -n file** Show file contents and line number.
	- **cat file.zip.0* > file.zip** also can be used to join seperated files.
	- **cat > tmpfile** cat without arguments will read contents from standred input. So type this command, and type contents then press ctrl-d. We can use this behavior to create short text files.
+ **sort** Sort lines of text
+ **uniq** Report or omit repeated lines
+ **grep** Print lines matching a pattern
	- **-i** Ignore case. Do not distinguish between upper and lower case characters. May also be specified --ignore-case.
	- **-v** Invert match. Normally, grep prints lines that contain a match. This option causes grep to print every line that does not contain a match. May also be specified --invert-match.
	- **-c** Print the number of matches (or non-matches if the -v option is also specified) instead of the lines themselves. May also be specified --count.
	- **-l** Print the name of each file that contains a match instead of the lines themselves. May also be specified --files-with-matches.
	- **-L** Like the -l option, but print only the names of files that do not contain matches. May also be specified --files-withoutmatch.
	- **-n** Prefix each matching line with the number of the line within the file. May also be specified --line-number.
	-	**-h** For multi-file searches, suppress the output of filenames. May also be specified --no-filename.
+ **wc** Print newline, word, and byte counts for each file
+ **head** Output the first part of a file
+ **tail** Output the last part of a file
	- **tail -f** real time watch appended contents
+ **tee** Read from standard input and write to standard output and files
	- **ls /usr/bin | tee ls.txt | grep zip** The tee program reads standard input and copies it to both standard output (allowing the data to continue down the pipeline) and to one or more files. This is useful for capturing a pipeline's contents at an intermediate stage of processing.

**> outputfile** Simply using the redirection operator with no command preceding it will truncate an existing file or create a new, empty file.

**ls 2> errorfile** The file descriptor 2 is placed immediately before the redirection operator to perform the redirection of standard error to the file.

**ls &> recfile** Redirect both standard output and standard error to the rec file.

**ls 2> /dev/null** Sometimes we don't want output from a command, we just want to throw it away.

**ls /bin /usr/bin | sort | less** Pipelines are often used to perform complex operations on data. It is possible to put several commands together into a pipeline. Frequently, the commands used this way are referred to as filters.

**ls /c /d /e /f | sort | uniq | grep program**

### 8 Seeing The World As The Shell Sees It

**echo** Display a line of text

#### Arithmetic Expansion

__echo $(($((5**2)) * 3))__ print 75

#### Brace Expansion

+ **echo Front-{A,B,C}-Back** Front-A-Back Front-B-Back Front-C-Back
+ **echo Number_{1..5}** Number_1 Number_2 Number_3 Number_4 Number_5
+ **echo {Z..A}** Z Y X W V U T S R Q P O N M L K J I H G F E D C B A
+ **echo a{A{1,2},B{3,4}}b** aA1b aA2b aB3b aB4b
+ **mkdir {2007..2009}-0{1..9} {2007..2009}-{10..12}**
	- ls
	- 2007-01 2007-07 2008-01 2008-07 2009-01 2009-07
	- 2007-02 2007-08 2008-02 2008-08 2009-02 2009-08
	- 2007-03 2007-09 2008-03 2008-09 2009-03 2009-09
	- 2007-04 2007-10 2008-04 2008-10 2009-04 2009-10
	- 2007-05 2007-11 2008-05 2008-11 2009-05 2009-11
	- 2007-06 2007-12 2008-06 2008-12 2009-06 2009-12

#### Parameter Expansion

**echo $PATH**

#### Command Substitution

+ **ls -l $(which cp)** -rwxr-xr-x 1 lenovo Administrators 116736 Apr 28  2010 /bin/cp.exe
+ **ls -l `which cp`** It uses back-quotes instead of the dollar sign and parentheses, which is an alternate syntax for command substitution in older shell programs which is also supported in bash.

#### Quotes

+ **echo $(cal)** February 2008 Su Mo Tu We Th Fr Sa 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
+ **Newlines are considered delimiters by the word-splitting mechanism**
+ **echo "$(cal)"**
+	February 2008
+	Su Mo Tu We Th Fr Sa
+	                1  2
+	 3  4  5  6  7  8  9
+	10 11 12 13 14 15 16
+	17 18 19 20 21 22 23
+	24 25 26 27 28 29
+ **echo text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER** text /home/me/ls-output.txt a b foo 4 me
+ **echo "text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER"** text ~/*.txt {a,b} foo 4 me
+ **echo 'text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER'** text ~/*.txt {a,b} $(echo foo) $((2+2)) $USER

### 9 Advanced Keyboard Tricks

+ **clear** Clear the screen
+ **history** Display the contents of the history list
  - **!88** will execute the 88th history command
  - **ctrl-r** will search previous commands

**same as emacs movements**

#### Cursor Movement

+ **Ctrl-a** Move cursor to the beginning of the line.
+ **Ctrl-e** Move cursor to the end of the line.
+ **Ctrl-f** Move cursor forward one character; same as the right arrow key.
+ **Ctrl-b** Move cursor backward one character; same as the left arrow key.
+ **Alt-f** Move cursor forward one word.
+ **Alt-b** Move cursor backward one word.
+ **Ctrl-l** Clear the screen and move the cursor to the top left corner. The
clear command does the same thing.

#### Modifying Text

+ **Ctrl-d** Delete the character at the cursor location
+ **Ctrl-t** Transpose (exchange) the character at the cursor location with the one preceding it.
+ **Alt-t** Transpose the word at the cursor location with the one preceding it.
+ **Alt-l** Convert the characters from the cursor location to the end of the
word to lowercase.
+ **Alt-u** Convert the characters from the cursor location to the end of the
word to uppercase.

#### Cutting And Pasting (Killing And Yanking) Text

+ **Ctrl-k** Kill text from the cursor location to the end of line.
+ **Ctrl-u** Kill text from the cursor location to the beginning of the line.
+ **Alt-d** Kill text from the cursor location to the end of the current word.
+ **Alt-Backspace** Kill text from the cursor location to the beginning of the current word. If the cursor is at the beginning of a word, kill the previous word.
+ **Ctrl-y** Yank text from the kill-ring and insert it at the cursor location.

#### search history

+ **Ctrl-p** Move to the previous history entry. Same action as the up arrow.
+ **Ctrl-n** Move to the next history entry. Same action as the down arrow.
+ **Alt-<** Move to the beginning (top) of the history list.
+ **Alt->** Move to the end (bottom) of the history list, i.e., the current
command line.
+ **Ctrl-r** Reverse incremental search. Searches incrementally from the
current command line up the history list.
+ **Alt-p** Reverse search, non-incremental. With this key, type in the search string and press enter before the search is performed.

#### History Expansion

+ **!!** Repeat the last command. It is probably easier to press up arrow
and enter.
+ **!number** Repeat history list item number.
+ **!string** Repeat last history list item starting with string.
+ **!?string** Repeat last history list item containing string.

### 10 Permissions

+ **id** Display user identity
+ **chmod** Change a file's mode
  - Be aware that only the file's owner or the superuser can change the mode of a file or directory. chmod supports two distinct ways of specifying mode changes: octal number representation, or symbolic representation. We will cover octal number representation first.
  - ls -l foo.txt
  - *-rw-rw-r-- 1 me me 0 2008-03-06 14:52 foo.txt*
  - **chmod 600 foo.txt**
  - ls -l foo.txt
  - *-rw------- 1 me me 0 2008-03-06 14:52 foo.txt*
  - **chmod also supports a symbolic notation for specifying file modes.**
		+ **u+x** Add execute permission for the owner.
		+ **u-x** Remove execute permission from the owner.
		+ **+x** Add execute permission for the owner, group, and world. Equivalent to a+x.
		+ **o-rw** Remove the read and write permission from anyone besides the owner and group owner.
		+ **go=rw** Set the group owner and anyone besides the owner to have read and write permission. If either the group owner or world previously had execute permissions, they are removed.
		+ **u+x,go=rx** Add execute permission for the owner and set the permissions for the group and others to read and execute. Multiple specifications may be separated by commas.
+ **umask** Set the default file permissions
+ **su** Run a shell as another user
+ **sudo** Execute a command as another user
	- The sudo command is like su in many ways, but has some important additional capabilities. The administrator can configure sudo to allow an ordinary user to execute commands as a different user (usually the superuser) in a very controlled way. In particular, a user may be restricted to one or more specific commands and no others. Another important difference is that the use of sudo does not require access to the superuser's password. To authenticate using sudo, the user uses his/her own password.
+ **chown** Change a file's owner
	- **Superuser privileges are required to use this command.**
	- **chown [owner][:[group]] file...**
		+ **bob** Changes the ownership of the file from its current owner to user bob.
		+ **bob:users** Changes the ownership of the file from its current owner to user bob and changes the file group owner to group users.
		+ **:admins** Changes the group owner to the group admins. The file owner is unchanged.
		+ **bob:** Change the file owner from the current owner to user bob and changes the group owner to the login group of user bob.
+ **chgrp** Change a file's group ownership
+ **passwd** Change a user's password

User accounts are defined in the /etc/passwd file and groups are defined in the /etc/group file. When user accounts and groups are created, these files are modified along with /etc/shadow which holds information about the user's password.

#### File Type

ls -l foo.txt
-rw-rw-r-- 1 me me 0 2008-03-06 14:52 foo.txt

+ **-** A regular file.
+ **d** A directory.
+ **l** A symbolic link. Notice that with symbolic links, the remaining file attributes are always "rwxrwxrwx" and are dummy values. The real file attributes are those of the file the symbolic link points to.
+ **c** A character special file. This file type refers to a device that handles data as a stream of bytes, such as a terminal or modem.
+ **b** A block special file. This file type refers to a device that handles data in blocks, such as a hard drive or CD-ROM drive.

#### Permission Attributes

+ **r**
  - **files** Allows a file to be opened and read.
  - **directories** Allows a directory's contents to be listed if the execute attribute is also set.
+ **w**
  - **files** Allows a file to be written to or truncated, however this attribute does not allow files to be renamed or deleted. The ability to delete or rename files is determined by directory attributes.
  - **directories** Allows files within a directory to be created, deleted, and renamed if the execute attribute is also set.
+ **x**
  - **files** Allows a file to be treated as a program and executed. Program files written in scripting languages must also be set as readable to be executed.
  - **directories** Allows a directory to be entered, e.g., cd directory.

### 11 Processes

+ **ps** Report a snapshot of current processes
	- **ps auxew**
+ **top** Display tasks
+ **jobs** List active jobs
+ **bg** Place a job in the background
	- To launch a program so that it is immediately placed in the background, we follow the command with an- "&" character.
+ **fg** Place a job in the foreground
+ **kill** Send a signal to a process
	- **kill [-signal] PID...**
	- **-1 HUP** Hangup. This is a vestige of the good old days when terminals were attached to remote computers with phone lines and modems. The signal is used to indicate to programs that the controlling terminal has "hung up." The effect of this signal can be demonstrated by closing a terminal session. The foreground program running on the terminal will be sent the signal and will terminate. This signal is also used by many daemon programs to cause a reinitialization. This means that when a daemon is sent this signal, it will restart and re-read its configuration file. The Apache web server is an example of a daemon that uses the HUP signal in this way.
	- **2 INT** Interrupt. Performs the same function as the Ctrl-c key sent from the terminal. It will usually terminate a program.
	- **9 KILL** Kill. This signal is special. Whereas programs may choose to handle signals sent to them in different ways, including ignoring them all together, the KILL signal is never actually sent to the target program. Rather, the kernel immediately terminates the process. When a process is terminated in this manner, it is given no opportunity to "clean up" after itself or save its work. For this reason, the KILL signal should only be used as a last resort when other termination signals fail.
	- **15 TERM** Terminate. This is the default signal sent by the kill command. If a program is still "alive" enough to receive signals, it will terminate.
	- **18 CONT** Continue. This will restore a process after a STOP signal.
	- **19 STOP** Stop. This signal causes a process to pause without terminating. Like the KILL signal, it is not sent to the target process, and thus it cannot beignored.
	- **3 QUIT** Quit.
	- **11 SEGV** Segmentation Violation. This signal is sent if a program makes illegal use of memory, that is, it tried to write somewhere it was not allowed to.
	- **20 TSTP** Terminal Stop. This is the signal sent by the terminal when the Ctrl-z key is pressed. Unlike the STOP signal, the TSTP signal is received by the process and may be ignored.
	- **28 WINCH** Window Change. This is a signal sent by the system when a window changes size. Some programs , like top and less will respond to this signal by redrawing themselves to fit the new window dimensions.
+ **killall** Kill processes by name
+ **shutdown** Shutdown or reboot the system
	- **shutdown -r now**
	- **shutdown -h now**

When the terminal receives one of these keystrokes, it sends a signal to the program in the foreground. In the case of Ctrl-c, a signal called INT (Interrupt) is sent; with Ctrl-z, a signal called TSTP (Terminal Stop.) Programs, in turn, "listen" for signals and may act upon them as they are received.

+ **pstree** Outputs a process list arranged in a tree-like pattern showing the parent/child relationships between processes.
+ **vmstat** Outputs a snapshot of system resource usage including, memory, swap and disk I/O. To see a continuous display, follow the command with a time delay (in seconds) for updates. For example: vmstat 5. Terminate the output with Ctrl-c.
+ **xload** A graphical program that draws a graph showing system load over time.
+ **tload** Similar to the xload program, but draws the graph in the terminal. Terminate the output with Ctrl-c.

### 12 The Environment

+ **printenv** Print part or all of the environment
+ **set** Set shell options
+ **export** Export environment to subsequently executed programs
+ **alias** Create an alias for a command

#### How Is The Environment Established?

When we log on to the system, the bash program starts, and reads a series of configuration scripts called startup files, which define the default environment shared by all users. This is followed by more startup files in our home directory that define our personal environment. The exact sequence depends on the type of shell session beingstarted. There are two kinds: a login shell session and a non-login shell session.


A login shell session is one in which we are prompted for our user name and password; when we start a virtual console session, for example.

+ **/etc/profile** A global configuration script that applies to all users.
+ **~/.bash_profile** A user's personal startup file. Can be used to extend or override settings in the global configuration script.
+ **~/.bash_login** If ~/.bash_profile is not found, bash attempts to read this script.
+ **~/.profile** If neither ~/.bash_profile nor ~/.bash_login is found, bash attempts to read this file. This is the default in Debian-based distributions, such as Ubuntu.

A non-login shell session typically occurs when we launch a terminal session in the GUI.

+ **/etc/bash.bashrc** A global configuration script that applies to all users.
+ **~/.bashrc** A user's personal startup file. Can be used to extend or override settings in the global configuration script.

The ~/.bashrc file is probably the most important startup file from the ordinary user's point of view, since it is almost always read. Non-login shells read it by default and most startup files for login shells are written in such a way as to read the ~/.bashrc file as well.

The changes we have made to our .bashrc will not take affect until we close our terminal session and start a new one, since the .bashrc file is only read at the beginning of a session. However, we can force bash to re-read the modified .bashrc file with the following command:
**source .bashrc**

### 13 A Gentle Introduction To vi

### 14 Customizing The Prompt

### 15 Package Management

Most distributions fall into one of two camps of packaging technologies: the Debian ".deb" camp and the Red Hat ".rpm" camp. There are some important exceptions such as Gentoo, Slackware, and Foresight, but most others use one of these two basic systems.

+ **Debian Style (.deb)** Debian, Ubuntu, Xandros, Linspire
	- **Low-Level Tools** dpkg
	- **High-Level Tools** apt-get, aptitude
+ **Red Hat Style (.rpm)** Fedora, CentOS, Red Hat Enterprise Linux, OpenSUSE, Mandriva, PCLinuxOS
	- **Low-Level Tools** rpm
	- **High-Level Tools** yum

#### Common Package Management Tasks

+ **Debian**
	- **search**
		+ apt-get update
		+ apt-cache search search_string
	- **install**
		+ apt-get update
		+ apt-get install package_name
	- **remove**
		+ apt-get remove package_name
	- **update**
		+ apt-get update
		+ apt-get upgrade
	- **list**
		+ dpkg --list
	- **if a package is installed**
		+ dpkg --status package_name
	- **info**
		+ apt-cache show package_name
+ **Red Hat**
	- **search**
		+ yum search search_string
	- **install**
		+ yum install package_name
	- **remove**
		+ yum erase package_name
	- **update**
		+ yum update
	- **list**
		+ rpm -qa
	- **if a package is installed**
		+ rpm -q package_name
	- **info**
		+ yum info package_name

##### Installing A Package From A Package File

If a package file has been downloaded from a source other than a repository, it can be installed directly (though without dependency resolution) using a low-level tool.

Example: If the emacs-22.1-7.fc7-i386.rpm package file had been downloaded from a non-repository site, it would be installed this way:
**rpm -i emacs-22.1-7.fc7-i386.rpm**

Note: Since this technique uses the low-level rpm program to perform the installation, no dependency resolution is performed. If rpm discovers a missing dependency, rpm will exit with an error.

+ **Debian** dpkg --install package_file
+ **Red Hat** rpm -i package_file
	
##### Upgrading A Package From A Package File

+ **Debian** dpkg --install package_file
+ **Red Hat** rpm -U package_file

##### Finding Which Package Installed A File

+ **Debian** dpkg --search file_name
+ **Red Hat** rpm -qf file_name

### 16 Storage Media

+ **mount** Mount a file system
+ **umount** Unmount a file system
+ **fsck** Check and repair a file system
+ **fdisk** Partition table manipulator
+ **mkfs** Create a file system
+ **fdformat** Format a floppy disk
+ **dd** Write block oriented data directly to a device
+ **genisoimage (mkisofs)** Create an ISO 9660 image file
+ **wodim (cdrecord)** Write data to optical storage media
+ **md5sum** Calculate an MD5 checksum

There is a file named /etc/fstab that lists the devices (typically hard disk partitions) that are to be mounted at boot time.

+ __/dev/fd*__ Floppy disk drives.
+ __/dev/hd*__ IDE (PATA) disks on older systems. Typical motherboards contain two IDE connectors or channels, each with a cable with two attachment points for drives. The first drive on the cable is called the master device and the second is called the slave device. The device names are ordered such that /dev/hda refers to the master device on the first channel, /dev/hdb is the slave device on the first channel; /dev/hdc, the master device on the second channel, and so on. A trailing digit indicates the partition number on the device. For example, /dev/hda1 refers to the first partition on the first hard drive on the system while / dev/hda refers to the entire drive.
+ __/dev/lp*__ Printers.
+ __/dev/sd*__ SCSI disks. On recent Linux systems, the kernel treats all disklike devices (including PATA/SATA hard disks, flash drives, and USB mass storage devices, such as portable music players and digital cameras) as SCSI disks. The rest of the naming system is similar to the older /dev/hd* naming scheme described above.
+ __/dev/sr*__ Optical drives (CD/DVD readers and burners)

#### Creating New File Systems

##### Manipulating Partitions With fdisk

The **fdisk** program allows us to interact directly with disk-like devices (such as hard disk drives and flash drives) at a very low level. With this tool we can edit, delete, and create partitions on the device. To work with our flash drive, we must first unmount it (if needed) and then invoke the fdisk program as follows:

+ **$ sudo umount /dev/sdb1**
+ **$ sudo fdisk /dev/sdb**

Notice that we must specify the device in terms of the entire device, not by partition number. After the program starts up, we will see the following prompt:

##### Creating A New File System With mkfs

With our partition editing done (lightweight though it might have been) it's time to create a new file system on our flash drive. To do this, we will use mkfs (short for "make file system"), which can create file systems in a variety of formats. To create an ext3 file system on the device, we use the "-t" option to specify the "ext3" system type, followed by the name of device containing the partition we wish to format:

+ **$ sudo mkfs -t ext3 /dev/sdb1**

##### Moving Data Directly To/From Devices

However, if we could treat a disk drive as simply a large collection
of data blocks, we could perform useful tasks, such as cloning devices.
The dd program performs this task. It copies blocks of data from one place to another. It uses a unique syntax (for historical reasons) and is usually used this way:

+ **dd if=input_file of=output_file [bs=block_size [count=blocks]]**
+ **dd if=/dev/sdb of=/dev/sdc**
+ **dd if=/dev/sdb of=flash_drive.img**

###### Creating An Image Copy Of A CD-ROM

+ **dd if=/dev/cdrom of=ubuntu.iso**

###### Creating An Image From A Collection Of Files

+ **genisoimage -o cd-rom.iso -R -J ~/cd-rom-files**

###### Mounting An ISO Image Directly

There is a trick that we can use to mount an iso image while it is still on our hard disk and treat it as though it was already on optical media. By adding the "-o loop" option to mount (along with the required "-t iso9660" file system type), we can mount the image file as though it were a device and attach it to the file system tree:

+ **mkdir /mnt/iso_image**
+ **mount -t iso9660 -o loop image.iso /mnt/iso_image**

###### Blanking A Re-Writable CD-ROM

+ **wodim dev=/dev/cdrw blank=fast**

###### Writing An Image

+ **wodim dev=/dev/cdrw image.iso**

### 17 Networking

+ **ping** - Send an ICMP ECHO_REQUEST to network hosts
+ **traceroute** - Print the route packets trace to a network host
+ **netstat** - Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships
	- **netstat -ie**
		+ **lo** is the loopback interface, a virtual interface that the system uses to "talk to itself".
		+ When performing causal network diagnostics, the important things to look for are the presence of the word "UP" at the beginning of the fourth line for each interface, indicating that the network interface is enabled, and the presence of a valid IP address in the inet addr field on the second line. For systems using DHCP (Dynamic Host Configuration Protocol), a valid IP address in this field will verify that the DHCP is working.
	-	**netstat -r**
		+ Kernel IP routing table
		+ Destination Gateway Genmask Flags MSS Window irtt Iface
		+ **192.168.1.0 * 255.255.255.0 U 0 0 0 eth0**
		+ **default 192.168.1.1 0.0.0.0 UG 0 0 0 eth0**
		+ In this simple example, we see a typical routing table for a client machine on a LAN (Local Area Network) behind a firewall/router. The first line of the listing shows the destination 192.168.1.0. IP addresses that end in zero refer to networks rather than individual hosts, so this destination means any host on the LAN. The next field, Gateway, is the name or IP address of the gateway (router) used to go from the current host to the destination network. An asterisk in this field indicates that no gateway is needed.
		+ The last line contains the destination default. This means any traffic destined for a network that is not otherwise listed in the table. In our example, we see that the gateway is defined as a router with the address of 192.168.1.1, which presumably knows what to do with the destination traffic.
+ **ftp** - Internet file transfer program
+ **wget** - Non-interactive network downloader
+ **ssh** - OpenSSH SSH client (remote login program)
	- **ssh -i id_rsa fairlyblank@192.168.26.152**
	- **sftp -i id_rsa fairlyblank@192.168.26.152**

### 18 C Searching For Files

+ **locate** Find files by name
+ **find** Search for files in a directory hierarchy
	- **find ~** show all files recusively in home dir
		+ **-depth** Direct find to process a directory's files before the directory itself. This option is automatically applied when the -delete action is specified.
		+ **-maxdepth levels** Set the maximum number of levels that find will descend into a directory tree when performing tests and actions.
		+ **-mindepth levels** Set the minimum number of levels that find will descend into a directory tree before applying tests and actions.
		+ **-mount** Direct find not to traverse directories that are mounted on other file systems.
		+ **-noleaf** Direct find not to optimize its search based on the assumption that it is searching a Unix-like file system. This is needed when scanning DOS/Windows file systems and CD-ROMs.
	- **find ~ | wc -l** show files&dir numbers in home dir
	- **find ~ -type d | wc -l** show dir numbers in home dir
	- **find ~ -type f | wc -l** show files numbers in home dir
	- **find ~ -type f -name "*.JPG" -size +1M | wc -l** look for all the regular files that match the wild card pattern "*.JPG" and are larger than one megabyte
		+ **-cmin n** Match files or directories whose content or attributes were last modified exactly n minutes ago. To specify less than n minutes ago, use -n and to specify more than n minutes ago, use +n.
		+ **-cnewer file** Match files or directories whose contents or attributes were last modified more recently than those of file.
		+ **-ctime n** Match files or directories whose contents or attributes were last modified n*24 hours ago.
		+ **-empty** Match empty files and directories.
		+ **-group name** Match file or directories belonging to group. group may be expressed as either a group name or as a numeric group ID.
		+ **-iname pattern** Like the -name test but case insensitive.
		+ **-inum n** Match files with inode number n. This is helpful for finding all the hard links to a particular inode.
		+ **-mmin n** Match files or directories whose contents were modified n minutes ago.
		+ **-mtime n** Match files or directories whose contents were modified n*24 hours ago.
		+ **-name pattern** Match files and directories with the specified wild card pattern.
		+ **-newer file** Match files and directories whose contents were modified more recently than the specified file. This is very useful when writing shell scripts that perform file backups. Each time you make a backup, update a file (such as a log), then use find to determine which files that have changed since the last update.
		+ **-nouser** Match file and directories that do not belong to a valid user. This can be used to find files belonging to deleted accounts or to detect activity by attackers.
		+ **-nogroup** Match files and directories that do not belong to a valid group.
		+ **-perm mode** Match files or directories that have permissions set to the specified mode. mode may be expressed by either octal or symbolic notation.
		+ **-samefile name** Similar to the -inum test. Matches files that share the same inode number as file name.
		+ **-size n** Match files of size n.
		+ **-type c** Match files of type c.
		+ **-user name** Match files or directories belonging to user name. The user may be expressed by a user name or by a numeric user ID.
	- **find ~ \( -type f -not -perm 0600 \) -or \( -type d -not -perm 0700 \)** look for all the files with permissions that are not 0600 and the directories with permissions that are not 0700.
		+ **-and** Match if the tests on both sides of the operator are true. May be shortened to -a. Note that when no operator is present, -and is implied by default.
		+ **-or** Match if a test on either side of the operator is true. May be shortened to -o.
		+ **-not** Match if the test following the operator is false. May be abbreviated with an exclamation point (!).
		+ **( )** Groups tests and operators together to form larger expressions. This is used to control the precedence of the logical evaluations. By default, find evaluates from left to right. It is often necessary to override the default evaluation order to obtain the desired result. Even if not needed, it is helpful sometimes to include the grouping characters to improve readability of the command. Note that since the parentheses characters have special meaning to the shell, they must be quoted when using them on the command line to allow them to be passed as arguments to find. Usually the backslash character is used to escape them.
	- **find ~ -type f -name '*.BAK' -delete** delete files that have the file extension ".BAK"
		+ **-delete** Delete the currently matching file.
		+ **-ls** Perform the equivalent of ls -dils on the matching file. Output is sent to standard output.
		+ **-print** Output the full pathname of the matching file to standard output. This is the default action if no other action is specified.
		+ **-quit** Quit once a match has been made.
	- **-exec rm '{}' ';'** command is the name of a command, {} is a symbolic representation of the current pathname and the semicolon is a required delimiter indicating the end of the command.
		+ **find ~ -type f -name 'foo*' -ok ls -l '{}' ';'** By using the -ok action in place of -exec, the user is prompted before execution of each specified command.
		+ **find ~ -type f -name 'foo*' -exec ls -l '{}' +** By changing the trailing semicolon character to a plus sign, we activate the ability of find to combine the results of the search into an argument list for a single execution of the desired command.
+ **xargs** It accepts input from standard input and converts it into an argument list for a specified command.
	- **find ~ -type f -name 'foo*' -print | xargs ls -l**
	- **find ~ -iname '*.jpg' -print0 | xargs --null ls -l** Unix-like systems allow embedded spaces (and even newlines!) in filenames. The find command provides the action -print0, which produces null separated output, and the xargs command has the --null option, which accepts null separated input.
+ **touch** Change file times
+ **stat** Display file or file system status

### 19 Archiving And Backup

+ **gzip** Compress or expand files
	- The gzip program is used to compress one or more files. When executed, it replaces the original file with a compressed version of the original. The corresponding gunzip program is used to restore compressed files to their original, uncompressed form.
	- **-c** Write output to standard output and keep original files. May also be specified with --stdout and --to-stdout.
	- **-d** Decompress. This causes gzip to act like gunzip. May also be specified with --decompress or --uncompress.
	- **-f** Force compression even if compressed version of the original file already exists. May also be specified with --force.
	- **-h** Display usage information. May also be specified with --help.
	- **-l** List compression statistics for each file compressed. May also be specified with --list.
	- **-r** If one or more arguments on the command line are directories, recursively compress files contained within them. May also be specified with --recursive.
	- **-t** Test the integrity of a compressed file. May also be specified with --test.
	- **-v** Display verbose messages while compressing. May also be specified with --verbose.
	- **-number** Set amount of compression. number is an integer in the range of 1 (fastest, least compression) to 9 (slowest, most compression). The values 1 and 9 may also be expressed as --fast and --best, respectively. The default value is 6.
+ **bzip2** A block sorting file compressor n
+ **tar** Tape archiving utility
	- **c** Create an archive from a list of files and/or directories.
	- **x** Extract an archive.
	- **r** Append specified pathnames to the end of an archive.
	- **t** List the contents of an archive.
+ **zip** Package and compress files
	- Unless we include the -r option for recursion, only the playground directory (but none of its contents) is stored.
+ **rsync** Remote file and directory synchronization

### 20 Regular Expressions

Regular expression metacharacters consist of the following:
__^ $ . [ ] { } - ? * + ( ) | \__

#### The Any Character **.**

#### Anchors **^** **$**

+ Note that the regular expression "^$" (a beginning and an end with nothing in between) will match blank lines.

#### Bracket Expressions And Character Classes **[]**

+ In addition to matching any character at a given position in our regular expression, we can also match a single character from a specified set of characters by using bracket expressions.

#### Alternation **|**

#### Quantifiers

+ **?** Match An Element Zero Or One Time
	- This quantifier means, in effect, make the preceding element optional.
+ __*__ Match An Element Zero Or More Times
	- __[[:upper:]][[:upper:][:lower:] ]*\.__
	- Let's say we wanted to see if a string was a sentence; that is, it starts with an uppercase letter, then contains any number of upper and lowercase letters and spaces, and ends with a period.
+ **+** Match An Element One Or More Times
	- __^([[:alpha:]]+ ?)+$__
	- Here is a regular expression that will only match lines consisting of groups of one or more alphabetic characters separated by single spaces.
+ **{}** Match An Element A Specific Number Of Times
	- **{n}** Match the preceding element if it occurs exactly n times.
	- **{n,m}** Match the preceding element if it occurs at least n times, but no more than m times.
	- **{n,}** Match the preceding element if it occurs n or more times.
	- **{,m}** Match the preceding element if it occurs no more than m times.
	
### 21 Text Processing

+ **cat** Concatenate files and print on the standard output
	- **-A** option, which is used to display nonprinting characters in the text.
	- **-n** which numbers lines, and **-s**, which suppresses the output of multiple blank lines.
+ **sort** Sort lines of text files
	- **-b** --ignore-leading-blanks By default, sorting is performed on the entire line, starting with the first character in the line. This option causes sort to ignore leading spaces in lines and calculates sorting based on the first non-whitespace character on the line.
	- **-f** --ignore-case Makes sorting case insensitive.
	- **-n** --numeric-sort Performs sorting based on the numeric evaluation of a string. Using this option allows sorting to be performed on numeric values rather than alphabetic values.
	- **-r** --reverse Sort in reverse order. Results are in descending rather than ascending order.
	- **-k** --key=field1[,field2] Sort based on a key field located from field1 to field2 rather than the entire line. See discussion below.
	- **-m** --merge Treat each each argument as the name of a presorted file. Merge multiple files into a single sorted result without performing any additional sorting.
	- **-o** --output=file Send sorted output to file rather than standard output.
	- **-t** --field-separator=char Define the field separator character. By default fields are separated by spaces or tabs.
+ **uniq** Report or omit repeated lines
	- **-c** Output a list of duplicate lines preceded by the number of times the line occurs.
	- **-d** Only output repeated lines, rather than unique lines.
	- **-f n** Ignore n leading fields in each line. Fields are separated by whitespace as they are in sort; however, unlike sort, uniq has no option for setting an alternate field separator.
	- **-i** Ignore case during the line comparisons.
	- **-s** n Skip (ignore) the leading n characters of each line.
	- **-u** Only output unique lines. This is the default.
+ **cut** Remove sections from each line of files
	- **-c char_list** Extract the portion of the line defined by char_list. The list may consist of one or more comma-separated numerical ranges.
	- **-f field_list** Extract one or more fields from the line as defined by field_list. The list may contain one or more fields or field ranges separated by commas.
	- **-d delim_char** When -f is specified, use delim_char as the field delimiting character. By default, fields must be separated by a single tab character.
	- **--complement** Extract the entire line of text, except for those portions specified by -c and/or -f.
+ **paste** Merge lines of files
+ **join** Join lines of two files on a common field
+ **comm** Compare two sorted files line by line
+ **diff** Compare files line by line
	- **normal format**
		+ **r1ar2** Append the lines at the position r2 in the second file to the position r1 in the first file.
		+ **r1cr2** Change (replace) the lines at position r1 with the lines at the position r2 in the second file.
		+ **r1dr2** Delete the lines in the first file at position r1, which would have appeared at range r2 in the second file
	- **context format**
		+ **blank** A line shown for context. It does not indicate a difference between the two files.
		+ **-** A line deleted. This line will appear in the first file but not in the second file.
		+ **+** A line added. This line will appear in the second file but not in the first file.
		+ **!** A line changed. The two versions of the line will be displayed, each in its respective section of the change group.
	- **unified format**
		+ **blank** This line is shared by both files.
		+ **-** This line was removed from the first file.
		+ **+** This line was added to the first file.
+ **patch** Apply a diff file to an original
+ **tr** Translate or delete characters
+ **sed** Stream editor for filtering and transforming text
+ **aspell** Interactive spell checker

### 22 Formatting Output

+ **nl** Number lines
+ **fold** Wrap each line to a specified length
+ **fmt** A simple text formatter
+ **pr** Prepare text for printing
+ **printf** Format and print data
+ **groff** A document formatting system

### 23 Printing

+ **pr** Convert text files for printing
+ **lpr** Print files
+ **a2ps** Format files for printing on a PostScript printer
+ **lpstat** Show printer status information
+ **lpq** Show printer queue status
+ **lprm** Cancel print jobs

### 24 Compiling Programs

+ **make** Utility To Maintain Programs

### 25 Writing Your First Script

### 26 Starting A Project

+ **a=z** Assign the string "z" to variable a.
+ **b="a string"** Embedded spaces must be within quotes.
+ **c="a string and $b"** Other expansions such as variables can be expanded into the assignment.
+ **d=$(ls -l foo.txt)** Results of a command.
+ **e=$((5 * 7))** Arithmetic expansion.
+ **f="\t\ta string\n"** Escape sequences such as tabs and newlines.

**a here document**

+ FTP_SERVER=ftp.nl.debian.org
+ FTP_PATH=/debian/dists/lenny/main/installer-i386/current/images/cdrom
+ REMOTE_FILE=debian-cd_info.tar.gz
+ ftp -n << _EOF_
+ open $FTP_SERVER
+ user anonymous me@linuxbox
+ cd $FTP_PATH
+ hash
+ get $REMOTE_FILE
+ bye
+ _EOF_
+ ls -l $REMOTE_FILE

If we change the redirection operator from "<<" to "<<-", the shell will ignore leading tab characters in the here document. This allows a here document to be indented, which can improve readability.

### 27 Top-Down Design

### 28 Flow Control: Branching With if

#### test expression

+ **file1 -ef file2** file1 and file2 have the same inode numbers (the two filenames refer to the same file by hard linking).
+ **file1 -nt file2** file1 is newer than file2.
+ **file1 -ot file2** file1 is older than file2.
+ **-b file** file exists and is a block special (device) file.
+ **-c file** file exists and is a character special (device) file.
+ **-d file** file exists and is a directory.
+ **-e file** file exists.
+ **-f file** file exists and is a regular file.
+ **-g file** file exists and is set-group-ID.
+ **-G file** file exists and is owned by the effective group ID.
+ **-k file** file exists and has its "sticky bit" set.
+ **-L file** file exists and is a symbolic link.
+ **-O file** file exists and is owned by the effective user ID.
+ **-p file** file exists and is a named pipe.
+ **-r file** file exists and is readable (has readable permission for the effective user).
+ **-s file** file exists and has a length greater than zero.
+ **-S file** file exists and is a network socket.
+ **-t fd** fd is a file descriptor directed to/from the terminal. This can be used to determine whether standard input/output/ error is being redirected.
+ **-u file** file exists and is setuid.
+ **-w file** file exists and is writable (has write permission for the effective user).
+ **-x file** file exists and is executable (has execute/search permission for the effective user).

#### String Expressions

+ **string** string is not null.
+ **-n string** The length of string is greater than zero.
+ **-z string** The length of string is zero.
+ **string1 = string2 string1 == string2** string1 and string2 are equal. Single or double equal signs may be used, but the use of double equal signs is greatly preferred.
+ **string1 != string2** string1 and string2 are not equal.
+ **string1 > string2** string1 sorts after string2.
+ **string1 < string2** string1 sorts before string2.

#### Integer Expressions

+ **integer1 -eq integer2** integer1 is equal to integer2.
+ **integer1 -ne integer2** integer1 is not equal to integer2.
+ **integer1 -le integer2** integer1 is less than or equal to integer2.
+ **integer1 -lt integer2** integer1 is less than integer2.
+ **integer1 -ge integer2** integer1 is greater than or equal to integer2.
+ **integer1 -gt integer2** integer1 is greater than integer2.

#### A More Modern Version Of test

+ **[[string1 =~ regex]]**
+ **(( ))** - Designed For Integers

### 29 Reading Keyboard Input

+ **read [-options] [variable...]** read Read Values From Standard Input
	- If no variables are listed after the read command, a shell variable, REPLY, will be assigned all the input.
	- **-a array** Assign the input to array, starting with index zero. We will cover arrays in Chapter 36.
	- **-d delimiter** The first character in the string delimiter is used to indicate end of input, rather than a newline character.
	- **-e** Use Readline to handle input. This permits input editing in the same manner as the command line.
	- **-n num** Read num characters of input, rather than an entire line.
	- **-p prompt** Display a prompt for input using the string prompt.
	- **-r** Raw mode. Do not interpret backslash characters as escapes.
	- **-s** Silent mode. Do not echo characters to the display as they are typed. This is useful when inputting passwords and other confidential information.
	- **-t seconds** Timeout. Terminate input after seconds. read returns a non-zero exit status if an input times out.
	- **-u fd** Use input from file descriptor fd, rather than standard input.

### 30 Flow Control: Looping With while / until

**while commands; do commands; done**

### 31 Troubleshooting

### 32 Flow Control: Branching With case

**case word in [pattern [| pattern]...) commands ;;]... esac**

### 33 Positional Parameters

#### Accessing The Command Line

+ **$0 ~ $9** The shell provides a set of variables called positional parameters that contain the individual words on the command line. The variables are named 0 through 9.
+ **$#** that yields the number of arguments on the command line.

#### shift Getting Access To Many Arguments

The shift command causes all of the parameters to "move down one" each time it is executed. In fact, by using shift, it is possible to get by with only one parameter (in addition to $0, which never changes)

+ __$*__ Expands into the list of positional parameters, starting with 1. When surrounded by double quotes, it expands into a double quoted string containing all of the positional parameters, each separated by the first character of the IFS shell variable (by default a space character).
+ __$@__ Expands into the list of positional parameters, starting with 1. When surrounded by double quotes, it expands each positional parameter into a separate word surrounded by double quotes.

### 34 Flow Control: Looping With for

+ **for variable [in words]; do commands done**
+ **for (( expression1; expression2; expression3 )); do commands done**

### 35 Strings And Numbers

#### Expansions To Manage Empty Variables

+ **${parameter:-word}** If parameter is unset (i.e., does not exist) or is empty, this expansion results in the value of word. If parameter is not empty, the expansion results in the value of parameter.
+ **${parameter:=word}**  If parameter is unset or empty, this expansion results in the value of word. In addition, the value of word is assigned to parameter. If parameter is not empty, the expansion results in the value of parameter.
+ **${parameter:?word}** If parameter is unset or empty, this expansion causes the script to exit with an error, and the contents of word are sent to standard error. If parameter is not empty, the expansion results in the value of parameter.
+ **${parameter:+word}** If parameter is unset or empty, the expansion results in nothing. If parameter is not empty, the value of word is substituted for parameter; however, the value of parameter is not changed.

#### Expansions That Return Variable Names

+ **${!prefix*} ${!prefix@}** This expansion returns the names of existing variables with names beginning with prefix.

#### String Operations

+ **${#parameter}** Expands into the length of the string contained by parameter.
+ **${parameter:offset}
${parameter:offset:length}**
This expansion is used to extract a portion of the string contained in parameter.
+ **${parameter#pattern}
${parameter##pattern}**
This expansion removes a leading portion of the string contained in parameter defined by pattern. pattern is a wildcard pattern like those used in pathname expansion. The difference in the two forms is that the # form removes the shortest match, while the ## form removes the longest match.
+ **${parameter%pattern}
${parameter%%pattern}**
These expansions are the same as the # and ## expansions above, except they remove text from the end of the string contained in parameter rather than from the beginning.
+ **${parameter/pattern/string}
${parameter//pattern/string}
${parameter/#pattern/string}
${parameter/%pattern/string}**
This expansion performs a search and replace upon the contents of parameter. If text is found matching wildcard pattern, it is replaced with the contents of string. In the normal form, only the first occurrence of pattern is replaced. In the // form, all occurrences are replaced. The /# form requires that the match occur at the beginning of the string, and the /% form requires the match to occur at the end of the string. /string may be omitted, which causes the text matched by pattern to be deleted.

#### Arithmetic Evaluation And Expansion

**$((expression))**

+ **number** By default, numbers without any notation are treated as decimal (base 10) integers.
+ **0number** In arithmetic expressions, numbers with a leading zero are considered octal.
+ **0xnumber** Hexadecimal notation
+ **base#number** number is in base

**Simple Arithmetic**

+ **+** Addition
+ **-** Subtraction
+ __*__ Multiplication
+ **/** Integer division
+ __**__ Exponentiation
+ **%** Modulo (remainder)

**Assignment**

+ **parameter = value** Simple assignment. Assigns value to parameter.
+ **parameter += value** Addition. Equivalent to parameter = parameter + value
+ **parameter -= value** Subtraction. Equivalent to parameter = parameter - value
+ **parameter *= value** Multiplication. Equivalent to parameter = parameter * value
+ **parameter /= value** Integer division. Equivalent to parameter = parameter / value
+ **parameter %= value** Modulo. Equivalent to parameter = parameter % value
+ **parameter++** Variable post-increment. Equivalent to parameter = parameter + 1 (however, see discussion below)
+ **parameter--** Variable post-decrement. Equivalent to parameter = parameter - 1
+ **++parameter** Variable pre-increment. Equivalent to parameter = parameter + 1
+ **--parameter** Variable pre-decrement. Equivalent to parameter = parameter - 1

**Bit Operations**

+ **~** Bitwise negation. Negate all the bits in a number.
+ **<<** Left bitwise shift. Shift all the bits in a number to the left.
+ **>>** Right bitwise shift. Shift all the bits in a number to the right.
+ **&** Bitwise AND. Perform an AND operation on all the bits in two numbers.
+ **|** Bitwise OR. Perform an OR operation on all the bits in two numbers.
+ **^** Bitwise XOR. Perform an exclusive OR operation on all the bits in two numbers.

**Logic**

+ **<=** Less than or equal to
+ **>=** Greater than or equal to
+ **<** Less than
+ **>** Greater than
+ **==** Equal to
+ **!=** Not equal to
+ **&&** Logical AND
+ **||** Logical OR
+ **expr1?expr2:expr3** Comparison (ternary) operator. If expression expr1 evaluates to be non-zero (arithmetic true) then expr2, else expr3.

### 36 Arrays

#### Assigning Values To An Array

**name[subscript]=value
name=(value1 value2 ...)**

**Outputting The Entire Contents Of An Array**

+ The behavior of notations **${animals[*]}** and **${animals[@]}** are identical until they are quoted. The * notation results in a single word containing the array¡¯s contents, while the @ notation results in three words, which matches the arrays "real" contents.

**echo ${#a[@]}** number of array elements

**echo ${#a[100]}** length of element 100

**${!array[*]} ${!array[@]}** As bash allows arrays to contain ¡°gaps¡± in the assignment of subscripts, it is sometimes useful to determine which elements actually exist. This can be done with a parameter expansion using the following forms.

**foo+=(d e f)** Adding Elements To The End Of An Array

**unset foo** Deleting An Array

### 37 Exotica



