package core

var GetUserNotFound = ErrorRes{
	Code:    "c_uh_get_0001",
	Message: "user not found",
}

var PostDecodeErr = ErrorRes{
	Code:    "c_uh_post_0001",
	Message: "invalid post body payload",
}

var PostBadData = ErrorRes{
	Code:    "c_uh_post_0002",
	Message: "user data is invalid",
}

var PostReachedCap = ErrorRes{
	Code:    "c_uh_post_0003",
	Message: "reached users capacity",
}

var PutBadStructure = ErrorRes{
	Code:    "c_uh_put_0001",
	Message: "invalid put body structure",
}

var PutBadData = ErrorRes{
	Code:    "c_uh_put_0002",
	Message: "invalid put body data",
}

var PutUserNotFound = ErrorRes{
	Code:    "c_uh_put_0003",
	Message: "did not find user to update",
}

var DeleteBadStructure = ErrorRes{
	Code:    "c_uh_delete_0001",
	Message: "invalid delete body payload",
}

var DeleteMissingId = ErrorRes{
	Code:    "c_uh_delete_0001",
	Message: "missing user id",
}

var DeleteUserNotFound = ErrorRes{
	Code:    "c_uh_delete_0002",
	Message: "did not find user to delete",
}
