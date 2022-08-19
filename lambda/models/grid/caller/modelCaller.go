package caller

import "lambda/lambda/models/grid"
import "github.com/lambda-platform/lambda/datagrid"

func GetMODEL(schema_id string) datagrid.Datagrid {

	switch schema_id {

	case "crud_grid":
		return grid.KrudGridDatagrid

	case "crud_log":
		return grid.CrudLogDatagrid

	case "analytic_grid":
		return grid.AnalyticGridDatagrid

	case "menu_grid":
		return grid.MenuGridDatagrid

	case "notification_target_grid":
		return grid.NotificationTargetDatagrid

	case "2":
		return grid.LutNewsType2Datagrid

	case "5":
		return grid.LutSocialType5Datagrid

	case "266":
		return grid.ViewNariivchilsanUnelgee266Datagrid

	case "307":
		return grid.LutUndsenTorol307Datagrid

	case "12":
		return grid.LutEvalutionSector12Datagrid

	case "14":
		return grid.LutConfirmStatus14Datagrid

	case "310":
		return grid.ViewDedUndsenTorol310Datagrid

	case "16":
		return grid.LutActivityType16Datagrid

	case "18":
		return grid.LutDisasterType18Datagrid

	case "269":
		return grid.ViewNariichivlsanUnelgeeAdmin269Datagrid

	case "20":
		return grid.LutDailyDisasterType20Datagrid

	case "313":
		return grid.ViewAvsanDedArgaHemjee313Datagrid

	case "22":
		return grid.LutGeneralEvalutionType22Datagrid

	case "24":
		return grid.LutDetailedEvalutionType24Datagrid

	case "26":
		return grid.LutDisasterIndicatorType26Datagrid

	case "271":
		return grid.GisConnection271Datagrid

	case "273":
		return grid.GisBaseMaps273Datagrid

	case "39":
		return grid.LutRatingType39Datagrid

	case "37":
		return grid.LutIndexType37Datagrid

	case "316":
		return grid.ViewSubSubSubDamageType316Datagrid

	case "41":
		return grid.LutMethodolgyType41Datagrid

	case "277":
		return grid.GisLayers277Datagrid

	case "275":
		return grid.GisCategory275Datagrid

	case "45":
		return grid.ViewEzBodlogiinBb45Datagrid

	case "57":
		return grid.HamtragchBaiguullaga57Datagrid

	case "281":
		return grid.ViewUsers281Datagrid

	case "318":
		return grid.LutHogjilBerhsheelTorol318Datagrid

	case "320":
		return grid.LutNiigmiinBaidal320Datagrid

	case "323":
		return grid.LutHumanAccidentType323Datagrid

	case "71":
		return grid.ViewAanbTzErhOlgoh71Datagrid

	case "59":
		return grid.HolbooBarih59Datagrid

	case "62":
		return grid.ViewBanner62Datagrid

	case "10":
		return grid.ViewSubSortingDoc10Datagrid

	case "81":
		return grid.ViewDocSort81Datagrid

	case "324":
		return grid.LutMAccidentType324Datagrid

	case "85":
		return grid.LutMaindisasterType85Datagrid

	case "75":
		return grid.ViewTzSungalt75Datagrid

	case "90":
		return grid.LutRiskAssessment90Datagrid

	case "92":
		return grid.LutRiskreduction92Datagrid

	case "79":
		return grid.ViewEronhiiUnelgee79Datagrid

	case "88":
		return grid.ViewSubMaindisasterType88Datagrid

	case "68":
		return grid.ViewTzAanb68Datagrid

	case "97":
		return grid.LutOrgSource97Datagrid

	case "95":
		return grid.ViewSubRiskreduction95Datagrid

	case "99":
		return grid.LutNDisasterReason99Datagrid

	case "103":
		return grid.LutEnvironment103Datagrid

	case "105":
		return grid.LutAccDisasterType105Datagrid

	case "107":
		return grid.LutFire107Datagrid

	case "110":
		return grid.ViewSubAccDisasterType110Datagrid

	case "328":
		return grid.ViewDedDedAyulTorol328Datagrid

	case "112":
		return grid.LutShaltgaanHunees112Datagrid

	case "115":
		return grid.ViewSubShaltgaanHunees115Datagrid

	case "117":
		return grid.LutShaltgaanGal117Datagrid

	case "120":
		return grid.ViewSubGalShaltgaan120Datagrid

	case "126":
		return grid.LutAnimals126Datagrid

	case "127":
		return grid.ViewSubAnimals127Datagrid

	case "131":
		return grid.ViewSubReason131Datagrid

	case "132":
		return grid.LutReason132Datagrid

	case "135":
		return grid.LutBiologyDisasterType135Datagrid

	case "137":
		return grid.LutInfoSource137Datagrid

	case "162":
		return grid.ViewSubIndicators162Datagrid

	case "146":
		return grid.ViewDedNervegdsenHun146Datagrid

	case "145":
		return grid.ViewDedDedNervegdsenHun145Datagrid

	case "159":
		return grid.LutLevelType159Datagrid

	case "157":
		return grid.ViewJuramGeo157Datagrid

	case "161":
		return grid.LutLevelType161Datagrid

	case "164":
		return grid.LutVibration164Datagrid

	case "218":
		return grid.ViewDedDedEdHorongo218Datagrid

	case "167":
		return grid.LutHeregtseeTorol167Datagrid

	case "298":
		return grid.LutUilAjilgaaniiChiglelAan298Datagrid

	case "187":
		return grid.LutMalTorol187Datagrid

	case "181":
		return grid.LutAlbanTushaal181Datagrid

	case "183":
		return grid.ViewHeregtseeDedTorol183Datagrid

	case "185":
		return grid.ViewHeregtseeDedDedTorol185Datagrid

	case "186":
		return grid.LutMeasure186Datagrid

	case "192":
		return grid.LutAvsanArgaHemjee192Datagrid

	case "189":
		return grid.ViewMalDedTorol189Datagrid

	case "190":
		return grid.LutMalNasTorol190Datagrid

	case "191":
		return grid.LutMalHuis191Datagrid

	case "194":
		return grid.LutTejeeverAmitniiTurul194Datagrid

	case "197":
		return grid.ViewEdHorongoDedTorol197Datagrid

	case "198":
		return grid.LutHohiroliinBaidal198Datagrid

	case "193":
		return grid.LutHorogdsonAmitan193Datagrid

	case "195":
		return grid.LutEdHorongo195Datagrid

	case "209":
		return grid.LutHunAmTorol209Datagrid

	case "302":
		return grid.ViewSubDamageType302Datagrid

	case "303":
		return grid.LutDamageType303Datagrid

	case "215":
		return grid.ViewHunAmDedDedTorol215Datagrid

	case "139":
		return grid.LutNervgdsenHun139Datagrid

	case "220":
		return grid.LutAjillasanBaidalGal220Datagrid

	case "223":
		return grid.ViewIndicator223Datagrid

	case "226":
		return grid.ViewUsTsagUur226Datagrid

	case "234":
		return grid.LutAngi234Datagrid

	case "235":
		return grid.ViewSalbar235Datagrid

	case "236":
		return grid.LutTsol236Datagrid

	case "239":
		return grid.LutAyultUzegdelTorol239Datagrid

	case "238":
		return grid.ViewAyultUzegdelDedTorol238Datagrid

	case "305":
		return grid.LutSubSubDamageType305Datagrid

	case "330":
		return grid.ViewHongiinMedee330Datagrid

	case "241":
		return grid.ViewMergejilZovlol241Datagrid

	case "243":
		return grid.ViewUndesniiZovlol243Datagrid

	case "246":
		return grid.ViewOronNutagZovlol246Datagrid

	case "249":
		return grid.ViewSendain249Datagrid

	case "251":
		return grid.ViewDundHugatsaaStrategy251Datagrid

	case "255":
		return grid.ViewHunAmDedTorol255Datagrid

	case "65":
		return grid.ViewNews65Datagrid

	}
	return datagrid.Datagrid{}

}
