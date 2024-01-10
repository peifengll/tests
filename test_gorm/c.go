package test_gorm


package data_obj

/*
 * @Author: SimingLiu siming.liu@linketech.cn
 * @Date:  2022-02-18 14:02:35
 * @LastEditors: SimingLiu siming.liu@linketech.cn
 * @LastEditTime: 2023-08-01 10:47:20
 * @FilePath: \zulinbao-server\infrastructure\data_obj\qjcg.go
 * @Description:
 *
 */
import "database/sql"

// ILocation 位置信息。
type ILocation struct {
	ID          uint64  `db:"id" gorm:"Column:id"`
	CarID       *uint64 `db:"car_id" gorm:"Column:car_id"`
	Tid         uint64  `db:"tid" gorm:"Column:tid"`
	Category    int8    `db:"category" gorm:"Column:category"`
	Latitude    int64   `db:"latitude" gorm:"Column:latitude"`
	Longitude   int64   `db:"longitude" gorm:"Column:longitude"`
	Clatitude   int64   `db:"clatitude" gorm:"Column:clatitude"`
	Clongitude  int64   `db:"clongitude" gorm:"Column:clongitude"`
	Address     string  `db:"address" gorm:"Column:address"`
	Altitude    int64   `db:"altitude" gorm:"Column:altitude"`
	Speed       int32   `db:"speed" gorm:"Column:speed"`
	Degree      float32 `db:"degree" gorm:"Column:degree"`
	LocateError int32   `db:"locate_error" gorm:"Column:locate_error"`
	Snr         uint32  `db:"snr" gorm:"Column:snr"`
	Mcc         *string `db:"mcc" gorm:"Column:mcc"`
	Mnc         *string `db:"mnc" gorm:"Column:mnc"`
	Lac         *string `db:"lac" gorm:"Column:lac"`
	CellID      *string `db:"cell_id" gorm:"Column:cell_id"`
	Timestamp   int64   `db:"timestamp" gorm:"Column:timestamp"`
	TType       *string `db:"t_type" gorm:"Column:t_type"`
	LocateType  int32   `db:"locate_type" gorm:"Column:locate_type"`
	RxLev       *int32  `db:"rxlev" gorm:"Column:rxlev"`
}

func (ILocation) TableName() string {
	return "i_location"
}

type IDMap struct {
	StrID string `db:"str_id" gorm:"Column:str_id"`
	U64ID uint64 `db:"u64_id" gorm:"Column:u64_id"`
}

func (IDMap) TableName() string {
	return "id_map"
}

// RenewCard SIM卡续费表
type RenewCard struct {
	ID          string  `db:"id" gorm:"Column:id"`
	Mobile      string  `db:"mobile" gorm:"Column:mobile"`
	Iccid       string  `db:"iccid" gorm:"Column:iccid"`
	Sn          *string `db:"sn" gorm:"Column:sn"`
	Operator    *string `db:"operator" gorm:"Column:operator"`
	RenewTime   *int64  `db:"renew_time" gorm:"Column:renew_time"`
	ExpiredTime *int64  `db:"expired_time" gorm:"Column:expired_time"`
	CorpName    *string `db:"corp_name" gorm:"Column:corp_name"`
	CreateTime  *int64  `db:"create_time" gorm:"Column:create_time"`
}

func (RenewCard) TableName() string {
	return "renew_card"
}

// RenewDevice 设备续费表
type RenewDevice struct {
	ID            string  `db:"id" gorm:"Column:id"`
	Sn            string  `db:"sn" gorm:"Column:sn"`
	Mobile        *string `db:"mobile" gorm:"Column:mobile"`
	Iccid         *string `db:"iccid" gorm:"Column:iccid"`
	Operator      *string `db:"operator" gorm:"Column:operator"`
	DeliveredTime *int64  `db:"delivered_time" gorm:"Column:delivered_time"`
	RenewTime     *int64  `db:"renew_time" gorm:"Column:renew_time"`
	ExpiredTime   *int64  `db:"expired_time" gorm:"Column:expired_time"`
	CorpName      *string `db:"corp_name" gorm:"Column:corp_name"`
	CreateTime    *int64  `db:"create_time" gorm:"Column:create_time"`
}

func (RenewDevice) TableName() string {
	return "renew_device"
}

type T808Sn struct {
	ID              string `db:"id" gorm:"Column:id"`
	EzeIdentifyCode string `db:"eze_identify_code" gorm:"Column:eze_identify_code"`
	Sn              string `db:"sn" gorm:"Column:sn"`
	CreateTime      int64  `db:"create_time" gorm:"Column:create_time"`
}

func (T808Sn) TableName() string {
	return "t_808_sn"
}

// AlarmRule 特殊事件告警规则
type AlarmRule struct {
	ID         int64   `db:"id" gorm:"Column:id"`
	Oid        string  `db:"oid" gorm:"Column:oid"`
	EventName  *string `db:"event_name" gorm:"Column:event_name"`
	EventNote  string  `db:"event_note" gorm:"Column:event_note"`
	EventType  int32   `db:"event_type" gorm:"Column:event_type"`
	Status     int32   `db:"status" gorm:"Column:status"`
	TimeLimit  *int32  `db:"time_limit" gorm:"Column:time_limit"`
	SpeedLimit *int32  `db:"speed_limit" gorm:"Column:speed_limit"`
	Rid        *int32  `db:"rid" gorm:"Column:rid"`
	StartTime1 *string `db:"start_time1" gorm:"Column:start_time1"`
	EndTime1   *string `db:"end_time1" gorm:"Column:end_time1"`
	StartTime2 *string `db:"start_time2" gorm:"Column:start_time2"`
	EndTime2   *string `db:"end_time2" gorm:"Column:end_time2"`
	StartTime3 *string `db:"start_time3" gorm:"Column:start_time3"`
	EndTime3   *string `db:"end_time3" gorm:"Column:end_time3"`
	StartTime4 *string `db:"start_time4" gorm:"Column:start_time4"`
	EndTime4   *string `db:"end_time4" gorm:"Column:end_time4"`
	StartTime5 *string `db:"start_time5" gorm:"Column:start_time5"`
	EndTime5   *string `db:"end_time5" gorm:"Column:end_time5"`
	StartTime6 *string `db:"start_time6" gorm:"Column:start_time6"`
	EndTime6   *string `db:"end_time6" gorm:"Column:end_time6"`
	StartTime7 *string `db:"start_time7" gorm:"Column:start_time7"`
	EndTime7   *string `db:"end_time7" gorm:"Column:end_time7"`
	Week       string  `db:"week" gorm:"Column:week"`
	Method     int32   `db:"method" gorm:"Column:method"`
}

func (AlarmRule) TableName() string {
	return "t_alarm_rule"
}

type Announcement struct {
	ID              string `db:"id" gorm:"Column:id"`
	Type            uint8  `db:"type" gorm:"Column:type"`
	Content         string `db:"content" gorm:"Column:content"`
	NoticeStartTime uint64 `db:"notice_start_time" gorm:"Column:notice_start_time"`
	NoticeEndTime   uint64 `db:"notice_end_time" gorm:"Column:notice_end_time"`
	CreateTime      uint64 `db:"create_time" gorm:"Column:create_time"`
	UpdateTime      uint64 `db:"update_time" gorm:"Column:update_time"`
	Uid             string `db:"uid" gorm:"Column:uid"`
}

func (Announcement) TableName() string {
	return "t_announcement"
}

type AppAddress struct {
	Target   string  `db:"target" gorm:"Column:target"`
	Ios      *string `db:"ios" gorm:"Column:ios"`
	Android  *string `db:"android" gorm:"Column:android"`
	Default  *string `db:"default" gorm:"Column:default"`
	BaseUrl  string  `db:"base_url" gorm:"Column:base_url"`
	CreateAt *int32  `db:"create_at" gorm:"Column:create_at"`
	Weixin   *string `db:"weixin" gorm:"Column:weixin"`
}

func (AppAddress) TableName() string {
	return "t_app_address"
}

// AreaCode 国家行政区划
type AreaCode struct {
	Code     string   `db:"code" gorm:"Column:code"`
	Pn       *string  `db:"pn" gorm:"Column:pn"`
	Cn       *string  `db:"cn" gorm:"Column:cn"`
	Dn       *string  `db:"dn" gorm:"Column:dn"`
	Name     *string  `db:"name" gorm:"Column:name"`
	FullName *string  `db:"full_name" gorm:"Column:full_name"`
	MinLat   *float64 `db:"min_lat" gorm:"Column:min_lat"`
	MaxLat   *float64 `db:"max_lat" gorm:"Column:max_lat"`
	MinLng   *float64 `db:"min_lng" gorm:"Column:min_lng"`
	MaxLng   *float64 `db:"max_lng" gorm:"Column:max_lng"`
	Polygons string   `db:"polygons" gorm:"Column:polygons"`
	AmapCode *string  `db:"amap_code" gorm:"Column:amap_code"`
}

func (AreaCode) TableName() string {
	return "t_area_code"
}

type AreaGroup struct {
	Cid    string `db:"cid" gorm:"Column:cid"`
	AreaNo string `db:"area_no" gorm:"Column:area_no"`
	Gid    string `db:"gid" gorm:"Column:gid"`
}

func (AreaGroup) TableName() string {
	return "t_area_group"
}

