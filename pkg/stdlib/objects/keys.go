package objects

import (
	"context"
	"sort"

	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

/*
 * Returns string array of object's keys
 * @params obj (Object) - The object whose keys you want to extract
 * @params sort (Boolean, optional) - If sort is true, then the returned keys will be sorted.
 * @returns (Array of String) - Array that contains object keys.
 */
func Keys(_ context.Context, args ...core.Value) (core.Value, error) {
	err := core.ValidateArgs(args, 1, 2)
	if err != nil {
		return values.None, err
	}

	err = core.ValidateType(args[0], core.ObjectType)
	if err != nil {
		return values.None, err
	}

	obj := args[0].(*values.Object)
	needSort := false

	if len(args) == 2 {
		err = core.ValidateType(args[1], core.BooleanType)
		if err != nil {
			return values.None, err
		}

		needSort = bool(args[1].(values.Boolean))
	}

	keys := sort.StringSlice(obj.Keys())
	keysArray := values.NewArray(len(keys))

	if needSort {
		keys.Sort()
	}

	for _, key := range keys {
		keysArray.Push(values.NewString(key))
	}

	return keysArray, nil
}
