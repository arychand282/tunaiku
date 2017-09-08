1. Enter to MySQL
2. Create a database MySQL "tunaiku"
3. Import sql file to "tunaiku" database
4. Enter "tunaiku" database
5. Make a custom function for 'converting number to string' from file number_to_string.txt in mysql
6. If you get an error: `#1418 - This function has none of DETERMINISTIC, NO SQL, or READS SQL DATA in its declaration and binary logging is enabled (you might want to use the less safe log_bin_trust_function_creators variable)`, run this query: `SET GLOBAL log_bin_trust_function_creators = 1;`
7. Please install following golang dependencies:

$ go get "github.com/go-sql-driver/mysql"
$ go get "github.com/gin-gonic/gin"



