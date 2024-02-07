package series

func variantToChassis(variant string) string {
	switch variant {
	case "adr-prime", "adr-primec", "adr-primei", "adr-primeg", "adr-a", "adr-ap", "adr-d", "adr-b", "adr-cinder", "adr-w", "adr-wl":
		return "Adder"
	case "ach-primec", "ach-primei", "ach-primes":
		return "Arctic Cheetah"
	case "cou-primes", "cou-prime":
		return "Cougar"
	case "inc-1s":
		return "Incubus"
	case "jr7-iicc", "jr7-iico":
		return "Jenner IIC"
	case "kfx-primeg", "kfx-primei", "kfx-prime":
		return "Kit Fox"
	case "mlx-primec", "mlx-primei", "mlx-primes":
		return "Mist Lynx"
	case "pir-1s":
		return "Piranha"
	case "a":
		return "Urbanmech IIC"
	case "wlf-c", "wlf-cl":
		return "Wolfhound IIC"
	case "b":
		return "Commando"
	case "c":
		return "Firestarter"
	case "d":
		return "Flea"
	case "e":
		return "Javelin"
	case "m":
		return "Jenner"
	case "f":
		return "Locust"
	case "g":
		return "Osiris"
	case "h":
		return "Panther"
	case "i":
		return "Raven"
	case "j":
		return "Spider"
	case "k":
		return "Urbanmech"
	case "l":
		return "Wolfhound"
	default:
		return variant
	}
}

func variantDedupe(variant string) string {
	switch variant {
	case "adr-primec", "adr-primeg", "adr-primei":
		return "adr-prime"
	case "adr-ap":
		return "adr-a"
	case "adr-wl":
		return "adr-w"
	case "ach-primec", "ach-primei", "ach-primes":
		return "ach-prime"
	case "ach-ep":
		return "ach-e"
	case "cou-primes":
		return "cou-prime"
	case "cou-hp":
		return "cou-h"
	case "inc-1s":
		return "inc-1"
	case "inc-4p":
		return "inc-4"
	case "inc-5c":
		return "inc-5"
	case "jr7-iicc", "jr7-iico":
		return "jr7-iic"
	case "jr7-iic-ap":
		return "jr7-iic-a"
	case "kfx-primeg", "kfx-primei":
		return "kfx-prime"
	case "kfx-dc", "kfx-dp":
		return "kfx-d"
	case "kfx-gl":
		return "kfx-g"
	case "kfx-ps":
		return "kfx-p"
	case "mlx-primec", "mlx-primei", "mlx-primes":
		return "mlx-prime"
	case "mlx-gp":
		return "mlx-g"
	case "pir-1s":
		return "pir-1"
	case "pir-2p":
		return "pir-2"
	case "pir-dl":
		return "pir-D"
	case "wlf-cl":
		return "wlf-c"
	default:
		return variant
	}
}
