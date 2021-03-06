package data

func getCallData() []byte {
	return getCallDataSample()
}

func getCallDataSample() []byte {
	data := `
[
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-16T22:24:19",
    "duration": 122
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-18T01:34:48",
    "duration": 514
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-21T20:21:17",
    "duration": 120
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-17T22:10:34",
    "duration": 144
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-25T20:24:59",
    "duration": 556
  }, 
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-24T09:38:52",
    "duration": 80
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-20T12:02:37",
    "duration": 9
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-28T06:12:49",
    "duration": 156
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-16T22:31:25",
    "duration": 543
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-19T00:12:47",
    "duration": 132
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-19T08:10:14",
    "duration": 76
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-18T22:47:52",
    "duration": 5356
  }
]
`
	return []byte(data)
}

func getCallDataAll() []byte {
	data := `
[
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-16T22:24:19",
    "duration": 122
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-18T01:34:48",
    "duration": 514
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-21T20:21:17",
    "duration": 120
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-17T22:10:34",
    "duration": 144
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-25T20:24:59",
    "duration": 556
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-23T22:23:25",
    "duration": 336
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+351 272 414 6570",
    "started_at": "2018-09-26T05:34:19",
    "duration": 405
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+263 498 495 0617",
    "started_at": "2018-09-25T22:58:02",
    "duration": 5665
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-23T08:55:18",
    "duration": 822
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+54 398 559 0423",
    "started_at": "2018-09-25T09:10:25",
    "duration": 8494
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-19T20:31:39",
    "duration": 12
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+57 629 420 5680",
    "started_at": "2018-09-17T10:47:21",
    "duration": 29
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+261 860 539 4754",
    "started_at": "2018-09-25T06:21:55",
    "duration": 2851
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-26T03:37:06",
    "duration": 573
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-19T08:19:36",
    "duration": 66
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+63 815 962 6097",
    "started_at": "2018-09-25T19:44:03",
    "duration": 3682
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-20T13:27:42",
    "duration": 32
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-20T00:56:31",
    "duration": 1434
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-18T18:47:17",
    "duration": 166
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-22T17:27:52",
    "duration": 112
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-16T03:38:09",
    "duration": 1142
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-23T14:57:25",
    "duration": 1665
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-27T05:08:53",
    "duration": 365
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-16T01:44:31",
    "duration": 96
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-20T12:27:48",
    "duration": 766
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+27 117 258 4149",
    "started_at": "2018-09-27T11:28:11",
    "duration": 710
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-18T23:24:30",
    "duration": 151
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-26T14:04:33",
    "duration": 5710
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+81 308 988 7153",
    "started_at": "2018-09-24T04:12:07",
    "duration": 9923
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+261 860 539 4754",
    "started_at": "2018-09-21T22:54:31",
    "duration": 4264
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-22T21:17:48",
    "duration": 202
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-17T22:55:06",
    "duration": 151
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-24T10:16:51",
    "duration": 2895
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-23T15:37:45",
    "duration": 251
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-25T11:34:35",
    "duration": 139
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+351 272 414 6570",
    "started_at": "2018-09-22T03:13:47",
    "duration": 140
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+62 107 530 7500",
    "started_at": "2018-09-16T19:18:32",
    "duration": 3660
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-21T21:42:00",
    "duration": 582
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-19T01:00:38",
    "duration": 141
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-24T03:16:48",
    "duration": 89
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-26T19:47:20",
    "duration": 21
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-26T23:47:19",
    "duration": 144
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-18T04:54:04",
    "duration": 163
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-19T02:11:53",
    "duration": 2681
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-23T14:14:42",
    "duration": 492
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+351 272 414 6570",
    "started_at": "2018-09-27T04:00:59",
    "duration": 384
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-17T05:58:16",
    "duration": 2575
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-27T18:02:22",
    "duration": 546
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-17T18:41:52",
    "duration": 869
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-16T04:41:12",
    "duration": 139
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-21T06:44:17",
    "duration": 53
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-27T12:32:32",
    "duration": 457
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-18T03:42:30",
    "duration": 157
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-21T21:20:56",
    "duration": 207
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-25T15:32:57",
    "duration": 500
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-17T23:45:04",
    "duration": 30
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-22T19:17:54",
    "duration": 161
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-22T02:01:08",
    "duration": 306
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+57 629 420 5680",
    "started_at": "2018-09-19T21:03:04",
    "duration": 129
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-22T07:55:23",
    "duration": 594
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-23T02:24:36",
    "duration": 125
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-26T09:21:22",
    "duration": 100
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-21T13:00:15",
    "duration": 172
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+63 815 962 6097",
    "started_at": "2018-09-16T23:11:52",
    "duration": 6789
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-19T07:41:23",
    "duration": 66
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-22T17:26:29",
    "duration": 950
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-16T23:28:04",
    "duration": 144
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+7 171 898 0853",
    "started_at": "2018-09-27T17:33:06",
    "duration": 4868
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+7 171 898 0853",
    "started_at": "2018-09-15T10:03:34",
    "duration": 6298
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-16T20:26:23",
    "duration": 2606
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-28T05:06:44",
    "duration": 886
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+54 398 559 0423",
    "started_at": "2018-09-23T04:48:41",
    "duration": 3458
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+27 117 258 4149",
    "started_at": "2018-09-26T22:32:19",
    "duration": 609
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-25T03:50:51",
    "duration": 68
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+351 272 414 6570",
    "started_at": "2018-09-23T09:20:33",
    "duration": 212
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-17T16:52:29",
    "duration": 156
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-18T09:21:38",
    "duration": 60
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-16T22:11:35",
    "duration": 681
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-21T15:12:44",
    "duration": 124
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-22T05:39:03",
    "duration": 124
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-23T18:08:42",
    "duration": 163
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-25T13:06:40",
    "duration": 45
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+62 107 530 7500",
    "started_at": "2018-09-24T05:05:43",
    "duration": 3924
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-17T21:40:47",
    "duration": 79
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-19T04:16:21",
    "duration": 79
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-16T03:22:12",
    "duration": 988
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-26T00:41:44",
    "duration": 164
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-20T18:55:03",
    "duration": 212
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-25T15:20:43",
    "duration": 283
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 690 597 4443",
    "started_at": "2018-09-26T00:06:19",
    "duration": 2357
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+351 272 414 6570",
    "started_at": "2018-09-22T13:27:01",
    "duration": 157
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-18T07:53:09",
    "duration": 295
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-17T07:13:25",
    "duration": 9460
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-15T22:08:23",
    "duration": 2308
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-23T02:24:15",
    "duration": 1018
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+263 498 495 0617",
    "started_at": "2018-09-26T17:22:34",
    "duration": 10499
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-16T18:17:49",
    "duration": 47
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-17T07:29:13",
    "duration": 1036
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+81 308 988 7153",
    "started_at": "2018-09-27T08:22:32",
    "duration": 6468
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-17T01:58:04",
    "duration": 77
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+263 498 495 0617",
    "started_at": "2018-09-19T23:16:49",
    "duration": 2519
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-15T07:02:32",
    "duration": 914
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+81 308 988 7153",
    "started_at": "2018-09-22T04:57:20",
    "duration": 4455
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+62 107 530 7500",
    "started_at": "2018-09-28T01:57:20",
    "duration": 5272
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-14T17:18:49",
    "duration": 1111
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-23T21:11:41",
    "duration": 290
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-25T22:57:07",
    "duration": 169
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+86 825 153 5518",
    "started_at": "2018-09-20T19:43:24",
    "duration": 276
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-17T14:43:15",
    "duration": 4
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-20T22:48:50",
    "duration": 655
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-17T13:54:54",
    "duration": 146
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-25T20:34:51",
    "duration": 22
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-24T21:14:56",
    "duration": 1177
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+7 690 597 4443",
    "started_at": "2018-09-20T21:01:36",
    "duration": 6227
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-23T11:04:36",
    "duration": 67
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-15T22:31:03",
    "duration": 71
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-19T02:01:36",
    "duration": 60
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-25T05:22:05",
    "duration": 2544
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+62 107 530 7500",
    "started_at": "2018-09-17T18:51:01",
    "duration": 7444
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+81 308 988 7153",
    "started_at": "2018-09-26T03:49:07",
    "duration": 1696
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+54 398 559 0423",
    "started_at": "2018-09-22T15:58:58",
    "duration": 9465
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-21T17:02:25",
    "duration": 442
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-21T16:46:05",
    "duration": 3928
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+7 690 597 4443",
    "started_at": "2018-09-24T08:11:36",
    "duration": 10309
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-24T15:41:04",
    "duration": 76
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-20T01:13:18",
    "duration": 997
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-24T08:53:14",
    "duration": 90
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-25T11:47:59",
    "duration": 124
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-15T21:44:00",
    "duration": 1122
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-21T17:24:29",
    "duration": 520
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+62 107 530 7500",
    "started_at": "2018-09-23T12:10:49",
    "duration": 3251
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-28T08:52:06",
    "duration": 599
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-18T06:03:28",
    "duration": 160
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-17T19:51:51",
    "duration": 79
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+7 171 898 0853",
    "started_at": "2018-09-21T12:08:28",
    "duration": 6538
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-21T02:19:22",
    "duration": 508
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-23T00:35:41",
    "duration": 809
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-16T15:16:47",
    "duration": 2457
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-17T09:48:54",
    "duration": 93
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+54 398 559 0423",
    "started_at": "2018-09-23T03:19:40",
    "duration": 1259
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-26T08:08:28",
    "duration": 6475
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-14T22:01:01",
    "duration": 547
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-23T20:04:51",
    "duration": 4941
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-16T03:35:49",
    "duration": 134
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+7 171 898 0853",
    "started_at": "2018-09-20T11:38:53",
    "duration": 9705
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-28T00:37:16",
    "duration": 6945
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-18T23:25:20",
    "duration": 3225
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-14T23:33:20",
    "duration": 27
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-19T07:52:06",
    "duration": 174
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-16T09:42:58",
    "duration": 143
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+63 815 962 6097",
    "started_at": "2018-09-19T02:17:35",
    "duration": 5413
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+81 308 988 7153",
    "started_at": "2018-09-26T06:19:03",
    "duration": 3552
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-20T08:28:51",
    "duration": 170
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-24T06:33:20",
    "duration": 108
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-25T10:37:27",
    "duration": 556
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-26T11:17:09",
    "duration": 1051
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-21T21:49:34",
    "duration": 802
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-23T11:41:10",
    "duration": 1903
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-21T15:26:33",
    "duration": 175
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-24T06:23:45",
    "duration": 166
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-25T04:08:43",
    "duration": 501
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-20T23:50:37",
    "duration": 52
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+30 419 575 7546",
    "started_at": "2018-09-24T01:11:29",
    "duration": 1959
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-17T18:43:42",
    "duration": 29
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-19T19:07:10",
    "duration": 3196
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-25T06:05:16",
    "duration": 4362
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+63 815 962 6097",
    "started_at": "2018-09-17T16:57:41",
    "duration": 4016
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+86 202 257 8619",
    "started_at": "2018-09-24T19:45:03",
    "duration": 422
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-23T20:00:30",
    "duration": 720
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-16T21:33:26",
    "duration": 402
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-22T06:56:03",
    "duration": 44
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-28T09:45:40",
    "duration": 78
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-27T13:01:03",
    "duration": 113
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+57 629 420 5680",
    "started_at": "2018-09-17T12:20:01",
    "duration": 128
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+36 318 105 5629",
    "started_at": "2018-09-26T08:07:59",
    "duration": 1171
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-20T19:41:35",
    "duration": 126
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-20T01:48:26",
    "duration": 176
  },
  {
    "caller_id": "+261 860 539 4754",
    "callee_id": "+7 690 597 4443",
    "started_at": "2018-09-14T23:40:40",
    "duration": 4389
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+7 414 625 3019",
    "started_at": "2018-09-17T23:24:18",
    "duration": 2391
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+86 922 760 0418",
    "started_at": "2018-09-15T04:32:48",
    "duration": 44
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+86 825 153 5518",
    "started_at": "2018-09-28T14:57:53",
    "duration": 160
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+63 815 962 6097",
    "started_at": "2018-09-16T07:34:42",
    "duration": 9303
  },
  {
    "caller_id": "+7 171 898 0853",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-23T01:37:19",
    "duration": 141
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+63 808 497 1769",
    "started_at": "2018-09-22T08:48:48",
    "duration": 83
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-15T12:12:59",
    "duration": 426
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-20T14:12:40",
    "duration": 51
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+54 398 559 0423",
    "started_at": "2018-09-21T22:45:47",
    "duration": 8861
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+33 614 339 0298",
    "started_at": "2018-09-23T22:18:08",
    "duration": 162
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+27 117 258 4149",
    "started_at": "2018-09-25T00:26:05",
    "duration": 69
  },
  {
    "caller_id": "+263 498 495 0617",
    "callee_id": "+62 533 266 3426",
    "started_at": "2018-09-20T03:03:46",
    "duration": 12
  },
  {
    "caller_id": "+62 107 530 7500",
    "callee_id": "+86 892 682 0628",
    "started_at": "2018-09-26T14:12:03",
    "duration": 68
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+48 195 624 2025",
    "started_at": "2018-09-26T21:20:35",
    "duration": 136
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+86 921 547 9004",
    "started_at": "2018-09-24T09:38:52",
    "duration": 80
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+1 254 875 4647",
    "started_at": "2018-09-20T12:02:37",
    "duration": 9
  },
  {
    "caller_id": "+81 308 988 7153",
    "callee_id": "+48 894 777 5173",
    "started_at": "2018-09-28T06:12:49",
    "duration": 156
  },
  {
    "caller_id": "+370 351 224 5176",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-16T22:31:25",
    "duration": 543
  },
  {
    "caller_id": "+7 690 597 4443",
    "callee_id": "+48 697 447 6933",
    "started_at": "2018-09-19T00:12:47",
    "duration": 132
  },
  {
    "caller_id": "+81 746 154 2598",
    "callee_id": "+351 515 605 7915",
    "started_at": "2018-09-19T08:10:14",
    "duration": 76
  },
  {
    "caller_id": "+54 398 559 0423",
    "callee_id": "+81 746 154 2598",
    "started_at": "2018-09-18T22:47:52",
    "duration": 5356
  },
  {
    "caller_id": "+63 815 962 6097",
    "callee_id": "+7 552 196 4096",
    "started_at": "2018-09-23T01:14:56",
    "duration": 53
  }
]
`
	return []byte(data)
}
