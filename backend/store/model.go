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
	Id                             int64  `json:"id"`
	CreatedAt                      int64  `json:"-"`
	UpdatedAt                      int64  `json:"-"`
	DateReceipt                    int64  `json:"dateReceipt"`
	DoctorId                       int64  `json:"doctorId"`
	DoctorName                     string `json:"doctorName" docx:"doctorName"`
	PatientId                      int64  `json:"patientId"`
	PatientName                    string `json:"patientName" docx:"patientName"`
	ReceiptKindId                  int64  `json:"receiptKindId"`
	ReceiptKindName                string `json:"-"`
	ReceiptDiagnosis               string `json:"receiptDiagnosis"`
	Alergo                         string `json:"alergo" docx:"alergo"`
	ContactInfected                string `json:"contactInfected" docx:"contactInfected"`
	Hiv                            string `json:"hiv" docx:"hiv"`
	Transfusion                    string `json:"transfusion" docx:"transfusion"`
	Dyscountry                     string `json:"dyscountry" docx:"dyscountry"`
	Smoking                        string `json:"smoking" docx:"smoking"`
	Drugs                          string `json:"drugs" docx:"drugs"`
	Inheritance                    string `json:"inheritance" docx:"inheritance"`
	Gyndiseases                    string `json:"gyndiseases" docx:"gyndiseases"`
	Paritet                        string `json:"paritet"`
	ParitetB                       string `json:"paritetB"`
	ParitetP                       string `json:"paritetP"`
	ParitetA                       string `json:"paritetA"`
	ParitetSV                      string `json:"paritetSV"`
	ParitetNB                      string `json:"paritetNB"`
	ParitetEB                      string `json:"paritetEB"`
	InfectionMarkersStateId        int64  `json:"infectionMarkersStateId"`
	InfectionMarkersStateName      string `json:"-"`
	InfectionMarkersDesc           string `json:"infectionMarkersDesc"`
	TromboflebiaStateId            int64  `json:"tromboflebiaStateId"`
	TromboflebiaStateName          string `json:"-"`
	TromboflebiaDesc               string `json:"tromboflebiaDesc"`
	FirstTrimester                 string `json:"firstTrimester" docx:"firstTrimester"`
	SecondTrimester                string `json:"secondTrimester" docx:"secondTrimester"`
	ThirdTrimester                 string `json:"thirdTrimester" docx:"thirdTrimester"`
	History                        string `json:"history" docx:"history"`
	Oprv                           string `json:"oprv"`
	OprvStateId                    int64  `json:"oprvStateId"`
	OprvStateName                  string `json:"-"`
	ExpByMenstruation              string `json:"expByMenstruation" docx:"expByMenstruation"`
	ExpByFirstVisit                string `json:"expByFirstVisit" docx:"expByFirstVisit"`
	ExpByUltraFirst                string `json:"expByUltraFirst"`
	ExpByUltraSecond               string `json:"expByUltraSecond"`
	ExpByUltraThird                string `json:"expByUltraThird"`
	HealthStateId                  int64  `json:"healthStateId"`
	HealthStateName                string `json:"-" docx:"health"`
	Claims                         string `json:"claims" docx:"claims"`
	Head                           string `json:"head" docx:"headache"`
	Vision                         string `json:"vision" docx:"vision"`
	SkinStateId                    int64  `json:"skinStateId"`
	SkinStateName                  string `json:"-" docx:"skinStateName"`
	Lymph                          string `json:"lymph" docx:"lymph"`
	BreathStateId                  int64  `json:"breathStateId"`
	BreathStateName                string `json:"-" docx:"breathStateName"`
	RaleStateId                    int64  `json:"raleStateId"`
	RaleStateName                  string `json:"-" docx:"raleStateName"`
	TonesStateId                   int64  `json:"tonesStateId"`
	TonesStateName                 string `json:"-" docx:"tonesStateName"`
	Pulse                          string `json:"pulse" docx:"pulseValue"`
	PulseType                      string `json:"pulseType" docx:"pulseType"`
	Pressure                       string `json:"pressure" docx:"pressure"`
	TongueClean                    bool   `json:"tongueClean"`
	TongueWet                      bool   `json:"tongueWet"`
	TongueDry                      bool   `json:"tongueDry"`
	TongueCoated                   bool   `json:"tongueCoated"`
	TongueUncoated                 bool   `json:"tongueUncoated"`
	Throat                         string `json:"throat" docx:"throat"`
	BellyPeriod                    string `json:"bellyPeriod" docx:"bellyPeriod"`
	BellyStateId                   int64  `json:"bellyStateId"`
	BellyStateName                 string `json:"-" docx:"bellyStateName"`
	EpigastriumStateUse            bool   `json:"epigastriumStateUse"`
	EpigastriumStateId             int64  `json:"epigastriumStateId"`
	EpigastriumStateName           string `json:"-"`
	ScarStateUse                   bool   `json:"scarStateUse"`
	ScarStateId                    int64  `json:"scarStateId"`
	ScarStateName                  string `json:"-"`
	Peritoneal                     string `json:"peritoneal" docx:"peritoneal"`
	Labors                         string `json:"labors" docx:"labors"`
	Dysuric                        bool   `json:"dysuric"`
	Bowel                          bool   `json:"bowel"`
	LimbSwelling                   string `json:"limbSwelling" docx:"limbSwelling"`
	UteruseStateId                 int64  `json:"uteruseStateId"`
	UteruseStateName               string `json:"-" docx:"uteruseStateName"`
	FetalPositionId                int64  `json:"fetalPositionId"`
	FetalPositionName              string `json:"-" docx:"fetalPositionName"`
	FetalPreviaId                  int64  `json:"fetalPreviaId"`
	FetalPreviaName                string `json:"-" docx:"fetalPreviaName"`
	FetalAlignId                   int64  `json:"fetalAlignId"`
	FetalAlignName                 string `json:"-" docx:"fetalAlignName"`
	FetalHeartbeatId               int64  `json:"fetalHeartbeatId"`
	FetalHeartbeatName             string `json:"-" docx:"fetalHeartbeatName"`
	HeartbeatRithmId               int64  `json:"heartbeatRithmId"`
	HeartbeatRithmName             string `json:"-" docx:"heartbeatRithmName"`
	FetalPulse                     string `json:"fetalPulse" docx:"fetalPulse"`
	ReproductiveDischargeTypeId    int64  `json:"reproductiveDischargeTypeId"`
	ReproductiveDischargeTypeName  string `json:"-" docx:"reproductiveDischargeTypeName"`
	ReproductiveDischargeStateId   int64  `json:"reproductiveDischargeStateId"`
	ReproductiveDischargeStateName string `json:"-" docx:"reproductiveDischargeStateName"`
	Vdm                            string `json:"vdm" docx:"vdm"`
	Oj                             string `json:"oj" docx:"oj"`
	Dspin                          string `json:"dspin" docx:"dspin"`
	Dcrist                         string `json:"dcrist" docx:"dcrist"`
	Dtroch                         string `json:"dtroch" docx:"dtroch"`
	Cext                           string `json:"cext" docx:"cext"`
	DevelOrgansId                  int64  `json:"develOrgansId"`
	DevelOrgansName                string `json:"-" docx:"develOrgansName"`
	GenitalAnomalies               string `json:"genitalAnomalies" docx:"genitalAnomalies"`
	VaginaStateId                  int64  `json:"vaginaStateId"`
	VaginaStateName                string `json:"-" docx:"vaginaStateName"`
	Bishop                         string `json:"bishop" docx:"bishop"`
	FetalBladderStateId            int64  `json:"fetalBladderStateId"`
	FetalBladderStateName          string `json:"-" docx:"fetalBladderStateName"`
	FetalBladderPreviaId           int64  `json:"fetalBladderPreviaId"`
	FetalBladderPreviaName         string `json:"-" docx:"fetalBladderPreviaName"`
	FetalBladderAlignId            int64  `json:"fetalBladderAlignId"`
	FetalBladderAlignName          string `json:"-" docx:"fetalBladderAlignName"`
	Arches                         string `json:"arches" docx:"arches"`
	Conjugate                      string `json:"conjugate" docx:"conjugate"`
	PelvisStateId                  int64  `json:"pelvisStateId"`
	PelvisStateName                string `json:"-" docx:"pelvisStateName"`
	PelvisExostosis                string `json:"pelvisExostosis" docx:"pelvisExostosis"`
	PelvisDischargeTypeId          int64  `json:"pelvisDischargeTypeId"`
	PelvisDischargeTypeName        string `json:"-" docx:"pelvisDischargeTypeName"`
	PelvisDischargeStateId         int64  `json:"pelvisDischargeStateId"`
	PelvisDischargeStateName       string `json:"-" docx:"pelvisDischargeStateName"`
	Diagnosis                      string `json:"diagnosis" docx:"diagnosis"`
	Conclusion                     string `json:"conclusion" docx:"conclusion"`
	BirthPlanUse                   bool   `json:"birthPlanUse"`
	BirthPlan                      string `json:"birthPlan"`
}
