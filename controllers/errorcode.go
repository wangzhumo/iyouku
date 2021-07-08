package controllers

const RequestOk = "ok"
const RequestError = "error"

const PhoneEmpty = "手机号为空"
const PhoneFormatError = "手机号格式有误"
const PhoneAlreadyRegister = "手机号已经注册"
const PasswordEmpty = "密码为空"
const PasswordFormatError = "至少8个字符，至少1个大写字母，1个小写字母和1个数字,不能包含特殊字符"

const RegisterSucceed = "注册成功"
const RegisterFail = "注册失败"
const LoginSucceed = "登录成功"
const LoginFail = "登录失败"

const NoChannelID = "需要指定频道ID"
const NoRegionID = "需要指定RegionId"
const NoVideoID = "需要指定VideoID"
const NoTypeID = "需要指定TypeID"
const VideoChannelError = "请求频道广告失败"
const VideoEpisodesError = "请求视频剧集失败"
const VideoChannelHotError = "请求频道热播失败"
const VideoChannelRecommendError = "请求推荐视频失败"
const VideoChannelTypeError = "请求推荐视频失败"
const ChannelRegionError = "获取地区失败"
const ChannelTypeError = "获取类型失败"
const ChannelVideoError = "获取视频失败"
