package nex

import (
	"fmt"
	"os"

	"github.com/ItzSwirlz/animal-crossing-new-leaf/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(2)
	globals.AuthenticationServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 10, 1))
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("d6f08b40")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Animal Crossing: New Leaf - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	// registerCommonAuthenticationServerProtocols() TODO: uncomment

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_ACNL_AUTHENTICATION_SERVER_PORT")))
}
