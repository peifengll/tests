package main

import "testing"

var s = `
"res": {
        "location": [
          {
            "category": 0,
            "car_id": "6eda84f3f1de4feaae9814e1b714977b",
            "cell_id": 65484,
            "type": 1,
            "degree": 322,
            "timestamp": 1574672308,
            "altitude": 1,
            "mcc": 460,
            "t_type": "ZJ210",
            "longitude": 408545640,
            "locate_error": 6,
            "tid": "f4a87b51313b4f758f9b4da38f6b59f9",
            "snr": 40,
            "address": "",
            "latitude": 80277480,
            "clongitude": 408588067,
            "clatitude": 80288988,
            "mnc": 0,
            "speed": 42,
            "locate_type": 1
          }
        ]
      },
      "packet_type": "S13"
}
`

func TestOp(t *testing.T) {

}
