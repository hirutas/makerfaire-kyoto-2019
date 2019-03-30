package main

import (
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	m1 := gpio.NewMotorDriver(r, "")
	m1.ForwardPin = "21"
	m1.BackwardPin = "22"

	m2 := gpio.NewMotorDriver(r, "")
	m2.ForwardPin = "23"
	m2.BackwardPin = "24"

	m3 := gpio.NewMotorDriver(r, "")
	m3.ForwardPin = "26"
	m3.BackwardPin = "29"

	m4 := gpio.NewMotorDriver(r, "")
	m4.ForwardPin = "31"
	m4.BackwardPin = "32"

	m5 := gpio.NewMotorDriver(r, "")
	m5.ForwardPin = "35"
	m5.BackwardPin = "33"

	swStart := gpio.NewButtonDriver(r, "37")

	sw1a := gpio.NewButtonDriver(r, "7")
	sw1b := gpio.NewButtonDriver(r, "8")
	sw2a := gpio.NewButtonDriver(r, "10")
	sw2b := gpio.NewButtonDriver(r, "12")
	sw3a := gpio.NewButtonDriver(r, "11")
	sw3b := gpio.NewButtonDriver(r, "13")
	sw4a := gpio.NewButtonDriver(r, "15")
	sw4b := gpio.NewButtonDriver(r, "16")
	sw5a := gpio.NewButtonDriver(r, "18")
	sw5b := gpio.NewButtonDriver(r, "19")
	fmt.Println("setup")

	err := m1.Direction("none")
	if err != nil {
		log.Fatal("direction error", err)
	}
	fmt.Printf("m1 direction: %v\n", m1.CurrentDirection)
	err = m2.Direction("none")
	if err != nil {
		log.Fatal("direction error", err)
	}
	fmt.Printf("m2 direction: %v\n", m2.CurrentDirection)
	err = m3.Direction("none")
	if err != nil {
		log.Fatal("direction error", err)
	}
	fmt.Printf("m3 direction: %v\n", m3.CurrentDirection)
	err = m4.Direction("none")
	if err != nil {
		log.Fatal("direction error", err)
	}
	fmt.Printf("m4 direction: %v\n", m4.CurrentDirection)
	err = m5.Direction("none")
	if err != nil {
		log.Fatal("direction error", err)
	}
	fmt.Printf("m5 direction: %v\n", m5.CurrentDirection)

	work := func() {
		err := swStart.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("swStart:", s)
			err := m1.Direction("forward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m1 direction: %v\n", m1.CurrentDirection)
			err = m2.Direction("forward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m2 direction: %v\n", m2.CurrentDirection)
			err = m3.Direction("forward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m3 direction: %v\n", m3.CurrentDirection)
			err = m4.Direction("forward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m4 direction: %v\n", m4.CurrentDirection)
			err = m5.Direction("forward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m5 direction: %v\n", m5.CurrentDirection)

			go func() {
				select {
				case <-time.After(15 * time.Second):
					fmt.Println("timeout")
					if m1.CurrentDirection == "forward" {
						err := m1.Direction("backward")
						if err != nil {
							log.Fatal("direction error", err)
						}
						fmt.Printf("m1 direction: %v\n", m1.CurrentDirection)
					}
					if m2.CurrentDirection == "forward" {
						err = m2.Direction("backward")
						if err != nil {
							log.Fatal("direction error", err)
						}
						fmt.Printf("m2 direction: %v\n", m2.CurrentDirection)
					}
					if m3.CurrentDirection == "forward" {
						err = m3.Direction("backward")
						if err != nil {
							log.Fatal("direction error", err)
						}
						fmt.Printf("m3 direction: %v\n", m3.CurrentDirection)
					}
					if m4.CurrentDirection == "forward" {
						err = m4.Direction("backward")
						if err != nil {
							log.Fatal("direction error", err)
						}
						fmt.Printf("m4 direction: %v\n", m4.CurrentDirection)
					}
					if m5.CurrentDirection == "forward" {
						err = m5.Direction("backward")
						if err != nil {
							log.Fatal("direction error", err)
						}
						fmt.Printf("m5 direction: %v\n", m5.CurrentDirection)
					}
				}
			}()
		})
		if err != nil {
			log.Fatal("swStart.On", err)
		}

		err = sw1a.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw1a:", s)
			if m1.CurrentDirection != "forward" {
				fmt.Println("m1 direction is not forward. ignore")
				return
			}
			err := m1.Direction("backward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m1 direction: %v\n", m1.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw1a.On", err)
		}

		err = sw1b.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw1b:", s)
			err := m1.Direction("none")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m1 direction: %v\n", m1.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw1b.On", err)
		}

		err = sw2a.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw2a:", s)
			if m2.CurrentDirection != "forward" {
				fmt.Println("m2 direction is not forward. ignore")
				return
			}
			err := m2.Direction("backward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m2 direction: %v\n", m2.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw2a.On", err)
		}

		err = sw2b.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw2b:", s)
			err := m2.Direction("none")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m2 direction: %v\n", m2.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw2b.On", err)
		}

		err = sw3a.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw3a:", s)
			if m3.CurrentDirection != "forward" {
				fmt.Println("m3 direction is not forward. ignore")
				return
			}
			err := m3.Direction("backward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m3 direction: %v\n", m3.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw3a.On", err)
		}

		err = sw3b.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw3b:", s)
			err := m3.Direction("none")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m3 direction: %v\n", m3.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw3b.On", err)
		}

		err = sw4a.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw4a:", s)
			if m4.CurrentDirection != "forward" {
				fmt.Println("m4 direction is not forward. ignore")
				return
			}
			err := m4.Direction("backward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m4 direction: %v\n", m4.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw4a.On", err)
		}

		err = sw4b.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw4b:", s)
			err := m4.Direction("none")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m4 direction: %v\n", m4.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw4b.On", err)
		}

		err = sw5a.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw5a:", s)
			if m5.CurrentDirection != "forward" {
				fmt.Println("m5 direction is not forward. ignore")
				return
			}
			err := m5.Direction("backward")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m5 direction: %v\n", m5.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw5a.On", err)
		}

		err = sw5b.On(gpio.ButtonPush, func(s interface{}) {
			fmt.Println("sw5b:", s)
			err := m5.Direction("none")
			if err != nil {
				log.Fatal("direction error", err)
			}
			fmt.Printf("m5 direction: %v\n", m5.CurrentDirection)
		})
		if err != nil {
			log.Fatal("sw5b.On", err)
		}
	}

	robot := gobot.NewRobot("tongari",
		[]gobot.Connection{r},
		[]gobot.Device{sw1a, sw1b, sw2a, sw2b, sw3a, sw3b, sw4a, sw4b, sw5a, sw5b, swStart, m1, m2, m3, m4, m5},
		work,
	)
	err = robot.Start()
	if err != nil {
		log.Fatal("robot.Start", err)
	}
}
