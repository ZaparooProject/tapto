/*
Zaparoo Core
Copyright (C) 2023 Gareth Jones
Copyright (C) 2023, 2024 Callan Barrett

This file is part of Zaparoo Core.

Zaparoo Core is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Zaparoo Core is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Zaparoo Core.  If not, see <http://www.gnu.org/licenses/>.
*/

package pn532_uart

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/hsanjuan/go-ndef"
)

var NdefEnd = []byte{0xFE}
var NdefStart = []byte{0x54, 0x02, 0x65, 0x6E}

var ErrNoNdef = fmt.Errorf("no NDEF record found")

func ParseRecordText(bs []byte) (string, error) {
	// sometimes there can be some read corruption and multiple copies of the
	// NDEF header get pulled in. we just flick through until the last one
	// TODO: is this going to mess up if there are multiple NDEF records?
	startIndex := bytes.LastIndex(bs, NdefStart)
	if startIndex == -1 {
		return "", ErrNoNdef
	}

	// check if there is another ndef start left, as it can mean we got come
	// corrupt data at the beginning
	if len(bs) > startIndex+8 {
		nextStart := bytes.Index(bs[startIndex+4:], NdefStart)
		if nextStart != -1 {
			startIndex += nextStart
		}
	}

	endIndex := bytes.Index(bs, NdefEnd)
	if endIndex == -1 {
		return "", fmt.Errorf("NDEF end not found: %x", bs)
	}

	if startIndex >= endIndex || startIndex+4 >= len(bs) {
		return "", fmt.Errorf("start index out of bounds: %d, %x", startIndex, bs)
	}

	if endIndex <= startIndex || endIndex >= len(bs) {
		return "", fmt.Errorf("end index out of bounds: %d, %x", endIndex, bs)
	}

	log.Debug().Msgf("NDEF start: %d, end: %d", startIndex, endIndex)
	log.Debug().Msgf("NDEF: %x", bs[startIndex:endIndex])

	tagText := string(bs[startIndex+4 : endIndex])

	return tagText, nil
}

func BuildMessage(text string) ([]byte, error) {
	msg := ndef.NewTextMessage(text, "en")
	var payload, err = msg.Marshal()
	if err != nil {
		return nil, err
	}

	header, err := CalculateNdefHeader(payload)
	if err != nil {
		return nil, err
	}
	payload = append(header, payload...)
	payload = append(payload, []byte{0xFE}...)
	return payload, nil
}

func CalculateNdefHeader(ndefRecord []byte) ([]byte, error) {
	var recordLength = len(ndefRecord)
	if recordLength < 255 {
		return []byte{0x03, byte(len(ndefRecord))}, nil
	}

	// NFCForum-TS-Type-2-Tag_1.1.pdf Page 9
	// > 255 Use three consecutive bytes format
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint16(recordLength))
	if err != nil {
		return nil, err
	}

	var header = []byte{0x03, 0xFF}
	return append(header, buf.Bytes()...), nil
}
