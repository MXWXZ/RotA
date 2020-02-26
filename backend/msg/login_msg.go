package msg

func LoginMsgInit() {
	Processor.Register(&Login{})
	Processor.Register(&Signup{})
	Processor.Register(&CheckToken{})
	Processor.Register(&NeedTokenRsp{})
}

/**
 * @api {WS} Login Login
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission client
 * @apiName Login
 * @apiDescription User Login
 *
 * @apiParam {string} UserName user name
 * @apiParam {string} UserPass user password
 * @apiSuccess {int} Status 0 for success, 1 for invalid, 2 for server error
 * @apiSuccess {int} ID user id
 * @apiSuccess {string} Token 64 length token for success
 */
type Login struct {
	UserName string
	UserPass string
}

type LoginRsp struct {
	Status int
	ID     int
	Token  string
}

/**
 * @api {WS} Signup Signup
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission client
 * @apiName Signup
 * @apiDescription User Sign up
 *
 * @apiParam {string} UserName user name
 * @apiParam {string} UserPass user password
 * @apiSuccess {int} Status 0 for success, 1 for exist, 2 for server error
 */
type Signup struct {
	UserName string
	UserPass string
}

type SignupRsp struct {
	Status int
}

/**
 * @api {WS} CheckToken CheckToken
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission client
 * @apiName CheckToken
 * @apiDescription Check user token
 *
 * @apiParam {int} ID user id
 * @apiParam {string} Token 64 length token
 * @apiSuccess {int} Code 0 for success, 1 for invalid
 * @apiSuccess {int} ID user id
 * @apiSuccess {string} Name user name
 * @apiSuccess {int} Room user room
 * @apiSuccess {int} Status user status
 */
type CheckToken struct {
	ID    int
	Token string
}

type CheckTokenRsp struct {
	Code   int
	ID     int
	Name   string
	Room   int
	Status int
}

/**
 * @api {WS} NeedTokenRsp NeedTokenRsp
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission server
 * @apiName NeedTokenRsp
 * @apiDescription Ask client to check token
 */
type NeedTokenRsp struct {
}
