package users

import (
	"fmt"
	"time"

	"github.com/angelik-mag/practice_go/modelos"
)

func AltaUsuario() {

	u := new(modelos.User)

	u.AddUser(10, "Angelica", time.Now(), true)

	fmt.Println(u)

}
