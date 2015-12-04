package run

import (
	"reflect"
	"testing"
)

func TestParseConf_IncorrectWatchSection(t *testing.T) {
	defer expectPanic(`Reading a conf file with incorrect "watch" section. Panic expected.`)
	parseConf("./testdata/configs/incorrect_watch.yml")
}

func TestParseConf_RunTextSection(t *testing.T) {
	defer expectPanic(`/run's first argument should be a section containing a list. Got text, panic expected.`)
	parseConf("./testdata/configs/run_not_list_section.yml")
}

func TestParseConf_StartTextSection(t *testing.T) {
	defer expectPanic(`/start's first argument should be a section containing a list. Got text, panic expected.`)
	parseConf("./testdata/configs/start_not_list_section.yml")
}

func TestParseConf_SingleTextSection(t *testing.T) {
	defer expectPanic(`/single's first argument should be a section containing a list. Got text, panic expected.`)
	parseConf("./testdata/configs/single_not_list_section.yml")
}

func TestParseConf_RunIncorrectArgsNum(t *testing.T) {
	defer expectPanic(`/run expects one argument. Got a few of them, panic expected.`)
	parseConf("./testdata/configs/run_incorrect_args_number.yml")
}

func TestParseConf_StartIncorrectArgsNum(t *testing.T) {
	defer expectPanic(`/start expects one argument. Got a few of them, panic expected.`)
	parseConf("./testdata/configs/start_incorrect_args_number.yml")
}

func TestParseConf_SingleIncorrectArgsNum(t *testing.T) {
	defer expectPanic(`/single expects one argument. Got a few of them, panic expected.`)
	parseConf("./testdata/configs/single_incorrect_args_number.yml")
}

func TestParseConf_PassIncorrectArgsNum(t *testing.T) {
	defer expectPanic(`/pass expects no arguments. Got a few of them, panic expected.`)
	parseConf("./testdata/configs/pass_incorrect_args_number.yml")
}

func TestParseConf_EmptyWatch(t *testing.T) {
	defer expectPanic(`No empty configuration files are allowed, panic expected.`)
	parseConf("./testdata/configs/empty_watch.yml")
}

func TestParseConf_LoopSection(t *testing.T) {
	defer expectPanic(`Loops in configuration files are not allowed. Panic expected.`)
	parseConf("./testdata/configs/loop_section.yml")
}

func TestParseConf(t *testing.T) {
	parseConf("./testdata/configs/correct_config.yml")
}

func TestParseTask(t *testing.T) {
	s := "goal run path/to/app"
	expN := "goal"
	expAs := []string{"run", "path/to/app"}
	if rn, ras := parseTask(s); rn != expN || !reflect.DeepEqual(ras, expAs) {
		t.Errorf(
			"Failed to parse task `%s`. Expected `%s`, `%v`; got `%s`, `%v`.", s, expN, expAs, rn, ras,
		)
	}
}
