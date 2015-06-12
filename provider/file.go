package provider

import (
	"fmt"
	"strings"
)

var mime_types map[string][]string
var file_extensions map[string][]string

func NewFile(locale, fallback_locale string) *File {
	application_mime_types := []string{

		"application/atom+xml", // Atom feeds
		"application/ecmascript",
		// ECMAScript/JavaScript; Defined in RFC 4329 (equivalent to application/javascript but with stricter processing rules)
		"application/EDI-X12",    // EDI X12 data; Defined in RFC 1767
		"application/EDIFACT",    // EDI EDIFACT data; Defined in RFC 1767
		"application/json",       // JavaScript Object Notation JSON; Defined in RFC 4627
		"application/javascript", // ECMAScript/JavaScript; Defined in RFC 4329 (equivalent to application/ecmascript
		//   but with looser processing rules) It is not accepted in IE 8
		//   or earlier - text/javascript is accepted but it is defined as obsolete in RFC 4329.
		//   The "type" attribute of the <script> tag in HTML5 is optional and in practice
		//   omitting the media type of JavaScript programs is the most interoperable
		//   solution since all browsers have always assumed the correct
		//   default even before HTML5.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        "application/octet-stream",  // Arbitrary binary data.[6] Generally speaking this type identifies files that are not associated with a specific application. Contrary to past assumptions by software packages such as Apache this is not a type that should be applied to unknown files. In such a case, a server or application should not indicate a content type, as it may be incorrect, but rather, should omit the type in order to allow the recipient to guess the type.[7]
		"application/ogg", // Ogg, a multimedia bitstream container format; Defined in RFC 5334
		"application/pdf", // Portable Document Format, PDF has been in use for document exchange
		//   on the Internet since 1993; Defined in RFC 3778
		"application/postscript", // PostScript; Defined in RFC 2046
		"application/rdf+xml",    // Resource Description Framework; Defined by RFC 3870
		"application/rss+xml",    // RSS feeds
		"application/soap+xml",   // SOAP; Defined by RFC 3902
		"application/font-woff",  // Web Open Font Format; (candidate recommendation; use application/x-font-woff
		//   until standard is official)
		"application/xhtml+xml", // XHTML; Defined by RFC 3236
		"application/xml-dtd",   // DTD files; Defined by RFC 3023
		"application/xop+xml",   // XOP
		"application/zip",       // ZIP archive files; Registered[8]
		"application/gzip",      // Gzip, Defined in RFC 6713
	}

	audio_mime_types := []string{
		"audio/basic",            // mulaw audio at 8 kHz, 1 channel; Defined in RFC 2046
		"audio/L24",              // 24bit Linear PCM audio at 8-48 kHz, 1-N channels; Defined in RFC 3190
		"audio/mp4",              // MP4 audio
		"audio/mpeg",             // MP3 or other MPEG audio; Defined in RFC 3003
		"audio/ogg",              // Ogg Vorbis, Speex, Flac and other audio; Defined in RFC 5334
		"audio/vorbis",           // Vorbis encoded audio; Defined in RFC 5215
		"audio/vnd.rn-realaudio", // RealAudio; Documented in RealPlayer Help[9]
		"audio/vnd.wave",         // WAV audio; Defined in RFC 2361
		"audio/webm",             // WebM open media format
	}

	image_mime_types := []string{
		"image/gif",  // GIF image; Defined in RFC 2045 and RFC 2046
		"image/jpeg", // JPEG JFIF image; Defined in RFC 2045 and RFC 2046
		"image/pjpeg",
		// JPEG JFIF image; Associated with Internet Explorer; Listed in ms775147(v=vs.85) - Progressive JPEG, initiated before global browser support for progressive JPEGs (Microsoft and Firefox).
		"image/png",                // Portable Network Graphics; Registered,[10] Defined in RFC 2083
		"image/svg+xml",            // SVG vector image; Defined in SVG Tiny 1.2 Specification Appendix M
		"image/tiff",               // Tag Image File Format (only for Baseline TIFF); Defined in RFC 3302
		"image/vnd.microsoft.icon", // ICO image; Registered[11]
	}

	message_mime_types := []string{
		"message/http",     // Defined in RFC 2616
		"message/imdn+xml", // IMDN Instant Message Disposition Notification; Defined in RFC 5438
		"message/partial",  // Email; Defined in RFC 2045 and RFC 2046
		"message/rfc822",   // Email; EML files, MIME files, MHT files, MHTML files; Defined in RFC 2045 and RFC 2046
	}

	model_mime_types := []string{
		"model/example",    // Defined in RFC 4735
		"model/iges",       // IGS files, IGES files; Defined in RFC 2077
		"model/mesh",       // MSH files, MESH files; Defined in RFC 2077, SILO files
		"model/vrml",       // WRL files, VRML files; Defined in RFC 2077
		"model/x3d+binary", // X3D ISO standard for representing 3D computer graphics, X3DB binary files
		"model/x3d+vrml",   // X3D ISO standard for representing 3D computer graphics, X3DV VRML files
		"model/x3d+xml",    // X3D ISO standard for representing 3D computer graphics, X3D XML files
	}

	multipart_mime_types := []string{
		"multipart/mixed",       // MIME Email; Defined in RFC 2045 and RFC 2046
		"multipart/alternative", // MIME Email; Defined in RFC 2045 and RFC 2046
		"multipart/related",     // MIME Email; Defined in RFC 2387 and used by MHTML (HTML mail)
		"multipart/form-data",   // MIME Webform; Defined in RFC 2388
		"multipart/signed",      // Defined in RFC 1847
		"multipart/encrypted",   // Defined in RFC 1847
	}

	text_mime_types := []string{
		"text/cmd",  // commands; subtype resident in Gecko browsers like Firefox 3.5
		"text/css",  // Cascading Style Sheets; Defined in RFC 2318
		"text/csv",  // Comma-separated values; Defined in RFC 4180
		"text/html", // HTML; Defined in RFC 2854
		"text/javascript",
		// (Obsolete): JavaScript; Defined in and obsoleted by RFC 4329 in order to discourage its usage in favor of application/javascript. However, text/javascript is allowed in HTML 4 and 5 and, unlike application/javascript, has cross-browser support. The "type" attribute of the <script> tag in HTML5 is optional and there is no need to use it at all since all browsers have always assumed the correct default (even in HTML 4 where it was required by the specification).
		"text/plain", // Textual data; Defined in RFC 2046 and RFC 3676
		"text/vcard", // vCard (contact information); Defined in RFC 6350
		"text/xml",   // Extensible Markup Language; Defined in RFC 3023
	}

	video_mime_types := []string{
		"video/mpeg",       // MPEG-1 video with multiplexed audio; Defined in RFC 2045 and RFC 2046
		"video/mp4",        // MP4 video; Defined in RFC 4337
		"video/ogg",        // Ogg Theora or other video (with audio); Defined in RFC 5334
		"video/quicktime",  // QuickTime video; Registered[12]
		"video/webm",       // WebM Matroska-based open media format
		"video/x-matroska", // Matroska open media format
		"video/x-ms-wmv",   // Windows Media Video; Documented in Microsoft KB 288102
		"video/x-flv",      // Flash video (FLV files)
	}

	mime_types = map[string][]string{
		"application": application_mime_types,
		"audio":       audio_mime_types,
		"image":       image_mime_types,
		"message":     message_mime_types,
		"model":       model_mime_types,
		"multipart":   multipart_mime_types,
		"text":        text_mime_types,
		"video":       video_mime_types,
	}

	audio_file_extenstions := []string{
		"flac",
		"mp3",
		"wav",
	}

	image_file_extenstions := []string{
		"bmp",
		"gif",
		"jpeg",
		"jpg",
		"png",
		"tiff",
	}

	text_file_extensions := []string{
		"css",
		"csv",
		"html",
		"js",
		"json",
		"txt",
	}

	video_file_extensions := []string{
		"mp4",
		"avi",
		"mov",
		"webm",
	}

	file_extensions = map[string][]string{
		"audio": audio_file_extenstions,
		"image": image_file_extenstions,
		"text":  text_file_extensions,
		"video": video_file_extensions,
	}

	return &File{Provider{locale, fallback_locale, "file"}}
}

