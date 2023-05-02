package common

type InfoPrinter interface {
	PrintInfo()
}

func PrintInfo(infoPrinter InfoPrinter) {
	infoPrinter.PrintInfo()
}
