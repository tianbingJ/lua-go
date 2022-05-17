--
-- Created by IntelliJ IDEA.
-- User: tianbing
-- Date: 2022/5/3
--


local t = {"a", "b", "c"}
t[2] = "B"
t["foo"] = "Bar"
local s = t[3] ..t[2] .. t[1] .. t["foo"] .. #t
