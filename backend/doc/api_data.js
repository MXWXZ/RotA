define({ "api": [
  {
    "type": "WS",
    "url": "GetRooms",
    "title": "GetRooms",
    "version": "1.0.0",
    "group": "Room",
    "permission": [
      {
        "name": "client"
      }
    ],
    "name": "GetRooms",
    "description": "<p>Get room info</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Offset",
            "defaultValue": "0",
            "description": "<p>room amount offset</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "size": "1-100",
            "optional": false,
            "field": "Limit",
            "defaultValue": "20",
            "description": "<p>amount limit</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Token",
            "description": "<p>32 length token</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "list",
            "optional": false,
            "field": "RoomInfo",
            "description": "<p>Room info, see <a href=\"#api-Room-NewRoomRsp\">NewRoomRsp</a></p>"
          }
        ]
      }
    },
    "filename": "msg/room_msg.go",
    "groupTitle": "Room",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "ServerError",
            "description": "<p>Server error</p>"
          }
        ]
      }
    }
  },
  {
    "type": "WS",
    "url": "NewRoom",
    "title": "NewRoom",
    "version": "1.0.0",
    "group": "Room",
    "permission": [
      {
        "name": "client"
      }
    ],
    "name": "NewRoom",
    "description": "<p>New room</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Name",
            "description": "<p>room name</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Type",
            "description": "<p>room type</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Token",
            "description": "<p>32 length token</p>"
          }
        ]
      }
    },
    "filename": "msg/room_msg.go",
    "groupTitle": "Room"
  },
  {
    "type": "WS",
    "url": "NewRoomRsp",
    "title": "NewRoomRsp",
    "version": "1.0.0",
    "group": "Room",
    "permission": [
      {
        "name": "server"
      }
    ],
    "name": "NewRoomRsp",
    "description": "<p>New room response</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "ID",
            "description": "<p>Room id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Name",
            "description": "<p>Room name</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Type",
            "description": "<p>Room type <br> 0 for solo</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Size",
            "description": "<p>Room size</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Capacity",
            "description": "<p>Room capacity</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Master",
            "description": "<p>Room master</p>"
          },
          {
            "group": "Parameter",
            "type": "int32",
            "optional": false,
            "field": "Status",
            "description": "<p>Room status <br> 0 for gaming <br> 1 for waiting</p>"
          }
        ]
      }
    },
    "filename": "msg/room_msg.go",
    "groupTitle": "Room"
  },
  {
    "type": "WS",
    "url": "Login",
    "title": "Login",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "client"
      }
    ],
    "name": "Login",
    "description": "<p>User Login</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UserName",
            "description": "<p>user name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UserPass",
            "description": "<p>user password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "Status",
            "description": "<p>0 for success, 1 for invalid</p>"
          },
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "ID",
            "description": "<p>user id</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "Token",
            "description": "<p>32 length token for success</p>"
          }
        ]
      }
    },
    "filename": "msg/login_msg.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "ServerError",
            "description": "<p>Server error</p>"
          }
        ]
      }
    }
  },
  {
    "type": "WS",
    "url": "Signup",
    "title": "Signup",
    "version": "1.0.0",
    "group": "User",
    "permission": [
      {
        "name": "client"
      }
    ],
    "name": "Signup",
    "description": "<p>User Sign up</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UserName",
            "description": "<p>user name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UserPass",
            "description": "<p>user password</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int32",
            "optional": false,
            "field": "Status",
            "description": "<p>0 for success, 1 for exist</p>"
          }
        ]
      }
    },
    "filename": "msg/login_msg.go",
    "groupTitle": "User",
    "error": {
      "fields": {
        "Error 400": [
          {
            "group": "Error 400",
            "optional": false,
            "field": "InvalidParam",
            "description": "<p>Invalid param</p>"
          }
        ],
        "Error 500": [
          {
            "group": "Error 500",
            "optional": false,
            "field": "ServerError",
            "description": "<p>Server error</p>"
          }
        ]
      }
    }
  }
] });
