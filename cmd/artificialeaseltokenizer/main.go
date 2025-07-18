// cmd/artificialeaseltokenizer/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaseltokenizer/internal/artificialeaseltokenizer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaseltokenizer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
