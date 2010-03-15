$ckage main

import (
        "fastweb"
        "os"
)

type Products struct {
        fastweb.Controller
        name string
        brand string
        features []string
        specifications []string
        image string
}

func (p *Products) View(id string) os.Error {
        if id == "ah64" {
                p.name = "RC Apache AH64 4-Channel Electric Helicoper"
                p.brand = "Colco"
                p.features = []string{
                        "4 channel radio control duel propeller system",
                        "Full movement controll: forward, backward, left, right, up and down",
                        "Replica design",
                        "Revolutionary co-axial rotor technology",
                }
                p.specifications = []string{
                        "Dimensions: L 16 Inches X W 5.5 Inches x H 6.5 Inches",
                        "Battery Duration: 10 min",
                        "Range: 120 Feet",
                }
                p.image = "/img/ah64.jpg"
        }
        return nil
}

func main() {
        a := fastweb.NewApplication()
        a.RegisterController(&Products{})
        a.Run(":80")
}
