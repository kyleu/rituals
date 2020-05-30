package util

import (
	"fmt"
	"github.com/gofrs/uuid"
	"logur.dev/logur"
	"sort"
	"strings"
)

func GetEntry(m map[string]interface{}, key string, logger logur.Logger) interface{} {
	retEntry, ok := m[key]
	if !ok {
		if logger != nil {
			keys := make([]string, 0, len(m))
			for k, _ := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			logger.Warn(fmt.Sprintf("no key [%v] in map from available keys [%v]", key, strings.Join(keys, ", ")))
		}
		return nil
	}
	return retEntry
}

func MapGetString(m map[string]interface{}, key string, logger logur.Logger) string {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(string)
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not string", key, retEntry))
		return ""
	}
	return ret
}

func MapGetMap(m map[string]interface{}, key string, logger logur.Logger) map[string]interface{} {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(map[string]interface{})
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not map[string]interface{}", key, retEntry))
		return nil
	}
	return ret
}

func MapGetBool(m map[string]interface{}, key string, logger logur.Logger) bool {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(bool)
	if !ok {
		logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not bool", key, retEntry))
		return false
	}
	return ret
}

func MapGetUUID(m map[string]interface{}, key string, logger logur.Logger) *uuid.UUID {
	retEntry := GetEntry(m, key, logger)
	ret, ok := retEntry.(uuid.UUID)
	if !ok {
		s, ok := retEntry.(string)
		if !ok {
			logger.Warn(fmt.Sprintf("key [%v] in map is type [%T], not uuid", key, retEntry))
			return nil
		}
		r, e := uuid.FromString(s)
		if e != nil {
			logger.Warn(fmt.Sprintf("key [%v] in map with value [%v] is not a valid uuid", key, s))
			return nil
		}
		ret = r
	}
	return &ret
}
