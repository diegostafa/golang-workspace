package main

import (
	ff "github.com/xfrr/goffmpeg/transcoder"
)

var ffmpegFormats = [...]string{
	"3dostr", "3g2", "3gp", "4xm", "a64", "aa", "aac", "aax", "ac3", "ace", "acm", "act", "adf", "adp", "ads", "adts", "adx", "aea", "afc", "aiff", "aix", "alaw", "alias_pix", "alp", "alsa", "amr", "amrnb", "amrwb", "amv", "anm", "apc", "ape", "apm", "apng", "aptx", "aptx_hd", "aqtitle", "argo_asf", "argo_brp", "argo_cvg", "asf", "asf_o", "asf_stream", "ass", "ast", "au", "av1", "avi", "avif", "avisynth", "avm2", "avr", "avs", "avs2", "avs3", "bethsoftvid", "bfi", "bfstm", "bin", "bink", "binka", "bit", "bitpacked", "bmp_pipe", "bmv", "boa", "brender_pix", "brstm", "c93", "caf", "cavsvideo", "cdg", "cdxl", "cine", "codec2", "codec2raw", "concat", "crc", "cri_pipe", "dash", "data", "daud", "dcstr", "dds_pipe", "derf", "dfa", "dfpwm", "dhav", "dirac", "dnxhd", "dpx_pipe", "dsf", "dsicin", "dss", "dts", "dtshd", "dv", "dvbsub", "dvbtxt", "dvd", "dxa", "ea", "ea_cdata", "eac3", "epaf", "exr_pipe", "f32be", "f32le", "f4v", "f64be", "f64le", "fbdev", "ffmetadata", "fifo", "fifo_test", "film_cpk", "filmstrip", "fits", "flac", "flic", "flv", "framecrc", "framehash", "framemd5", "frm", "fsb", "fwse", "g722", "g723_1", "g726", "g726le", "g729", "gdv", "gem_pipe", "genh", "gif", "gif_pipe", "gsm", "gxf", "h261", "h263", "h264", "hash", "hca", "hcom", "hds", "hevc", "hls", "hnm", "ico", "idcin", "idf", "iec61883", "iff", "ifv", "ilbc", "image2", "image2pipe", "imf", "ingenient", "ipmovie", "ipod", "ipu", "ircam", "ismv", "iss", "iv8", "ivf", "ivr", "j2k_pipe", "jack", "jacosub", "jpeg_pipe", "jpegls_pipe", "jpegxl_pipe", "jv", "kmsgrab", "kux", "kvag", "latm", "lavfi", "libmodplug", "live_flv", "lmlm4", "loas", "lrc", "luodat", "lvf", "lxf", "m4v", "matroska", "matroska", "webm", "mca", "mcc", "md5", "mgsts", "microdvd", "mjpeg", "mjpeg_2000", "mkvtimestamp_v2", "mlp", "mlv", "mm", "mmf", "mods", "moflex", "mov", "mov", "mp4", "m4a", "3gp", "3g2", "mj2", "mp2", "mp3", "mp4", "mpc", "mpc8", "mpeg", "mpeg1video", "mpeg2video", "mpegts", "mpegtsraw", "mpegvideo", "mpjpeg", "mpl2", "mpsub", "msf", "msnwctcp", "msp", "mtaf", "mtv", "mulaw", "musx", "mv", "mvi", "mxf", "mxf_d10", "mxf_opatom", "mxg", "nc", "nistsphere", "nsp", "nsv", "null", "nut", "nuv", "obu", "oga", "ogg", "ogv", "oma", "opengl", "opus", "oss", "paf", "pam_pipe", "pbm_pipe", "pcx_pipe", "pfm_pipe", "pgm_pipe", "pgmyuv_pipe", "pgx_pipe", "phm_pipe", "photocd_pipe", "pictor_pipe", "pjs", "pmp", "png_pipe", "pp_bnk", "ppm_pipe", "psd_pipe", "psp", "psxstr", "pulse", "pva", "pvf", "qcp", "qdraw_pipe", "qoi_pipe", "r3d", "rawvideo", "realtext", "redspark", "rl2", "rm", "roq", "rpl", "rsd", "rso", "rtp", "rtp_mpegts", "rtsp", "s16be", "s16le", "s24be", "s24le", "s32be", "s32le", "s337m", "s8", "sami", "sap", "sbc", "sbg", "scc", "scd", "sdl", "sdl2", "sdp", "sdr2", "sds", "sdx", "segment", "ser", "sga", "sgi_pipe", "shn", "siff", "simbiosis_imx", "sln", "smjpeg", "smk", "smoothstreaming", "smush", "sol", "sox", "spdif", "spx", "srt", "stl", "stream_segment", "ssegment", "streamhash", "subviewer", "subviewer1", "sunrast_pipe", "sup", "svag", "svcd", "svg_pipe", "svs", "swf", "tak", "tedcaptions", "tee", "thp", "tiertexseq", "tiff_pipe", "tmv", "truehd", "tta", "ttml", "tty", "txd", "ty", "u16be", "u16le", "u24be", "u24le", "u32be", "u32le", "u8", "uncodedframecrc", "v210", "v210x", "vag", "vbn_pipe", "vc1", "vc1test", "vcd", "vidc", "video4linux2", "v4l2", "vividas", "vivo", "vmd", "vob", "vobsub", "voc", "vpk", "vplayer", "vqf", "w64", "wav", "wc3movie", "webm", "webm_chunk", "webm_dash_manifest", "webp", "webp_pipe", "webvtt", "wsaud", "wsd", "wsvqa", "wtv", "wv", "wve", "x11grab", "xa", "xbin", "xbm_pipe", "xmv", "xpm_pipe", "xv", "xvag", "xwd_pipe", "xwma", "yop", "yuv4mpegpipe"}

func initTranscoder(in string, out string) *ff.Transcoder {
	trans := new(ff.Transcoder)
	err := trans.Initialize(in, out)
	checkError(err)
	trans.MediaFile().SetPreset("ultrafast")
	return trans
}

func isFormatAvaiable(value string) bool {
	for _, format := range ffmpegFormats {
		if value == format {
			return true
		}
	}
	return false
}

func stringifyFormats() string {
	tmp := ""
	for _, format := range ffmpegFormats {
		tmp += format + "    "
	}
	return tmp
}
