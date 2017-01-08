package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/austinov/gyn/backend/store"
	"github.com/nguyenthenguyen/docx"
)

type FillDocxCallback func(doc *docx.Docx) error

func FillDocx(ap store.Appointment, templatePath string, cb FillDocxCallback) (*os.File, error) {
	doc, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		return nil, err
	}
	defer doc.Close()
	tmpDocx := doc.Editable()
	tmpDocx.Replace("[dateReceipt]", fmt.Sprintf("%d", ap.DateReceipt), -1)
	tmpDocx.Replace("[doctorName]", ap.DoctorName, -1)
	tmpDocx.Replace("[howReceipt]", ap.HowReceipt, -1)
	tmpDocx.Replace("[alergo]", ap.Alergo, -1)
	tmpDocx.Replace("[contactInfected]", ap.ContactInfected, -1)
	tmpDocx.Replace("[hiv]", ap.Hiv, -1)
	tmpDocx.Replace("[transfusion]", ap.Transfusion, -1)
	tmpDocx.Replace("[dyscountry]", ap.Dyscountry, -1)
	tmpDocx.Replace("[smoking]", ap.Smoking, -1)
	tmpDocx.Replace("[drugs]", ap.Drugs, -1)
	tmpDocx.Replace("[inheritance]", ap.Inheritance, -1)
	tmpDocx.Replace("[diseases]", ap.Diseases, -1)
	tmpDocx.Replace("[gyndiseases]", ap.Gyndiseases, -1)
	tmpDocx.Replace("[history]", ap.History, -1)
	tmpDocx.Replace("[paritet]", ap.Paritet, -1)
	tmpDocx.Replace("[pregnancy]", ap.Pregnancy, -1)
	tmpDocx.Replace("[firstTrimester]", ap.FirstTrimester, -1)
	tmpDocx.Replace("[secondTrimester]", ap.SecondTrimester, -1)
	tmpDocx.Replace("[thirdTrimester]", ap.ThirdTrimester, -1)

	if cb != nil {
		if err = cb(tmpDocx); err != nil {
			return nil, err
		}
	}

	file, err := ioutil.TempFile(os.TempDir(), "")
	if err != nil {
		return nil, err
	}
	tmpDocx.WriteToFile(file.Name())
	return file, nil
}
