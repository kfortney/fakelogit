package lib

import (
	"testing"
	"time"
)

func TestNewLog(t *testing.T) {
	type args struct {
		format string
		t      time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLog(tt.args.format, tt.args.t); got != tt.want {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC3164Log(tt.args.t); got != tt.want {
				t.Errorf("NewRFC3164Log() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC5424Log(tt.args.t); got != tt.want {
				t.Errorf("NewRFC5424Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFidelisLog1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFidelisLog(tt.args.t); got != tt.want {
				t.Errorf("NewFidelisLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLancopeLog1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLancopeLog(tt.args.t); got != tt.want {
				t.Errorf("NewLancopeLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPaloThreatLog1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaloThreatLog(tt.args.t); got != tt.want {
				t.Errorf("NewPaloThreatLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPaloTrafficLog1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaloTrafficLog(tt.args.t); got != tt.want {
				t.Errorf("NewPaloTrafficLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRFC3164Log1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC3164Log(tt.args.t); got != tt.want {
				t.Errorf("NewRFC3164Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRFC5424Log1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRFC5424Log(tt.args.t); got != tt.want {
				t.Errorf("NewRFC5424Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTrendMicroLog1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTrendMicroLog(tt.args.t); got != tt.want {
				t.Errorf("NewTrendMicroLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
