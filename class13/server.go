package class13

import (
	"fmt"
	"net/http"
)

func PlayersServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
