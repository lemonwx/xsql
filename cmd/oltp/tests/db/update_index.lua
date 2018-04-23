pathtest = string.match(test, "(.*/)")

if pathtest then
   dofile(pathtest .. "common.lua")
else
   require("common")
end

function thread_init()
   set_vars()
end

function ff(e) 
	db_query("rollback")
end

function event()
	xpcall(e, ff)
end

function e()
   local table_name
   table_name = "sbtest".. sb_rand_uniform(1, oltp_tables_count)
   db_query("begin")
   rs = db_query("UPDATE ".. table_name .." SET k=k+1 WHERE id=" .. sb_rand(1, oltp_table_size))
   rs = db_query("UPDATE ".. table_name .." SET k=k-1 WHERE id=" .. sb_rand(1, oltp_table_size))
   db_query("commit")
end
