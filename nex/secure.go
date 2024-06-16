package nex

import (
	"fmt"
	"os"

	"github.com/ItzSwirlz/animal-crossing-new-leaf/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(2)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 10, 1))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("d6f08b40")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Animal Crossing: New Leaf - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	// registerCommonSecureServerProtocols() TODO: uncomment
	// registerSecureServerNEXProtocols() TODO: uncomment

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_ACNL_SECURE_SERVER_PORT")))
}
