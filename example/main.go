package main

import (
	"github.com/rentziass/wallet"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	pass := &wallet.Pass{
		FormatVersion:       1,
		// This has to be the identifier used when
		// creating certificates on Apple Developer Portal
		PassTypeIdentifier:  "pass.com.company.pass_name",
		SerialNumber:        "nmyuxofgna",
		// This is your team identifier, you grab yours here
		// https://developer.apple.com/account/#/membership ,
		// you will find it under 'Team ID'
		TeamIdentifier:      "XXXXXXXXXX",
		WebServiceURL:       "https://example.com/passes/",
		AuthenticationToken: "vxwxd7J8AlNNFPS8k0a0FfUFtq0ewzFdc",
		RelevantDate:        "2018-10-09T16:00Z",
		Barcode:             wallet.NewPassBarcode("123456789", "PKBarcodeFormatQR"),
		OrganizationName:    "SpaceX",
		Description:         "SpaceX event ticket",
		ForegroundColor:     "rgb(255, 255, 255)",
		BackgroundColor:     "rgb(60, 65, 76)",
		LogoText:            "SpaceX",
		LabelColor:          "rgb(255, 255, 255)",
		EventTicketDetails: &wallet.EventTicketDetails{
			PrimaryFields: []*wallet.Field{
				{
					Key:   "event",
					Label: "EVENT",
					Value: "Mars, here we come!",
				},
			},
			SecondaryFields: []*wallet.Field{
				{
					Key:   "loc",
					Label: "PLANET",
					Value: "Mars",
				},
			},
			AuxiliaryFields: []*wallet.Field{
				{
					Key:   "address",
					Label: "ADDRESS",
					Value: "Third rock to the left",
				},
				{
					Key:   "time",
					Label: "AT",
					Value: "17:00",
				},
			},
		},
	}

	// You need to provide Apple WWDR certificate along with the certificate and key you
	// you created specifically for your pass on Apple Developer Portal.
	// More details on how to do that here:
	// https://www.raywenderlich.com/2855-beginning-passbook-in-ios-6-part-1-2
	w, err := wallet.NewWriter(pass, "./WWDR.pem", "./passcertificate.pem", "./passkey.pem", "password")
	if err != nil {
		panic(err)
	}

	err = filepath.Walk("./pass/", func(path string, info os.FileInfo, err error) error {
		if info.Name() == ".DS_Store" {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		err = w.AddFile(info.Name(), b)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	buf, err := w.Close()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("mars.pkpass")
	if err != nil {
		panic(err)
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
}
