package formModels

import (
	"github.com/lambda-platform/lambda/DB"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

type JuramGeologyGaubm133 struct {
	Aimagid                          *int       `gorm:"column:aimagid" json:"aimagid"`
	AlbanTushaalID                   *int       `gorm:"column:alban_tushaal_id" json:"alban_tushaal_id"`
	AngiID                           *int       `gorm:"column:angi_id" json:"angi_id"`
	BaGarinUseg                      *string    `gorm:"column:ba_garin_useg" json:"ba_garin_useg"`
	BaHuleenAvsanOperator            *string    `gorm:"column:ba_huleen_avsan_operator" json:"ba_huleen_avsan_operator"`
	BaNer                            *string    `gorm:"column:ba_ner" json:"ba_ner"`
	BaOvog                           *string    `gorm:"column:ba_ovog" json:"ba_ovog"`
	Bagid                            *int       `gorm:"column:bagid" json:"bagid"`
	Bairshil                         *string    `gorm:"column:bairshil" json:"bairshil"`
	Baishin                          *string    `gorm:"column:baishin" json:"baishin"`
	BatAngiID                        *int       `gorm:"column:bat_angi_id" json:"bat_angi_id"`
	BatBaiguullagaID                 *int       `gorm:"column:bat_baiguullaga_id" json:"bat_baiguullaga_id"`
	BatGariinUseg                    *string    `gorm:"column:bat_gariin_useg" json:"bat_gariin_useg"`
	BatNer                           *string    `gorm:"column:bat_ner" json:"bat_ner"`
	BatOvog                          *string    `gorm:"column:bat_ovog" json:"bat_ovog"`
	BatSalbarID                      *int       `gorm:"column:bat_salbar_id" json:"bat_salbar_id"`
	BatalAlbanTushaalID              *int       `gorm:"column:batal_alban_tushaal_id" json:"batal_alban_tushaal_id"`
	BatalAngiID                      *int       `gorm:"column:batal_angi_id" json:"batal_angi_id"`
	BatalGarUseg                     *string    `gorm:"column:batal_gar_useg" json:"batal_gar_useg"`
	BatalNer                         *string    `gorm:"column:batal_ner" json:"batal_ner"`
	BatalOvog                        *string    `gorm:"column:batal_ovog" json:"batal_ovog"`
	BatalSalbarID                    *int       `gorm:"column:batal_salbar_id" json:"batal_salbar_id"`
	BatalTsolID                      *int       `gorm:"column:batal_tsol_id" json:"batal_tsol_id"`
	BatalgaAlbanTushaalID            *int       `gorm:"column:batalga_alban_tushaal_id" json:"batalga_alban_tushaal_id"`
	BatalgaajuulaltAlbanTushaalID    *int       `gorm:"column:batalgaajuulalt_alban_tushaal_id" json:"batalgaajuulalt_alban_tushaal_id"`
	BatalgaajuulaltGarUseg           *string    `gorm:"column:batalgaajuulalt_gar_useg" json:"batalgaajuulalt_gar_useg"`
	BatalgaajuulaltNer               *string    `gorm:"column:batalgaajuulalt_ner" json:"batalgaajuulalt_ner"`
	BatalgaajuulaltOgnoo             *DB.Date   `gorm:"column:batalgaajuulalt_ognoo" json:"batalgaajuulalt_ognoo"`
	BatalgaajuulaltOvog              *string    `gorm:"column:batalgaajuulalt_ovog" json:"batalgaajuulalt_ovog"`
	BioDisasterTypeID                *int       `gorm:"column:bio_disaster_type_id" json:"bio_disaster_type_id"`
	ChandralsanBairshil              *string    `gorm:"column:chandralsan_bairshil" json:"chandralsan_bairshil"`
	CreatedAt                        *time.Time `gorm:"column:created_at" json:"created_at"`
	FormID                           *int       `gorm:"column:form_id" json:"form_id"`
	GamshigBolsonEseh                *int       `gorm:"column:gamshig_bolson_eseh" json:"gamshig_bolson_eseh"`
	GariinUseg                       *string    `gorm:"column:gariin_useg" json:"gariin_useg"`
	Gudamj                           *string    `gorm:"column:gudamj" json:"gudamj"`
	GurvanugHayag                    *string    `gorm:"column:gurvanug_hayag" json:"gurvanug_hayag"`
	HNutagHemjeeToo                  *int       `gorm:"column:h_nutag_hemjee_too" json:"h_nutag_hemjee_too"`
	HNutagHemjihNegjID               *int       `gorm:"column:h_nutag_hemjih_negj_id" json:"h_nutag_hemjih_negj_id"`
	Haashaa                          *string    `gorm:"column:haashaa" json:"haashaa"`
	HaldvartOvchinEseh               *int       `gorm:"column:haldvart_ovchin_eseh" json:"haldvart_ovchin_eseh"`
	HamarsanNurgiinHemjee            float32    `gorm:"column:hamarsan_nurgiin_hemjee" json:"hamarsan_nurgiin_hemjee"`
	HamrahHuree                      *int       `gorm:"column:hamrah_huree" json:"hamrah_huree"`
	HamrahHureeHemjihNegjID          *int       `gorm:"column:hamrah_huree_hemjih_negj_id" json:"hamrah_huree_hemjih_negj_id"`
	HamrahHureeHoyorDeeshEseh        *int       `gorm:"column:hamrah_huree_hoyor_deesh_eseh" json:"hamrah_huree_hoyor_deesh_eseh"`
	HeregjuulsenArga                 *string    `gorm:"column:heregjuulsen_arga" json:"heregjuulsen_arga"`
	Hohirol                          *string    `gorm:"column:hohirol" json:"hohirol"`
	HomroglonUstagsanMalHemjihNegjID *int       `gorm:"column:homroglon_ustagsan_mal_hemjih_negj_id" json:"homroglon_ustagsan_mal_hemjih_negj_id"`
	HorioTogtoosonOgnoo              *time.Time `gorm:"column:horio_togtooson_ognoo" json:"horio_togtooson_ognoo"`
	HorioTsutsalsanOgnoo             *time.Time `gorm:"column:horio_tsutsalsan_ognoo" json:"horio_tsutsalsan_ognoo"`
	HorogdsonMalAravDeeshEseh        *int       `gorm:"column:horogdson_mal_arav_deesh_eseh" json:"horogdson_mal_arav_deesh_eseh"`
	HoyorBaDeeshEseh                 *int       `gorm:"column:hoyor_ba_deesh_eseh" json:"hoyor_ba_deesh_eseh"`
	HuHuleenavsanOgnoo               *time.Time `gorm:"column:hu_huleenavsan_ognoo" json:"hu_huleenavsan_ognoo"`
	HyzgaarlahDeglemTogtoosonOgnoo   *time.Time `gorm:"column:hyzgaarlah_deglem_togtooson_ognoo" json:"hyzgaarlah_deglem_togtooson_ognoo"`
	HyzgaarlahDeglemTsutsalsanOgnoo  *time.Time `gorm:"column:hyzgaarlah_deglem_tsutsalsan_ognoo" json:"hyzgaarlah_deglem_tsutsalsan_ognoo"`
	ID                               int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	InfoSouceID                      *int       `gorm:"column:info_souce_id" json:"info_souce_id"`
	MOHuis                           *string    `gorm:"column:m_o_huis" json:"m_o_huis"`
	MONer                            *string    `gorm:"column:m_o_ner" json:"m_o_ner"`
	MOOvog                           *string    `gorm:"column:m_o_ovog" json:"m_o_ovog"`
	MOTovchUtga                      *string    `gorm:"column:m_o_tovch_utga" json:"m_o_tovch_utga"`
	MOUtas                           *int       `gorm:"column:m_o_utas" json:"m_o_utas"`
	MailEzlehHuvi                    float32    `gorm:"column:mail_ezleh_huvi" json:"mail_ezleh_huvi"`
	MalOvchilsonEseh                 *int       `gorm:"column:mal_ovchilson_eseh" json:"mal_ovchilson_eseh"`
	MalOvchilsonMaliinToo            float32    `gorm:"column:mal_ovchilson_maliin_too" json:"mal_ovchilson_maliin_too"`
	MalUchirsanHohirolHoyorEseh      *int       `gorm:"column:mal_uchirsan_hohirol_hoyor_eseh" json:"mal_uchirsan_hohirol_hoyor_eseh"`
	MeasureID                        *int       `gorm:"column:measure_id" json:"measure_id"`
	NasBarsanHun                     float32    `gorm:"column:nas_barsan_hun" json:"nas_barsan_hun"`
	NasBarsanHunHemjihNegjID         *int       `gorm:"column:nas_barsan_hun_hemjih_negj_id" json:"nas_barsan_hun_hemjih_negj_id"`
	NbCode                           *string    `gorm:"column:nb_code" json:"nb_code"`
	NbFile                           *string    `gorm:"column:nb_file" json:"nb_file"`
	NervegdesenHunEseh               *int       `gorm:"column:nervegdesen_hun_eseh" json:"nervegdesen_hun_eseh"`
	NervegdsenHunToo                 *int       `gorm:"column:nervegdsen_hun_too" json:"nervegdsen_hun_too"`
	NhHemjihNegjID                   *int       `gorm:"column:nh_hemjih_negj_id" json:"nh_hemjih_negj_id"`
	NutagDevsgerEseh                 *int       `gorm:"column:nutag_devsger_eseh" json:"nutag_devsger_eseh"`
	Ognoo                            *DB.Date   `gorm:"column:ognoo" json:"ognoo"`
	OnchingarsanOgnoo                *time.Time `gorm:"column:onchingarsan_ognoo" json:"onchingarsan_ognoo"`
	Orgorog                          *string    `gorm:"column:orgorog" json:"orgorog"`
	OrgorogH                         *string    `gorm:"column:orgorog_h" json:"orgorog_h"`
	OrgorogHoyor                     *string    `gorm:"column:orgorog_hoyor" json:"orgorog_hoyor"`
	OvchilsonMalHemjihNegjID         *int       `gorm:"column:ovchilson_mal_hemjih_negj_id" json:"ovchilson_mal_hemjih_negj_id"`
	OvchinGarsanEseh                 *int       `gorm:"column:ovchin_garsan_eseh" json:"ovchin_garsan_eseh"`
	OvchinNerCode                    *string    `gorm:"column:ovchin_ner_code" json:"ovchin_ner_code"`
	OvogNer                          *string    `gorm:"column:ovog_ner" json:"ovog_ner"`
	SalbarID                         *int       `gorm:"column:salbar_id" json:"salbar_id"`
	ShalguurUzuulelt                 *string    `gorm:"column:shalguur_uzuulelt" json:"shalguur_uzuulelt"`
	ShaltgaanDedReasonID             *int       `gorm:"column:shaltgaan_ded_reason_id" json:"shaltgaan_ded_reason_id"`
	ShaltgaanReasonID                *int       `gorm:"column:shaltgaan_reason_id" json:"shaltgaan_reason_id"`
	ShinjIlersenOgnoo                *time.Time `gorm:"column:shinj_ilersen_ognoo" json:"shinj_ilersen_ognoo"`
	SourceBioID                      *int       `gorm:"column:source_bio_id" json:"source_bio_id"`
	Sumid                            *int       `gorm:"column:sumid" json:"sumid"`
	TavDeeshEseh                     *int       `gorm:"column:tav_deesh_eseh" json:"tav_deesh_eseh"`
	TestID                           *int       `gorm:"column:test_id" json:"test_id"`
	Toim                             *string    `gorm:"column:toim" json:"toim"`
	TosovtEzelehHuvi                 float32    `gorm:"column:tosovt_ezeleh_huvi" json:"tosovt_ezeleh_huvi"`
	TsolID                           *int       `gorm:"column:tsol_id" json:"tsol_id"`
	TsolsID                          *int       `gorm:"column:tsols_id" json:"tsols_id"`
	UchirsanBoditHohirolEseh         *int       `gorm:"column:uchirsan_bodit_hohirol_eseh" json:"uchirsan_bodit_hohirol_eseh"`
	UchirsanBoditHohirolHemjihNegjID *int       `gorm:"column:uchirsan_bodit_hohirol_hemjih_negj_id" json:"uchirsan_bodit_hohirol_hemjih_negj_id"`
	UpdatedAt                        *time.Time `gorm:"column:updated_at" json:"updated_at"`
	Urtrag                           *string    `gorm:"column:urtrag" json:"urtrag"`
	UrtragH                          *string    `gorm:"column:urtrag_h" json:"urtrag_h"`
	UrtragNeg                        *string    `gorm:"column:urtrag_neg" json:"urtrag_neg"`
	UstagsanMalAravEseh              *int       `gorm:"column:ustagsan_mal_arav_eseh" json:"ustagsan_mal_arav_eseh"`
	UstgaliinBairshil                *string    `gorm:"column:ustgaliin_bairshil" json:"ustgaliin_bairshil"`
	YronhiiMedeelel                  *string    `gorm:"column:yronhii_medeelel" json:"yronhii_medeelel"`
	Zardal                           *string    `gorm:"column:zardal" json:"zardal"`
}

//  TableName sets the insert table name for this struct type
func (j *JuramGeologyGaubm133) TableName() string {
	return "juram_geology_gaubm"
}
