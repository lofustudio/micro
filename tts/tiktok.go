package tts

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"

	"github.com/matthew-balzan/dca"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

type TiktokTTS struct{}

// Name implements TTS.
func (*TiktokTTS) Name() string {
	return "tiktok"
}

// Voices implements TTS.
func (*TiktokTTS) Voices() []string {
	// https://github.com/oscie57/tiktok-voice/issues/1#issuecomment-1636963982
	return []string{"en_uk_001", "en_uk_003", "en_female_emotional", "en_au_001", "en_au_002", "en_us_002", "en_us_006", "en_us_007", "en_us_009", "en_us_010", "en_female_samc", "en_male_cody", "en_male_narration", "en_male_funny", "en_male_jarvis", "en_male_santa_narration", "en_female_betty", "en_female_makeup", "en_female_richgirl", "en_male_cupid", "en_female_shenna", "en_male_ghosthost", "en_female_grandma", "en_male_ukneighbor", "en_male_wizard", "en_male_trevor", "en_male_deadpool", "en_male_ukbutler", "en_male_petercullen", "en_male_pirate", "en_male_santa", "en_male_santa_effect", "en_female_pansino", "en_male_grinch", "en_us_ghostface", "en_us_chewbacca", "en_us_c3po", "en_us_stormtrooper", "en_us_stitch", "en_us_rocket", "en_female_madam_leota", "en_male_sing_deep_jingle", "en_male_m03_classical", "en_female_f08_salut_damour", "en_male_m2_xhxs_m03_christmas", "en_female_f08_warmy_breeze", "en_female_ht_f08_halloween", "en_female_ht_f08_glorious", "en_male_sing_funny_it_goes_up", "en_male_m03_lobby", "en_female_ht_f08_wonderful_world", "en_female_ht_f08_newyear", "en_male_sing_funny_thanksgiving", "en_male_m03_sunshine_soon", "en_female_f08_twinkle", "en_male_m2_xhxs_m03_silly", "fr_001", "fr_002", "de_001", "de_002", "id_male_darma", "id_female_icha", "id_female_noor", "id_male_putra", "it_male_m18", "jp_001", "jp_003", "jp_005", "jp_006", "jp_male_osada", "jp_male_matsuo", "jp_female_machikoriiita", "jp_male_matsudake", "jp_male_shuichiro", "jp_female_rei", "jp_male_hikakin", "jp_female_yagishaki", "kr_002", "kr_004", "kr_003", "br_003", "br_004", "br_005", "pt_female_lhays", "pt_female_laizza", "pt_male_transformer", "es_002", "es_male_m3", "es_female_f6", "es_female_fp1", "es_mx_002", "es_mx_male_transformer", "es_mx_female_supermom"}
}

// Run implements TTS.
func (*TiktokTTS) Run(request TTSRequest) (dca.OpusReader, error) {
	// Limit text size to 200
	if utf8.RuneCountInString(request.Text) > 200 {
		return nil, errors.New("text is too long")
	}

	host, err := url.Parse("https://api16-normal-c-useast2a.tiktokv.com/media/api/text/speech/invoke/")
	if err != nil {
		return nil, err
	}

	q := host.Query()
	if request.Voice != "" {
		q.Add("text_speaker", request.Voice)
	} else {
		q.Add("text_speaker", "en_us_001")
	}
	q.Add("req_text", request.Text)
	q.Add("speaker_map_type", "0")
	q.Add("aid", "1233")
	host.RawQuery = q.Encode()

	client := &http.Client{
		Timeout: time.Second * 3,
	}
	req, err := http.NewRequest(http.MethodPost, host.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "com.zhiliaoapp.musically/2022600030 (Linux; U; Android 7.1.2; es_ES; SM-G988N; Build/NRD90M;tt-ok/3.12.13.1)")
	req.Header.Add("Cookie", "sessionid=1e6150e05b0e7e26930e9d0d3a66d4a8")
	req.Header.Add("Accept-Encoding", "deflate,compress,json")

	log.Trace().Str("url", host.String()).Msg("Tiktok request")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("status: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()

	jsonStr := string(body)

	vStr := gjson.Get(jsonStr, "data.v_str").String()
	decoded, err := base64.StdEncoding.DecodeString(vStr)
	if err != nil {
		return nil, err
	}
	vStrReader := bytes.NewReader(decoded)

	if len(vStr) < 1 {
		return nil, errors.New("empty response")
	}

	encoded, err := EncodeToOpus(vStrReader)
	if err != nil {
		return nil, err
	}

	return encoded, nil
}

var _ TTS = (*TiktokTTS)(nil)
