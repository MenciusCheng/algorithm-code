package main

import "testing"

func TestUnmarshalPromotionV2Data(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				extension: "{\"advertiser_id\":1761605854914567,\"budget\":80000,\"cpa_bid\":599.99,\"deep_cpabid\":0,\"material_score_info\":{\"low_quality_material_list\":{\"low_quality_image_ids\":null,\"low_quality_video_ids\":null},\"material_advice\":null,\"score_num_of_material\":\"\",\"score_type_of_material\":\"\",\"score_value_of_material\":\"\"},\"opt_status\":\"DISABLE\",\"project_id\":7239635813243322424,\"promotion_create_time\":\"2023-06-06 17:07:44\",\"promotion_id\":7241496563691962428,\"promotion_materials\":{\"call_to_action_buttons\":[\"极速下载\"],\"product_info\":{\"image_ids\":[\"web.business.image/202212065d0d815571613f464f38bd3f\"],\"selling_points\":[\"御姐萝莉在线聊天！\",\"美女在线陪你聊天！\",\"美女真人在线聊天！\"],\"titles\":[\"焦糖\"]},\"title_material_list\":[{\"title\":\"来焦糖，每天都有超多有趣小姐姐在线早安，还能连麦聊天！\",\"word_list\":null},{\"title\":\"失恋难过、无聊消遣，快来焦糖找个甜甜的小姐姐陪你聊天吧！\",\"word_list\":null}],\"video_material_list\":[{\"image_mode\":\"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL\",\"material_id\":7236684947607502851,\"material_status\":\"MATERIAL_STATUS_PROMOTION_DISABLE\",\"video_cover_id\":\"tos-cn-p-0051/osbMtCuAAABDVDgjr5neDkIrbEr6F3NZd8eD1h\",\"video_id\":\"v03033g10000chmttbjc77uase2c1vf0\",\"item_id\":0}],\"anchor_material_list\":[{\"anchor_type\":\"APP_INTERNET_SERVICE\",\"anchor_id\":\"7231371030764571450\"}],\"component_material_list\":null,\"web_url_material_list\":null},\"promotion_modify_time\":\"2023-06-06 20:02:10\",\"promotion_name\":\"HD-41996-IOS-御姐-（焦糖）不准不回改-面男-06.06\\n\",\"roi_goal\":0,\"status\":\"DISABLED\",\"learning_phase\":\"DEFAULT\",\"native_setting\":{\"aweme_id\":\"59625275203\",\"is_feed_and_fav_see\":\"ON\",\"anchor_related_type\":\"SELECT\"},\"materials_type\":\"\",\"source\":\"\",\"is_comment_disable\":\"ON\",\"ad_download_status\":\"\"}",
			},
		},
		{
			args: args{extension: "{}"},
		},
		{
			args: args{extension: "null"},
		},
		{
			args:    args{extension: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalPromotionV2Data(tt.args.extension); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalPromotionV2Data() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshal2PromotionV2Data(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				extension: "{\"advertiser_id\":1761605854914567,\"budget\":80000,\"cpa_bid\":599.99,\"deep_cpabid\":0,\"material_score_info\":{\"low_quality_material_list\":{\"low_quality_image_ids\":null,\"low_quality_video_ids\":null},\"material_advice\":null,\"score_num_of_material\":\"\",\"score_type_of_material\":\"\",\"score_value_of_material\":\"\"},\"opt_status\":\"DISABLE\",\"project_id\":7239635813243322424,\"promotion_create_time\":\"2023-06-06 17:07:44\",\"promotion_id\":7241496563691962428,\"promotion_materials\":{\"call_to_action_buttons\":[\"极速下载\"],\"product_info\":{\"image_ids\":[\"web.business.image/202212065d0d815571613f464f38bd3f\"],\"selling_points\":[\"御姐萝莉在线聊天！\",\"美女在线陪你聊天！\",\"美女真人在线聊天！\"],\"titles\":[\"焦糖\"]},\"title_material_list\":[{\"title\":\"来焦糖，每天都有超多有趣小姐姐在线早安，还能连麦聊天！\",\"word_list\":null},{\"title\":\"失恋难过、无聊消遣，快来焦糖找个甜甜的小姐姐陪你聊天吧！\",\"word_list\":null}],\"video_material_list\":[{\"image_mode\":\"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL\",\"material_id\":7236684947607502851,\"material_status\":\"MATERIAL_STATUS_PROMOTION_DISABLE\",\"video_cover_id\":\"tos-cn-p-0051/osbMtCuAAABDVDgjr5neDkIrbEr6F3NZd8eD1h\",\"video_id\":\"v03033g10000chmttbjc77uase2c1vf0\",\"item_id\":0}],\"anchor_material_list\":[{\"anchor_type\":\"APP_INTERNET_SERVICE\",\"anchor_id\":\"7231371030764571450\"}],\"component_material_list\":null,\"web_url_material_list\":null},\"promotion_modify_time\":\"2023-06-06 20:02:10\",\"promotion_name\":\"HD-41996-IOS-御姐-（焦糖）不准不回改-面男-06.06\n\",\"roi_goal\":0,\"status\":\"DISABLED\",\"learning_phase\":\"DEFAULT\",\"native_setting\":{\"aweme_id\":\"59625275203\",\"is_feed_and_fav_see\":\"ON\",\"anchor_related_type\":\"SELECT\"},\"materials_type\":\"\",\"source\":\"\",\"is_comment_disable\":\"ON\",\"ad_download_status\":\"\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unmarshal2PromotionV2Data(tt.args.extension); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal2PromotionV2Data() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnmarshal2PromotionV2DataXX(t *testing.T) {
	type args struct {
		extension string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				extension: "{\"advertiser_id\":1761605854914567,\"budget\":80000,\"cpa_bid\":599.99,\"deep_cpabid\":0,\"material_score_info\":{\"low_quality_material_list\":{\"low_quality_image_ids\":null,\"low_quality_video_ids\":null},\"material_advice\":null,\"score_num_of_material\":\"\",\"score_type_of_material\":\"\",\"score_value_of_material\":\"\"},\"opt_status\":\"DISABLE\",\"project_id\":7239635813243322424,\"promotion_create_time\":\"2023-06-06 17:07:44\",\"promotion_id\":7241496563691962428,\"promotion_materials\":{\"call_to_action_buttons\":[\"极速下载\"],\"product_info\":{\"image_ids\":[\"web.business.image/202212065d0d815571613f464f38bd3f\"],\"selling_points\":[\"御姐萝莉在线聊天！\",\"美女在线陪你聊天！\",\"美女真人在线聊天！\"],\"titles\":[\"焦糖\"]},\"title_material_list\":[{\"title\":\"来焦糖，每天都有超多有趣小姐姐在线早安，还能连麦聊天！\",\"word_list\":null},{\"title\":\"失恋难过、无聊消遣，快来焦糖找个甜甜的小姐姐陪你聊天吧！\",\"word_list\":null}],\"video_material_list\":[{\"image_mode\":\"CREATIVE_IMAGE_MODE_VIDEO_VERTICAL\",\"material_id\":7236684947607502851,\"material_status\":\"MATERIAL_STATUS_PROMOTION_DISABLE\",\"video_cover_id\":\"tos-cn-p-0051/osbMtCuAAABDVDgjr5neDkIrbEr6F3NZd8eD1h\",\"video_id\":\"v03033g10000chmttbjc77uase2c1vf0\",\"item_id\":0}],\"anchor_material_list\":[{\"anchor_type\":\"APP_INTERNET_SERVICE\",\"anchor_id\":\"7231371030764571450\"}],\"component_material_list\":null,\"web_url_material_list\":null},\"promotion_modify_time\":\"2023-06-06 20:02:10\",\"promotion_name\":\"HD-41996-IOS-御姐-（焦糖）不准不回改-面男-06.06\",\"roi_goal\":0,\"status\":\"DISABLED\",\"learning_phase\":\"DEFAULT\",\"native_setting\":{\"aweme_id\":\"59625275203\",\"is_feed_and_fav_see\":\"ON\",\"anchor_related_type\":\"SELECT\"},\"materials_type\":\"\",\"source\":\"\",\"is_comment_disable\":\"ON\",\"ad_download_status\":\"\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unmarshal2PromotionV2DataXX(tt.args.extension); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal2PromotionV2DataXX() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMapToJson(t *testing.T) {
	MapToJsonToMapToJson()
}
