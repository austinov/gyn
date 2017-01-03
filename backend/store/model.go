package store

type Profile struct {
	Login    string `json:"login"`
	UserName string `json:"username"`
}

type Dictionary struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Appointment struct {
	Id                      int64  `json:"id"`
	Date                    int64  `json:"date"`
	DoctorId                int64  `json:"doctorId"`
	DoctorName              string `json:"doctorName"`
	PatientId               int64  `json:"patientId"`
	PatientName             string `json:"patientName"`
	HowReceipt              string `json:"howReceipt"`
	Alergo                  string `json:"alergo"`
	ContactInfectied        string `json:"contactInfectied"`
	Hiv                     string `json:"hiv"`
	Transfusion             string `json:"transfusion"`
	Dyscountry              string `json:"dyscountry"`
	Smoking                 string `json:"smoking"`
	Drugs                   string `json:"drugs"`
	Inheritance             string `json:"inheritance"`
	Diseases                string `json:"diseases"`
	Gyndiseases             string `json:"gyndiseases"`
	Paritet                 string `json:"paritet"`
	Pregnancy               string `json:"pregnancy"`
	FirstTrimester          string `json:"firstTrimester"`
	SecondTrimester         string `json:"secondTrimester"`
	ThirdTrimester          string `json:"thirdTrimester"`
	History                 string `json:"history"`
	ExpByMenstruation       string `json:"expByMenstruation"`
	ExpByFirstVisit         string `json:"expByFirstVisit"`
	ExpByUltra              string `json:"expByUltra"`
	HealthStateId           int64  `json:"healthStateId"`
	Claims                  string `json:"claims"`
	Head                    string `json:"head"`
	Vision                  string `json:"vision"`
	SkinStateId             int64  `json:"skinStateId"`
	Lymph                   string `json:"lymph"`
	Breath                  string `json:"breath"`
	Rale                    string `json:"rale"`
	Tones                   string `json:"tones"`
	Pulse                   string `json:"pulse"`
	PulseType               string `json:"pulseType"`
	Pressure                string `json:"pressure"`
	TongueClean             bool   `json:"tongueClean"`
	TongueWet               bool   `json:"tongueWet"`
	TongueDry               bool   `json:"tongueDry"`
	TongueCoated            bool   `json:"tongueCoated"`
	TongueUncoated          bool   `json:"tongueUncoated"`
	Throat                  string `json:"throat"`
	Belly                   string `json:"belly"`
	Peritoneal              string `json:"peritoneal"`
	Labors                  string `json:"labors"`
	Dysuric                 bool   `json:"dysuric"`
	Bowel                   bool   `json:"bowel"`
	LimbSwelling            string `json:"limbSwelling"`
	FaceSwelling            string `json:"faceSwelling"`
	UteruseStateId          int64  `json:"uteruseStateId"`
	FetalPositionId         int64  `json:"fetalPositionId"`
	FetalPreviaId           int64  `json:"fetalPreviaId"`
	FetalAlignId            int64  `json:"fetalAlignId"`
	FetalHeartbeatId        int64  `json:"fetalHeartbeatId"`
	FetalPulse              string `json:"fetalPulse"`
	ReproductiveDischargeId int64  `json:"reproductiveDischargeId"`
	Vdm                     string `json:"vdm"`
	Oj                      string `json:"oj"`
	Dspin                   string `json:"dspin"`
	Dcrist                  string `json:"dcrist"`
	Dtroch                  string `json:"dtroch"`
	Cext                    string `json:"cext"`
	DevelOrgansId           int64  `json:"develOrgansId"`
	GenitalAnomalies        string `json:"genitalAnomalies"`
	VaginaStateId           int64  `json:"vaginaStateId"`
	LenghtCervix            string `json:"lenghtCervix"`
	TruncateCervix          string `json:"truncateCervix"`
	OuterThroatStateId      int64  `json:"outerThroatStateId"`
	ChannelCervix           string `json:"channelCervix"`
	FetalBladder            string `json:"fetalBladder"`
	FetalBladderPreviaId    int64  `json:"fetalBladderPreviaId"`
	FetalBladderAlignId     int64  `json:"fetalBladderAlignId"`
	Arches                  string `json:"arches"`
	Conjugate               string `json:"conjugate"`
	PelvisStateId           int64  `json:"pelvisStateId"`
	PelvisDischarge         string `json:"pelvisDischarge"`
	Diagnosis               string `json:"diagnosis"`
	Conclusion              string `json:"conclusion"`
	BirthPlan               string `json:"birthPlan"`
}