type BakLocation struct {
	Sn          string  `db:"sn" gorm:"Column:sn"`
	GPSLid      *int64  `db:"gps_lid" gorm:"Column:gps_lid"`
	CellLid     *int64  `db:"cell_lid" gorm:"Column:cell_lid"`
	WiFiLid     *int64  `db:"wifi_lid" gorm:"Column:wifi_lid"`
	ID          uint64  `db:"id" gorm:"Column:id"`
	CarID       *string `db:"car_id" gorm:"Column:car_id"`
	Tid         string  `db:"tid" gorm:"Column:tid"`
	Category    int8    `db:"category" gorm:"Column:category"`
	Latitude    int64   `db:"latitude" gorm:"Column:latitude"`
	Longitude   int64   `db:"longitude" gorm:"Column:longitude"`
	Clatitude   int64   `db:"clatitude" gorm:"Column:clatitude"`
	Clongitude  int64   `db:"clongitude" gorm:"Column:clongitude"`
	Address     string  `db:"address" gorm:"Column:address"`
	Altitude    int64   `db:"altitude" gorm:"Column:altitude"`
	Speed       int32   `db:"speed" gorm:"Column:speed"`
	Degree      float32 `db:"degree" gorm:"Column:degree"`
	LocateError int32   `db:"locate_error" gorm:"Column:locate_error"`
	Snr         uint32  `db:"snr" gorm:"Column:snr"`
	Mcc         *string `db:"mcc" gorm:"Column:mcc"`
	Mnc         *string `db:"mnc" gorm:"Column:mnc"`
	Lac         *string `db:"lac" gorm:"Column:lac"`
	CellID      *string `db:"cell_id" gorm:"Column:cell_id"`
	Timestamp   int64   `db:"timestamp" gorm:"Column:timestamp"`
	TType       *string `db:"t_type" gorm:"Column:t_type"`
	LocateType  int32   `db:"locate_type" gorm:"Column:locate_type"`
	RxLev       *int32  `db:"rxlev" gorm:"Column:rxlev"`
}

func (BakLocation) TableName() string {
	return "t_bak_location"
}

// Batch 批次表
type Batch struct {
	ID        int32   `db:"id" gorm:"Column:id"`
	Bid       string  `db:"bid" gorm:"Column:bid"`
	Iid       string  `db:"iid" gorm:"Column:iid"`
	Name      string  `db:"name" gorm:"Column:name"`
	Timestamp int64   `db:"timestamp" gorm:"Column:timestamp"`
	Backup    *string `db:"backup" gorm:"Column:backup"`
	BatchNo   string  `db:"batch_no" gorm:"Column:batch_no"`
}

func (Batch) TableName() string {
	return "t_batch"
}

// BatchTerminal 批次与设备关系表
type BatchTerminal struct {
	Bid string `db:"bid" gorm:"Column:bid"`
	Tid string `db:"tid" gorm:"Column:tid"`
}

func (BatchTerminal) TableName() string {
	return "t_batch_terminal"
}

// BindLog 绑定解绑表。记录终端的绑定和解绑动作
type BindLog struct {
	ID           uint64  `db:"id" gorm:"Column:id"`
	MainOid      *string `db:"main_oid" gorm:"Column:main_oid"`
	MainMobile   *string `db:"main_mobile" gorm:"Column:main_mobile"`
	Tid          *string `db:"tid" gorm:"Column:tid"`
	Sn           *string `db:"sn" gorm:"Column:sn"`
	Iccid        *string `db:"iccid" gorm:"Column:iccid"`
	OpType       uint8   `db:"op_type" gorm:"Column:op_type"`
	OpTime       *uint64 `db:"op_time" gorm:"Column:op_time"`
	Vin          *string `db:"vin" gorm:"Column:vin"`
	Cnum         *string `db:"cnum" gorm:"Column:cnum"`
	UnbindReason *string `db:"unbind_reason" gorm:"Column:unbind_reason"`
}

func (BindLog) TableName() string {
	return "t_bind_log"
}

// Car 车辆表
type Car struct {
	ID                 uint64  `db:"id" gorm:"Column:id"`
	CarID              string  `db:"car_id" gorm:"Column:car_id"`
	Cid                *string `db:"cid" gorm:"Column:cid"`
	Vin                *string `db:"vin" gorm:"Column:vin"`
	Cnum               *string `db:"cnum" gorm:"Column:cnum"`
	ModelName          *string `db:"model_name" gorm:"Column:model_name"`
	Alias              *string `db:"alias" gorm:"Column:alias"`
	Type               *string `db:"type" gorm:"Column:type"`
	Pid                *string `db:"pid" gorm:"Column:pid"`
	Icon               *int32  `db:"icon" gorm:"Column:icon"`
	OilVolume          float64 `db:"oil_volume" gorm:"Column:oil_volume"`
	Oil                *int32  `db:"oil" gorm:"Column:oil"`
	Status             *string `db:"status" gorm:"Column:status"`
	CreateTime         *int64  `db:"create_time" gorm:"Column:create_time"`
	AlertMobile        *string `db:"alert_mobile" gorm:"Column:alert_mobile"`
	EmergencyStatus    int16   `db:"emergency_status" gorm:"Column:emergency_status"`
	Emergentable       int16   `db:"emergentable" gorm:"Column:emergentable"`
	Matchable          int8    `db:"matchable" gorm:"Column:matchable"`
	Style              int32   `db:"style" gorm:"Column:style"`
	Mileage            *int64  `db:"mileage" gorm:"Column:mileage"`
	TotalMileage       *int64  `db:"total_mileage" gorm:"Column:total_mileage"`
	CarKey             *int8   `db:"car_key" gorm:"Column:car_key"`
	Remark             *string `db:"remark" gorm:"Column:remark"`
	CipherLock         *int32  `db:"cipher_lock" gorm:"Column:cipher_lock"`
	CipherLockDoStatus *int8   `db:"cipher_lock_do_status" gorm:"Column:cipher_lock_do_status"`
	DriverFrontDoor    *int8   `db:"driver_front_door" gorm:"Column:driver_front_door"`
	CopilotFrontDoor   *int8   `db:"copilot_front_door" gorm:"Column:copilot_front_door"`
	DriverBackDoor     *int8   `db:"driver_back_door" gorm:"Column:driver_back_door"`
	CopilotBackDoor    *int8   `db:"copilot_back_door" gorm:"Column:copilot_back_door"`
	DoorLock           *int8   `db:"door_lock" gorm:"Column:door_lock"`
	CarDoorDoStatus    *int8   `db:"car_door_do_status" gorm:"Column:car_door_do_status"`
	Rmp                *int32  `db:"rmp" gorm:"Column:rmp"`
	Gear               *int32  `db:"gear" gorm:"Column:gear"`
	Brake              *int32  `db:"brake" gorm:"Column:brake"`
	Parking            *int32  `db:"parking" gorm:"Column:parking"`
	DoorWindowLf       *int8   `db:"door_window_lf" gorm:"Column:door_window_lf"`
	DoorWindowRf       *int8   `db:"door_window_rf" gorm:"Column:door_window_rf"`
	DoorWindowLb       *int8   `db:"door_window_lb" gorm:"Column:door_window_lb"`
	DoorWindowRb       *int8   `db:"door_window_rb" gorm:"Column:door_window_rb"`
	LowBeam            *int8   `db:"low_beam" gorm:"Column:low_beam"`
	HighBeam           *int8   `db:"high_beam" gorm:"Column:high_beam"`
	PositionLamp       *int8   `db:"position_lamp" gorm:"Column:position_lamp"`
	EmergencyLamp      *int8   `db:"emergency_lamp" gorm:"Column:emergency_lamp"`
	FoglightFront      *int8   `db:"foglight_front" gorm:"Column:foglight_front"`
	FoglightBack       *int8   `db:"foglight_back" gorm:"Column:foglight_back"`
	Trunk              *int8   `db:"trunk" gorm:"Column:trunk"`
	Speed              int32   `db:"speed" gorm:"Column:speed"`
}

func (Car) TableName() string {
	return "t_car"
}

type CarOperator struct {
	Oid   string `db:"oid" gorm:"Column:oid"`
	CarID string `db:"car_id" gorm:"Column:car_id"`
}

func (CarOperator) TableName() string {
	return "t_car_operator"
}

type CarTerBindLog struct {
	ID      int64   `db:"id" gorm:"Column:id"`
	Cid     string  `db:"cid" gorm:"Column:cid"`
	Vin     *string `db:"vin" gorm:"Column:vin"`
	Sn      string  `db:"sn" gorm:"Column:sn"`
	DevType string  `db:"dev_type" gorm:"Column:dev_type"`
	OpType  string  `db:"op_type" gorm:"Column:op_type"`
	OpTime  int32   `db:"op_time" gorm:"Column:op_time"`
	OSn     *string `db:"o_sn" gorm:"Column:o_sn"`
}

func (CarTerBindLog) TableName() string {
	return "t_car_ter_bind_log"
}

type CarTerminal struct {
	CarID string `db:"car_id" gorm:"Column:car_id"`
	Tid   string `db:"tid" gorm:"Column:tid"`
}

func (CarTerminal) TableName() string {
	return "t_car_terminal"
}

// Card 手机卡表
type Card struct {
	ID           int64   `db:"id" gorm:"Column:id"`
	Iccid        string  `db:"iccid" gorm:"Column:iccid"`
	Mobile       *string `db:"mobile" gorm:"Column:mobile"`
	Imsi         *string `db:"imsi" gorm:"Column:imsi"`
	Type         *int8   `db:"type" gorm:"Column:type"`
	Supplier     *string `db:"supplier" gorm:"Column:supplier"`
	Attribution  *string `db:"attribution" gorm:"Column:attribution"`
	InputTime    *int64  `db:"input_time" gorm:"Column:input_time"`
	Buyer        string  `db:"buyer" gorm:"Column:buyer"`
	ActivateTime *uint64 `db:"activate_time" gorm:"Column:activate_time"`
	Validity     uint8   `db:"validity" gorm:"Column:validity"`
}

func (Card) TableName() string {
	return "t_card"
}

