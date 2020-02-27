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
 * @apiParam {int} Offset=0 room amount offset
 * @apiParam {int{1-20}} Limit=20 amount limit
 */
type GetRooms struct {
	Offset int
	Limit  int
}

/**
 * @api {WS} GetRoomsRsp GetRoomsRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName GetRoomsRsp
 * @apiDescription Get room info rsp
 *
 * @apiParam {int} Total total room number
 * @apiParam {list} RoomInfo Room info, see [NewRoomRsp](#api-Room-NewRoomRsp)
 */
type GetRoomsRsp struct {
	Total    int
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
 * @apiParam {int} Type room type
 */
type NewRoom struct {
	Name string
	Type int
}

/**
 * @api {WS} NewRoomRsp NewRoomRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName NewRoomRsp
 * @apiDescription New room response
 *
 * @apiParam {int} ID Room id
 * @apiParam {string} Name Room name
 * @apiParam {int} Type Room type <br> 1 for solo
 * @apiParam {int} Size Room size
 * @apiParam {int} Capacity Room capacity
 * @apiParam {int} Master Room master
 * @apiParam {int} Status Room status <br> 0 for waiting <br> 1 for gaming
 */
type NewRoomRsp RoomInfo

type RoomType int

const (
	Room_Solo RoomType = 1
)

type RoomInfo struct {
	ID       int
	Name     string
	Type     RoomType
	Size     int
	Capacity int
	Master   int
	Status   int
	Members  []int `json:"-"`
}

/**
 * @api {WS} DeleteRoomRsp DeleteRoomRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName DeleteRoomRsp
 * @apiDescription Delete room
 *
 * @apiParam {int} ID room id
 */
type DeleteRoomRsp struct {
	ID int
}

/**
 * @api {WS} RoomInfoRsp RoomInfoRsp
 * @apiVersion 1.0.0
 * @apiGroup Room
 * @apiPermission server
 * @apiName RoomInfoRsp
 * @apiDescription Room current info
 *
 * @apiParam {-} - see [NewRoomRsp](#api-Room-NewRoomRsp)
 */
type RoomInfoRsp RoomInfo
