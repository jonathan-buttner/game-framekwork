// Code generated by "stringer -type=ResourceType"; DO NOT EDIT.

package resource

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Yellow-0]
	_ = x[Red-1]
	_ = x[Green-2]
	_ = x[Brown-3]
}

const _ResourceType_name = "YellowRedGreenBrown"

var _ResourceType_index = [...]uint8{0, 6, 9, 14, 19}

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceType_index)-1) {
		return "ResourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResourceType_name[_ResourceType_index[i]:_ResourceType_index[i+1]]
}
