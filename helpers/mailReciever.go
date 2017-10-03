package helpers

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/veqryn/go-email/email"
)

// RecieveMessage Get message from os.Stdin and return github.com/veqryn/go-email/email
// Message struct
func RecieveMessage() *email.Message {
	reader := bufio.NewReader(os.Stdin)
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	err := ioutil.WriteFile("/home/watch/mail.txt", buf.Bytes(), 0644)
	check(err)

	r := bytes.NewReader(buf.Bytes())
	m, err := email.ParseMessage(r)
	check(err)

	return m
}

// TestVideoLost For testing and app building, predefined email message
func TestVideoLost() *email.Message {

	s := `From watch@watch.it  Wed Feb  1 14:10:46 2017
Received: from Lodzska-PC (unknown [128.0.183.218])
        by mail.watch.it (Postfix) with ESMTP id 8CBDF19F3C4
        for <watch@watch.it>; Wed,  1 Feb 2017 14:10:46 +0100 (CET)
From: "watch@watch.it" <watch@watch.it>
To: watch@watch.it
Subject: =?utf-8?Q?cctv=5Fkotorska1752?=
Content-Type: text/plain;
        charset=utf-8
Date: Wed, 1 Feb 2017 14:10:42 +0100

Video Lost  Notice!!
Lost Video Camera :
vytah 46
on Wed, 01 Feb 17 14:10:42 Strední Evropa (be<9e>ný cas)`

	return testMessage(s)
}

// TestParadoxTest For testing and app building, predefined email message
func TestParadoxTest() *email.Message {
	s := `From paradox@watch.it  Wed Feb  1 19:34:13 2017
Received: from paradox.com (177-111-243-80.cust.centrio.cz [80.243.111.177])
        by mail.watch.it (Postfix) with ESMTP id 8FC0019F3C4
        for <paradox@watch.it>; Wed,  1 Feb 2017 19:34:13 +0100 (CET)
From: =?UTF-8?B?QVBzeXN0ZW0gTWFrb3Zza2VobzExNDA=?=<paradox@watch.it>
To: paradox@watch.it
Subject: =?UTF-8?B?VGVzdG92YWNpIGVtYWls?=
Content-Type: text/plain; charset="UTF-8"

Email zaslany z modulu Paradox IP150
Misto: APsystem Kotorska1752
Zprava: Testovaci email
SMTP Server: 193.165.81.174
Port: 25
Vyzadovano overeni Ano
`
	return testMessage(s)
}

// TestOnOffParadox For testing and app building, predefined email message
func TestOnOffParadox() *email.Message {
	s := `From paradox@watch.it  Wed Feb  1 19:34:13 2017
Received: from paradox.com (177-111-243-80.cust.centrio.cz [80.243.111.177])
        by mail.watch.it (Postfix) with ESMTP id 8FC0019F3C4
        for <paradox@watch.it>; Wed,  1 Feb 2017 19:34:13 +0100 (CET)
From: =?UTF-8?B?QVBzeXN0ZW0gTWFrb3Zza2VobzExNDA=?=<paradox@watch.it>
To: paradox@watch.it
Subject: =?UTF-8?B?VGVzdG92YWNpIGVtYWls?=
Content-Type: text/plain; charset="UTF-8"

Email zaslany z modulu Paradox IP150
Misto: APsystem Kotorska1752
Zprava: Zapnuto do ostrahy
Podsystem: 1 - Vchod 1572
z: Kardas
Cas: 11 Zari 2014 14:05`
	return testMessage(s)
}

// TestPorts Run test
func TestPorts() {
	// Todo; test local port??
}

func testMessage(s string) *email.Message {
	fmt.Println("testMessage(s string)")
	r := strings.NewReader(s)
	m, err := email.ParseMessage(r)
	check(err)

	return m
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