// Corp 集团表
type Corp struct {
	ID                        int32   `db:"id" gorm:"Column:id"`
	Cid                       string  `db:"cid" gorm:"Column:cid"`
	Name                      string  `db:"name" gorm:"Column:name"`
	Mobile                    string  `db:"mobile" gorm:"Column:mobile"`
	AlertMobile               string  `db:"alert_mobile" gorm:"Column:alert_mobile"`
	Linkman                   string  `db:"linkman" gorm:"Column:linkman"`
	Address                   *string `db:"address" gorm:"Column:address"`
	Email                     *string `db:"email" gorm:"Column:email"`
	CreateTime                int64   `db:"create_time" gorm:"Column:create_time"`
	NameShow                  int8    `db:"name_show" gorm:"Column:name_show"`
	SpeedLimit                int32   `db:"speed_limit" gorm:"Column:speed_limit"`
	LongStopMin               int32   `db:"long_stop_min" gorm:"Column:long_stop_min"`
	LongStopMax               int32   `db:"long_stop_max" gorm:"Column:long_stop_max"`
	WirelessOfflineThreshold  uint32  `db:"wireless_offline_threshold" gorm:"Column:wireless_offline_threshold"`
	WiredOfflineThreshold     uint32  `db:"wired_offline_threshold" gorm:"Column:wired_offline_threshold"`
	EmergencyOfflineThreshold uint32  `db:"emergency_offline_threshold" gorm:"Column:emergency_offline_threshold"`
	WiredHbi                  *string `db:"wired_hbi" gorm:"Column:wired_hbi"`
	WirelessHbi               *string `db:"wireless_hbi" gorm:"Column:wireless_hbi"`
	MileageThreshold          *string `db:"mileage_threshold" gorm:"Column:mileage_threshold"`
}

func (Corp) TableName() string {
	return "t_corp"
}

// CorpVins 集团用户车架号信息表
type CorpVins struct {
	ID                   int32  `db:"id" gorm:"Column:id"`
	Mobile               string `db:"mobile" gorm:"Column:mobile"`
	Vin                  string `db:"vin" gorm:"Column:vin"`
	AllowMutipleTerminal uint8  `db:"allow_mutiple_terminal" gorm:"Column:allow_mutiple_terminal"`
	Status               uint8  `db:"status" gorm:"Column:status"`
	UpdateTime           int32  `db:"update_time" gorm:"Column:update_time"`
	CreateTime           int32  `db:"create_time" gorm:"Column:create_time"`
}

func (CorpVins) TableName() string {
	return "t_corp_vins"
}

// CountDownTask 倒计时任务
type CountDownTask struct {
	ID      uint64 `db:"id" gorm:"Column:id"`
	TraceID string `db:"trace_id" gorm:"Column:trace_id"`
	EndAt   int32  `db:"end_at" gorm:"Column:end_at"`
	Action  uint8  `db:"action" gorm:"Column:action"`
	ReplyTo string `db:"reply_to" gorm:"Column:reply_to"`
	Key     string `db:"key" gorm:"Column:key"`
	Carrier string `db:"carrier" gorm:"Column:carrier"`
	Status  int8   `db:"status" gorm:"Column:status"`
}

func (CountDownTask) TableName() string {
	return "t_count_down_task"
}

type CriticalWarning struct {
	ID        int64  `db:"id" gorm:"Column:id"`
	Cnum      string `db:"cnum" gorm:"Column:cnum"`
	Vin       string `db:"vin" gorm:"Column:vin"`
	Cid       string `db:"cid" gorm:"Column:cid"`
	Sn        string `db:"sn" gorm:"Column:sn"`
	TypeName  string `db:"type_name" gorm:"Column:type_name"`
	Category  int32  `db:"category" gorm:"Column:category"`
	Eid       string `db:"eid" gorm:"Column:eid"`
	Timestamp int32  `db:"timestamp" gorm:"Column:timestamp"`
}

func (CriticalWarning) TableName() string {
	return "t_critical_warning"
}

// DailyMileage 按日统计的里程即移动停留时间
type DailyMileage struct {
	ID            uint64  `db:"id" gorm:"Column:id"`
	Tid           string  `db:"tid" gorm:"Column:tid"`
	Timestamp     int64   `db:"timestamp" gorm:"Column:timestamp"`
	IncMileage    float32 `db:"inc_mileage" gorm:"Column:inc_mileage"`
	DriveTimeLong int32   `db:"drive_time_long" gorm:"Column:drive_time_long"`
	StopTimeLong  int32   `db:"stop_time_long" gorm:"Column:stop_time_long"`
}

func (DailyMileage) TableName() string {
	return "t_daily_mileage"
}

// DeviceDailyCounter 设备计数表(按天)
type DeviceDailyCounter struct {
	ID           int32  `db:"id" gorm:"Column:id"`
	Sn           string `db:"sn" gorm:"Column:sn"`
	Iccid        string `db:"iccid" gorm:"Column:iccid"`
	DayTimestamp uint32 `db:"day_timestamp" gorm:"Column:day_timestamp"`
	Category     uint16 `db:"category" gorm:"Column:category"`
	Count        uint32 `db:"count" gorm:"Column:count"`
}

func (DeviceDailyCounter) TableName() string {
	return "t_device_daily_counter"
}

// DeviceMonthlyCounter 设备计数表(按月)
type DeviceMonthlyCounter struct {
	ID             int64  `db:"id" gorm:"Column:id"`
	Sn             string `db:"sn" gorm:"Column:sn"`
	Iccid          string `db:"iccid" gorm:"Column:iccid"`
	MonthTimestamp uint32 `db:"month_timestamp" gorm:"Column:month_timestamp"`
	Category       uint16 `db:"category" gorm:"Column:category"`
	Count          uint32 `db:"count" gorm:"Column:count"`
}

func (DeviceMonthlyCounter) TableName() string {
	return "t_device_monthly_counter"
}

type District struct {
	Name     *string `db:"name" gorm:"Column:name"`
	Adcode   string  `db:"adcode" gorm:"Column:adcode"`
	CityCode *string `db:"city_code" gorm:"Column:city_code"`
}

func (District) TableName() string {
	return "t_district"
}

// EmergencyAutoExitMask 24小时自动退出紧急模式黑名单
type EmergencyAutoExitMask struct {
	Sn  string `db:"sn" gorm:"Column:sn"`
	Cid string `db:"cid" gorm:"Column:cid"`
}

func (EmergencyAutoExitMask) TableName() string {
	return "t_emergency_auto_exit_mask"
}

// Event 事件信息
type Event struct {
	ID         int32   `db:"id" gorm:"Column:id"`
	Eid        string  `db:"eid" gorm:"Column:eid"`
	Cid        string  `db:"cid" gorm:"Column:cid"`
	Tid        string  `db:"tid" gorm:"Column:tid"`
	Timestamp  int64   `db:"timestamp" gorm:"Column:timestamp"`
	Lid        *int64  `db:"lid" gorm:"Column:lid"`
	Pbat       *string `db:"pbat" gorm:"Column:pbat"`
	Category   *int8   `db:"category" gorm:"Column:category"`
	Rid        int64   `db:"rid" gorm:"Column:rid"`
	Latitude   *int64  `db:"latitude" gorm:"Column:latitude"`
	Longitude  *int64  `db:"longitude" gorm:"Column:longitude"`
	Altitude   *int64  `db:"altitude" gorm:"Column:altitude"`
	Clatitude  *int64  `db:"clatitude" gorm:"Column:clatitude"`
	Clongitude *int64  `db:"clongitude" gorm:"Column:clongitude"`
	RxLev      int32   `db:"rxlev" gorm:"Column:rxlev"`
	Name       string  `db:"name" gorm:"Column:name"`
	Speed      *int32  `db:"speed" gorm:"Column:speed"`
	EscapeTime uint64  `db:"escape_time" gorm:"Column:escape_time"`
}

func (Event) TableName() string {
	return "t_event"
}

// Feedback 反馈表
type Feedback struct {
	ID        uint64 `db:"id" gorm:"Column:id"`
	Contact   string `db:"contact" gorm:"Column:contact"`
	Mobile    string `db:"mobile" gorm:"Column:mobile"`
	Email     string `db:"email" gorm:"Column:email"`
	Num       *int32 `db:"num" gorm:"Column:num"`
	Content   string `db:"content" gorm:"Column:content"`
	Timestamp int64  `db:"timestamp" gorm:"Column:timestamp"`
	Reply     string `db:"reply" gorm:"Column:reply"`
	ReplyTime int64  `db:"reply_time" gorm:"Column:reply_time"`
	Isreplied uint8  `db:"isreplied" gorm:"Column:isreplied"`
	Category  int32  `db:"category" gorm:"Column:category"`
}

func (Feedback) TableName() string {
	return "t_feedback"
}

type Fence struct {
	ID         uint32  `db:"id" gorm:"Column:id"`
	Name       *string `db:"name" gorm:"Column:name"`
	Shape      *int8   `db:"shape" gorm:"Column:shape"`
	ShapeInfo  string  `db:"shape_info" gorm:"Column:shape_info"`
	EnterAlarm int8    `db:"enter_alarm" gorm:"Column:enter_alarm"`
	OutAlarm   int8    `db:"out_alarm" gorm:"Column:out_alarm"`
	CreateTime *int64  `db:"create_time" gorm:"Column:create_time"`
	Cid        *string `db:"cid" gorm:"Column:cid"`
}

func (Fence) TableName() string {
	return "t_fence"
}

type FenceCar struct {
	FenceID  *int32 `db:"fence_id" gorm:"Column:fence_id"`
	Type     int8   `db:"type" gorm:"Column:type"`
	TargetID string `db:"target_id" gorm:"Column:target_id"`
}

func (FenceCar) TableName() string {
	return "t_fence_car"
}

