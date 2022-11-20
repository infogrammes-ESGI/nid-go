package main

var LIST_PROTOCOLS = []string{"tcp", "udp", "icmp", "http", "ip", "ssh", "tls", "ssl"}
var LIST_ACTIONS = []string{"alert", "log", "pass"}

type Rule struct {
	action string
	protocol string
	in_network string
	in_port int // -1 if 'any'
	out_network string
	out_port int // -1 if 'any'
	condition string // TODO: change to RuleCondition[]
}

var rules = make([]Rule, 0)

func Add_New_Rule(new_rule Rule) {
	rules = append(rules, new_rule)
}
