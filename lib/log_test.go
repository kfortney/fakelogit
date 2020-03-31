package lib

import (
	"testing"
	"time"
)

func TestNewFidelisLog(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewFidelisLog", args{t: LogTimeFormat()}, NewFidelisLog(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewFidelisLog(tt.args.t); got == tt.want {
				t.Errorf("NewFidelisLog() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewLancopeLog(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewLancopeLog", args{t: LogTimeFormat()}, NewLancopeLog(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewLancopeLog(tt.args.t); got == tt.want {
				t.Errorf("NewLancopeLog() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewPaloThreatLog(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewPaloThreatLog", args{t: LogTimeFormat()}, NewPaloThreatLog(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaloThreatLog(tt.args.t); got == tt.want {
				t.Errorf("NewPaloThreatLog() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewPaloTrafficLog(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewPaloTrafficLog", args{t: LogTimeFormat()}, NewPaloTrafficLog(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaloTrafficLog(tt.args.t); got == tt.want {
				t.Errorf("NewPaloTrafficLog() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewRFC3164Log(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewRFC3164Log", args{t: LogTimeFormat()}, NewRFC3164Log(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC3164Log(tt.args.t); got == tt.want {
				t.Errorf("NewRFC3164Log() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewRFC5424Log(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewRFC5424Log", args{t: LogTimeFormat()}, NewRFC5424Log(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC5424Log(tt.args.t); got == tt.want {
				t.Errorf("NewRFC5424Log() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}

func TestNewTrendMicroLog(t *testing.T) {

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNewTrendMicroLog", args{t: LogTimeFormat()}, NewTrendMicroLog(LogTimeFormat())},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrendMicroLog(tt.args.t); got == tt.want {
				t.Errorf("NewTrendMicroLog() = \n%s, \n%s", got, tt.want)
			}
		})
	}
}
