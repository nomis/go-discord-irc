package bridge

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"

	"github.com/thoj/go-ircevent"
)

func setupIRCConnection(con *irc.Connection, webIRCPass, hostname, ip string) {
	con.UseTLS = true
	con.TLSConfig = &tls.Config{}

	// this requires a modification to thoj/go-ircevent
	con.InitialCommand = fmt.Sprintf("WEBIRC %s discord %s %s", webIRCPass, hostname, ip)
}

// Leftpad is from github.com/douglarek/leftpad
func Leftpad(s string, length int, ch ...rune) string {
	c := ' '
	if len(ch) > 0 {
		c = ch[0]
	}
	l := length - utf8.RuneCountInString(s)
	if l > 0 {
		s = strings.Repeat(string(c), l) + s
	}
	return s
}

// SnowflakeToIP takes a snowflake and the first half of an IP to make an IP suitable for WEBIRC
func SnowflakeToIP(base string, snowflake string) string {
	num, err := strconv.ParseUint(snowflake, 10, 64)
	if err != nil {
		panic(errors.Wrapf(err, "could not convert snowflake to uint: %s", snowflake))
	}

	for i, c := range Leftpad(strconv.FormatUint(num, 16), 16, '0') {
		if (i % 4) == 0 {
			base += ":"
		}
		base += string(c)
	}

	return base
}
