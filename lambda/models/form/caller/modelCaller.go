package caller

import "lambda/lambda/models/form"
import "github.com/lambda-platform/lambda/dataform"

func GetMODEL(schema_id string) dataform.Dataform {

	switch schema_id {

	case "crud_form":
		return form.KrudDataform()

	case "notification_target_form":
		return form.NotificationTargetsDataform()

	case "menu_form":
		return form.MenuFormDataform()

	case "user_form":
		return form.UserFormDataform()

	case "user_profile":
		return form.UserProfile()

	case "user_password":
		return form.UsersDataform()

	case "1":
		return form.LutNewsType1Dataform()

	case "6":
		return form.LutDocumentSorting6Dataform()

	case "308":
		return form.LutDedUndsenTorol308Dataform()

	case "264":
		return form.NariivchilsanUnelgee264Dataform()

	case "8":
		return form.LutSubDocumentSorting8Dataform()

	case "11":
		return form.LutEvalutionSector11Dataform()

	case "13":
		return form.LutConfirmStatus13Dataform()

	case "311":
		return form.LutAvsanArgaDedHemjee311Dataform()

	case "17":
		return form.LutDisasterType17Dataform()

	case "19":
		return form.LutDailyDisasterType19Dataform()

	case "15":
		return form.LutActivityType15Dataform()

	case "21":
		return form.LutGeneralEvalutionType21Dataform()

	case "263":
		return form.SubNariinUnelgeeDugnelt263Dataform()

	case "27":
		return form.LutIndicator27Dataform()

	case "23":
		return form.LutDetailedEvalutionType23Dataform()

	case "29":
		return form.LutSubIndicator29Dataform()

	case "272":
		return form.GisBaseMaps272Dataform()

	case "36":
		return form.LutIndexType36Dataform()

	case "278":
		return form.GisLegends278Dataform()

	case "276":
		return form.GisLayers276Dataform()

	case "38":
		return form.LutRatingType38Dataform()

	case "42":
		return form.EzBodlogiinBb42Dataform()

	case "46":
		return form.UndesniiZowlol46Dataform()

	case "274":
		return form.GisCategory274Dataform()

	case "317":
		return form.LutHogjilBerhsheelTorol317Dataform()

	case "49":
		return form.MergejliinZowlol49Dataform()

	case "52":
		return form.SendainUa52Dataform()

	case "54":
		return form.DundHugtsaaStrategy54Dataform()

	case "279":
		return form.Users279Dataform()

	case "319":
		return form.LutNiigmiinBaidal319Dataform()

	case "56":
		return form.HamtragchBaiguullaga56Dataform()

	case "25":
		return form.LutDisasterIndicatorType25Dataform()

	case "58":
		return form.HolbooBarih58Dataform()

	case "321":
		return form.LutMAccidentType321Dataform()

	case "270":
		return form.GisConnection270Dataform()

	case "322":
		return form.LutHumanAccidentType322Dataform()

	case "60":
		return form.Banner60Dataform()

	case "63":
		return form.News63Dataform()

	case "66":
		return form.TzBuhiiAanb66Dataform()

	case "69":
		return form.Users69Dataform()

	case "4":
		return form.LutSocialType4Dataform()

	case "40":
		return form.LutMethodolgyType40Dataform()

	case "72":
		return form.TzHugtsaaSungah72Dataform()

	case "84":
		return form.LutMaindisasterType84Dataform()

	case "91":
		return form.LutRiskreduction91Dataform()

	case "286":
		return form.SubNariinUnelgeeBagBureldehuun286Dataform()

	case "76":
		return form.EronhiiUnelgee76Dataform()

	case "304":
		return form.LutSubSubDamageType304Dataform()

	case "86":
		return form.LutSubMaindisasterType86Dataform()

	case "89":
		return form.LutRiskAssessment89Dataform()

	case "93":
		return form.LutSubRiskreduction93Dataform()

	case "96":
		return form.LutOrgSource96Dataform()

	case "98":
		return form.LutNDisasterReason98Dataform()

	case "314":
		return form.LutSubSubSubDamageType314Dataform()

	case "102":
		return form.LutEnvironment102Dataform()

	case "326":
		return form.LutSubSubMaindisasterType326Dataform()

	case "325":
		return form.HongiinMedee325Dataform()

	case "108":
		return form.LutSubAccDisasterType108Dataform()

	case "106":
		return form.LutFire106Dataform()

	case "111":
		return form.LutShaltgaanHunees111Dataform()

	case "113":
		return form.LutSubReasonHunees113Dataform()

	case "116":
		return form.LutShaltgaanGal116Dataform()

	case "118":
		return form.LutSubReasonGal118Dataform()

	case "123":
		return form.LutAnimals123Dataform()

	case "124":
		return form.LutSubAnimals124Dataform()

	case "128":
		return form.LutReason128Dataform()

	case "129":
		return form.LutSubReason129Dataform()

	case "134":
		return form.LutBiologyDisasterType134Dataform()

	case "136":
		return form.LutInfoSource136Dataform()

	case "133":
		return form.JuramGeologyGaubm133Dataform()

	case "140":
		return form.LutDedNervegdsenHun140Dataform()

	case "144":
		return form.LutSubSubNervgdsenHun144Dataform()

	case "138":
		return form.LutNervgdsenHun138Dataform()

	case "296":
		return form.NariivchilsanUnelgee296Dataform()

	case "153":
		return form.SubJuramGeoHorogdsonMal153Dataform()

	case "154":
		return form.SubJuramGeoTejeeverAmitan154Dataform()

	case "151":
		return form.SubJuramGeoHeregtseeToim151Dataform()

	case "155":
		return form.SubJuramGeoEvdersenEdZuils155Dataform()

	case "148":
		return form.SubJuramGeoNervegdsenAmitan148Dataform()

	case "160":
		return form.LutLevelType160Dataform()

	case "150":
		return form.SubJuramGeoUchirsanBoditHorhirol150Dataform()

	case "297":
		return form.LutUilAjilgaaniiChiglelAan297Dataform()

	case "147":
		return form.SubJuramGeoNervegdsenHunAm147Dataform()

	case "104":
		return form.LutAccDisasterType104Dataform()

	case "163":
		return form.LutVibration163Dataform()

	case "165":
		return form.LutAlbanTushaal165Dataform()

	case "166":
		return form.LutHeregtseeTorol166Dataform()

	case "170":
		return form.LutMeasure170Dataform()

	case "171":
		return form.LutMalTorol171Dataform()

	case "169":
		return form.LutHeregtseeDedDedTorol169Dataform()

	case "172":
		return form.LutMalDedTorol172Dataform()

	case "173":
		return form.LutMalNasTorol173Dataform()

	case "174":
		return form.LutMalHuis174Dataform()

	case "175":
		return form.LutAvsanArgaHemjee175Dataform()

	case "77":
		return form.SubDugnelt77Dataform()

	case "176":
		return form.LutHorogdsonAmitan176Dataform()

	case "177":
		return form.LutTejeeverAmitniiTurul177Dataform()

	case "178":
		return form.LutEdHorongo178Dataform()

	case "179":
		return form.LutDedEdHorongo179Dataform()

	case "180":
		return form.LutHohiroliinBaidal180Dataform()

	case "149":
		return form.SubJuramGeoAvrahSergeeh149Dataform()

	case "299":
		return form.LutDamageType299Dataform()

	case "168":
		return form.LutHeregtseeDedTorol168Dataform()

	case "300":
		return form.LutSubDamageType300Dataform()

	case "208":
		return form.LutHunAmTorol208Dataform()

	case "207":
		return form.SubJuramGeoNervegdsenAmitan207Dataform()

	case "211":
		return form.LutHunAmDedDedTorol211Dataform()

	case "216":
		return form.LutDedDedEdHorongo216Dataform()

	case "224":
		return form.UsTsagUurGaraltaiAub224Dataform()

	case "221":
		return form.LutAjillasanBaidalGal221Dataform()

	case "219":
		return form.SubJuramGeoGamshigtOrtsonHunMedeelel219Dataform()

	case "227":
		return form.LutAyultUzegdelTorol227Dataform()

	case "228":
		return form.LutAyultUzegdelDedTorol228Dataform()

	case "229":
		return form.LutEhSurvalj229Dataform()

	case "230":
		return form.LutAngi230Dataform()

	case "231":
		return form.LutSalbar231Dataform()

	case "232":
		return form.LutTsol232Dataform()

	case "253":
		return form.LutHunAmDedTorol253Dataform()

	case "244":
		return form.OronNutgiinZowlol244Dataform()

	case "247":
		return form.SubSendainUaZorilt247Dataform()

	case "262":
		return form.NariivchilsanUnelgee262Dataform()

	case "306":
		return form.LutUndsenTorol306Dataform()

	}
	return dataform.Dataform{}

}
