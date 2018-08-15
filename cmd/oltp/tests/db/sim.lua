pathtest = string.match(test, "(.*/)")

if pathtest then
   dofile(pathtest .. "common.lua")
else
   require("common")
end

function thread_init()
   set_vars()
end

function event()
   local table_name
   db_query("begin")
   table_name = "sbtest".. sb_rand_uniform(1, oltp_tables_count)
   rs = db_query("SELECT pad FROM ".. table_name .." WHERE id=" .. sb_rand(1, oltp_table_size))
   rs = db_query("UPDATE " .. table_name .. " SET k=k+1 WHERE id=" .. sb_rand(1, oltp_table_size))
   rs = db_query("UPDATE " .. table_name .. " SET k=k+1 WHERE id=" .. sb_rand(1, oltp_table_size))
   c_val = sb_rand_str("###########-###########-###########-###########-###########-###########-###########-###########-###########-###########")
   query = "UPDATE " .. table_name .. " SET c='" .. c_val .. "' WHERE id=" .. sb_rand(1, oltp_table_size)
   rs = db_query(query)
   i = sb_rand(1, oltp_table_size)
   rs = db_query("DELETE FROM " .. table_name .. " WHERE id=" .. i)

   c_val = sb_rand_str([[
###########-###########-###########-###########-###########-###########-###########-###########-###########-###########]])
   pad_val = sb_rand_str([[
###########-###########-###########-###########-###########]])

   rs = db_query("INSERT INTO " .. table_name ..  " (id, k, c, pad) VALUES " .. string.format("(%d, %d, '%s', '%s')",i, sb_rand(1, oltp_table_size) , c_val, pad_val))
   db_query("commit")
end

