package docker

//import (
//	"fmt"
//	"strings"
//	"testing"
//)
//
//func TestVolumeMap(t *testing.T) {
//	tests := []struct {
//		item          string
//		wantHost      string
//		wantContainer string
//		wantMode      string
//		wantErr       bool
//	}{
//		{item: "/var/lib/mysql", wantHost: "/var/lib/mysql", wantErr: false},
//		{item: "/opt/data:/var/lib/mysql", wantHost: "/opt/data", wantContainer: "/var/lib/mysql", wantErr: false},
//		{item: "./cache:/tmp/cache", wantHost: "./cache", wantContainer: "/tmp/cache", wantErr: false},
//		{item: "~/configs:/udp", wantHost: "~/configs", wantContainer: "/udp", wantErr: false},
//		{item: "~/configs:/etc/configs/:ro", wantHost: "~/configs", wantContainer: "/etc/configs/", wantMode: "ro", wantErr: false},
//		{item: "datavolume:/var/lib/mysql", wantHost: "datavolume", wantContainer: "/var/lib/mysql", wantErr: false},
//		{item: "datavolume::ro", wantHost: "datavolume", wantContainer: "", wantMode: "ro", wantErr: true},
//		{item: ":/var/lib/mysql", wantHost: "", wantContainer: "/var/lib/mysql", wantMode: "", wantErr: true},
//	}
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
//			// MarshalYaml
//			if !tt.wantErr {
//				item := VolumeMap{Host: tt.wantHost, Container: tt.wantContainer, Mode: tt.wantMode}
//				content, _ := MarshalYaml(item)
//				content = strings.TrimRight(content, "\n")
//				if content != tt.item {
//					t.Logf("%d %d", len(content), len(tt.item))
//					t.Errorf("VolumeMap.MarshalYAML() content = %v, wantContent %v", content, tt.item)
//					return
//				}
//			}
//			// UnmarshalYaml
//			var item VolumeMap
//			err := UnmarshalYaml(tt.item, &item)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("VolumeMap.UnarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if item.Host != tt.wantHost {
//				t.Errorf("VolumeMap.UnarshalYAML() host = %v, wantHost %v", item.Host, tt.wantHost)
//				return
//			}
//			if item.Container != tt.wantContainer {
//				t.Errorf("Image.UnarshalYAML() container = %v, wantContainer %v", item.Container, tt.wantContainer)
//				return
//			}
//			if item.Mode != tt.wantMode {
//				t.Errorf("VolumeMap.UnarshalYAML() mode = %v, wantMode %v", item.Mode, tt.wantMode)
//				return
//			}
//		})
//	}
//}
