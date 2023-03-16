package localcache

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Cache
	}{
		{
			name: "TestNew",
			want: &cache{
				cacheMap: make(map[string]cacheData),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Set string",
			args: args{
				k: "hello",
				v: "world",
			},
		},
		{
			name: "Test Set int",
			args: args{
				k: "hello",
				v: 123,
			},
		},
		{
			name: "Test Set struct",
			args: args{
				k: "hello",
				v: struct {
					Name string
					Age  int
				}{
					Name: "John",
					Age:  18,
				},
			},
		},
		{
			name: "Test Set nil",
			args: args{
				k: "hello",
				v: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.Set(tt.args.k, tt.args.v)

			cacheMap := c.CacheMap()

			if got := cacheMap[tt.args.k].Data; !reflect.DeepEqual(got, tt.args.v) || reflect.TypeOf(got) != reflect.TypeOf(tt.args.v) {
				t.Errorf("Get() = %v %v, want %v %v", got, reflect.TypeOf(got), tt.args.v, reflect.TypeOf(tt.args.v))
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test Get string",
			args: args{
				k: "A",
				v: "world",
			},
			want: "world",
		},
		{
			name: "Test Get int",
			args: args{
				k: "B",
				v: 123,
			},
			want: 123,
		},
		{
			name: "Test Get struct",
			args: args{
				k: "C",
				v: struct {
					Name string
					Age  int
				}{
					Name: "John",
					Age:  18,
				},
			},
			want: struct {
				Name string
				Age  int
			}{
				Name: "John",
				Age:  18,
			},
		},

		{
			name: "Test Get nil",
			args: args{
				k: "D",
				v: nil,
			},
			want: nil,
		},

		{
			name: "Test Get empty string",
			args: args{
				k: "E",
				v: "",
			},
			want: "",
		},

		{
			name: "Test Get empty struct",
			args: args{
				k: "F",
				v: struct{}{},
			},
			want: struct{}{},
		},
	}

	mockMap := make(map[string]cacheData)

	// insert test data
	for _, tt := range tests {
		mockMap[tt.args.k] = cacheData{
			Data:     tt.args.v,
			ExpireAt: time.Now().Add(time.Hour),
		}
	}

	// mock initCacheMap
	initCacheMap = func() map[string]cacheData {
		return mockMap
	}

	c := New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTTL(t *testing.T) {
	type args struct {
		k string
		v any
	}
	tests := []struct {
		name     string
		args     args
		timePass time.Duration
		want     any
	}{
		{
			name: "Test Get before TTL",
			args: args{
				k: "hello",
				v: "world",
			},
			timePass: (cacheTTL - 1) * time.Second,
			want:     "world",
		},
		{
			name: "Test Get after TTL",
			args: args{
				k: "hello",
				v: "world",
			},
			timePass: (cacheTTL + 1) * time.Second,
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock time start
			timeNow = func() time.Time { return time.Unix(1629446406, 0) }

			c := New()
			c.Set(tt.args.k, tt.args.v)

			// mock time passes
			timeNow = func() time.Time { return time.Unix(1629446406, 0).Add(tt.timePass) }

			if got := c.Get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyNotExists(t *testing.T) {
	c := New()
	if got := c.Get("hello"); got != nil {
		t.Errorf("Get() = %v, want nil", got)
	}
}