type Geocoding struct {
	ID         int32  `db:"id" gorm:"Column:id"`
	Latitude   int32  `db:"latitude" gorm:"Column:latitude"`
	Longitude  int32  `db:"longitude" gorm:"Column:longitude"`
	Geohash    string `db:"geohash" gorm:"Column:geohash"`
	Address    string `db:"address" gorm:"Column:address"`
	UpdateTime int32  `db:"update_time" gorm:"Column:update_time"`
}

func (Geocoding) TableName() string {
	return "t_geocoding"
}

// GPSInfo 安装检测阶段的gps信息
type GPSInfo struct {
	ID              uint64  `db:"id" gorm:"Column:id"`
	Sn              string  `db:"sn" gorm:"Column:sn"`
	Iccid           string  `db:"iccid" gorm:"Column:iccid"`
	Category        int8    `db:"category" gorm:"Column:category"`
	Latitude        float32 `db:"latitude" gorm:"Column:latitude"`
	Longitude       float32 `db:"longitude" gorm:"Column:longitude"`
	Altitude        int64   `db:"altitude" gorm:"Column:altitude"`
	Speed           int32   `db:"speed" gorm:"Column:speed"`
	Degree          *int16  `db:"degree" gorm:"Column:degree"`
	LocateError     int32   `db:"locate_error" gorm:"Column:locate_error"`
	Snr             uint32  `db:"snr" gorm:"Column:snr"`
	Mcc             *int32  `db:"mcc" gorm:"Column:mcc"`
	Mnc             *int32  `db:"mnc" gorm:"Column:mnc"`
	Lac             *int32  `db:"lac" gorm:"Column:lac"`
	CellID          *int32  `db:"cell_id" gorm:"Column:cell_id"`
	Timestamp       int64   `db:"timestamp" gorm:"Column:timestamp"`
	FirmwareVersion *string `db:"firmware_version" gorm:"Column:firmware_version"`
	RxLev           *int32  `db:"rxlev" gorm:"Column:rxlev"`
}

func (GPSInfo) TableName() string {
	return "t_gps_info"
}

// Group 集团终端分组表
type Group struct {
	ID         int32  `db:"id" gorm:"Column:id"`
	Gid        string `db:"gid" gorm:"Column:gid"`
	Cid        string `db:"cid" gorm:"Column:cid"`
	Pgid       string `db:"pgid" gorm:"Column:pgid"`
	Name       string `db:"name" gorm:"Column:name"`
	Type       int8   `db:"type" gorm:"Column:type"`
	CreateTime int64  `db:"create_time" gorm:"Column:create_time"`
	Level      int32  `db:"level" gorm:"Column:level"`
}

func (Group) TableName() string {
	return "t_group"
}

// GroupCar 集团车辆表
type GroupCar struct {
	GroupID string `db:"group_id" gorm:"Column:group_id"`
	CarID   string `db:"car_id" gorm:"Column:car_id"`
}

func (GroupCar) TableName() string {
	return "t_group_car"
}

// GroupOperator 组与操作员对应关系表
type GroupOperator struct {
	GroupID string `db:"group_id" gorm:"Column:group_id"`
	OperID  string `db:"oper_id" gorm:"Column:oper_id"`
}

func (GroupOperator) TableName() string {
	return "t_group_operator"
}

// GroupPath 组父子线索表
type GroupPath struct {
	ID       int64  `db:"id" gorm:"Column:id"`
	Ancestor string `db:"ancestor" gorm:"Column:ancestor"`
	Distance int8   `db:"distance" gorm:"Column:distance"`
	Gid      string `db:"gid" gorm:"Column:gid"`
}

func (GroupPath) TableName() string {
	return "t_group_path"
}

type HelpDetail struct {
	ID          string  `db:"id" gorm:"Column:id"`
	Name        string  `db:"name" gorm:"Column:name"`
	Sequence    uint8   `db:"sequence" gorm:"Column:sequence"`
	Status      uint8   `db:"status" gorm:"Column:status"`
	Description *string `db:"description" gorm:"Column:description"`
	FileName    string  `db:"file_name" gorm:"Column:file_name"`
	HelpMenuID  string  `db:"help_menu_id" gorm:"Column:help_menu_id"`
	CreateTime  uint64  `db:"create_time" gorm:"Column:create_time"`
}

func (HelpDetail) TableName() string {
	return "t_help_detail"
}

type HelpMenu struct {
	ID          string  `db:"id" gorm:"Column:id"`
	Name        string  `db:"name" gorm:"Column:name"`
	Sequence    uint8   `db:"sequence" gorm:"Column:sequence"`
	Status      uint8   `db:"status" gorm:"Column:status"`
	Description *string `db:"description" gorm:"Column:description"`
	Type        uint64  `db:"type" gorm:"Column:type"`
	CreateTime  uint64  `db:"create_time" gorm:"Column:create_time"`
}

func (HelpMenu) TableName() string {
	return "t_help_menu"
}

type HybridLocatorWiFiLocation struct {
	ID           int64    `db:"id" gorm:"Column:id"`
	GenerateTime int64    `db:"generate_time" gorm:"Column:generate_time"`
	Mac          string   `db:"mac" gorm:"Column:mac"`
	Lng          *float64 `db:"lng" gorm:"Column:lng"`
	Lat          *float64 `db:"lat" gorm:"Column:lat"`
	Supplier     *string  `db:"supplier" gorm:"Column:supplier"`
	Tag          *string  `db:"tag" gorm:"Column:tag"`
	Address      *string  `db:"address" gorm:"Column:address"`
}

func (HybridLocatorWiFiLocation) TableName() string {
	return "t_hybrid_locator_wifi_location"
}

// InstallCheckMode 终端设备安装检查模式表
type InstallCheckMode struct {
	ID                uint32  `db:"id" gorm:"Column:id"`
	Iid               string  `db:"iid" gorm:"Column:iid"`
	Cid               string  `db:"cid" gorm:"Column:cid"`
	Sn                string  `db:"sn" gorm:"Column:sn"`
	ProtocolName      string  `db:"protocol_name" gorm:"Column:protocol_name"`
	IsPowerUp         uint8   `db:"is_power_up" gorm:"Column:is_power_up"`
	IsOnline          uint8   `db:"is_online" gorm:"Column:is_online"`
	IsLocate          uint8   `db:"is_locate" gorm:"Column:is_locate"`
	IsPowerConnect    int8    `db:"is_power_connect" gorm:"Column:is_power_connect"`
	IsVoltageNormal   int8    `db:"is_voltage_normal" gorm:"Column:is_voltage_normal"`
	CanEnterEmergency int8    `db:"can_enter_emergency" gorm:"Column:can_enter_emergency"`
	IsLightSense      int8    `db:"is_light_sense" gorm:"Column:is_light_sense"`
	Battery           int8    `db:"battery" gorm:"Column:battery"`
	EnterTime         uint32  `db:"enter_time" gorm:"Column:enter_time"`
	CheckResult       uint8   `db:"check_result" gorm:"Column:check_result"`
	CheckCount        uint8   `db:"check_count" gorm:"Column:check_count"`
	FuelCut           *string `db:"fuel_cut" gorm:"Column:fuel_cut"`
}

func (InstallCheckMode) TableName() string {
	return "t_install_check_mode"
}

// Installer 安装者表
type Installer struct {
	ID           int32   `db:"id" gorm:"Column:id"`
	Iid          string  `db:"iid" gorm:"Column:iid"`
	Cid          string  `db:"cid" gorm:"Column:cid"`
	Loginname    string  `db:"loginname" gorm:"Column:loginname"`
	Password     string  `db:"password" gorm:"Column:password"`
	Responsible  *string `db:"responsible" gorm:"Column:responsible"`
	Gid          string  `db:"gid" gorm:"Column:gid"`
	PwExpireTime *int64  `db:"pw_expire_time" gorm:"Column:pw_expire_time"`
	WarehouseID  *string `db:"warehouse_id" gorm:"Column:warehouse_id"`
}

func (Installer) TableName() string {
	return "t_installer"
}

// JobState 任务执行情况
type JobState struct {
	ID        int32  `db:"id" gorm:"Column:id"`
	JobID     string `db:"job_id" gorm:"Column:job_id"`
	State     string `db:"state" gorm:"Column:state"`
	UpdatedAt uint32 `db:"updated_at" gorm:"Column:updated_at"`
}

func (JobState) TableName() string {
	return "t_job_state"
}

// Location 位置信息。
type Location struct {
	ID          uint64  `db:"id" gorm:"Column:id"`
	CarID       *string `db:"car_id" gorm:"Column:car_id"`
	Tid         string  `db:"tid" gorm:"Column:tid"`
	Category    int8    `db:"category" gorm:"Column:category"`
	Latitude    int64   `db:"latitude" gorm:"Column:latitude"`
	Longitude   int64   `db:"longitude" gorm:"Column:longitude"`
	Clatitude   int64   `db:"clatitude" gorm:"Column:clatitude"`
	Clongitude  int64   `db:"clongitude" gorm:"Column:clongitude"`
	Address     string  `db:"address" gorm:"Column:address"`
	Altitude    int64   `db:"altitude" gorm:"Column:altitude"`
	Speed       int32   `db:"speed" gorm:"Column:speed"`
	Degree      float32 `db:"degree" gorm:"Column:degree"`
	LocateError int32   `db:"locate_error" gorm:"Column:locate_error"`
	Snr         uint32  `db:"snr" gorm:"Column:snr"`
	Mcc         *string `db:"mcc" gorm:"Column:mcc"`
	Mnc         *string `db:"mnc" gorm:"Column:mnc"`
	Lac         *string `db:"lac" gorm:"Column:lac"`
	CellID      *string `db:"cell_id" gorm:"Column:cell_id"`
	Timestamp   int64   `db:"timestamp" gorm:"Column:timestamp"`
	TType       *string `db:"t_type" gorm:"Column:t_type"`
	LocateType  int32   `db:"locate_type" gorm:"Column:locate_type"`
	RxLev       *int32  `db:"rxlev" gorm:"Column:rxlev"`
}

