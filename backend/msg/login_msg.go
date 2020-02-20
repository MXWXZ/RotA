package msg

func LoginMsgInit() {
	Processor.Register(&Login{})
	Processor.Register(&Signup{})
	Processor.Register(&CheckToken{})
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
 * @apiSuccess {int32} Status 0 for success, 1 for invalid
 * @apiSuccess {int32} ID user id
 * @apiSuccess {string} Token 64 length token for success
 * @apiUse InvalidParam
 * @apiUse ServerError
 */
type Login struct {
	UserName string
	UserPass string
}

type LoginRsp struct {
	Status int32
	ID     int32
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
 * @apiSuccess {int32} Status 0 for success, 1 for exist
 * @apiUse InvalidParam
 * @apiUse ServerError
 */
type Signup struct {
	UserName string
	UserPass string
}

type SignupRsp struct {
	Status int32
}

/**
 * @api {WS} CheckToken CheckToken
 * @apiVersion 1.0.0
 * @apiGroup User
 * @apiPermission client
 * @apiName CheckToken
 * @apiDescription Check user token
 *
 * @apiParam {int32} ID user id
 * @apiParam {string} Token 64 length token
 * @apiSuccess {int32} Status 0 for success, 1 for invalid
 * @apiUse InvalidParam
 * @apiUse ServerError
 */
type CheckToken struct {
	ID    int32
	Token string
}

type CheckTokenRsp struct {
	Status int32
}
