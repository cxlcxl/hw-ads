package serviceauth

import (
	"bs.mobgi.cc/app/model"
	"bs.mobgi.cc/app/validator/v_data"
	"bs.mobgi.cc/app/vars"
	"strconv"
)

func RoleUpdate(role *v_data.VRoleUpdate) error {
	d := map[string]interface{}{
		"role_name": role.RoleName,
		"state":     role.State,
	}
	ps, err := model.NewPermission(vars.DBMysql).FindPermissionsByPers(role.Permissions)
	if err != nil {
		return err
	}
	var aps []*model.PR
	for _, v := range ps {
		aps = append(aps, &model.PR{
			PType: "p",
			V0:    strconv.Itoa(int(role.Id)),
			V1:    v.Per,
			V2:    v.Method,
		})
	}
	return model.NewRole(vars.DBMysql).UpdateRole(d, role.Id, aps)
}