func (file *File) MimeType(category string) string {
	mime_category := file.mimeTypeCategory(category)
	return mime_types[mime_category][file.Provider.Number(len(mime_types[mime_category]))]
}

func (file *File) Name(category string) string {
	// For now lets use the Person provider, while word generator is not implemented
	return strings.ToLower(fmt.Sprintf("%v.%v", NewPerson(file.Provider.Locale, file.Provider.FallbackLocale).FirstName(), file.Extension(category)))
}

func (file *File) Extension(category string) string {
	file_category := file.fileCategory(category)
	return file_extensions[file_category][file.Provider.Number(len(file_extensions[file_category]))]
}

func (file *File) mimeTypeCategory(mime_type string) string {
	if _, ok := mime_types[mime_type]; ok && mime_type != "" {
		return mime_type
	} else {
		types := make([]string, 0, len(mime_types))
		for k := range mime_types {
			types = append(types, k)
		}
		return types[file.Provider.Number(len(types))]
	}
}

func (file *File) fileCategory(file_category string) string {
	if _, ok := file_extensions[file_category]; ok && file_category != "" {
		return file_category
	} else {
		types := make([]string, 0, len(file_extensions))
		for k := range file_extensions {
			types = append(types, k)
		}
		return types[file.Provider.Number(len(types))]
	}
}