func (Location) TableName() string {
	return "t_location"
}

type LocationDelaySummary struct {
	CreateTime uint64 `db:"create_time" gorm:"Column:create_time"`
	Sum        uint64 `db:"sum" gorm:"Column:sum"`
	Avg        uint32 `db:"avg" gorm:"Column:avg"`
	Mid        uint32 `db:"mid" gorm:"Column:mid"`
	Min        uint32 `db:"min" gorm:"Column:min"`
	Max        uint32 `db:"max" gorm:"Column:max"`
	Serious    uint32 `db:"serious" gorm:"Column:serious"`
}

func (LocationDelaySummary) TableName() string {
	return "t_location_delay_summary"
}

// LocationExtend 扩展位置表
type LocationExtend struct {
	Lid         uint64 `db:"lid" gorm:"Column:lid"`
	ReceiveTime int64  `db:"receive_time" gorm:"Column:receive_time"`
	Sn          string `db:"sn" gorm:"Column:sn"`
	Delay       uint16 `db:"delay" gorm:"Column:delay"`
	Geohash     string `db:"geohash" gorm:"Column:geohash"`
}

func (LocationExtend) TableName() string {
	return "t_location_extend"
}

// LocationWiFi wifi定位报文信息
type LocationWiFi struct {
	ID           uint64 `db:"id" gorm:"Column:id"`
	Lid          int64  `db:"lid" gorm:"Column:lid"`
	Mode         uint8  `db:"mode" gorm:"Column:mode"`
	Ssid         string `db:"ssid" gorm:"Column:ssid"`
	Mac          string `db:"mac" gorm:"Column:mac"`
	Basestations string `db:"basestations" gorm:"Column:basestations"`
}

func (LocationWiFi) TableName() string {
	return "t_location_wifi"
}

// LoginLog 用户登陆日志
type LoginLog struct {
	ID        uint64 `db:"id" gorm:"Column:id"`
	Oid       string `db:"oid" gorm:"Column:oid"`
	Method    uint8  `db:"method" gorm:"Column:method"`
	Timestamp uint64 `db:"timestamp" gorm:"Column:timestamp"`
}

func (LoginLog) TableName() string {
	return "t_login_log"
}

type LogisticsLog struct {
	ID            string  `db:"id" gorm:"Column:id"`
	Sn            string  `db:"sn" gorm:"Column:sn"`
	CorpID        string  `db:"corp_id" gorm:"Column:corp_id"`
	WarehouseID   string  `db:"warehouse_id" gorm:"Column:warehouse_id"`
	Operation     int8    `db:"operation" gorm:"Column:operation"`
	OperationTime int64   `db:"operation_time" gorm:"Column:operation_time"`
	Remark        *string `db:"remark" gorm:"Column:remark"`
}

func (LogisticsLog) TableName() string {
	return "t_logistics_log"
}

// Lua 终端升级脚本表
type Lua struct {
	ID          int64   `db:"id" gorm:"Column:id"`
	Version     string  `db:"version" gorm:"Column:version"`
	VersionName *string `db:"version_name" gorm:"Column:version_name"`
	Description *string `db:"description" gorm:"Column:description"`
	Timestamp   int64   `db:"timestamp" gorm:"Column:timestamp"`
	FileSize    *int64  `db:"file_size" gorm:"Column:file_size"`
	FileName    string  `db:"file_name" gorm:"Column:file_name"`
	Uploader    *string `db:"uploader" gorm:"Column:uploader"`
	Status      int8    `db:"status" gorm:"Column:status"`
	SnRegexp    *string `db:"sn_regexp" gorm:"Column:sn_regexp"`
}

func (Lua) TableName() string {
	return "t_lua"
}

// Maintainer 维护人员信息。
type Maintainer struct {
	ID     uint32 `db:"id" gorm:"Column:id"`
	Mid    string `db:"mid" gorm:"Column:mid"`
	Name   string `db:"name" gorm:"Column:name"`
	Mobile string `db:"mobile" gorm:"Column:mobile"`
	Email  string `db:"email" gorm:"Column:email"`
	Remark string `db:"remark" gorm:"Column:remark"`
	Valid  uint8  `db:"valid" gorm:"Column:valid"`
}

func (Maintainer) TableName() string {
	return "t_maintainer"
}

// MapApiKey 地图开发者key存储库
type MapApiKey struct {
	MapType string  `db:"map_type" gorm:"Column:map_type"`
	ApiKey  string  `db:"api_key" gorm:"Column:api_key"`
	Remain  int32   `db:"remain" gorm:"Column:remain"`
	Remark  *string `db:"remark" gorm:"Column:remark"`
}

func (MapApiKey) TableName() string {
	return "t_map_api_key"
}

// Mileage 终端里程统计表，仅统计符合[第三方推送]规则的设备里程
type Mileage struct {
	ID          uint64   `db:"id" gorm:"Column:id"`
	Tid         *string  `db:"tid" gorm:"Column:tid"`
	Sn          *string  `db:"sn" gorm:"Column:sn"`
	Mileage     *float32 `db:"mileage" gorm:"Column:mileage"`
	IncMileage  *float32 `db:"inc_mileage" gorm:"Column:inc_mileage"`
	BaseMileage *float32 `db:"base_mileage" gorm:"Column:base_mileage"`
	Timestamp   *int64   `db:"timestamp" gorm:"Column:timestamp"`
}

func (Mileage) TableName() string {
	return "t_mileage"
}

type NybAddcarFootPrints struct {
	Ip        *string `db:"ip" gorm:"Column:ip"`
	Timestamp *int64  `db:"timestamp" gorm:"Column:timestamp"`
	Request   *string `db:"request" gorm:"Column:request"`
	Message   *string `db:"message" gorm:"Column:message"`
}

func (NybAddcarFootPrints) TableName() string {
	return "t_nyb_addcar_foot_prints"
}

// Operator 群组操作员列表
type Operator struct {
	ID         int32   `db:"id" gorm:"Column:id"`
	Oid        string  `db:"oid" gorm:"Column:oid"`
	Cid        *string `db:"cid" gorm:"Column:cid"`
	Mobile     string  `db:"mobile" gorm:"Column:mobile"`
	Password   string  `db:"password" gorm:"Column:password"`
	Name       *string `db:"name" gorm:"Column:name"`
	Email      *string `db:"email" gorm:"Column:email"`
	Address    *string `db:"address" gorm:"Column:address"`
	CreateTime *int64  `db:"create_time" gorm:"Column:create_time"`
	Status     int32   `db:"status" gorm:"Column:status"`
	Type       int32   `db:"type" gorm:"Column:type"`
	Source     int32   `db:"source" gorm:"Column:source"`
	Privilege  string  `db:"privilege" gorm:"Column:privilege"`
	AuthType   int32   `db:"auth_type" gorm:"Column:auth_type"`
	OrgID      *int32  `db:"org_id" gorm:"Column:org_id"`
	DeptID     *int32  `db:"dept_id" gorm:"Column:dept_id"`
}

func (Operator) TableName() string {
	return "t_operator"
}

// Package 装箱表
type Package struct {
	ID              string `db:"id" gorm:"Column:id"`
	PackageNo       string `db:"package_no" gorm:"Column:package_no"`
	BoxingTimestamp uint64 `db:"boxing_timestamp" gorm:"Column:boxing_timestamp"`
}

func (Package) TableName() string {
	return "t_package"
}

type PackageHasSn struct {
	PackageID string `db:"package_id" gorm:"Column:package_id"`
	Serial    uint8  `db:"serial" gorm:"Column:serial"`
	Sn        string `db:"sn" gorm:"Column:sn"`
}

func (PackageHasSn) TableName() string {
	return "t_package_has_sn"
}

// PrimeStatistics 初步(简单)统计表
type PrimeStatistics struct {
	ID         int32   `db:"id" gorm:"Column:id"`
	Oid        *string `db:"oid" gorm:"Column:oid"`
	Cid        *string `db:"cid" gorm:"Column:cid"`
	Content    *string `db:"content" gorm:"Column:content"`
	CreateTime *int32  `db:"create_time" gorm:"Column:create_time"`
	Stype      *string `db:"stype" gorm:"Column:stype"`
	Version    *int32  `db:"version" gorm:"Column:version"`
}

func (PrimeStatistics) TableName() string {
	return "t_prime_statistics"
}

// Provider 第三方供货商表
type Provider struct {
	ID           uint32  `db:"id" gorm:"Column:id"`
	AppID        *string `db:"app_id" gorm:"Column:app_id"`
	AppKey       *string `db:"app_key" gorm:"Column:app_key"`
	ProviderNo   int32   `db:"provider_no" gorm:"Column:provider_no"`
	ProviderName *string `db:"provider_name" gorm:"Column:provider_name"`
	Protocol     *int8   `db:"protocol" gorm:"Column:protocol"`
	CreateTime   int64   `db:"create_time" gorm:"Column:create_time"`
}

func (Provider) TableName() string {
	return "t_provider"
}

type PushableCache struct {
	ID      int32   `db:"id" gorm:"Column:id"`
	Mobile  *string `db:"mobile" gorm:"Column:mobile"`
	ObjID   *string `db:"obj_id" gorm:"Column:obj_id"`
	ObjType *int8   `db:"obj_type" gorm:"Column:obj_type"`
	Content string  `db:"content" gorm:"Column:content"`
	OptTime *int32  `db:"opt_time" gorm:"Column:opt_time"`
}

