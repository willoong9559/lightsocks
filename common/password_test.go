package common_test

import (
	"sort"
	"testing"

	"github.com/willoong9559/lightsocks/common"
)

type lspasswdWarp struct {
	*common.Lspasswd
}

func (l *lspasswdWarp) Len() int {
	return len(l.Lspasswd)
}

func (l *lspasswdWarp) Less(i, j int) bool {
	return l.Lspasswd[i] < l.Lspasswd[j]
}

func (l *lspasswdWarp) Swap(i, j int) {
	l.Lspasswd[i], l.Lspasswd[j] = l.Lspasswd[j], l.Lspasswd[i]
}

func TestRandPassword(t *testing.T) {
	passwordStr := common.NewRandPasswdStr()
	randPassword, err := common.Atp(passwordStr)
	randPasswordWarp := &lspasswdWarp{randPassword}
	if err != nil {
		t.Error(err)
	}
	sort.Sort(randPasswordWarp)
	for i := 0; i < len(randPassword); i++ {
		if randPassword[i] != byte(i) {
			t.Error("不能出现任何一个重复的byte位")
		}
	}
}
