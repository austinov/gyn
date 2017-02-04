package store

type Profile struct {
	Id       int32  `json:"-"`
	Login    string `json:"login"`
	UserName string `json:"username"`
}

type Dictionary struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

//go:generate $GOPATH/src/github.com/austinov/go-recipes/genorm/genorm -dst-path ./pg .

//genorm:appointments
type Appointment struct {
	Id                           int64  `json:"id" genorm:"id,pk"`
	CreatedAt                    int64  `json:"-" genorm:"created_at"`
	UpdatedAt                    int64  `json:"-" genorm:"updated_at"`
	DateReceipt                  int64  `json:"dateReceipt" genorm:"date_receipt"`
	DoctorId                     int32  `json:"doctorId" genorm:"doctor_id"`
	PatientId                    int32  `json:"patientId" genorm:"patient_id"`
	ReceiptKindId                int32  `json:"receiptKindId" genorm:"receipt_kind_id"`
	ReceiptDiagnosis             string `json:"receiptDiagnosis" genorm:"receipt_diagnosis"`
	Alergo                       string `json:"alergo" docx:"alergo" genorm:"alergo"`
	ContactInfected              string `json:"contactInfected" docx:"contactInfected" genorm:"contact_infected"`
	Hiv                          string `json:"hiv" docx:"hiv" genorm:"hiv"`
	Transfusion                  string `json:"transfusion" docx:"transfusion" genorm:"transfusion"`
	Dyscountry                   string `json:"dyscountry" docx:"dyscountry" genorm:"dyscountry"`
	Smoking                      string `json:"smoking" docx:"smoking" genorm:"smoking"`
	Drugs                        string `json:"drugs" docx:"drugs" genorm:"drugs"`
	Inheritance                  string `json:"inheritance" docx:"inheritance" genorm:"inheritance"`
	Gyndiseases                  string `json:"gyndiseases" docx:"gyndiseases" genorm:"gyndiseases"`
	Paritet                      string `json:"paritet" genorm:"paritet"`
	ParitetB                     string `json:"paritetB" genorm:"paritet_b"`
	ParitetP                     string `json:"paritetP" genorm:"paritet_p"`
	ParitetA                     string `json:"paritetA" genorm:"paritet_a"`
	ParitetSV                    string `json:"paritetSV" genorm:"paritet_sv"`
	ParitetNB                    string `json:"paritetNB" genorm:"paritet_nb"`
	ParitetEB                    string `json:"paritetEB" genorm:"paritet_eb"`
	InfectionMarkersStateId      int32  `json:"infectionMarkersStateId" genorm:"infection_markers_state_id"`
	InfectionMarkersDesc         string `json:"infectionMarkersDesc" genorm:"infection_markers_desc"`
	TromboflebiaStateId          int32  `json:"tromboflebiaStateId" genorm:"tromboflebia_state_id"`
	TromboflebiaDesc             string `json:"tromboflebiaDesc" genorm:"tromboflebia_desc"`
	FirstTrimester               string `json:"firstTrimester" docx:"firstTrimester" genorm:"first_trimester"`
	SecondTrimester              string `json:"secondTrimester" docx:"secondTrimester" genorm:"second_trimester"`
	ThirdTrimester               string `json:"thirdTrimester" docx:"thirdTrimester" genorm:"third_trimester"`
	History                      string `json:"history" docx:"history" genorm:"history"`
	Oprv                         string `json:"oprv" genorm:"oprv"`
	OprvStateId                  int32  `json:"oprvStateId" genorm:"oprv_state_id"`
	ExpByMenstruation            string `json:"expByMenstruation" docx:"expByMenstruation" genorm:"exp_by_menstruation"`
	ExpByFirstVisit              string `json:"expByFirstVisit" docx:"expByFirstVisit" genorm:"exp_by_first_visit"`
	ExpByUltraFirst              string `json:"expByUltraFirst" genorm:"exp_by_ultra_first"`
	ExpByUltraSecond             string `json:"expByUltraSecond" genorm:"exp_by_ultra_second"`
	ExpByUltraThird              string `json:"expByUltraThird" genorm:"exp_by_ultra_third"`
	HealthStateId                int32  `json:"healthStateId" genorm:"health_state_id"`
	Claims                       string `json:"claims" docx:"claims" genorm:"claims"`
	Head                         string `json:"head" docx:"headache" genorm:"head"`
	Vision                       string `json:"vision" docx:"vision" genorm:"vision"`
	SkinStateId                  int32  `json:"skinStateId" genorm:"skin_state_id"`
	Lymph                        string `json:"lymph" docx:"lymph" genorm:"lymph"`
	BreathStateId                int32  `json:"breathStateId" genorm:"breath_state_id"`
	RaleStateId                  int32  `json:"raleStateId" genorm:"rale_state_id"`
	TonesStateId                 int32  `json:"tonesStateId" genorm:"tones_state_id"`
	Pulse                        string `json:"pulse" docx:"pulseValue" genorm:"pulse"`
	PulseType                    string `json:"pulseType" docx:"pulseType" genorm:"pulse_type"`
	Pressure                     string `json:"pressure" docx:"pressure" genorm:"pressure"`
	TongueClean                  bool   `json:"tongueClean" genorm:"tongue_clean"`
	TongueWet                    bool   `json:"tongueWet" genorm:"tongue_wet"`
	TongueDry                    bool   `json:"tongueDry" genorm:"tongue_dry"`
	TongueCoated                 bool   `json:"tongueCoated" genorm:"tongue_coated"`
	TongueUncoated               bool   `json:"tongueUncoated" genorm:"tongue_uncoated"`
	Throat                       string `json:"throat" docx:"throat" genorm:"throat"`
	BellyPeriod                  string `json:"bellyPeriod" docx:"bellyPeriod" genorm:"belly_period"`
	BellyStateId                 int32  `json:"bellyStateId" genorm:"belly_state_id"`
	EpigastriumStateUse          bool   `json:"epigastriumStateUse" genorm:"epigastrium_state_use"`
	EpigastriumStateId           int32  `json:"epigastriumStateId" genorm:"epigastrium_state_id"`
	ScarStateUse                 bool   `json:"scarStateUse" genorm:"scar_state_use"`
	ScarStateId                  int32  `json:"scarStateId" genorm:"scar_state_id"`
	Peritoneal                   string `json:"peritoneal" docx:"peritoneal" genorm:"peritoneal"`
	Labors                       string `json:"labors" docx:"labors" genorm:"labors"`
	Dysuric                      bool   `json:"dysuric" genorm:"dysuric"`
	Bowel                        bool   `json:"bowel" genorm:"bowel"`
	LimbSwelling                 string `json:"limbSwelling" docx:"limbSwelling" genorm:"limb_swelling"`
	UteruseStateId               int32  `json:"uteruseStateId" genorm:"uteruse_state_id"`
	FetalPositionId              int32  `json:"fetalPositionId" genorm:"fetal_position_id"`
	FetalPreviaId                int32  `json:"fetalPreviaId" genorm:"fetal_previa_id"`
	FetalAlignId                 int32  `json:"fetalAlignId" genorm:"fetal_align_id"`
	FetalHeartbeatId             int32  `json:"fetalHeartbeatId" genorm:"fetal_heartbeat_id"`
	HeartbeatRithmId             int32  `json:"heartbeatRithmId" genorm:"heartbeat_rithm_id"`
	FetalPulse                   string `json:"fetalPulse" docx:"fetalPulse" genorm:"fetal_pulse"`
	ReproductiveDischargeTypeId  int32  `json:"reproductiveDischargeTypeId" genorm:"reproductive_discharge_type_id"`
	ReproductiveDischargeStateId int32  `json:"reproductiveDischargeStateId" genorm:"reproductive_discharge_state_id"`
	Vdm                          string `json:"vdm" docx:"vdm" genorm:"vdm"`
	Oj                           string `json:"oj" docx:"oj" genorm:"oj"`
	Dspin                        string `json:"dspin" docx:"dspin" genorm:"dspin"`
	Dcrist                       string `json:"dcrist" docx:"dcrist" genorm:"dcrist"`
	Dtroch                       string `json:"dtroch" docx:"dtroch" genorm:"dtroch"`
	Cext                         string `json:"cext" docx:"cext" genorm:"cext"`
	DevelOrgansId                int32  `json:"develOrgansId" genorm:"devel_organs_id"`
	GenitalAnomalies             string `json:"genitalAnomalies" docx:"genitalAnomalies" genorm:"genital_anomalies"`
	VaginaStateId                int32  `json:"vaginaStateId" genorm:"vagina_state_id"`
	Bishop                       string `json:"bishop" docx:"bishop" genorm:"bishop"`
	FetalBladderStateId          int32  `json:"fetalBladderStateId" genorm:"fetal_bladder_state_id"`
	FetalBladderPreviaId         int32  `json:"fetalBladderPreviaId" genorm:"fetal_bladder_previa_id"`
	FetalBladderAlignId          int32  `json:"fetalBladderAlignId" genorm:"fetal_bladder_align_id"`
	Arches                       string `json:"arches" docx:"arches" genorm:"arches"`
	Conjugate                    string `json:"conjugate" docx:"conjugate" genorm:"conjugate"`
	PelvisStateId                int32  `json:"pelvisStateId" genorm:"pelvis_state_id"`
	PelvisExostosis              string `json:"pelvisExostosis" docx:"pelvisExostosis" genorm:"pelvis_exostosis"`
	PelvisDischargeTypeId        int32  `json:"pelvisDischargeTypeId" genorm:"pelvis_discharge_type_id"`
	PelvisDischargeStateId       int32  `json:"pelvisDischargeStateId" genorm:"pelvis_discharge_state_id"`
	Diagnosis                    string `json:"diagnosis" docx:"diagnosis" genorm:"diagnosis"`
	Conclusion                   string `json:"conclusion" docx:"conclusion" genorm:"conclusion"`
	BirthPlanUse                 bool   `json:"birthPlanUse" genorm:"birth_plan_use"`
	BirthPlan                    string `json:"birthPlan" genorm:"birth_plan"`
}

