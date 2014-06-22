package db

import (
	"labix.org/v2/mgo/bson"
)

type M_Mem struct {
	//desc in KB
	Total  int64 //总内存大小
	Free   int64 //空闲内存
	Cached int64 //缓存内存

	Allocated int64   //总预分配大小
	Percent   float32 //规定的预分配标准
	//memery box amount
	//Box1G  int8
	//Box5G  int8
	//Box10G int8
}

type M_CPU struct {
}

type M_Net struct {
}

type M_Disk struct {
}

//describe a machine's status
type Machine struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Mtype  string        //机器类型
	Host   string        //主机名
	Status bool          //主机状态，是否被封禁
	Mroom  string        //机房
	Logic  string        //逻辑机房
	Region string        //地域
	Idc    string        //IDC
	Mem    M_Mem         //内存描述
	Cpu    M_CPU         //CPU描述
	Net    M_Net         //网络描述
	Disk   M_Disk        //硬盘描述
}

//describe a service
type Service struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Stype   string        //服务类型
	Host    string        //服务所在机器
	Port    int32         //服务端口
	MaxMem  int64         //服务申请的最大内存
	DirName string        //服务的dir
	UsedMem int64         //使用的内存
	Role    string        //角色
	Version string        //版本号

}
