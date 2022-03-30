package blame

import (
	. "gopkg.in/check.v1"
)

type ShareMgrSuite struct{}

var _ = Suite(&ShareMgrSuite{})

func (ShareMgrSuite) TestTssShareMgr(c *C) {
	mgr := NewTssShareMgr()
	mgr.Set("test1")
	ret := mgr.QueryAndDelete("test3")
	c.Assert(ret, Equals, false)
	ret = mgr.QueryAndDelete("test1")
	c.Assert(ret, Equals, true)
	ret = mgr.QueryAndDelete("test1")
	c.Assert(ret, Equals, false)
}
