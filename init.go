package boolean

import "regexp"

var rNum *regexp.Regexp
var rTru *regexp.Regexp
var rFal *regexp.Regexp
var rNeg *regexp.Regexp


func init() {
  rNum = regexp.MustCompile(`(?i)^[-+]?\d+$`)
  rTru = regexp.MustCompile(`(?i)\b(?:t|y|1)\b|\b(?:\+|ay|go|on|up)|(?:tru|acc|asc|day|for|hot|inc|joy|new|pos|top|win|yes|dawn|full|safe|grow|high|just|real|some|know|live|love|open|pure|shin|warm|wis[de]|activ|admit|advan|agree|begin|brigh|build|creat|early|enter|float|f(?:i|ou)nd|grant|light|north|prett|prese|publi|start|succe|victr)`)
  rFal = regexp.MustCompile(`(?i)\b(?:f|n|0)\b|(?:fal|off|dim|end|low|old|back|cold|cool|dark|dead|decr|desc|dirt|down|dull|dusk|exit|late|sink|ugly|absen|botto|close|finis|night|priva|south|wrong)`)
  rNeg = regexp.MustCompile(`(?i)\b(?:-|na|no|un|in|aft|bad|dis|lie|non|ben[dt]|den[iy]|empt|fail|fake|hate|los[es]|stop|decli|defea|destr|never|negat|refus|rejec|forget|shr[iu]nk|against|is.?nt|can.?(?:no)?t)|(?:hind)`)
}
