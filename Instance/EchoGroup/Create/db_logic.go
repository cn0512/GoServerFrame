package main

import (
	RPC "MSvrs/Core/RPC"
)

func CheckDatabase(u RPC.CreateUserReq) RPC.CreateUserRep {
	ret := RPC.CreateUserRep{
		Uuid: u.Uuid,
		Ret:  0,
		ID:   1000,
	}

	return ret
}