func (PushableCache) TableName() string {
	return "t_pushable_cache"
}

type Recycle struct {
	ID            string  `db:"id" gorm:"Column:id"`
	WaybillNumber *string `db:"waybill_number" gorm:"Column:waybill_number"`
	Timestamp     *uint64 `db:"timestamp" gorm:"Column:timestamp"`
	Remark        *string `db:"remark" gorm:"Column:remark"`
}

func (Recycle) TableName() string {
	return "t_recycle"
}

type Redelivery struct {
	ID            string  `db:"id" gorm:"Column:id"`
	WaybillNumber *string `db:"waybill_number" gorm:"Column:waybill_number"`
	Timestamp     *uint64 `db:"timestamp" gorm:"Column:timestamp"`
	Remark        *string `db:"remark" gorm:"Column:remark"`
}

func (Redelivery) TableName() string {
	return "t_redelivery"
}

type RedoThirdpartyData struct {
	ID       int32         `db:"id" gorm:"Column:id"`
	Cid      string        `db:"cid" gorm:"Column:cid"`
	Content  string        `db:"content" gorm:"Column:content"`
	Time     sql.NullInt64 `db:"time" gorm:"Column:time"`
	ThreadID *int32        `db:"thread_id" gorm:"Column:thread_id"`
	Remark   *string       `db:"remark" gorm:"Column:remark"`
}

func (RedoThirdpartyData) TableName() string {
	return "t_redo_thirdparty_data"
}

// Region 电子围栏
type Region struct {
	ID        int32   `db:"id" gorm:"Column:id"`
	Name      *string `db:"name" gorm:"Column:name"`
	Longitude *int64  `db:"longitude" gorm:"Column:longitude"`
	Latitude  *int64  `db:"latitude" gorm:"Column:latitude"`
	Radius    *int32  `db:"radius" gorm:"Column:radius"`
	Points    string  `db:"points" gorm:"Column:points"`
	Shape     int8    `db:"shape" gorm:"Column:shape"`
	Cid       string  `db:"cid" gorm:"Column:cid"`
	Uid       string  `db:"uid" gorm:"Column:uid"`
}

func (Region) TableName() string {
	return "t_region"
}

// RegionTerminal 围栏与终端绑定关系表
type RegionTerminal struct {
	Rid int64  `db:"rid" gorm:"Column:rid"`
	Tid string `db:"tid" gorm:"Column:tid"`
}

func (RegionTerminal) TableName() string {
	return "t_region_terminal"
}

type Rework struct {
	ID           string  `db:"id" gorm:"Column:id"`
	Sn           string  `db:"sn" gorm:"Column:sn"`
	Status       int8    `db:"status" gorm:"Column:status"`
	RecycleCause *string `db:"recycle_cause" gorm:"Column:recycle_cause"`
	FaultCause   *string `db:"fault_cause" gorm:"Column:fault_cause"`
	RecycleID    *string `db:"recycle_id" gorm:"Column:recycle_id"`
	RedeliveryID *string `db:"redelivery_id" gorm:"Column:redelivery_id"`
	CreateTime   *int64  `db:"create_time" gorm:"Column:create_time"`
}

func (Rework) TableName() string {
	return "t_rework"
}

// SConfig 系统运行配置表
type SConfig struct {
	ID      uint32  `db:"id" gorm:"Column:id"`
	Node    string  `db:"node" gorm:"Column:node"`
	Key     string  `db:"key" gorm:"Column:key"`
	Value   *string `db:"value" gorm:"Column:value"`
	EnvType string  `db:"env_type" gorm:"Column:env_type"`
	KeyNote *string `db:"key_note" gorm:"Column:key_note"`
}

func (SConfig) TableName() string {
	return "t_s_config"
}

type SConfigDevelop struct {
	ID      uint32  `db:"id" gorm:"Column:id"`
	Node    string  `db:"node" gorm:"Column:node"`
	Key     string  `db:"key" gorm:"Column:key"`
	Value   *string `db:"value" gorm:"Column:value"`
	KeyType string  `db:"key_type" gorm:"Column:key_type"`
	KeyNote *string `db:"key_note" gorm:"Column:key_note"`
}

func (SConfigDevelop) TableName() string {
	return "t_s_config_develop"
}

type SConfigToOnline struct {
	ID      uint32  `db:"id" gorm:"Column:id"`
	Node    string  `db:"node" gorm:"Column:node"`
	Key     string  `db:"key" gorm:"Column:key"`
	Value   *string `db:"value" gorm:"Column:value"`
	KeyType string  `db:"key_type" gorm:"Column:key_type"`
	KeyNote *string `db:"key_note" gorm:"Column:key_note"`
}

func (SConfigToOnline) TableName() string {
	return "t_s_config_to_online"
}

// Salesman 销售人员表
type Salesman struct {
	ID     uint32  `db:"id" gorm:"Column:id"`
	Name   string  `db:"name" gorm:"Column:name"`
	Email  string  `db:"email" gorm:"Column:email"`
	Mobile *string `db:"mobile" gorm:"Column:mobile"`
	Remark *string `db:"remark" gorm:"Column:remark"`
}

func (Salesman) TableName() string {
	return "t_salesman"
}

// Script 终端脚本版本控制
type Script struct {
	ID        uint64  `db:"id" gorm:"Column:id"`
	Version   string  `db:"version" gorm:"Column:version"`
	Filename  string  `db:"filename" gorm:"Column:filename"`
	Timestamp uint64  `db:"timestamp" gorm:"Column:timestamp"`
	Author    *string `db:"author" gorm:"Column:author"`
	Islocked  int8    `db:"islocked" gorm:"Column:islocked"`
}

func (Script) TableName() string {
	return "t_script"
}

// ScriptDownload 版本信息下载表
type ScriptDownload struct {
	ID          uint64  `db:"id" gorm:"Column:id"`
	Tid         string  `db:"tid" gorm:"Column:tid"`
	Versionname *string `db:"versionname" gorm:"Column:versionname"`
	Timestamp   uint64  `db:"timestamp" gorm:"Column:timestamp"`
}

func (ScriptDownload) TableName() string {
	return "t_script_download"
}

type Sn struct {
	ID                            int64   `db:"id" gorm:"Column:id"`
	Sn                            string  `db:"sn" gorm:"Column:sn"`
	ActivateCode                  string  `db:"activate_code" gorm:"Column:activate_code"`
	PreInstall                    int8    `db:"pre_install" gorm:"Column:pre_install"`
	Type                          string  `db:"type" gorm:"Column:type"`
	CreateTime                    uint64  `db:"create_time" gorm:"Column:create_time"`
	Iccid                         *string `db:"iccid" gorm:"Column:iccid"`
	Timestamp                     *int64  `db:"timestamp" gorm:"Column:timestamp"`
	WarehouseID                   *string `db:"warehouse_id" gorm:"Column:warehouse_id"`
	IntoWarehouseTime             *int64  `db:"into_warehouse_time" gorm:"Column:into_warehouse_time"`
	Agriculture2dCode             *string `db:"agriculture_2d_code" gorm:"Column:agriculture_2d_code"`
	Agriculture2dCodeAllocateTime *int64  `db:"agriculture_2d_code_allocate_time" gorm:"Column:agriculture_2d_code_allocate_time"`
	HasProvider                   int8    `db:"has_provider" gorm:"Column:has_provider"`
	ProviderName                  *string `db:"provider_name" gorm:"Column:provider_name"`
	GuessCnum                     *string `db:"guess_cnum" gorm:"Column:guess_cnum"`
	ProviderNo                    int32   `db:"provider_no" gorm:"Column:provider_no"`
	VehicleType                   *int32  `db:"vehicle_type" gorm:"Column:vehicle_type"`
	AgricultureUpdateTime         *int64  `db:"agriculture_update_time" gorm:"Column:agriculture_update_time"`
}

func (Sn) TableName() string {
	return "t_sn"
}

type SnArea struct {
	Sn     string `db:"sn" gorm:"Column:sn"`
	AreaNo string `db:"area_no" gorm:"Column:area_no"`
}

func (SnArea) TableName() string {
	return "t_sn_area"
}

type SubuserObj struct {
	ID          int32   `db:"id" gorm:"Column:id"`
	Oid         *string `db:"oid" gorm:"Column:oid"`
	CreateTime  int32   `db:"create_time" gorm:"Column:create_time"`
	CarID       *string `db:"car_id" gorm:"Column:car_id"`
	Gid         *string `db:"gid" gorm:"Column:gid"`
	OperateType int8    `db:"operate_type" gorm:"Column:operate_type"`
	Active      int8    `db:"active" gorm:"Column:active"`
	ParentID    *string `db:"parent_id" gorm:"Column:parent_id"`
	CacheStatus *int8   `db:"cache_status" gorm:"Column:cache_status"`
}

func (SubuserObj) TableName() string {
	return "t_subuser_obj"
}

// SysSettings 系统动态配置表
type SysSettings struct {
	Key   string `db:"key" gorm:"Column:key"`
	Value string `db:"value" gorm:"Column:value"`
}

func (SysSettings) TableName() string {
	return "t_sys_settings"
}

// TermInfoHistory 设备上报的部分信息，历史记录表。
type TermInfoHistory struct {
	Tid        string `db:"tid" gorm:"Column:tid"`
	Timestamp  int64  `db:"timestamp" gorm:"Column:timestamp"`
	RfID       string `db:"rf_id" gorm:"Column:rf_id"`
	RfTempRssi string `db:"rf_temp_rssi" gorm:"Column:rf_temp_rssi"`
	RfLost     *int8  `db:"rf_lost" gorm:"Column:rf_lost"`
	RfSwitch   *int8  `db:"rf_switch" gorm:"Column:rf_switch"`
	RfMode     *int8  `db:"rf_mode" gorm:"Column:rf_mode"`
}

