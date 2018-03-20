package Config

import (
	"MSvrs/Core/Utils"
	"strconv"

	"github.com/stathat/go"
)

/*
	use consistent router db
	1,table_x = table_(user_id % 8)
	2,db_x = db_consistent(n)
*/

//mysql
const (
	max              = 8
	mysql_src        = "root:@(localhost:3306)/"
	db_prefix        = "msvrs"
	db_suffix        = "?charset=utf8"
	tlb_users_prefix = "users"
)

//redis-mq
const (
	Redis_addr = "localhost:6379"

	Topic_echo    = "TP_ECHO"
	Topic_gate    = "TP_Gate"
	Topic_storage = "TP_Storage"

	sub = "_Sub"
	pub = "_Pub"
)

//json-rpc
const (
	Json_rpc_addr = "127.0.0.1:8080"
)

/*use jwt tols https://jwt.io/*/
const (
	Port_Auth  = 8080
	Port_Login = 8081
	Port_Reg   = 8082

	JWT_rs256_public = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDdlatRjRjogo3WojgGHFHYLugd
UWAY9iR3fy4arWNA1KoS8kVw33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQs
HUfQrSDv+MuSUMAe8jzKE4qW+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5D
o2kQ+X5xK9cipRgEKwIDAQAB
-----END PUBLIC KEY-----`

	JWT_rs256_private = `-----BEGIN RSA PRIVATE KEY----- 
MIICWwIBAAKBgQDdlatRjRjogo3WojgGHFHYLugdUWAY9iR3fy4arWNA1KoS8kVw
33cJibXr8bvwUAUparCwlvdbH6dvEOfou0/gCFQsHUfQrSDv+MuSUMAe8jzKE4qW
+jK+xQU9a03GUnKHkkle+Q0pX/g6jXZ7r1/xAK5Do2kQ+X5xK9cipRgEKwIDAQAB
AoGAD+onAtVye4ic7VR7V50DF9bOnwRwNXrARcDhq9LWNRrRGElESYYTQ6EbatXS
3MCyjjX2eMhu/aF5YhXBwkppwxg+EOmXeh+MzL7Zh284OuPbkglAaGhV9bb6/5Cp
uGb1esyPbYW+Ty2PC0GSZfIXkXs76jXAu9TOBvD0ybc2YlkCQQDywg2R/7t3Q2OE
2+yo382CLJdrlSLVROWKwb4tb2PjhY4XAwV8d1vy0RenxTB+K5Mu57uVSTHtrMK0
GAtFr833AkEA6avx20OHo61Yela/4k5kQDtjEf1N0LfI+BcWZtxsS3jDM3i1Hp0K
Su5rsCPb8acJo5RO26gGVrfAsDcIXKC+bQJAZZ2XIpsitLyPpuiMOvBbzPavd4gY
6Z8KWrfYzJoI/Q9FuBo6rKwl4BFoToD7WIUS+hpkagwWiz+6zLoX1dbOZwJACmH5
fSSjAkLRi54PKJ8TFUeOP15h9sQzydI8zJU+upvDEKZsZc/UhT/SySDOxQ4G/523
Y0sz/OZtSWcol/UMgQJALesy++GdvoIDLfJX5GBQpuFgFenRiRDabxrE9MNUZ2aP
FaFp+DyAe+b4nDwuJaW2LURbr8AEZga7oQj0uYxcYw==
-----END RSA PRIVATE KEY-----`
)

//svrs reg
const (
	Topic_Svrs_Reg   = "TP_Svrs_Reg"
	Topic_Svrs_Unreg = "TP_Svrs_Unreg"

	Topic_Svrs_CreateUser_Req = "TP_Svrs_CreateUser_Req"
	Topic_Svrs_CreateUser_Rep = "TP_Svrs_CreateUser_Rep"

	Type_Svrs_Echo_Auth  = ""
	Type_Svrs_Echo_Login = ""
	Type_Svrs_Echo_Chat  = ""
)

func init() {
	stathat.PostEZCount("user created", "cn0512@126.com", 1)
}

func GetMysqlSrc(id int) string {
	return mysql_src + getDBPrefix(id) + db_suffix
}
func getDBPrefix(id int) string {
	return db_prefix + strconv.Itoa(Utils.GetConsistentDB(id))
}

func GetTableUsers(userId int64) string {
	return tlb_users_prefix + strconv.FormatInt(userId%max, 10)
}