//genorm:vw_appointments:view
type AppointmentView struct {
	Appointment                    `genorm:",embed"`
	DoctorName                     string `json:"doctorName" docx:"doctorName" genorm:"doctor_name"`
	PatientName                    string `json:"patientName" docx:"patientName" genorm:"patient_name"`
	ReceiptKindName                string `json:"-" genorm:"receipt_kind_name"`
	InfectionMarkersStateName      string `json:"-" genorm:"infection_markers_state_name"`
	TromboflebiaStateName          string `json:"-" genorm:"tromboflebia_state_name"`
	OprvStateName                  string `json:"-" genorm:"oprv_state_name"`
	HealthStateName                string `json:"-" docx:"health" genorm:"health_state_name"`
	SkinStateName                  string `json:"-" docx:"skinStateName" genorm:"skin_state_name"`
	BreathStateName                string `json:"-" docx:"breathStateName" genorm:"breath_state_name"`
	RaleStateName                  string `json:"-" docx:"raleStateName" genorm:"rale_state_name"`
	TonesStateName                 string `json:"-" docx:"tonesStateName" genorm:"tones_state_name"`
	BellyStateName                 string `json:"-" docx:"bellyStateName" genorm:"belly_state_name"`
	EpigastriumStateName           string `json:"-" genorm:"epigastrium_state_name"`
	ScarStateName                  string `json:"-" genorm:"scar_state_name"`
	UteruseStateName               string `json:"-" docx:"uteruseStateName" genorm:"uteruse_state_name"`
	FetalPositionName              string `json:"-" docx:"fetalPositionName" genorm:"fetal_position_name"`
	FetalPreviaName                string `json:"-" docx:"fetalPreviaName" genorm:"fetal_previa_name"`
	FetalAlignName                 string `json:"-" docx:"fetalAlignName" genorm:"fetal_align_name"`
	FetalHeartbeatName             string `json:"-" docx:"fetalHeartbeatName" genorm:"fetal_heartbeat_name"`
	HeartbeatRithmName             string `json:"-" docx:"heartbeatRithmName" genorm:"heartbeat_rithm_name"`
	ReproductiveDischargeTypeName  string `json:"-" docx:"reproductiveDischargeTypeName" genorm:"reproductive_discharge_type_name"`
	ReproductiveDischargeStateName string `json:"-" docx:"reproductiveDischargeStateName" genorm:"reproductive_discharge_state_name"`
	DevelOrgansName                string `json:"-" docx:"develOrgansName" genorm:"devel_organs_name"`
	VaginaStateName                string `json:"-" docx:"vaginaStateName" genorm:"vagina_state_name"`
	FetalBladderStateName          string `json:"-" docx:"fetalBladderStateName" genorm:"fetal_bladder_state_name"`
	FetalBladderPreviaName         string `json:"-" docx:"fetalBladderPreviaName" genorm:"fetal_bladder_previa_name"`
	FetalBladderAlignName          string `json:"-" docx:"fetalBladderAlignName" genorm:"fetal_bladder_align_name"`
	PelvisStateName                string `json:"-" docx:"pelvisStateName" genorm:"pelvis_state_name"`
	PelvisDischargeTypeName        string `json:"-" docx:"pelvisDischargeTypeName" genorm:"pelvis_discharge_type_name"`
	PelvisDischargeStateName       string `json:"-" docx:"pelvisDischargeStateName" genorm:"pelvis_discharge_state_name"`
}