func (TermInfoHistory) TableName() string {
	return "t_term_info_history"
}

// TerminalControlHistory 终端控制信息下发记录表
type TerminalControlHistory struct {
	ID         int64  `db:"id" gorm:"Column:id"`
	Sn         string `db:"sn" gorm:"Column:sn"`
	Type       string `db:"type" gorm:"Column:type"`
	Action     string `db:"action" gorm:"Column:action"`
	IdentifyID uint16 `db:"identify_id" gorm:"Column:identify_id"`
	Status     int8   `db:"status" gorm:"Column:status"`
	CreateTime uint64 `db:"create_time" gorm:"Column:create_time"`
	UpdateTime uint64 `db:"update_time" gorm:"Column:update_time"`
	Remark     string `db:"remark" gorm:"Column:remark"`
}

func (TerminalControlHistory) TableName() string {
	return "t_terminal_control_history"
}

type TerminalEventHistory struct {
	ID              int64   `db:"id" gorm:"Column:id"`
	Sn              string  `db:"sn" gorm:"Column:sn"`
	CreateAt        int64   `db:"create_at" gorm:"Column:create_at"`
	EventType       int16   `db:"event_type" gorm:"Column:event_type"`
	Name            string  `db:"name" gorm:"Column:name"`
	Remark          string  `db:"remark" gorm:"Column:remark"`
	LocateErrorInfo *string `db:"locate_error_info" gorm:"Column:locate_error_info"`
}

func (TerminalEventHistory) TableName() string {
	return "t_terminal_event_history"
}

// TerminalExtend 终端扩展表，用于记录某些终端额外增加的客户关心的字段，可能车管系统本身并不太关心。
type TerminalExtend struct {
	Tid                         *string `db:"tid" gorm:"Column:tid"`
	ProducerID                  *string `db:"producer_id" gorm:"Column:producer_id"`
	TerminalVersion             *string `db:"terminal_version" gorm:"Column:terminal_version"`
	TerminalID                  *string `db:"terminal_id" gorm:"Column:terminal_id"`
	VehicleIdentificationNumber *string `db:"vehicle_identification_number" gorm:"Column:vehicle_identification_number"`
}

func (TerminalExtend) TableName() string {
	return "t_terminal_extend"
}

// TerminalInfo 终端表
type TerminalInfo struct {
	ID                        int64   `db:"id" gorm:"Column:id"`
	Tid                       string  `db:"tid" gorm:"Column:tid"`
	Cid                       string  `db:"cid" gorm:"Column:cid"`
	Iccid                     *string `db:"iccid" gorm:"Column:iccid"`
	Sn                        string  `db:"sn" gorm:"Column:sn"`
	Alias                     string  `db:"alias" gorm:"Column:alias"`
	LoginTime                 uint64  `db:"login_time" gorm:"Column:login_time"`
	OfflineTime               uint64  `db:"offline_time" gorm:"Column:offline_time"`
	Domain                    string  `db:"domain" gorm:"Column:domain"`
	Vibchk                    string  `db:"vibchk" gorm:"Column:vibchk"`
	Vibl                      int8    `db:"vibl" gorm:"Column:vibl"`
	ServiceStatus             uint8   `db:"service_status" gorm:"Column:service_status"`
	Begintime                 uint64  `db:"begintime" gorm:"Column:begintime"`
	Endtime                   uint64  `db:"endtime" gorm:"Column:endtime"`
	AlertInt                  int32   `db:"alert_int" gorm:"Column:alert_int"`
	UseScene                  int32   `db:"use_scene" gorm:"Column:use_scene"`
	GSM                       uint32  `db:"gsm" gorm:"Column:gsm"`
	GPS                       uint32  `db:"gps" gorm:"Column:gps"`
	Pbat                      *uint32 `db:"pbat" gorm:"Column:pbat"`
	Login                     int32   `db:"login" gorm:"Column:login"`
	DeviceMode                int8    `db:"device_mode" gorm:"Column:device_mode"`
	UpdateTime                int64   `db:"update_time" gorm:"Column:update_time"`
	Exstatus                  int32   `db:"exstatus" gorm:"Column:exstatus"`
	StopInterval              uint32  `db:"stop_interval" gorm:"Column:stop_interval"`
	ActivateCode              string  `db:"activate_code" gorm:"Column:activate_code"`
	Hbi                       string  `db:"hbi" gorm:"Column:hbi"`
	BatteryVoltage            int32   `db:"battery_voltage" gorm:"Column:battery_voltage"`
	Unbind                    int8    `db:"unbind" gorm:"Column:unbind"`
	TerminalMode              string  `db:"terminal_mode" gorm:"Column:terminal_mode"`
	ParamInterval             string  `db:"param_interval" gorm:"Column:param_interval"`
	GPSStatus                 string  `db:"gps_status" gorm:"Column:gps_status"`
	Type                      int32   `db:"type" gorm:"Column:type"`
	ChargeStatus              int32   `db:"charge_status" gorm:"Column:charge_status"`
	Status                    int32   `db:"status" gorm:"Column:status"`
	ActivateTime              int64   `db:"activate_time" gorm:"Column:activate_time"`
	Temp                      *int32  `db:"temp" gorm:"Column:temp"`
	Acc                       *int32  `db:"acc" gorm:"Column:acc"`
	AccSwitch                 int8    `db:"acc_switch" gorm:"Column:acc_switch"`
	Sl                        *int32  `db:"sl" gorm:"Column:sl"`
	Rl                        *int32  `db:"rl" gorm:"Column:rl"`
	Ce                        *int32  `db:"ce" gorm:"Column:ce"`
	ProductVersion            string  `db:"product_version" gorm:"Column:product_version"`
	ProtocolVersion           *string `db:"protocol_version" gorm:"Column:protocol_version"`
	FirmwareVersion           *string `db:"firmware_version" gorm:"Column:firmware_version"`
	FirmwareC                 *string `db:"firmware_c" gorm:"Column:firmware_c"`
	FirmwareL                 *string `db:"firmware_l" gorm:"Column:firmware_l"`
	LoginReason               *int32  `db:"login_reason" gorm:"Column:login_reason"`
	PreInstall                int8    `db:"pre_install" gorm:"Column:pre_install"`
	Delay                     int32   `db:"delay" gorm:"Column:delay"`
	Rdest                     int32   `db:"rdest" gorm:"Column:rdest"`
	Vsens                     string  `db:"vsens" gorm:"Column:vsens"`
	Acd                       string  `db:"acd" gorm:"Column:acd"`
	WorkType                  int32   `db:"work_type" gorm:"Column:work_type"`
	U9Time                    int64   `db:"u9_time" gorm:"Column:u9_time"`
	Mdest                     int32   `db:"mdest" gorm:"Column:mdest"`
	Wakeup                    int32   `db:"wakeup" gorm:"Column:wakeup"`
	Tint                      int32   `db:"tint" gorm:"Column:tint"`
	Lvol                      int32   `db:"lvol" gorm:"Column:lvol"`
	Imsi                      *string `db:"imsi" gorm:"Column:imsi"`
	Alarm                     int64   `db:"alarm" gorm:"Column:alarm"`
	HbTime                    int64   `db:"hb_time" gorm:"Column:hb_time"`
	LastPktTime               int64   `db:"last_pkt_time" gorm:"Column:last_pkt_time"`
	InstallTime               int32   `db:"install_time" gorm:"Column:install_time"`
	GPSLid                    *int64  `db:"gps_lid" gorm:"Column:gps_lid"`
	CellLid                   *int64  `db:"cell_lid" gorm:"Column:cell_lid"`
	Power                     int8    `db:"power" gorm:"Column:power"`
	Remark                    *string `db:"remark" gorm:"Column:remark"`
	RemarkType                uint8   `db:"remark_type" gorm:"Column:remark_type"`
	InstallUrl                *string `db:"install_url" gorm:"Column:install_url"`
	WakeupReason              *int32  `db:"wakeup_reason" gorm:"Column:wakeup_reason"`
	Light                     int8    `db:"light" gorm:"Column:light"`
	WiFi                      int8    `db:"wifi" gorm:"Column:wifi"`
	WiFiApTime                int16   `db:"wifi_ap_time" gorm:"Column:wifi_ap_time"`
	WiFiStatus                int8    `db:"wifi_status" gorm:"Column:wifi_status"`
	LastPositionType          *int8   `db:"last_position_type" gorm:"Column:last_position_type"`
	LightStatus               int8    `db:"light_status" gorm:"Column:light_status"`
	WiFiLid                   *int64  `db:"wifi_lid" gorm:"Column:wifi_lid"`
	EmergencyReason           *int8   `db:"emergency_reason" gorm:"Column:emergency_reason"`
	EmergencyAt               uint32  `db:"emergency_at" gorm:"Column:emergency_at"`
	OfflineThreshold          uint32  `db:"offline_threshold" gorm:"Column:offline_threshold"`
	EmergencyOfflineThreshold uint32  `db:"emergency_offline_threshold" gorm:"Column:emergency_offline_threshold"`
	Reboot                    int8    `db:"reboot" gorm:"Column:reboot"`
	WiredFuelExpStatus        *int8   `db:"wired_fuel_exp_status" gorm:"Column:wired_fuel_exp_status"`
	WiredFuelExeStatus        *int8   `db:"wired_fuel_exe_status" gorm:"Column:wired_fuel_exe_status"`
	WiredFuelStatus           *int8   `db:"wired_fuel_status" gorm:"Column:wired_fuel_status"`
	DormantFuelExpStatus      *int8   `db:"dormant_fuel_exp_status" gorm:"Column:dormant_fuel_exp_status"`
	DormantFuelExeStatus      *int8   `db:"dormant_fuel_exe_status" gorm:"Column:dormant_fuel_exe_status"`
	DormantFuelStatus         *int8   `db:"dormant_fuel_status" gorm:"Column:dormant_fuel_status"`
	FuelCutLock               *int8   `db:"fuel_cut_lock" gorm:"Column:fuel_cut_lock"`
	CarLed                    int8    `db:"car_led" gorm:"Column:car_led"`
	CarLedDoStatus            int8    `db:"car_led_do_status" gorm:"Column:car_led_do_status"`
	CarSpeaker                int8    `db:"car_speaker" gorm:"Column:car_speaker"`
	CarSpeakerDoStatus        int8    `db:"car_speaker_do_status" gorm:"Column:car_speaker_do_status"`
	LastGPSTime               int32   `db:"last_gps_time" gorm:"Column:last_gps_time"`
	Engine                    int8    `db:"engine" gorm:"Column:engine"`
	CarMode                   int8    `db:"car_mode" gorm:"Column:car_mode"`
	ConfigChange              int8    `db:"config_change" gorm:"Column:config_change"`
	Alcohol                   int8    `db:"alcohol" gorm:"Column:alcohol"`
	Concentration             uint16  `db:"concentration" gorm:"Column:concentration"`
	AGPS                      string  `db:"agps" gorm:"Column:agps"`
	LastMileageID             *int64  `db:"last_mileage_id" gorm:"Column:last_mileage_id"`
	Endurance                 int64   `db:"endurance" gorm:"Column:endurance"`
	Rf                        *int8   `db:"rf" gorm:"Column:rf"`
	Rfid                      *string `db:"rfid" gorm:"Column:rfid"`
	AccChangeTime             int32   `db:"acc_change_time" gorm:"Column:acc_change_time"`
	StatusChangeTime          int32   `db:"status_change_time" gorm:"Column:status_change_time"`
}

