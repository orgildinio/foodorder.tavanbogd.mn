package form

import (
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/dataform"
	"github.com/lambda-platform/lambda/models"
	"github.com/thedevsaddam/govalidator"
	"lambda/lambda/models/form/formModels"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

func UsTsagUurGaraltaiAub224Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Ус  цаг уурын гаралтай АҮБ",
		Identity: "id",
		Table:    "us_tsag_uur_garaltai_aub",
		Model:    new(formModels.UsTsagUurGaraltaiAub224),
		FieldTypes: map[string]string{
			"id":                                    "Text",
			"ayult_u_torol_id":                      "Select",
			"ayult_u_ded_torol_id":                  "Select",
			"aimagid":                               "Select",
			"sumid":                                 "Select",
			"bagid":                                 "Select",
			"bairshil":                              "Geographic",
			"ugt_hayg":                              "Text",
			"ovog":                                  "Text",
			"ner":                                   "Text",
			"utas":                                  "Number",
			"huis":                                  "Radio",
			"source_id":                             "Select",
			"info_source_id":                        "Select",
			"tovch_utga":                            "Textarea",
			"medee_avsan_ognoo":                     "DateTime",
			"ayult_e_ognoo":                         "DateTime",
			"ayult_d_ognoo":                         "DateTime",
			"eren_haih_ajilgaa_e_ognoo":             "DateTime",
			"eren_haih_ajilgaa_d_ognoo":             "DateTime",
			"hoishluulashgu_uil_a_e_ognoo":          "DateTime",
			"hoishluulashgu_uil_a_d_ognoo":          "DateTime",
			"user_id":                               "Text",
			"angi_id":                               "Select",
			"salbar_id":                             "Select",
			"tsol_id":                               "Select",
			"ovogs":                                 "Text",
			"ners":                                  "Text",
			"gar_useg":                              "File",
			"salhi_hurd":                            "Text",
			"salhi_hemjih_negj_id":                  "Select",
			"mondor":                                "Text",
			"m_h_negj_id":                           "Select",
			"hur_tundas":                            "Text",
			"h_h_negj_id":                           "Select",
			"agaar_tempratur":                       "Text",
			"a_te_h_negj_id":                        "Select",
			"us_tuvshin":                            "Text",
			"us_h_negj_id":                          "Select",
			"avsan_arga_file":                       "File",
			"medeelel_oruulsan_baidal":              "Textarea",
			"b_angiid":                              "Select",
			"b_salbarid":                            "Select",
			"b_tsol_id":                             "Select",
			"b_ovog":                                "Text",
			"b_ner":                                 "Text",
			"b_gar_useg":                            "File",
			"ba_angiid":                             "Select",
			"ba_salbarid":                           "Select",
			"ba_tsolid":                             "Select",
			"ba_ovog":                               "Text",
			"ba_ner":                                "Text",
			"ba_gar_useg":                           "File",
			"bat_angiid":                            "Select",
			"bat_salbarid":                          "Select",
			"bat_albantushaalid":                    "Select",
			"bat_tsolid":                            "Select",
			"bat_ovog":                              "Text",
			"bat_ner":                               "Text",
			"bat_gar_useg":                          "File",
			"hu_hu_salhi_dundaj_hurd_tunees":        "Checkbox",
			"hu_hu_salhi_dundaj_hurd":               "Number",
			"hu_hu_salhi_dundaj_hurd_h_negj_id":     "Select",
			"hu_hu_salhi_nervegdsen_hun":            "Checkbox",
			"hu_hu_salhi_nervegdsen_hun_h_negj_id":  "Select",
			"hu_hu_salhi_nas_barsan_hun":            "Checkbox",
			"hu_hu_salhi_nas_barsan_hun_too":        "Number",
			"hu_hu_salhi_nas_barsan_hun_h_negj_id":  "Select",
			"hu_hu_salhi_hamarsan_nutag_eseh":       "Checkbox",
			"hu_hu_salhi_hamarsan_nutag_huvi_too":   "Number",
			"hu_hu_salhi_hamarsan_nutag_h_negj_id":  "Select",
			"hu_hu_salhi_ulsad_hohirol_eseh":        "Checkbox",
			"hu_hu_salhi_ded_butets_eseh":           "Checkbox",
			"hu_hu_salhi_uchirsan_nutsg_tosov_eseh": "Checkbox",
			"hu_hu_salhi_uchirsan_nutsg_tosov_too":  "Number",
			"hu_hu_salhi_uchirsan_nutsg_tosov_h_negj_id": "Select",
			"hu_hu_nervegdsen_hun_too":                   "Number",
			"batalgaa_angi_id":                           "Select",
			"batalgaa_salbar_id":                         "Select",
			"batalgaa_alban_tushaal_id":                  "Select",
			"batalgaa_tsol_id":                           "Select",
			"batalgaa_ovog_id":                           "Text",
			"batalgaa_ner_id":                            "Text",
			"batalgaa_gar_useg":                          "File",
			"batalgaaj_alban_tushaal_id":                 "Select",
			"batalgaaj_ovog":                             "Text",
			"batalgaaj_ner":                              "Text",
			"batalgaaj_gar_useg":                         "File",
			"batalgaaj_ognoo":                            "Date",
			"batalgaajilt_angi_id":                       "Select",
			"batalgaajilt_salbar_id":                     "Select",
			"batalgaajilt_tsol_id":                       "Select",
			"batalgaajilt_ovog":                          "Text",
			"batalgaajilt_ner":                           "Text",
			"batalgaajilt_duudlaga_ognoo":                "DateTime",
			"batalgaajilt_gar_useg":                      "File",
			"hui_salhi":                                  "Checkbox",
			"tsas_shoroo_shuurga_eseh":                   "Checkbox",
			"ayanga_eseh":                                "Checkbox",
			"uyr_eseh":                                   "Checkbox",
			"gan_eseh":                                   "Checkbox",
			"zud_eseh":                                   "Checkbox",
			"tsas_salhi_arav_eseh":                       "Checkbox",
			"tsas_salhi_dundaj_hurd_too":                 "Number",
			"tsas_salhi_h_negjid":                        "Select",
			"tsas_als_baraa_eseh":                        "Checkbox",
			"tsas_als_baraa":                             "Number",
			"tsas_als_baraa_h_negj_id":                   "Select",
			"tsas_urgeljleh_hugatsaa_eseh":               "Checkbox",
			"tsas_urgeljleh_hugatsaa":                    "Number",
			"tsas_urgeljleh_hugatsaa_h_negj_id":          "Select",
			"tsas_nervegdsen_arav_eseh":                  "Checkbox",
			"tsas_nervegdsen_arav":                       "Number",
			"tsas_nervegdsen_arav_h_negj_id":             "Select",
			"tsas_nas_barsan_eseh":                       "Checkbox",
			"tsas_nas_barsan":                            "Number",
			"tsas_nas_barsan_h_negj_id":                  "Select",
			"tsas_hamarsan_nutag_eseh":                   "Checkbox",
			"tsas_hamarsan_nutag":                        "Number",
			"tsas_hamarsan_nutag_h_negj_id":              "Select",
			"tsas_uls_onts_eseh":                         "Checkbox",
			"tsas_ded_butets_eseh":                       "Checkbox",
			"tsas_uchirsan_nodit_eseh":                   "Checkbox",
			"ayanga_nervegdsen_hun_arav_eseh":            "Checkbox",
			"ayanga_nas_barsan_eseh":                     "Checkbox",
			"ayanga_nas_barsan":                          "Number",
			"ayanga_nas_barsan_h_negj_id":                "Select",
			"ayanga_uchirsan_b_h_hoyor_eseh":             "Checkbox",
			"ayanga_uchirsan_b_h_hoyor":                  "Number",
			"ayanga_uchirsan_b_h_hoyor_h_negj_id":        "Select",
			"uyer_gol_moron_eseh":                        "Checkbox",
			"uyer_nervegdsen_hun_eseh":                   "Checkbox",
			"uyer_nervegdsen_hun":                        "Number",
			"uyer_nervegdsen_hun_h_negj_id":              "Select",
			"uyer_nb_hun_eseh":                           "Checkbox",
			"uyer_nb_hun":                                "Number",
			"uyer_nb_hun_h_negj_id":                      "Select",
			"uyer_nervegdsen_orh_eseh":                   "Checkbox",
			"uyer_nervegdsen_orh":                        "Number",
			"uyer_nervegdsen_orh_h_negj_id":              "Select",
			"uyer_nervegdsen_aj_eseh":                    "Checkbox",
			"uyer_nervegdsen_aj":                         "Number",
			"uyer_nervegdsen_aj_h_negj_id":               "Select",
			"uyer_gemtel_uchirsan_eseh":                  "Checkbox",
			"uyer_gemtel_uchirsan":                       "Number",
			"uyer_gemtel_uchirsan_h_negj_id":             "Select",
			"uyer_ded_butets_eseh":                       "Checkbox",
			"uyer_uchirsan_eseh":                         "Checkbox",
			"uyer_uchirsan":                              "Number",
			"uyer_uchirsan_h_negj_id":                    "Select",
			"gan_jil_eseh":                               "Checkbox",
			"gan_jil":                                    "Number",
			"gan_jil_h_negj_id":                          "Select",
			"gan_agaar_chiig_eseh":                       "Checkbox",
			"gan_agaar_chiig":                            "Number",
			"gan_agaar_chiig_h_negj_id":                  "Select",
			"gan_urgelj_hugatsaa_eseh":                   "Checkbox",
			"gan_urgelj_hugatsaa":                        "Number",
			"gan_urgelj_hugatsaa_h_negj_id":              "Select",
			"gan_ner_hun_eseh":                           "Checkbox",
			"gan_ner_hun":                                "Number",
			"gan_ner_hun_h_negj_id":                      "Select",
			"gan_horogdson_mal_eseh":                     "Checkbox",
			"gan_horogdson_mal":                          "Number",
			"gan_horogdson_mal_h_negj_id":                "Select",
			"gan_hamarsan_nutag_eseh":                    "Checkbox",
			"gan_hamarsan_nutag":                         "Number",
			"gan_hamarsan_nutag_h_negj_id":               "Select",
			"gan_tuhain_oron_eseh":                       "Checkbox",
			"gan_tuhain_oron":                            "Number",
			"gan_tuhain_oron_h_negj_id":                  "Select",
			"gan_th_aj_eseh":                             "Checkbox",
			"gan_th_aj":                                  "Number",
			"gan_th_aj_h_negj_id":                        "Select",
			"gan_uchirsan_hohirol_eseh":                  "Checkbox",
			"gan_uchirsan_hohirol":                       "Number",
			"gan_uchirsan_hohirol_h_negj_id":             "Select",
			"zud_arav_honog_eseh":                        "Checkbox",
			"zud_arav_honog":                             "Number",
			"zud_arav_honog_h_negj_id":                   "Select",
			"zud_belcheer_eseh":                          "Checkbox",
			"zud_belcheer":                               "Number",
			"zud_belcheer_h_negj_id":                     "Select",
			"zud_tsas_dundaj_eseh":                       "Checkbox",
			"zud_tsas_dundaj":                            "Number",
			"zud_tsas_dundaj_h_negj_id":                  "Select",
			"zud_urgelj_hugatsaa_eseh":                   "Checkbox",
			"zud_urgelj_hugatsaa":                        "Number",
			"zud_urgelj_hugatsaa_h_negj_id":              "Select",
			"zud_ner_hun_eseh":                           "Checkbox",
			"zud_ner_hun":                                "Number",
			"zud_ner_hun_h_negj_id":                      "Select",
			"zud_horogdson_mal_eseh":                     "Checkbox",
			"zud_horogdson_mal":                          "Number",
			"zud_horogdson_mal_h_negj_id":                "Select",
			"zud_hamarsan_nutag_eseh":                    "Checkbox",
			"zud_hamarsan_nutag":                         "Number",
			"zud_hamarsan_nutag_h_negj_id":               "Select",
			"zud_uchirsan_eseh":                          "Checkbox",
			"zud_uchirsan":                               "Number",
			"zud_uchirsan_h_negj_id":                     "Select",
			"created_at":                                 "Text",
			"updated_at":                                 "Text",
			"deleted_at":                                 "Text",
			"tsas_uchirsan_bodit":                        "Number",
			"tsas_uchirsan_bodit_h_negj_id":              "Select",
			"ayanga_ner_hun":                             "Number",
			"ayanga_ner_hun_h_negj_id":                   "Select",
			"sub_juram_geo_nervegdsen_hun_am":            "SubForm",
			"sub_juram_geo_avrah_sergeeh":                "SubForm",
			"sub_juram_geo_uchirsan_bodit_horhirol":      "SubForm",
			"sub_juram_geo_heregtsee_toim":               "SubForm",
			"form_id":                                    "Select",
			"sub_juram_geo_gamshigt_ortson_hun_medeelel": "SubForm",
			"sub_juram_geo_horogdson_mal":                "SubForm",
			"sub_juram_geo_tejeever_amitan":              "SubForm",
			"sub_juram_geo_evdersen_ed_zuils":            "SubForm",
		},
		Formulas: []models.Formula{
			models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "hu_hu_salhi_dundaj_hurd_tunees",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_dundaj_hurd",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_dundaj_hurd_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_nervegdsen_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_nervegdsen_hun_too",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_nervegdsen_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_nas_barsan_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_nas_barsan_hun_too",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_nas_barsan_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_hamarsan_nutag_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_hamarsan_nutag_huvi_too",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_hamarsan_nutag_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_ulsad_hohirol_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_ded_butets_eseh",
						Prop:  "hidden",
					},
				},
				Template: "'{hui_salhi}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "hu_hu_salhi_uchirsan_nutsg_tosov_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_uchirsan_nutsg_tosov_too",
						Prop:  "hidden",
					}, models.Target{
						Field: "hu_hu_salhi_uchirsan_nutsg_tosov_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{hui_salhi}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "tsas_salhi_arav_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_salhi_dundaj_hurd_too",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_salhi_h_negjid",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_als_baraa_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_als_baraa",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_als_baraa_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_urgeljleh_hugatsaa_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_urgeljleh_hugatsaa",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_urgeljleh_hugatsaa_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nervegdsen_arav_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nervegdsen_arav",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nervegdsen_arav_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nas_barsan_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nas_barsan",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_nas_barsan_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_hamarsan_nutag_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_hamarsan_nutag",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_hamarsan_nutag_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_uls_onts_eseh",
						Prop:  "hidden",
					},
				},
				Template: "'{tsas_shoroo_shuurga_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "tsas_ded_butets_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_uchirsan_nodit_eseh",
						Prop:  "hidden",
					},
				},
				Template: "'{tsas_shoroo_shuurga_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "ayanga_nervegdsen_hun_arav_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_nas_barsan_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_nas_barsan",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_nas_barsan_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_uchirsan_b_h_hoyor_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_uchirsan_b_h_hoyor",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_uchirsan_b_h_hoyor_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{ayanga_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "uyer_gol_moron_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_hun_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nb_hun_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nb_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nb_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_orh_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_orh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_orh_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_aj_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_aj",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_nervegdsen_aj_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{uyr_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "uyer_gemtel_uchirsan_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_gemtel_uchirsan",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_gemtel_uchirsan_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_ded_butets_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_uchirsan_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_uchirsan",
						Prop:  "hidden",
					}, models.Target{
						Field: "uyer_uchirsan_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{uyr_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "gan_horogdson_mal",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_horogdson_mal_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_hamarsan_nutag_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_hamarsan_nutag",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_hamarsan_nutag_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_tuhain_oron_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_tuhain_oron",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_tuhain_oron_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_th_aj_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_th_aj",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_th_aj_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_uchirsan_hohirol_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_uchirsan_hohirol",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_uchirsan_hohirol_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{gan_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "zud_arav_honog_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_arav_honog",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_arav_honog_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_belcheer_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_belcheer",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_belcheer_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_tsas_dundaj_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_tsas_dundaj",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_tsas_dundaj_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_urgelj_hugatsaa_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_urgelj_hugatsaa",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_urgelj_hugatsaa_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_ner_hun_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_ner_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_ner_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_horogdson_mal_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_horogdson_mal",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_horogdson_mal_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{zud_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "zud_horogdson_mal",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_horogdson_mal_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_hamarsan_nutag_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_hamarsan_nutag",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_hamarsan_nutag_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_uchirsan_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_uchirsan",
						Prop:  "hidden",
					}, models.Target{
						Field: "zud_uchirsan_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{zud_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "tsas_uchirsan_bodit",
						Prop:  "hidden",
					}, models.Target{
						Field: "tsas_uchirsan_bodit_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{tsas_shoroo_shuurga_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "ayanga_ner_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "ayanga_ner_hun_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{ayanga_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "zud_tsas_dundaj_eseh",
						Prop:  "hidden",
					},
				},
				Template: "'{zud_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "gan_jil_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_jil",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_jil_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_agaar_chiig_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_agaar_chiig",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_agaar_chiig_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_urgelj_hugatsaa_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_urgelj_hugatsaa",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_urgelj_hugatsaa_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_ner_hun_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_ner_hun",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_ner_hun_h_negj_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_horogdson_mal_eseh",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_horogdson_mal",
						Prop:  "hidden",
					}, models.Target{
						Field: "gan_horogdson_mal_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{gan_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "gan_urgelj_hugatsaa_h_negj_id",
						Prop:  "hidden",
					},
				},
				Template: "'{gan_eseh}' != '1'",
				Model:    "",
				Form:     "main",
			},
		},
		ValidationRules: govalidator.MapData{

			"utas": []string{}},
		ValidationMessages: govalidator.MapData{

			"utas": []string{"mongolianMobileNumber:8 оронтой утасны дугаар оруулна уу!"}},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_nervegdsen_hun_am",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoNervegdsenHunAm147Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoNervegdsenHunAm147),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_avrah_sergeeh",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoAvrahSergeeh149Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoAvrahSergeeh149),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_uchirsan_bodit_horhirol",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoUchirsanBoditHorhirol150Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoUchirsanBoditHorhirol150),
			},
			map[string]interface{}{
				"connection_field": "",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_heregtsee_toim",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoHeregtseeToim151Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoHeregtseeToim151),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_gamshigt_ortson_hun_medeelel",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoGamshigtOrtsonHunMedeelel219Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoGamshigtOrtsonHunMedeelel219),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_horogdson_mal",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoHorogdsonMal153Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoHorogdsonMal153),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_tejeever_amitan",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoTejeeverAmitan154Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoTejeeverAmitan154),
			},
			map[string]interface{}{
				"connection_field": "us_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_juram_geo_evdersen_ed_zuils",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubJuramGeoEvdersenEdZuils155Dataform(),
				"subFormArray":     new([]formModels.SubJuramGeoEvdersenEdZuils155),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
