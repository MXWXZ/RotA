package msg

func ChatMsgInit() {
	Processor.Register(&Chat{})
}

/**
 * @api {WS} Chat Chat
 * @apiVersion 1.0.0
 * @apiGroup Chat
 * @apiPermission client
 * @apiName Chat
 * @apiDescription Chat msg
 *
 * @apiParam {int} Channel Chat channel <br> 1 for room
 * @apiParam {string} Msg Chat msg
 */
type Chat struct {
	Channel ChatChan
	Msg     string
}

type ChatChan int

const (
	Chat_Room ChatChan = 1
)

/**
 * @api {WS} ChatRsp ChatRsp
 * @apiVersion 1.0.0
 * @apiGroup ChatRsp
 * @apiPermission server
 * @apiName Chat
 * @apiDescription New chat msg
 *
 * @apiParam {string} Name user name
 * @apiParam {int} Channel Chat channel <br> 1 for local
 * @apiParam {string} Msg Chat msg
 */
type ChatRsp struct {
	Name    string
	Channel ChatChan
	Msg     string
}