func (TerminalInfo) TableName() string {
	return "t_terminal_info"
}

// TerminalOfflineHistory 终端离线历史表
type TerminalOfflineHistory struct {
	ID              int64   `db:"id" gorm:"Column:id"`
	Cid             *string `db:"cid" gorm:"Column:cid"`
	Sn              string  `db:"sn" gorm:"Column:sn"`
	FirmwareVersion string  `db:"firmware_version" gorm:"Column:firmware_version"`
	CreateAt        int32   `db:"create_at" gorm:"Column:create_at"`
	BatteryVoltage  int32   `db:"battery_voltage" gorm:"Column:battery_voltage"`
	Temp            *int32  `db:"temp" gorm:"Column:temp"`
	ChargeStatus    int32   `db:"charge_status" gorm:"Column:charge_status"`
	GSM             uint32  `db:"gsm" gorm:"Column:gsm"`
	GPS             uint32  `db:"gps" gorm:"Column:gps"`
	Pbat            *uint32 `db:"pbat" gorm:"Column:pbat"`
}

func (TerminalOfflineHistory) TableName() string {
	return "t_terminal_offline_history"
}

// TerminalPairlog 终端配对历史记录
type TerminalPairlog struct {
	ID          int64  `db:"id" gorm:"Column:id"`
	TidWired    string `db:"tid_wired" gorm:"Column:tid_wired"`
	TidWireless string `db:"tid_wireless" gorm:"Column:tid_wireless"`
	Type        int8   `db:"type" gorm:"Column:type"`
	PairTime    int64  `db:"pair_time" gorm:"Column:pair_time"`
	Remark      string `db:"remark" gorm:"Column:remark"`
}

func (TerminalPairlog) TableName() string {
	return "t_terminal_pairlog"
}

// UserFootPrint 短信密码登录用户的脚印表，最好只保留1年的数据
type UserFootPrint struct {
	Cid       *string `db:"cid" gorm:"Column:cid"`
	Mobile    *string `db:"mobile" gorm:"Column:mobile"`
	Name      *string `db:"name" gorm:"Column:name"`
	Time      *int32  `db:"time" gorm:"Column:time"`
	Vin       *string `db:"vin" gorm:"Column:vin"`
	OptType   *string `db:"opt_type" gorm:"Column:opt_type"`
	Ip        *string `db:"ip" gorm:"Column:ip"`
	Address   *string `db:"address" gorm:"Column:address"`
	Param     *string `db:"param" gorm:"Column:param"`
	LoginFrom *string `db:"login_from" gorm:"Column:login_from"`
}

func (UserFootPrint) TableName() string {
	return "t_user_foot_print"
}

type Warehouse struct {
	ID         string `db:"id" gorm:"Column:id"`
	CorpID     string `db:"corp_id" gorm:"Column:corp_id"`
	Name       string `db:"name" gorm:"Column:name"`
	CreateTime int64  `db:"create_time" gorm:"Column:create_time"`
	UpdateTime *int64 `db:"update_time" gorm:"Column:update_time"`
}

func (Warehouse) TableName() string {
	return "t_warehouse"
}

// WarningOption 是否推送告警到前台页面
type WarningOption struct {
	ID               int64  `db:"id" gorm:"Column:id"`
	Oid              string `db:"oid" gorm:"Column:oid"`
	Voice            uint8  `db:"voice" gorm:"Column:voice"`
	LowPower         uint8  `db:"low_power" gorm:"Column:low_power"`
	Offline          uint8  `db:"offline" gorm:"Column:offline"`
	PullOut          uint8  `db:"pull_out" gorm:"Column:pull_out"`
	Emergency        uint8  `db:"emergency" gorm:"Column:emergency"`
	Wait4emergency   uint8  `db:"wait4emergency" gorm:"Column:wait4emergency"`
	FullPower        uint8  `db:"full_power" gorm:"Column:full_power"`
	SelfTestError    uint8  `db:"self_test_error" gorm:"Column:self_test_error"`
	LowVoltage       int8   `db:"low_voltage" gorm:"Column:low_voltage"`
	LongStop         int8   `db:"long_stop" gorm:"Column:long_stop"`
	OverSpeed        int8   `db:"over_speed" gorm:"Column:over_speed"`
	LightSensed      int8   `db:"light_sensed" gorm:"Column:light_sensed"`
	AccOn            int8   `db:"acc_on" gorm:"Column:acc_on"`
	AccOff           int8   `db:"acc_off" gorm:"Column:acc_off"`
	FenceIn          int8   `db:"fence_in" gorm:"Column:fence_in"`
	FenceOut         int8   `db:"fence_out" gorm:"Column:fence_out"`
	OverMileageLimit int8   `db:"over_mileage_limit" gorm:"Column:over_mileage_limit"`
	DeviceRemoveRisk int8   `db:"device_remove_risk" gorm:"Column:device_remove_risk"`
	LightRemoveRisk  int8   `db:"light_remove_risk" gorm:"Column:light_remove_risk"`
}

func (WarningOption) TableName() string {
	return "t_warning_option"
}

// YhInstallersTerminal 一嗨定制版"安装员终端记录"表
type YhInstallersTerminal struct {
	ID                     int32   `db:"id" gorm:"Column:id"`
	Tid                    *string `db:"tid" gorm:"Column:tid"`
	InstallTime            *int32  `db:"install_time" gorm:"Column:install_time"`
	Installers             *string `db:"installers" gorm:"Column:installers"`
	Position               *string `db:"position" gorm:"Column:position"`
	ReplacedSn             *string `db:"replaced_sn" gorm:"Column:replaced_sn"`
	Sn                     *string `db:"sn" gorm:"Column:sn"`
	ActivateCode           *string `db:"activate_code" gorm:"Column:activate_code"`
	Vin                    *string `db:"vin" gorm:"Column:vin"`
	TypeName               *string `db:"type_name" gorm:"Column:type_name"`
	Style                  int32   `db:"style" gorm:"Column:style"`
	BatchName              *string `db:"batch_name" gorm:"Column:batch_name"`
	Loginname              *string `db:"loginname" gorm:"Column:loginname"`
	ReplacedReason         string  `db:"replaced_reason" gorm:"Column:replaced_reason"`
	ReplacedCode           int8    `db:"replaced_code" gorm:"Column:replaced_code"`
	BatchNo                string  `db:"batch_no" gorm:"Column:batch_no"`
	InstallModeCheckEscape int32   `db:"install_mode_check_escape" gorm:"Column:install_mode_check_escape"`
}

func (YhInstallersTerminal) TableName() string {
	return "t_yh_installers_terminal"
}

// YhSecondHandUnbind 二手车解绑历史
type YhSecondHandUnbind struct {
	ID           int32   `db:"id" gorm:"Column:id"`
	UnbindTime   int32   `db:"unbind_time" gorm:"Column:unbind_time"`
	Sn           string  `db:"sn" gorm:"Column:sn"`
	ActivateCode string  `db:"activate_code" gorm:"Column:activate_code"`
	Vin          string  `db:"vin" gorm:"Column:vin"`
	TypeName     string  `db:"type_name" gorm:"Column:type_name"`
	Loginname    *string `db:"loginname" gorm:"Column:loginname"`
	UnbindReason string  `db:"unbind_reason" gorm:"Column:unbind_reason"`
	UnbindCode   int8    `db:"unbind_code" gorm:"Column:unbind_code"`
	Operator     *string `db:"operator" gorm:"Column:operator"`
}

func (YhSecondHandUnbind) TableName() string {
	return "t_yh_second_hand_unbind"
}

type WechatUsers struct {
	ID     int64  `db:"id" gorm:"Column:id"`
	OpenID string `db:"open_id" gorm:"Column:open_id"`
	Oid    string `db:"oid" gorm:"Column:oid"`
}

func (WechatUsers) TableName() string {
	return "wechat_users"
}

