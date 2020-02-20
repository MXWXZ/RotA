package msg

func RoomMsgInit() {
	Processor.Register(&GetRooms{})
	Processor.Register(&NewRoom{})
}

/**
 * @api {WS} GetRooms GetRooms
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission client
 * @apiName GetRooms
 * @apiDescription Get room info
 *
 * @apiParam {int32} Offset=0 room amount offset
 * @apiParam {int32{1-100}} Limit=20 amount limit
 */
type GetRooms struct {
	Offset int32
	Limit  int32
}

/**
 * @api {WS} GetRoomsRsp GetRoomsRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName GetRoomsRsp
 * @apiDescription Get room info rsp
 *
 * @apiParam {list} RoomInfo Room info, see [NewRoomRsp](#api-Room-NewRoomRsp)
 */
type GetRoomsRsp struct {
	RoomInfo []RoomInfo
}

/**
 * @api {WS} NewRoom NewRoom
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission client
 * @apiName NewRoom
 * @apiDescription New room
 *
 * @apiParam {string} Name room name
 * @apiParam {int32} Type room type
 */
type NewRoom struct {
	Name string
	Type int32
}

/**
 * @api {WS} NewRoomRsp NewRoomRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName NewRoomRsp
 * @apiDescription New room response
 *
 * @apiParam {int32} ID Room id
 * @apiParam {string} Name Room name
 * @apiParam {int32} Type Room type <br> 0 for solo
 * @apiParam {int32} Size Room size
 * @apiParam {int32} Capacity Room capacity
 * @apiParam {int32} Master Room master
 * @apiParam {int32} Status Room status <br> 0 for gaming <br> 1 for waiting
 */
type NewRoomRsp RoomInfo

type RoomType int32

const (
	Room_Solo RoomType = 0
)

type RoomInfo struct {
	ID       int32
	Name     string
	Type     RoomType
	Size     int32
	Capacity int32
	Master   int32
	Status   int32
}
