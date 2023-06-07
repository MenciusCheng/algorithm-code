package main

// 头条V2.0 广告系列维度结构体
type PromotionV2Data struct {
	AdvertiserId      int64   `json:"advertiser_id"`
	Budget            float64 `json:"budget"`
	CpaBid            float64 `json:"cpa_bid"`
	DeepCpaBid        float64 `json:"deep_cpabid"`
	MaterialScoreInfo struct {
		LowQualityMaterialList struct {
			LowQualityImageIds []string `json:"low_quality_image_ids"`
			LowQualityVideoIds []string `json:"low_quality_video_ids"`
		} `json:"low_quality_material_list"`
		MaterialAdvice       []string `json:"material_advice"`
		ScoreNumOfMaterial   string   `json:"score_num_of_material"`
		ScoreTypeOfMaterial  string   `json:"score_type_of_material"`
		ScoreValueOfMaterial string   `json:"score_value_of_material"`
	} `json:"material_score_info"`
	OptStatus           string `json:"opt_status"`
	ProjectId           int64  `json:"project_id"`
	PromotionCreateTime string `json:"promotion_create_time"`
	PromotionId         int64  `json:"promotion_id"`
	PromotionMaterials  struct {
		CallToActionButtons []string `json:"call_to_action_buttons"`
		ProductInfo         struct {
			ImageIds      []string `json:"image_ids"`
			SellingPoints []string `json:"selling_points"`
			Titles        []string `json:"titles"`
		} `json:"product_info"`
		TitleMaterialList []struct {
			Title    string  `json:"title"`
			WordList []int64 `json:"word_list"`
		} `json:"title_material_list"`
		VideoMaterialList []struct {
			ImageMode      string `json:"image_mode"`
			MaterialId     int64  `json:"material_id"`
			MaterialStatus string `json:"material_status"`
			VideoCoverId   string `json:"video_cover_id"`
			VideoId        string `json:"video_id"`
			ItemId         int64  `json:"item_id"`
		} `json:"video_material_list"`
		AnchorMaterialList []struct {
			AnchorType string `json:"anchor_type"`
			AnchorId   string `json:"anchor_id"`
		} `json:"anchor_material_list"`
		ComponentMaterialList []struct {
			ComponentId             int64    `json:"component_id"`
			ExternalUrlMaterialList []string `json:"external_url_material_list"`
		} `json:"component_material_list"`
		WebUrlMaterialList []string `json:"web_url_material_list"`
	} `json:"promotion_materials"`
	PromotionModifyTime string  `json:"promotion_modify_time"`
	PromotionName       string  `json:"promotion_name"`
	RoiGoal             float64 `json:"roi_goal"`
	Status              string  `json:"status"`
	// 下面是修正字段
	LearningPhase string `json:"learning_phase"`
	NativeSetting struct {
		AwEmeId           string `json:"aweme_id"`
		IsFeedAndFavSee   string `json:"is_feed_and_fav_see"`
		AnchorRelatedType string `json:"anchor_related_type"`
	} `json:"native_setting"`
	MaterialsType    string `json:"materials_type"`
	Source           string `json:"source"`
	IsCommentDisable string `json:"is_comment_disable"`
	AdDownloadStatus string `json:"ad_download_status"`
}
