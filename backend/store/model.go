package store

type Profile struct {
	Id       int64  `json:"-"`
	Login    string `json:"login"`
	UserName string `json:"username"`
}

type Dictionary struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Appointment struct {
	Id                        int64  `json:"id"`
	DateReceipt               int64  `json:"dateReceipt"`
	DoctorId                  int64  `json:"doctorId"`
	DoctorName                string `json:"doctorName" docx:"[doctorName]"`
	PatientId                 int64  `json:"patientId"`
	PatientName               string `json:"patientName" docx:"[patientName]"`
	HowReceipt                string `json:"howReceipt" docx:"[howReceipt]"`
	Alergo                    string `json:"alergo" docx:"[alergo]"`
	ContactInfected           string `json:"contactInfected" docx:"[contactInfected]"`
	Hiv                       string `json:"hiv" docx:"[hiv]"`
	Transfusion               string `json:"transfusion" docx:"[transfusion]"`
	Dyscountry                string `json:"dyscountry" docx:"[dyscountry]"`
	Smoking                   string `json:"smoking" docx:"[smoking]"`
	Drugs                     string `json:"drugs" docx:"[drugs]"`
	Inheritance               string `json:"inheritance" docx:"[inheritance]"`
	Diseases                  string `json:"diseases" docx:"[diseases]"`
	Gyndiseases               string `json:"gyndiseases" docx:"[gyndiseases]"`
	Paritet                   string `json:"paritet" docx:"[paritet]"`
	Pregnancy                 string `json:"pregnancy" docx:"[pregnancy]"`
	FirstTrimester            string `json:"firstTrimester" docx:"[firstTrimester]"`
	SecondTrimester           string `json:"secondTrimester" docx:"[secondTrimester]"`
	ThirdTrimester            string `json:"thirdTrimester" docx:"[thirdTrimester]"`
	History                   string `json:"history" docx:"[history]"`
	ExpByMenstruation         string `json:"expByMenstruation" docx:"[expByMenstruation]"`
	ExpByFirstVisit           string `json:"expByFirstVisit" docx:"[expByFirstVisit]"`
	ExpByUltra                string `json:"expByUltra" docx:"[expByUltra]"`
	HealthStateId             int64  `json:"healthStateId"`
	HealthStateName           string `json:"-" docx:"[healthStateName]"`
	Claims                    string `json:"claims" docx:"[claims]"`
	Head                      string `json:"head" docx:"[head]"`
	Vision                    string `json:"vision" docx:"[vision]"`
	SkinStateId               int64  `json:"skinStateId"`
	SkinStateName             string `json:"-" docx:"[skinStateName]"`
	Lymph                     string `json:"lymph" docx:"[lymph]"`
	Breath                    string `json:"breath" docx:"[breath]"`
	Rale                      string `json:"rale" docx:"[rale]"`
	Tones                     string `json:"tones" docx:"[tones]"`
	Pulse                     string `json:"pulse" docx:"[pulse]"`
	PulseType                 string `json:"pulseType" docx:"[pulseType]"`
	Pressure                  string `json:"pressure" docx:"[pressure]"`
	TongueClean               bool   `json:"tongueClean"`
	TongueWet                 bool   `json:"tongueWet"`
	TongueDry                 bool   `json:"tongueDry"`
	TongueCoated              bool   `json:"tongueCoated"`
	TongueUncoated            bool   `json:"tongueUncoated"`
	Throat                    string `json:"throat" docx:"[throat]"`
	Belly                     string `json:"belly" docx:"[belly]"`
	Peritoneal                string `json:"peritoneal" docx:"[peritoneal]"`
	Labors                    string `json:"labors" docx:"[labors]"`
	Dysuric                   bool   `json:"dysuric"`
	Bowel                     bool   `json:"bowel"`
	LimbSwelling              string `json:"limbSwelling" docx:"[limbSwelling]"`
	FaceSwelling              string `json:"faceSwelling" docx:"[faceSwelling]"`
	UteruseStateId            int64  `json:"uteruseStateId"`
	UteruseStateName          string `json:"-" docx:"[uteruseStateName]"`
	FetalPositionId           int64  `json:"fetalPositionId"`
	FetalPositionName         string `json:"-" docx:"[fetalPositionName]"`
	FetalPreviaId             int64  `json:"fetalPreviaId"`
	FetalPreviaName           string `json:"-" docx:"[fetalPreviaName]"`
	FetalAlignId              int64  `json:"fetalAlignId"`
	FetalAlignName            string `json:"-" docx:"[fetalAlignName]"`
	FetalHeartbeatId          int64  `json:"fetalHeartbeatId"`
	FetalHeartbeatName        string `json:"-" docx:"[fetalHeartbeatName]"`
	FetalPulse                string `json:"fetalPulse" docx:"[fetalPulse]"`
	ReproductiveDischargeId   int64  `json:"reproductiveDischargeId"`
	ReproductiveDischargeName string `json:"" docx:"[reproductiveDischargeName]"`
	Vdm                       string `json:"vdm" docx:"[vdm]"`
	Oj                        string `json:"oj" docx:"[oj]"`
	Dspin                     string `json:"dspin" docx:"[dspin]"`
	Dcrist                    string `json:"dcrist" docx:"[dcrist]"`
	Dtroch                    string `json:"dtroch" docx:"[dtroch]"`
	Cext                      string `json:"cext" docx:"[cext]"`
	DevelOrgansId             int64  `json:"develOrgansId"`
	DevelOrgansName           string `json:"-" docx:"[develOrgansName]"`
	GenitalAnomalies          string `json:"genitalAnomalies" docx:"[genitalAnomalies]"`
	VaginaStateId             int64  `json:"vaginaStateId"`
	VaginaStateName           string `json:"-" docx:"[vaginaStateName]"`
	LenghtCervix              string `json:"lenghtCervix" docx:"[lenghtCervix]"`
	TruncateCervix            string `json:"truncateCervix" docx:"[truncateCervix]"`
	OuterThroatStateId        int64  `json:"outerThroatStateId"`
	OuterThroatStateName      string `json:"-" docx:"[outerThroatStateName]"`
	ChannelCervix             string `json:"channelCervix" docx:"[channelCervix]"`
	FetalBladder              string `json:"fetalBladder" docx:"[fetalBladder]"`
	FetalBladderPreviaId      int64  `json:"fetalBladderPreviaId"`
	FetalBladderPreviaName    string `json:"-" docx:"[fetalBladderPreviaName]"`
	FetalBladderAlignId       int64  `json:"fetalBladderAlignId"`
	FetalBladderAlignName     string `json:"-" docx:"[fetalBladderAlignName]"`
	Arches                    string `json:"arches" docx:"[arches]"`
	Conjugate                 string `json:"conjugate" docx:"[conjugate]"`
	PelvisStateId             int64  `json:"pelvisStateId"`
	PelvisStateName           string `json:"-" docx:"[pelvisStateName]"`
	PelvisDischarge           string `json:"pelvisDischarge" docx:"[pelvisDischarge]"`
	Diagnosis                 string `json:"diagnosis" docx:"[diagnosis]"`
	Conclusion                string `json:"conclusion" docx:"[conclusion]"`
	BirthPlan                 string `json:"birthPlan" docx:"[birthPlan]"`
}
